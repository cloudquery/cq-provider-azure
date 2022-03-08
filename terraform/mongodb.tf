
resource "azurerm_resource_group" "mongodb" {
  name     = "mongodb"
  location = "East US"
}

resource "azurerm_cosmosdb_account" "test" {
  name                = "cq-provider-azure-mongodb"
  resource_group_name = azurerm_resource_group.mongodb.name
  location = azurerm_resource_group.mongodb.location

  offer_type          = "Standard"
  kind                = "MongoDB"

  enable_automatic_failover = true

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["DELETE", "GET", "HEAD", "MERGE", "POST", "OPTIONS", "PUT", "PATCH"]
    allowed_origins = ["*"]
    exposed_headers = ["*"]
    max_age_in_seconds = 730
  }

  capabilities {
    name = "EnableAggregationPipeline"
  }

  capabilities {
    name = "mongoEnableDocLevelTTL"
  }

  capabilities {
    name = "MongoDBv3.4"
  }

  capabilities {
    name = "EnableMongo"
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 300
    max_staleness_prefix    = 100000
  }

  geo_location {
    location          = azurerm_resource_group.mongodb.location
    failover_priority = 0
  }
  tags = var.tags
}

resource "azurerm_cosmosdb_mongo_database" "test" {
  name                = "cq-provider-azure-mongodb"
  resource_group_name = azurerm_cosmosdb_account.test.resource_group_name
  account_name        = azurerm_cosmosdb_account.test.name
  throughput          = 400
  // autoscale_settings {
  //   max_throughput = 4000
  // }
}


resource "azurerm_private_endpoint" "cosmodb-private-endpoint" {
  name                = "cosmodb_private_endpoint"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  subnet_id           = azurerm_subnet.endpoint.id

  private_service_connection {
    name                           = "tfex-cosmodb-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_cosmosdb_account.test.id
    subresource_names              = ["MongoDB"]
  }
}