module "monitor_activity_log_alerts" {
  source              = "./modules/monitor_activity_log_alerts"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  resource_group_id   = azurerm_resource_group.resource_group.id
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
  storage_account_id  = module.storage_accounts.storage_account_id
}

module "monitor_diagnostic_settings" {
  source             = "./modules/monitor_diagnostic_settings"
  test_prefix        = var.test_prefix
  test_suffix        = var.test_suffix
  storage_account_id = module.storage_accounts.storage_account_id
  network_id         = module.networks.network1_id
}

module "monitor_log_profiles" {
  source             = "./modules/monitor_log_profiles"
  test_prefix        = var.test_prefix
  test_suffix        = var.test_suffix
  storage_account_id = module.storage_accounts.storage_account_id
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
}


