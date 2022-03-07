
resource "random_password" "mysql" {
  length           = 16
  special          = true
}

resource "azurerm_resource_group" "mysql" {
  name     = "mysql"
  location = "East US"
}

resource "azurerm_mysql_server" "test" {
  name                = "cq-provider-azure-mysql-server"
  location            = azurerm_resource_group.mysql.location
  resource_group_name = azurerm_resource_group.mysql.name

  administrator_login          = "mysqladminun"
  administrator_login_password = random_password.mysql.result

  sku_name   = "GP_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = false
  backup_retention_days             = 7
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = true
  public_network_access_enabled     = false
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
  tags = var.tags
}

resource "azurerm_mysql_database" "db" {
  name                = "db"
  resource_group_name = azurerm_resource_group.mysql.name
  server_name         = azurerm_mysql_server.test.name
  charset             = "utf8"
  collation           = "utf8_unicode_ci"
}

resource "azurerm_private_endpoint" "example" {
  name                = "mysql_private_endpoint"
  location            = azurerm_resource_group.mysql.location
  resource_group_name = azurerm_resource_group.mysql.name
  subnet_id           = azurerm_subnet.endpoint.id

  private_service_connection {
    name                           = "tfex-mysql-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_mysql_server.test.id
    subresource_names              = ["mysqlServer"]
  }
}