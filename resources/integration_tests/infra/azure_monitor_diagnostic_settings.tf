resource "azurerm_storage_account" "diagnostics_settings_storage_account" {
  name                     = "${substr(var.test_suffix,1,-1)}dssa"
  resource_group_name      = azurerm_resource_group.resource_group.name
  location                 = azurerm_resource_group.resource_group.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}


resource "azurerm_monitor_diagnostic_setting" "diagnostic_settings_ds" {
  name               = "${var.test_prefix}-${var.test_suffix}-ds"
  target_resource_id = azurerm_virtual_network.main.id
  storage_account_id = azurerm_storage_account.diagnostics_settings_storage_account.id

  log {
    category = "VMProtectionAlerts"
    enabled  = true
    retention_policy {
      days    = 1
      enabled = false
    }
  }

  metric {
    category = "AllMetrics"
    enabled  = true

    retention_policy {
      enabled = false
    }
  }
}