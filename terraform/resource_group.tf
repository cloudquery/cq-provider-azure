resource "azurerm_resource_group" "cq-provider-azure" {
  name     = "cq-provider-azure"
  location = "East US"
}

resource "azurerm_resource_group" "example" {
  name     = "cq-provider-azure-example"
  location = "East US"
}