module "mysql_servers" {
  source              = "./modules/mysql_servers"
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  subnet_id           = module.networks.subnet_id
}