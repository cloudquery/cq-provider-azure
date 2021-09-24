resource "azurerm_virtual_network" "main" {
  name                = "${var.test_suffix}-${var.test_suffix}-network"
  address_space       = [
    "10.0.0.0/16"
  ]
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
}

resource "azurerm_subnet" "internal" {
  name                 = "${var.test_prefix}-${var.test_suffix}-internal"
  resource_group_name  = azurerm_resource_group.resource_group.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = [
    "10.0.2.0/24"
  ]
}
