
# Table: azure_compute_virtual_machine_resource_network_interfaces
NetworkInterfaceReference describes a network interface reference
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_resource_id|uuid|Unique ID of azure_compute_virtual_machine_resources table (FK)|
|network_interface_reference_properties_primary|boolean|Specifies the primary network interface in case the virtual machine has more than 1 network interface|
|resource_id|text|Resource Id|
