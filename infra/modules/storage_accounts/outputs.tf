output "storage_account_id" {
  description = "id of storage account"
  value       = azurerm_storage_account.storage_accounts_storage_account.id
}

output "storage_account_endpoint" {
  description = "endpoint of storage account"
  value       = azurerm_storage_account.storage_accounts_storage_account.primary_blob_endpoint
}


output "storage_account_access_key" {
  description = "access key of storage account"
  value       = azurerm_storage_account.storage_accounts_storage_account.primary_access_key
}