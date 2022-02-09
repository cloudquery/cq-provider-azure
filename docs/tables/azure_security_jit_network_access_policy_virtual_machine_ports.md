
# Table: azure_security_jit_network_access_policy_virtual_machine_ports
Describes port rule of vm
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|jit_network_access_policy_virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_security_jit_network_access_policy_virtual_machines table (FK)|
|number|integer|Port number|
|protocol|text|Possible values include: 'TCP', 'UDP', 'All'|
|allowed_source_address_prefix|text|Mutually exclusive with the "allowedSourceAddressPrefixes" parameter|
|allowed_source_address_prefixes|text[]|Mutually exclusive with the "allowedSourceAddressPrefix" parameter.|
|max_request_access_duration|text|Maximum duration requests can be made for|
