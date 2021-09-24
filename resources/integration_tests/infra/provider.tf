provider "azurerm" {
  features {
    key_vault {
      purge_soft_delete_on_destroy = true
    }
  }
}

resource "azurerm_resource_group" "resource_group" {
  name     = "resource-group-${var.test_prefix}${var.test_suffix}"
  location = "West Europe"
  tags = {
    TestId = var.test_suffix
    Type   = "integration_test"
  }
}

data "azurerm_client_config" "current" {}

