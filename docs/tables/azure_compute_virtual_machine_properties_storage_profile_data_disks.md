
# Table: azure_compute_virtual_machine_properties_storage_profile_data_disks
DataDisk describes a data disk
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|lun|integer|Specifies the logical unit number of the data disk This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM|
|name|text|The disk name|
|vhd_uri|text|Specifies the virtual hard disk's uri|
|image_uri|text|Specifies the virtual hard disk's uri|
|caching|text|Specifies the caching requirements <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage ReadOnly for Premium storage** Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'|
|write_accelerator_enabled|boolean|Specifies whether writeAccelerator should be enabled or disabled on the disk|
|create_option|text|Specifies how the virtual machine should be created<br><br> Possible values are:<br><br> **Attach** \u2013 This value is used when you are using a specialized disk to create the virtual machine<br><br> **FromImage** \u2013 This value is used when you are using an image to create the virtual machine If you are using a platform image, you also use the imageReference element described above If you are using a marketplace image, you  also use the plan element previously described Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'|
|disk_size_gb|integer|Specifies the size of an empty data disk in gigabytes This element can be used to overwrite the size of the disk in a virtual machine image <br><br> This value cannot be larger than 1023 GB|
|managed_disk_storage_account_type|text|Specifies the storage account type for the managed disk NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk Possible values include: 'StorageAccountTypesStandardLRS', 'StorageAccountTypesPremiumLRS', 'StorageAccountTypesStandardSSDLRS', 'StorageAccountTypesUltraSSDLRS', 'StorageAccountTypesPremiumZRS', 'StorageAccountTypesStandardSSDZRS'|
|managed_disk_disk_encryption_set_id|text|Resource Id|
|managed_disk_id|text|Resource Id|
|to_be_detached|boolean|Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset|
|disk_iops_read_write|bigint|Specifies the Read-Write IOPS for the managed disk when StorageAccountType is UltraSSD_LRS Returned only for VirtualMachine ScaleSet VM disks Can be updated only via updates to the VirtualMachine Scale Set|
|disk_mb_ps_read_write|bigint|Specifies the bandwidth in MB per second for the managed disk when StorageAccountType is UltraSSD_LRS Returned only for VirtualMachine ScaleSet VM disks Can be updated only via updates to the VirtualMachine Scale Set|
|detach_option|text|Specifies the detach behavior to be used while detaching a disk or which is already in the process of detachment from the virtual machine Supported values: **ForceDetach** <br><br> detachOption: **ForceDetach** is applicable only for managed data disks If a previous detachment attempt of the data disk did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use force-detach as a last resort option to detach the disk forcibly from the VM All writes might not have been flushed when using this detach behavior <br><br> This feature is still in preview mode and is not supported for VirtualMachineScaleSet To force-detach a data disk update toBeDetached to 'true' along with setting detachOption: 'ForceDetach' Possible values include: 'ForceDetach'|
