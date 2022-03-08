terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.98.0"
    }
  }
  backend "azurerm" {
    resource_group_name  = "default"
    storage_account_name = "cqproviderazureterraform"
    container_name       = "tfstate"
    key                  = "cq_provider_azure.terraform.tfstate"
  }
}
