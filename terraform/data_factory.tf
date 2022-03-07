resource "azurerm_resource_group" "data_factory" {
  name     = "data-factory"
  location = "East US 2"
}

resource "azurerm_data_factory" "data_factory" {
  name                = "cq-provider-azure-data-factory"
  location            = azurerm_resource_group.data_factory.location
  resource_group_name = azurerm_resource_group.data_factory.name
  tags = var.tags
}