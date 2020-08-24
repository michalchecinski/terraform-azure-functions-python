# Azure Functions Python and Terraform

[![Build Status](https://dev.azure.com/michalchecinski/Terraform%20Azure%20Functions%20with%20Python/_apis/build/status/CD?branchName=master)](https://dev.azure.com/michalchecinski/Terraform%20Azure%20Functions%20with%20Python/_build/latest?definitionId=20&branchName=master)

## 📝 Overview

This is my pet project using Terraform and python with Azure Functions. I want to try out new technologies and tools that I don't use on a daily basis in my job.

## 👩‍💻 Technologies

As you may read in the title of this project uses mainly two things: Azure Functions to host Python functions and Terraform as an IaC (Infrastructure as a Code) tool.

## 🚀 CI/CD

CI/CD is based on the Azure Pipelines. You can find pipeline definitions in the [azure-pipelines folder](https://github.com/michalchecinski/terraform-azure-functions-python/tree/master/azure-pipelines) in this repo. That folder contains:

- PR check pipeline - for PR validation purposes. Azure DevOps view can be found here: [PR check pipeline link](https://dev.azure.com/michalchecinski/Terraform%20Azure%20Functions%20with%20Python/_build?definitionId=19);
- deployment pipeline - for CD of this project to the Azure environment. Azure DevOps view can be found here: [CD pipeline link](https://dev.azure.com/michalchecinski/Terraform%20Azure%20Functions%20with%20Python/_build?definitionId=20)

Also CD pipeline uses Azure Pipeline Environments to report status of deployment: [DevEnv deployment view in Azure DevOps Pipelines](https://dev.azure.com/michalchecinski/Terraform%20Azure%20Functions%20with%20Python/_environments/1?view=resources
)

## ⚙ Project Setup

If you want to setup the project on your own computer and try to play with it locally, you need to do a couple of things:

0. Clone this repo to your machine.
1. Install [terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) and [az cli](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) on your machine.
2. Run `az login` command and login to your Azure account.
3. Create a Azure Storage Account for Terraform state using [tfstate_storage_setup.sh](https://github.com/michalchecinski/terraform-azure-functions-python/blob/master/scripts/tfstate_storage_setup.sh) script.
4. Set the `ARM_ACCESS_KEY` environment variable to the storage access key that you got from script output from the previous point of this guide.
5.  Change a name of the storage in the [main.tf file in the 9th line](https://github.com/michalchecinski/terraform-azure-functions-python/blob/master/src/infrastructure/main.tf#L9) from `tfstate6671` to a name that you got from script output from the 3rd point of this guide.
6. Go to a `src/infrastructure` folder, eg. `cd src/infrastructure`.
7. Run a couple of Terraform commands:
```sh
terraform init
terraform plan # This will output the changes on the Azure env that terraform will make
terrafrom apply -auto-approve
```

Now you have prepared the Azure environment to start hacking.

Next thing is to publish Azure Functions to Azure Function App that was created. For this I highly recommend using VS Code with Azure functions extension as this is the easiest approach. You can stick to this [Microsoft Tutorial: Create and deploy serverless Azure Functions in Python with Visual Studio Code](https://docs.microsoft.com/en-us/azure/developer/python/tutorial-vs-code-serverless-python-01#azure-functions-core-tools).
