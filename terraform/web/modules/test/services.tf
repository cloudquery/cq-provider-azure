resource "azurerm_resource_group" "search" {
  name     = "${var.prefix}-search"
  location = "West Europe"
}

resource "azurerm_service_plan" "example" {
  name                = "${var.prefix}-example"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West Europe"
  os_type             = "Linux"
  sku_name            = "P1v2"
}

resource "azurerm_linux_web_app" "example" {
  name                = "${var.prefix}-example"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_service_plan.example.location
  service_plan_id     = azurerm_service_plan.example.id

  site_config {}
}