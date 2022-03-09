terraform {
  backend "azurerm" {
    resource_group_name  = "default"
    storage_account_name = "cqproviderazureterraform"
    container_name       = "tfstate"
    key                  = "cq_keyvault.terraform.tfstate"
  }
}
