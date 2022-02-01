
# Table: azure_security_jit_network_access_policy_virtual_machine_ports
JitNetworkAccessPortRule - Describes port rule of vm
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|jit_network_access_policy_virtual_machine_cq_id|uuid|Unique CloudQuery ID of azure_security_jit_network_access_policy_virtual_machines table (FK)|
|number|integer|Number - Port number|
|protocol|text|Protocol - Possible values include: 'TCP', 'UDP', 'All'|
|allowed_source_address_prefix|text|AllowedSourceAddressPrefix - Mutually exclusive with the "allowedSourceAddressPrefixes" parameter|
|allowed_source_address_prefixes|text[]|AllowedSourceAddressPrefixes - Mutually exclusive with the "allowedSourceAddressPrefix" parameter.|
|max_request_access_duration|text|MaxRequestAccessDuration - Maximum duration requests can be made for|
