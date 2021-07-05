
# Table: azure_compute_virtual_machine_available_patch_summary_error_details
APIErrorBase api error base
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|code|text|The error code|
|target|text|The target of the particular error|
|message|text|The error message|
