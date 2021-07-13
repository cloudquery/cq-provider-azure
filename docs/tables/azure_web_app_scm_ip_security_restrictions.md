
# Table: azure_web_app_scm_ip_security_restrictions
IPSecurityRestriction IP security restriction on an app
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|ip_address|text|IP address the security restriction is valid for It can be in form of pure ipv4 address (required SubnetMask property) or CIDR notation such as ipv4/mask (leading bit match) For CIDR, SubnetMask property must not be specified|
|subnet_mask|text|Subnet mask for the range of IP addresses the restriction is valid for|
|vnet_subnet_resource_id|text|Virtual network resource id|
|vnet_traffic_tag|integer|(internal) Vnet traffic tag|
|subnet_traffic_tag|integer|(internal) Subnet traffic tag|
|action|text|Allow or Deny access for this IP range|
|tag|text|Defines what this IP filter will be used for This is to support IP filtering on proxies Possible values include: 'Default', 'XffProxy', 'ServiceTag'|
|priority|integer|Priority of IP restriction rule|
|name|text|IP restriction rule name|
|description|text|IP restriction rule description|
|headers|jsonb|IP restriction rule headers X-Forwarded-Host (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host#Examples) The matching logic is  - If the property is null or empty (default), all hosts(or lack of) are allowed - A value is compared using ordinal-ignore-case (excluding port number) - Subdomain wildcards are permitted but don't match the root domain For example, *contosocom matches the subdomain foocontosocom  but not the root domain contosocom or multi-level foobarcontosocom - Unicode host names are allowed but are converted to Punycode for matching X-Forwarded-For (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#Examples) The matching logic is  - If the property is null or empty (default), any forwarded-for chains (or lack of) are allowed - If any address (excluding port number) in the chain (comma separated) matches the CIDR defined by the property X-Azure-FDID and X-FD-HealthProbe The matching logic is exact match|
