resource "azurerm_resource_group" "test" {
  name     = "servicebus"
  location = "East US"
}

resource "azurerm_servicebus_namespace" "example" {
  name                = "cq-provider-azure-servicebus"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "Standard"

  tags = var.tags
}