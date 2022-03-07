provider "azurerm" {
    features {
      key_vault {
        purge_soft_delete_on_destroy = true
      }
    }
}

data "azurerm_subscription" "current" {}
data "azurerm_client_config" "current" {}
