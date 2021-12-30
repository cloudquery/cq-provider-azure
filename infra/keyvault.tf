module "keyvault_vaults" {
  source         = "./modules/keyvault_vaults"
  resource_group = azurerm_resource_group.resource_group.name
  location       = azurerm_resource_group.resource_group.location
  test_prefix    = var.test_prefix
  test_suffix    = var.test_suffix
}
