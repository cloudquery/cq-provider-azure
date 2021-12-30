resource "azurerm_network_security_group" "network_security_groups_nsg1" {
  name                = "${var.test_prefix}-${var.test_suffix}-nsg"
  location            = var.location
  resource_group_name = var.resource_group_name

  security_rule {
    name                       = "test12f23"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "test12223"
    priority                   = 121
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "3389"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "test312223"
    priority                   = 120
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "1-4000"
    destination_port_range     = "*"
    source_address_prefix      = "internet"
    destination_address_prefix = "*"
  }

  tags = {
    environment = "Production"
  }
}

resource "azurerm_log_analytics_workspace" "network_security_groups_analytics_workspace" {
  name                = "${var.test_prefix}-${var.test_suffix}-aw"
  location            = var.location
  resource_group_name = var.resource_group_name
  sku                 = "PerGB2018"
}

resource "azurerm_network_watcher_flow_log" "network_security_groups_flow_log" {
  network_watcher_name = var.network_watcher_name
  resource_group_name  = var.resource_group_name

  network_security_group_id = azurerm_network_security_group.network_security_groups_nsg1.id
  storage_account_id        = var.storage_account_id
  enabled                   = true

  retention_policy {
    enabled = true
    days    = 7
  }

  traffic_analytics {
    enabled               = true
    workspace_id          = azurerm_log_analytics_workspace.network_security_groups_analytics_workspace.workspace_id
    workspace_region      = azurerm_log_analytics_workspace.network_security_groups_analytics_workspace.location
    workspace_resource_id = azurerm_log_analytics_workspace.network_security_groups_analytics_workspace.id
    interval_in_minutes   = 10
  }
}