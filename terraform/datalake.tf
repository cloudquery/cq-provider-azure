resource "azurerm_resource_group" "data_lake" {
  name     = "data_lake"
  location = "East US 2"
}

resource "azurerm_data_lake_store" "test" {
  name                = "cqproviderazuredatalake"
  resource_group_name = azurerm_resource_group.data_lake.name
  location            = azurerm_resource_group.data_lake.location
  encryption_state    = "Enabled"
  encryption_type     = "ServiceManaged"
  tags = var.tags
}

resource "azurerm_data_lake_analytics_account" "example" {
  name                = "analyticsaccounttest"
  resource_group_name = azurerm_resource_group.data_lake.name
  location            = azurerm_resource_group.data_lake.location

  default_store_account_name = azurerm_data_lake_store.test.name
  tags = var.tags
}


resource "azurerm_data_lake_analytics_firewall_rule" "example" {
  name                = "ip-range"
  account_name        = azurerm_data_lake_analytics_account.example.name
  resource_group_name = azurerm_resource_group.data_lake.name
  start_ip_address    = "10.0.0.0"
  end_ip_address      = "10.0.0.1"
}

resource "azurerm_data_lake_store_firewall_rule" "example" {
  name                = "office-ip-range"
  account_name        = azurerm_data_lake_store.test.name
  resource_group_name = azurerm_resource_group.data_lake.name
  start_ip_address    = "10.0.0.0"
  end_ip_address      = "10.0.0.1"
}