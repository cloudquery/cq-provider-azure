resource "azurerm_storage_account" "log_profiles_storage_account" {
  name                     =  "${substr(var.test_suffix,1,-1)}lpsa"
  resource_group_name      = azurerm_resource_group.resource_group.name
  location                 = azurerm_resource_group.resource_group.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_eventhub_namespace" "log_profiles_eventhub" {
  name                = "${var.test_prefix}-${var.test_suffix}-lp-eh"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  sku                 = "Standard"
  capacity            = 2
}

resource "azurerm_monitor_log_profile" "log_profiles_log_profile" {
  name = "${var.test_prefix}-${var.test_suffix}-log-profile"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  # RootManageSharedAccessKey is created by default with listen, send, manage permissions
  servicebus_rule_id = "${azurerm_eventhub_namespace.log_profiles_eventhub.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.log_profiles_storage_account.id

  retention_policy {
    enabled = true
    days    = 7
  }
}