
# Table: azure_compute_virtual_machine_windows_configuration_additional_unattend_contents
AdditionalUnattendContent specifies additional XML formatted information that can be included in the Unattendxml file, which is used by Windows Setup Contents are defined by setting name, component name, and the pass in which the content is applied
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|pass_name|text|The pass name Currently, the only allowable value is OobeSystem Possible values include: 'OobeSystem'|
|component_name|text|The component name Currently, the only allowable value is Microsoft-Windows-Shell-Setup Possible values include: 'MicrosoftWindowsShellSetup'|
|setting_name|text|Specifies the name of the setting to which the content applies Possible values are: FirstLogonCommands and AutoLogon Possible values include: 'AutoLogon', 'FirstLogonCommands'|
|content|text|Specifies the XML formatted content that is added to the unattendxml file for the specified path and component The XML must be less than 4KB and must include the root element for the setting or feature that is being inserted|
