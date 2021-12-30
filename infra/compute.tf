module "compute_disks" {
  source              = "./modules/compute_disks"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  test_prefix         = var.test_prefix
  test_suffix         = var.test_suffix
}


module "compute_virtual_machines" {
  source                  = "./modules/compute_virtual_machines"
  resource_group_name     = azurerm_resource_group.resource_group.name
  location                = azurerm_resource_group.resource_group.location
  test_prefix             = var.test_prefix
  test_suffix             = var.test_suffix
  subnet_id               = module.networks.subnet_id
  keyvault_id             = module.keyvault_vaults.keyvault_id
  keyvault_cert_secret_id = module.keyvault_vaults.keyvault_secret_id
}