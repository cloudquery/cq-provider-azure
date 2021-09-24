
resource "azurerm_managed_disk" "compute_disks_disk" {
  name = "disks-${var.test_prefix}-${var.test_suffix}"
  location = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  storage_account_type = "Standard_LRS"
  create_option = "Empty"
  disk_size_gb = "10"

#  disk_encryption_set_id = azurerm_disk_encryption_set.compute_disks_enc_set.id

  tags = {
    environment = "staging"
  }
}
#
#resource "azurerm_disk_encryption_set" "compute_disks_enc_set" {
#  name                = "disks-enc-set-${var.test_prefix}-${var.test_suffix}"
#  resource_group_name = azurerm_resource_group.resource_group.name
#  location            = azurerm_resource_group.resource_group.location
#  key_vault_key_id    = azurerm_key_vault_key.compute_disks_disk_enc_key.id
#
#  identity {
#    type = "SystemAssigned"
#  }
#}
#
#resource "azurerm_key_vault_key" "compute_disks_disk_enc_key" {
#  name                 = "disks-key1-${var.test_prefix}-${var.test_suffix}"
#  key_vault_id         = azurerm_key_vault.keyvaults_keyvault.id
#  key_type             = "RSA"
#  key_size             = 2048
#
#  key_opts = [
#    "decrypt",
#    "encrypt",
#    "sign",
#    "unwrapKey",
#    "verify",
#    "wrapKey",
#  ]
#}
#
#resource "azurerm_key_vault_access_policy" "example-disk" {
#  key_vault_id = azurerm_key_vault.keyvaults_keyvault.id
#
#  tenant_id = azurerm_disk_encryption_set.compute_disks_enc_set.identity.0.tenant_id
#  object_id = azurerm_disk_encryption_set.compute_disks_enc_set.identity.0.principal_id
#
#  key_permissions = [
#    "Get",
#    "WrapKey",
#    "UnwrapKey",
#    "Purge"
#  ]
#}
#
#resource "azurerm_key_vault" "keyvaults_keyvault" {
#  name                        = "disks-valut-${substr(var.test_prefix,-5,-1)}-${substr(var.test_suffix,-5,-1)}"
#  location                    = azurerm_resource_group.resource_group.location
#  resource_group_name         = azurerm_resource_group.resource_group.name
#  tenant_id                   = data.azurerm_client_config.current.tenant_id
#  sku_name                    = "standard"
#  soft_delete_retention_days  = 7
#  enabled_for_disk_encryption = true
#  purge_protection_enabled    = true
#
#  access_policy {
#    tenant_id = data.azurerm_client_config.current.tenant_id
#    object_id = data.azurerm_client_config.current.object_id
#
#    certificate_permissions = [
#      "create",
#      "recover",
#      "delete",
#      "deleteissuers",
#      "get",
#      "getissuers",
#      "import",
#      "list",
#      "listissuers",
#      "managecontacts",
#      "manageissuers",
#      "purge",
#      "setissuers",
#      "update",
#    ]
#
#    key_permissions = [
#      "backup",
#      "recover",
#      "create",
#      "decrypt",
#      "delete",
#      "encrypt",
#      "get",
#      "import",
#      "list",
#      "purge",
#      "recover",
#      "restore",
#      "sign",
#      "unwrapKey",
#      "update",
#      "verify",
#      "wrapKey",
#    ]
#
#    secret_permissions = [
#      "backup",
#      "delete",
#      "recover",
#      "get",
#      "list",
#      "purge",
#      "recover",
#      "restore",
#      "set",
#    ]
#  }
#}

