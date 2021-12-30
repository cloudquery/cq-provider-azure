output "keyvault_id" {
  description = "id of the keyvault"
  value       = azurerm_key_vault.keyvaults_keyvault.id
}

output "keyvault_secret_id" {
  description = "secret if of keyvaults cert"
  value       = azurerm_key_vault_certificate.keyvaults_cert.secret_id
}