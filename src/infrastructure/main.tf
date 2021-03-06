provider "azurerm" {
  version = "~>2.5.0"
  features {}
}

terraform {
  backend "azurerm" {
    resource_group_name  = "tfstate"
    storage_account_name = "tfstate6671"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

resource "azurerm_resource_group" "rg" {
  name     = var.rg_name
  location = var.location
}

resource "azurerm_application_insights" "functions" {
  name                = var.function_app_insights_name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  application_type    = "web"
}

resource "azurerm_storage_account" "functions" {
  name                     = var.function_storage_name
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_app_service_plan" "functions" {
  name                = var.function_service_plan_name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  kind                = "FunctionApp"
  reserved            = true

  sku {
    tier = "Dynamic"
    size = "Y1"
  }
}

resource "azurerm_function_app" "functions" {
  name                      = var.function_app_name
  location                  = azurerm_resource_group.rg.location
  resource_group_name       = azurerm_resource_group.rg.name
  app_service_plan_id       = azurerm_app_service_plan.functions.id
  storage_connection_string = azurerm_storage_account.functions.primary_connection_string
  os_type                   = "linux"

  version = "~3"

  identity {
    type = "SystemAssigned"
  }

  app_settings = {
    https_only                     = true
    FUNCTIONS_WORKER_RUNTIME       = "python"
    FUNCTION_APP_EDIT_MODE         = "readonly"
    APPINSIGHTS_INSTRUMENTATIONKEY = azurerm_application_insights.functions.instrumentation_key
    storage_name                   = azurerm_storage_account.storage.name
  }
}

resource "azurerm_storage_account" "storage" {
  name                     = var.storage_name
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_role_assignment" "storage" {
  scope                = azurerm_storage_account.storage.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azurerm_function_app.functions.identity[0].principal_id
  depends_on = [
    azurerm_function_app.functions
  ]
}
