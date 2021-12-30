resource "azurerm_monitor_action_group" "log_alerts_action_group" {
  name                = "${var.test_prefix}-${var.test_suffix}-actiongroup"
  resource_group_name         = var.resource_group_name
  short_name          = "p0action"

  webhook_receiver {
    name        = "callmyapi"
    service_uri = "http://example.com/alert"
  }
}

resource "azurerm_monitor_activity_log_alert" "log_alerts_log_alert" {
  name                = "${var.test_prefix}-${var.test_suffix}-activitylogalert"
  resource_group_name         = var.resource_group_name
  scopes              = [var.resource_group_id]
  description         = "This alert will monitor a specific storage account updates."

  criteria {
    resource_id    = var.storage_account_id
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