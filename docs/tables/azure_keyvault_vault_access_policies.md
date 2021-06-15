
# Table: azure_keyvault_vault_access_policies
AccessPolicyEntry an identity that have access to the key vault
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_id|uuid|Unique ID of azure_keyvault_vaults table (FK)|
|tenant_id|uuid|TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault|
|object_id|text|ObjectID - The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault The object ID must be unique for the list of access policies|
|application_id|uuid|ApplicationID -  Application ID of the client making request on behalf of a principal|
|permissions_keys|text[]|Keys - Permissions to keys|
|permissions_secrets|text[]|Secrets - Permissions to secrets|
|permissions_certificates|text[]|Certificates - Permissions to certificates|
|permissions_storage|text[]|Storage - Permissions to storage accounts|
