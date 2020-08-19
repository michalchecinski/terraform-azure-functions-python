variable rg_name {
  type    = string
  default = "rg-python-function-mch01"
}

variable location {
  type    = string
  default = "westeurope"
}

variable "function_app_name" {
    type = string
    default = "func-python-mch01"
}

variable "function_service_plan_name" {
    type = string
    default = "plan-func-python-mch01"
}

variable "function_storage_name" {
    type = string
    default = "stfuncpythonmch01"
}

variable "function_app_insights_name" {
    type = string
    default = "appi-unc-python-mch01"
}
