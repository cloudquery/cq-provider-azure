
# Table: azure_compute_virtual_machine_vm_agent_extension_handlers
VirtualMachineExtensionHandlerInstanceView the instance view of a virtual machine extension handler
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|type|text|an example is "CustomScriptExtension"|
|type_handler_version|text|Specifies the version of the script handler|
|status_code|text|The status code|
|status_level|text|The level code Possible values include: 'Info', 'Warning', 'Error'|
|status_display_status|text|The short localizable label for the status|
|status_message|text|The detailed status message, including for alerts and error messages|
|status_time|timestamp without time zone||
