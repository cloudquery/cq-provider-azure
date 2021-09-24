resource "azurerm_monitor_action_group" "log_alerts_action_group" {
  name                = "${var.test_prefix}-${var.test_suffix}-actiongroup"
  resource_group_name = azurerm_resource_group.resource_group.name
  short_name          = "p0action"

  webhook_receiver {
    name        = "callmyapi"
    service_uri = "http://example.com/alert"
  }
}

resource "azurerm_storage_account" "log_alerts_storage_account" {
  name                     = "${substr(var.test_prefix,-9,-1)}${substr(var.test_suffix,-9,-1)}sa"
  resource_group_name      = azurerm_resource_group.resource_group.name
  location                 = azurerm_resource_group.resource_group.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_monitor_activity_log_alert" "log_alerts_log_alert" {
  name                = "${var.test_prefix}-${var.test_suffix}-activitylogalert"
  resource_group_name = azurerm_resource_group.resource_group.name
  scopes              = [azurerm_resource_group.resource_group.id]
  description         = "This alert will monitor a specific storage account updates."

  criteria {
    resource_id    = azurerm_storage_account.log_alerts_storage_account.id
    operation_name = "Microsoft.Storage/storageAccounts/write"
    category       = "Recommendation"
  }

  tags = {
    test = "test"
  }

  action {
    action_group_id = azurerm_monitor_action_group.log_alerts_action_group.id

    webhook_properties = {
      from = "terraform"
    }
  }
}