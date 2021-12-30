module "security_auto_provisioning_settings" {
  source = "./modules/security_auto_provisioning_settings"
}

module "security_contacts" {
  source      = "./modules/security_contacts"
  test_prefix = var.test_prefix
  test_suffix = var.test_suffix
}

module "security_pricings" {
  source = "./modules/security_pricings"
}

module "security_settings" {
  source = "./modules/security_settings"
}


