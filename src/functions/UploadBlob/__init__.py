import logging
import azure.functions as func
import os
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient, BlobClient, ContainerClient
from azure.core import ResourceExistsError


def main(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')

    files = req.files
    files_no = len(files)

    if files_no == 0:
        return func.HttpResponse(f"There are no files in request.", status_code=500)

    storage_name = "stfilesmch01" # os.environ["storage_name"]
    account_url = f"https://{storage_name}.blob.core.windows.net"

    credential = DefaultAzureCredential()

    blob_service_client = BlobServiceClient(account_url, credential)

    container_name = "http-files"

    try:
        blob_service_client.create_container(container_name)
    except ResourceExistsError:
        print("Container already exists.")

    for file_key in files:
        file = files[file_key]
        blob_client = blob_service_client.get_blob_client(container=container_name, blob=file.filename)
        blob_client.upload_blob(file.stream)

    return func.HttpResponse(f"Uploaded {files_no} files from request.")