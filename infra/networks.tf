module "networks" {
  source              = "./modules/networks"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
}

module "network_public_ip_addresses" {
  source              = "./modules/network_public_ip_addresses"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
}

module "network_security_groups" {
  source               = "./modules/network_security_groups"
  resource_group_name  = azurerm_resource_group.resource_group.name
  location             = azurerm_resource_group.resource_group.location
  test_prefix          = var.test_prefix
  test_suffix          = var.test_suffix
  storage_account_id   = module.storage_accounts.storage_account_id
  network_watcher_name = module.networks.network_watcher_name
}



