
# Table: azure_compute_virtual_machine_resources
VirtualMachineExtension describes a Virtual Machine Extension
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|virtual_machine_extension_properties_force_update_tag|text|How the extension handler should be forced to update even if the extension configuration has not changed|
|virtual_machine_extension_properties_publisher|text|The name of the extension handler publisher|
|virtual_machine_extension_properties_type|text|an example is "CustomScriptExtension"|
|virtual_machine_extension_properties_type_handler_version|text|Specifies the version of the script handler|
|virtual_machine_extension_properties_auto_upgrade_minor_version|boolean|Indicates whether the extension should use a newer minor version if one is available at deployment time Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true|
|virtual_machine_extension_properties_enable_automatic_upgrade|boolean|Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available|
|virtual_machine_extension_properties_settings|jsonb|Json formatted public settings for the extension|
|virtual_machine_extension_properties_protected_settings|jsonb|The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all|
|virtual_machine_extension_properties_provisioning_state|text|The provisioning state, which only appears in the response|
|virtual_machine_extension_properties_instance_view_name|text|The virtual machine extension name|
|virtual_machine_extension_properties_instance_view_type|text|an example is "CustomScriptExtension"|
|virtual_machine_extension_properties_instance_view_type_handler_version|text|Specifies the version of the script handler|
|resource_id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
