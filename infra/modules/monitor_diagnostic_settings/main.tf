resource "azurerm_monitor_diagnostic_setting" "diagnostic_settings_ds" {
  name               = "${var.test_prefix}-${var.test_suffix}-ds"
  target_resource_id = var.network_id
  storage_account_id = var.storage_account_id

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