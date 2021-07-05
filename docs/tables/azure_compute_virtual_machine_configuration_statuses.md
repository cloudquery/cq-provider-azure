
# Table: azure_compute_virtual_machine_configuration_statuses
InstanceViewStatus instance view status
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|code|text|The status code|
|level|text|The level code Possible values include: 'Info', 'Warning', 'Error'|
|display_status|text|The short localizable label for the status|
|message|text|The detailed status message, including for alerts and error messages|
|time|timestamp without time zone||
