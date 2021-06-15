
# Table: azure_storage_account_network_rule_set_virtual_network_rules
VirtualNetworkRule virtual Network rule. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|uuid|Unique ID of azure_storage_accounts table (FK)|
|virtual_network_resource_id|text|VirtualNetworkResourceID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/MicrosoftNetwork/virtualNetworks/{vnetName}/subnets/{subnetName}|
|action|text|Action - The action of virtual network rule Possible values include: 'Allow'|
|state|text|State - Gets the state of virtual network rule Possible values include: 'StateProvisioning', 'StateDeprovisioning', 'StateSucceeded', 'StateFailed', 'StateNetworkSourceDeleted'|
