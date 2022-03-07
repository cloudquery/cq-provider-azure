resource "random_password" "sql-server" {
  length           = 16
  special          = true
}

resource "azurerm_sql_server" "example" {
  name                         = "cqproviderazuresqlserver"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  version                      = "12.0"
  administrator_login          = "mradministrator"
  administrator_login_password = random_password.sql-server.result

  tags = var.tags
  threat_detection_policy {
    state = "Disabled"
    retention_days = 1
    email_addresses = ["jobs@cloudquery.io"]
  }
}

resource "azurerm_private_endpoint" "sqlserver-private-endpoint" {
  name                = "sqlserver_private_endpoint"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  subnet_id           = azurerm_subnet.endpoint.id

  private_service_connection {
    name                           = "sql-server-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_sql_server.example.id
    subresource_names              = ["sqlServer"]
  }
}