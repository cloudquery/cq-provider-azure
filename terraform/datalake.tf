resource "azurerm_data_lake_store" "azurerm_data_lake_store" {
  name                = "cqinttestsdatalake"
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  location            = azurerm_resource_group.cq_int_tests.location
  encryption_state    = "Enabled"
  encryption_type     = "ServiceManaged"
}