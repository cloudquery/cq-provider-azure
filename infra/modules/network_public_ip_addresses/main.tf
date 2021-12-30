resource "azurerm_public_ip" "public_ips_ip" {
  name                = "${var.test_prefix}-${var.test_suffix}-ip"
  resource_group_name = var.resource_group_name
  location            = var.location
  allocation_method   = "Static"

  tags = {
    environment = "Production"
  }
}