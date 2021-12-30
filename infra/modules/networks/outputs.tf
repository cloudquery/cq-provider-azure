output "network1_id" {
  value       = azurerm_virtual_network.network1.id
}

output "network2_id" {
  value       = azurerm_virtual_network.network2.id
}

output "subnet_id" {
  value       = azurerm_subnet.internal.id
}
output "network_watcher_name" {
  value       = azurerm_network_watcher.network_watcher.name
}