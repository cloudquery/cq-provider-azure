
# Table: azure_compute_virtual_machine_disk_encryption_settings
DiskEncryptionSettings describes a Encryption Settings for a Disk
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_disk_id|uuid|Unique ID of azure_compute_virtual_machine_disks table (FK)|
|disk_encryption_key_secret_url|text|The URL referencing a secret in a Key Vault|
|disk_encryption_key_source_vault_id|text|Resource Id|
|key_encryption_key_key_url|text|The URL referencing a key encryption key in Key Vault|
|key_encryption_key_source_vault_id|text|Resource Id|
|enabled|boolean|Specifies whether disk encryption should be enabled on the virtual machine|
