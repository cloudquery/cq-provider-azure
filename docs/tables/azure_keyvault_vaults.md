
# Table: azure_keyvault_vaults
Azure ketvault vault
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|resource_id|text|Fully qualified identifier of the key vault resource|
|name|text|Name of the key vault resource|
|type|text|Resource type of the key vault resource|
|location|text|Location - Azure location of the key vault resource|
|tags|jsonb|Tags - Tags assigned to the key vault resource|
|tenant_id|uuid|TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault|
|sku_family|text|Family - SKU family name|
|sku_name|text|Name - SKU name to specify whether the key vault is a standard vault or a premium vault Possible values include: 'Standard', 'Premium'|
|vault_uri|text|VaultURI - The URI of the vault for performing operations on keys and secrets|
|enabled_for_deployment|boolean|EnabledForDeployment - Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault|
|enabled_for_disk_encryption|boolean|EnabledForDiskEncryption - Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys|
|enabled_for_template_deployment|boolean|EnabledForTemplateDeployment - Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault|
|enable_soft_delete|boolean|EnableSoftDelete - Property to specify whether the 'soft delete' functionality is enabled for this key vault If it's not set to any value(true or false) when creating new key vault, it will be set to true by default Once set to true, it cannot be reverted to false|
|soft_delete_retention_in_days|integer|SoftDeleteRetentionInDays - softDelete data retention days It accepts >=7 and <=90|
|enable_rbac_authorization|boolean|EnableRbacAuthorization - Property that controls how data actions are authorized When true, the key vault will use Role Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties will be  ignored (warning: this is a preview feature) When false, the key vault will use the access policies specified in vault properties, and any policy stored on Azure Resource Manager will be ignored If null or not specified, the vault is created with the default value of false Note that management actions are always authorized with RBAC|
|create_mode|text|CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not Possible values include: 'CreateModeRecover', 'CreateModeDefault'|
|enable_purge_protection|boolean|EnablePurgeProtection - Property specifying whether protection against purge is enabled for this vault Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion The setting is effective only if soft delete is also enabled Enabling this functionality is irreversible - that is, the property does not accept false as its value|
|network_acls_bypass|text|Bypass - Tells what traffic can bypass network rules This can be 'AzureServices' or 'None'  If not specified the default is 'AzureServices' Possible values include: 'AzureServices', 'None'|
|network_acls_default_action|text|DefaultAction - The default action when no rule from ipRules and from virtualNetworkRules match This is only used after the bypass property has been evaluated Possible values include: 'Allow', 'Deny'|
|network_acls_ip_rules|text[]|IPRules - The list of IP address rules|
|network_acls_virtual_network_rules|text[]|VirtualNetworkRules - The list of virtual network rules|
