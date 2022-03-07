
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
}

resource "azurerm_cosmosdb_mongo_database" "test" {
  name                = "cq-provider-azure-mongodb"
  resource_group_name = azurerm_cosmosdb_account.test.resource_group_name
  account_name        = azurerm_cosmosdb_account.test.name
  throughput          = 400
}