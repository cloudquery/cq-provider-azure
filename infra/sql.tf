module "sql_servers" {
  source                     = "./modules/sql_servers"
  resource_group_name        = azurerm_resource_group.resource_group.name
  location                   = azurerm_resource_group.resource_group.location
  test_prefix                = var.test_prefix
  test_suffix                = var.test_suffix
  subnet_id                  = module.networks.subnet_id
  storage_account_access_key = module.storage_accounts.storage_account_access_key
  storage_account_endpoint   = module.storage_accounts.storage_account_endpoint
}
