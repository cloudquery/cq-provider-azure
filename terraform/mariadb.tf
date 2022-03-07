resource "random_password" "mariadb" {
  length           = 16
  special          = true
}

resource "azurerm_mariadb_server" "example" {
  name                = "cq-provider-azure-maria-server"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name

  sku_name = "GP_Gen5_2"

  storage_mb                   = 51200
  backup_retention_days        = 7
  geo_redundant_backup_enabled = false

  administrator_login          = "acctestun"
  administrator_login_password = random_password.mariadb.result
  version                      = "10.2"
  ssl_enforcement_enabled      = true
}

resource "azurerm_mariadb_database" "example" {
  name                = "mariadb_database"
  resource_group_name = azurerm_resource_group.test.name
  server_name         = azurerm_mariadb_server.example.name
  charset             = "utf8"
  collation           = "utf8_general_ci"
}

resource "azurerm_private_endpoint" "mariadb-private-endpoint" {
  name                = "mariadb_private_endpoint"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  subnet_id           = azurerm_subnet.endpoint.id

  private_service_connection {
    name                           = "tfex-mysql-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_mariadb_server.example.id
    subresource_names              = ["mariadbServer"]
  }
}