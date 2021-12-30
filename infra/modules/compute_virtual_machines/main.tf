resource "azurerm_network_interface" "main" {
  name                = "${var.test_prefix}-${var.test_suffix}-vms-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "${var.test_prefix}-${var.test_suffix}-vms-ip-conifg"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}


resource "azurerm_virtual_machine" "main" {
  name                  = "${var.test_prefix}-${var.test_suffix}-vm"
  location              = var.location
  resource_group_name   = var.resource_group_name
  network_interface_ids = [
    azurerm_network_interface.main.id
  ]
  vm_size               = "Standard_B1ls"

  # Uncomment this line to delete the OS disk automatically when deleting the VM
  delete_os_disk_on_termination = true

  # Uncomment this line to delete the data disks automatically when deleting the VM
  delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name              = "${var.test_prefix}-${var.test_suffix}-osdisk"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  os_profile {
    computer_name  = "hostname"
    admin_username = "testadmin"
    admin_password = "Password1234!"

  }

  os_profile_secrets {
    source_vault_id = var.keyvault_id
    vault_certificates {
      certificate_url = var.keyvault_cert_secret_id
    }
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags = {
    environment = "staging"
  }
}

resource "azurerm_virtual_machine_extension" "virtual_machines_vm_extension" {
  name                 = "vm-extension-${var.test_prefix}-${var.test_suffix}"
  virtual_machine_id   = azurerm_virtual_machine.main.id
  publisher            = "Microsoft.Azure.Extensions"
  type                 = "CustomScript"
  type_handler_version = "2.0"
  tags                 = {
    test = "test"
  }

  settings = <<SETTINGS
    {
        "commandToExecute": "hostname && uptime"
    }
SETTINGS
}


// windows virtual machines
resource "azurerm_network_interface" "virtual_machines_w_netowrk_interface" {
  name                = "${var.test_prefix}-${var.test_suffix}-vms-w-nic"
  location            = var.location
  resource_group_name = var.resource_group_name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_windows_virtual_machine" "virtual_machines_w_vm" {
  name                = "${var.test_prefix}-${var.test_suffix}-w-vm"
  location            = var.location
  resource_group_name = var.resource_group_name
  size                = "Standard_B1ls"
  admin_username      = "adminuser"
  admin_password      = "P@$$w0rd1234!"

  network_interface_ids = [
    azurerm_network_interface.virtual_machines_w_netowrk_interface.id,
  ]

  computer_name = substr(var.test_suffix, 0, 15)
  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2016-Datacenter"
    version   = "latest"
  }

  winrm_listener {
    protocol = "Http"
  }
}