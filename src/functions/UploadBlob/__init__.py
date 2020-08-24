import logging
import azure.functions as func
import os
from azure.identity import DefaultAzureCredential
from azure.storage.blob import BlobServiceClient, BlobClient, ContainerClient


def main(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')

    files = req.files
    files_no = len(files)

    if files_no == 0:
        return func.HttpResponse(f"There are no files in request.", status_code=500)

    storage_name = os.environ["storage_name"]
    account_url = f"https://{storage_name}.blob.core.windows.net"

    credential = DefaultAzureCredential()

    blob_service_client = BlobServiceClient(account_url, credential)

    container_name = "quickstart"

    blob_service_client.get_container_client(container_name)

    for file in files:
        blob_client = blob_service_client.get_blob_client(container=container_name, blob=file.filename)
        blob_client.upload_blob(str(file))

    return func.HttpResponse(f"There are {files_no} files in request.")