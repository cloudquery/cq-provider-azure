
# Table: azure_compute_virtual_machine_extensions
VirtualMachineExtensionInstanceView the instance view of a virtual machine extension
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|name|text|The virtual machine extension name|
|type|text|an example is "CustomScriptExtension"|
|type_handler_version|text|Specifies the version of the script handler|
