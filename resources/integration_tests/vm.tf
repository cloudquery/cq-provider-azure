variable "prefix" {
  default = "tfvmex"
}

data "azurerm_client_config" "current" {}

resource "azurerm_virtual_network" "main" {
  name = "${var.prefix}-network"
  address_space = [
    "10.0.0.0/16"]
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_subnet" "internal" {
  name = "internal"
  resource_group_name = azurerm_resource_group.example.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes = [
    "10.0.2.0/24"]
}

resource "azurerm_network_interface" "main" {
  name = "${var.prefix}-nic"
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  ip_configuration {
    name = "testconfiguration1"
    subnet_id = azurerm_subnet.internal.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "main" {
  name = "${var.prefix}-vm"
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  network_interface_ids = [
    azurerm_network_interface.main.id]
  vm_size = "Standard_E2a_v4"

  # Uncomment this line to delete the OS disk automatically when deleting the VM
  # delete_os_disk_on_termination = true

  # Uncomment this line to delete the data disks automatically when deleting the VM
  # delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer = "UbuntuServer"
    sku = "16.04-LTS"
    version = "latest"
  }

  storage_os_disk {
    name = "osdisk2"
    caching = "ReadWrite"
    create_option = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_data_disk {
    name = azurerm_managed_disk.example3.name
    caching = "ReadWrite"
    managed_disk_id = azurerm_managed_disk.example3.id
    create_option = "Attach"
    lun = 2
  disk_size_gb = "10"
  }

  storage_data_disk {
    name = "data-disk1"
    managed_disk_type = "Standard_LRS"
    create_option = "Empty"
    lun = 0
  disk_size_gb = "10"
  }


  os_profile {
    computer_name = "hostname"
    admin_username = "testadmin"
    admin_password = "Password1234!"
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags = {
    environment = "staging"
  }

  depends_on = [
    azurerm_managed_disk.example3,
    azurerm_managed_disk.example]
}


resource "azurerm_managed_disk" "example" {
  name = "unattached-disk2"
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option = "Empty"
  disk_size_gb = "10"

  tags = {
    environment = "staging"
  }
}


resource "azurerm_managed_disk" "example3" {
  name = "attached-disk2"
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option = "Empty"
  disk_size_gb = "10"

  tags = {
    environment = "staging"
  }
}

resource "azurerm_virtual_machine_data_disk_attachment" "example" {
  managed_disk_id = azurerm_managed_disk.example.id
  virtual_machine_id = azurerm_virtual_machine.main.id
  lun = "10"
  caching = "ReadWrite"
  depends_on = [
    azurerm_managed_disk.example3,
    azurerm_managed_disk.example]
}


resource "azurerm_key_vault" "example" {
  name = "des-example-keyvault1123"
  location = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  tenant_id = data.azurerm_client_config.current.tenant_id
  sku_name = "premium"
  enabled_for_disk_encryption = true
  soft_delete_enabled = true
  purge_protection_enabled = true
}

resource "azurerm_key_vault_key" "example" {
  name = "des-example-key111232"
  key_vault_id = azurerm_key_vault.example.id
  key_type = "RSA"
  key_size = 2048

  depends_on = [
    azurerm_key_vault_access_policy.example-user
  ]

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}

resource "azurerm_disk_encryption_set" "example" {
  name = "des1123"
  resource_group_name = azurerm_resource_group.example.name
  location = azurerm_resource_group.example.location
  key_vault_key_id = azurerm_key_vault_key.example.id

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_key_vault_access_policy" "example-disk" {
  key_vault_id = azurerm_key_vault.example.id

  tenant_id = azurerm_disk_encryption_set.example.identity.0.tenant_id
  object_id = azurerm_disk_encryption_set.example.identity.0.principal_id

  key_permissions = [
    "Get",
    "WrapKey",
    "UnwrapKey"
  ]
}

resource "azurerm_key_vault_access_policy" "example-user" {
  key_vault_id = azurerm_key_vault.example.id

  tenant_id = data.azurerm_client_config.current.tenant_id
  object_id = data.azurerm_client_config.current.object_id

  key_permissions = [
    "get",
    "create",
    "delete"
  ]
}

resource "azurerm_virtual_machine_extension" "example" {
  name = "test123"
  virtual_machine_id = azurerm_virtual_machine.main.id
  publisher = "Microsoft.Azure.Extensions"
  type = "CustomScript"
  type_handler_version = "2.0"

  settings = <<SETTINGS
    {
        "commandToExecute": "hostname && uptime"
    }
SETTINGS

}
