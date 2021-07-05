package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeVirtualMachines() *schema.Table {
	return &schema.Table{
		Name:        "azure_compute_virtual_machines",
		Description: "VirtualMachine describes a Virtual Machine",
		Resolver:    fetchComputeVirtualMachines,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "plan_name",
				Description: "The plan ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Name"),
			},
			{
				Name:        "plan_publisher",
				Description: "The publisher ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Publisher"),
			},
			{
				Name:        "plan_product",
				Description: "Specifies the product of the image from the marketplace This is the same value as Offer under the imageReference element",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.Product"),
			},
			{
				Name:        "plan_promotion_code",
				Description: "The promotion code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Plan.PromotionCode"),
			},
			{
				Name:        "virtual_machine_properties_hardware_profile_vm_size",
				Description: "Specifies the size of the virtual machine <br><br> The enum data type is currently deprecated and will be removed by December 23rd 2023 <br><br> Recommended way to get the list of available sizes is using these APIs: <br><br> [List all available virtual machine sizes in an availability set](https://docsmicrosoftcom/rest/api/compute/availabilitysets/listavailablesizes) <br><br> [List all available virtual machine sizes in a region]( https://docsmicrosoftcom/en-us/rest/api/compute/resourceskus/list) <br><br> [List all available virtual machine sizes for resizing](https://docsmicrosoftcom/rest/api/compute/virtualmachines/listavailablesizes) For more information about virtual machine sizes, see [Sizes for virtual machines](https://docsmicrosoftcom/en-us/azure/virtual-machines/sizes) <br><br> The available VM sizes depend on region and availability set Possible values include: 'BasicA0', 'BasicA1', 'BasicA2', 'BasicA3', 'BasicA4', 'StandardA0', 'StandardA1', 'StandardA2', 'StandardA3', 'StandardA4', 'StandardA5', 'StandardA6', 'StandardA7', 'StandardA8', 'StandardA9', 'StandardA10', 'StandardA11', 'StandardA1V2', 'StandardA2V2', 'StandardA4V2', 'StandardA8V2', 'StandardA2mV2', 'StandardA4mV2', 'StandardA8mV2', 'StandardB1s', 'StandardB1ms', 'StandardB2s', 'StandardB2ms', 'StandardB4ms', 'StandardB8ms', 'StandardD1', 'StandardD2', 'StandardD3', 'StandardD4', 'StandardD11', 'StandardD12', 'StandardD13', 'StandardD14', 'StandardD1V2', 'StandardD2V2', 'StandardD3V2', 'StandardD4V2', 'StandardD5V2', 'StandardD2V3', 'StandardD4V3', 'StandardD8V3', 'StandardD16V3', 'StandardD32V3', 'StandardD64V3', 'StandardD2sV3', 'StandardD4sV3', 'StandardD8sV3', 'StandardD16sV3', 'StandardD32sV3', 'StandardD64sV3', 'StandardD11V2', 'StandardD12V2', 'StandardD13V2', 'StandardD14V2', 'StandardD15V2', 'StandardDS1', 'StandardDS2', 'StandardDS3', 'StandardDS4', 'StandardDS11', 'StandardDS12', 'StandardDS13', 'StandardDS14', 'StandardDS1V2', 'StandardDS2V2', 'StandardDS3V2', 'StandardDS4V2', 'StandardDS5V2', 'StandardDS11V2', 'StandardDS12V2', 'StandardDS13V2', 'StandardDS14V2', 'StandardDS15V2', 'StandardDS134V2', 'StandardDS132V2', 'StandardDS148V2', 'StandardDS144V2', 'StandardE2V3', 'StandardE4V3', 'StandardE8V3', 'StandardE16V3', 'StandardE32V3', 'StandardE64V3', 'StandardE2sV3', 'StandardE4sV3', 'StandardE8sV3', 'StandardE16sV3', 'StandardE32sV3', 'StandardE64sV3', 'StandardE3216V3', 'StandardE328sV3', 'StandardE6432sV3', 'StandardE6416sV3', 'StandardF1', 'StandardF2', 'StandardF4', 'StandardF8', 'StandardF16', 'StandardF1s', 'StandardF2s', 'StandardF4s', 'StandardF8s', 'StandardF16s', 'StandardF2sV2', 'StandardF4sV2', 'StandardF8sV2', 'StandardF16sV2', 'StandardF32sV2', 'StandardF64sV2', 'StandardF72sV2', 'StandardG1', 'StandardG2', 'StandardG3', 'StandardG4', 'StandardG5', 'StandardGS1', 'StandardGS2', 'StandardGS3', 'StandardGS4', 'StandardGS5', 'StandardGS48', 'StandardGS44', 'StandardGS516', 'StandardGS58', 'StandardH8', 'StandardH16', 'StandardH8m', 'StandardH16m', 'StandardH16r', 'StandardH16mr', 'StandardL4s', 'StandardL8s', 'StandardL16s', 'StandardL32s', 'StandardM64s', 'StandardM64ms', 'StandardM128s', 'StandardM128ms', 'StandardM6432ms', 'StandardM6416ms', 'StandardM12864ms', 'StandardM12832ms', 'StandardNC6', 'StandardNC12', 'StandardNC24', 'StandardNC24r', 'StandardNC6sV2', 'StandardNC12sV2', 'StandardNC24sV2', 'StandardNC24rsV2', 'StandardNC6sV3', 'StandardNC12sV3', 'StandardNC24sV3', 'StandardNC24rsV3', 'StandardND6s', 'StandardND12s', 'StandardND24s', 'StandardND24rs', 'StandardNV6', 'StandardNV12', 'StandardNV24'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.HardwareProfile.VMSize"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_publisher",
				Description: "The image publisher",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.Publisher"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_offer",
				Description: "Specifies the offer of the platform image or marketplace image used to create the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.Offer"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_sku",
				Description: "The image SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.Sku"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_version",
				Description: "Specifies the version of the platform image or marketplace image used to create the virtual machine The allowed formats are MajorMinorBuild or 'latest' Major, Minor, and Build are decimal numbers Specify 'latest' to use the latest version of an image available at deploy time Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.Version"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_exact_version",
				Description: "Specifies in decimal numbers, the version of platform image or marketplace image used to create the virtual machine This readonly field differs from 'version', only if the value specified in 'version' field is 'latest'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.ExactVersion"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_image_reference_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.ImageReference.ID"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_os_type",
				Description: "This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux** Possible values include: 'Windows', 'Linux'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.OsType"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_encryption_settings_disk_encryption_key_secret_url",
				Description: "The URL referencing a secret in a Key Vault",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SecretURL"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_encryption_settings_disk_encryption_key_source_vault_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.DiskEncryptionKey.SourceVault.ID"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_encryption_settings_key_encryption_key_key_url",
				Description: "The URL referencing a key encryption key in Key Vault",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.KeyURL"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_encryption_settings_key_encryption_key_source_vault_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.KeyEncryptionKey.SourceVault.ID"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_encryption_settings_enabled",
				Description: "Specifies whether disk encryption should be enabled on the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.EncryptionSettings.Enabled"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_name",
				Description: "The disk name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.Name"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_vhd_uri",
				Description: "Specifies the virtual hard disk's uri",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.Vhd.URI"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_image_uri",
				Description: "Specifies the virtual hard disk's uri",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.Image.URI"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_caching",
				Description: "Specifies the caching requirements <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None** for Standard storage **ReadOnly** for Premium storage Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.Caching"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_write_accelerator_enabled",
				Description: "Specifies whether writeAccelerator should be enabled or disabled on the disk",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.WriteAcceleratorEnabled"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_diff_disk_settings_option",
				Description: "Specifies the ephemeral disk settings for operating system disk Possible values include: 'Local'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.DiffDiskSettings.Option"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_diff_disk_settings_placement",
				Description: "Specifies the ephemeral disk placement for operating system disk<br><br> Possible values are: <br><br> **CacheDisk** <br><br> **ResourceDisk** <br><br> Default: **CacheDisk** if one is configured for the VM size otherwise **ResourceDisk** is used<br><br> Refer to VM size documentation for Windows VM at https://docsmicrosoftcom/en-us/azure/virtual-machines/windows/sizes and Linux VM at https://docsmicrosoftcom/en-us/azure/virtual-machines/linux/sizes to check which VM sizes exposes a cache disk Possible values include: 'CacheDisk', 'ResourceDisk'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.DiffDiskSettings.Placement"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_create_option",
				Description: "Specifies how the virtual machine should be created<br><br> Possible values are:<br><br> **Attach** \\u2013 This value is used when you are using a specialized disk to create the virtual machine<br><br> **FromImage** \\u2013 This value is used when you are using an image to create the virtual machine If you are using a platform image, you also use the imageReference element described above If you are using a marketplace image, you  also use the plan element previously described Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.CreateOption"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_disk_size_gb",
				Description: "Specifies the size of an empty data disk in gigabytes This element can be used to overwrite the size of the disk in a virtual machine image <br><br> This value cannot be larger than 1023 GB",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.DiskSizeGB"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_managed_disk_storage_account_type",
				Description: "Specifies the storage account type for the managed disk NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk Possible values include: 'StorageAccountTypesStandardLRS', 'StorageAccountTypesPremiumLRS', 'StorageAccountTypesStandardSSDLRS', 'StorageAccountTypesUltraSSDLRS', 'StorageAccountTypesPremiumZRS', 'StorageAccountTypesStandardSSDZRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.ManagedDisk.StorageAccountType"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_managed_disk_disk_encryption_set_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.ManagedDisk.DiskEncryptionSet.ID"),
			},
			{
				Name:        "virtual_machine_properties_storage_profile_os_disk_managed_disk_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.StorageProfile.OsDisk.ManagedDisk.ID"),
			},
			{
				Name:        "virtual_machine_properties_additional_capabilities_ultra_s_s_d_enabled",
				Description: "The flag that enables or disables a capability to have one or more managed data disks with UltraSSD_LRS storage account type on the VM or VMSS Managed disks with storage account type UltraSSD_LRS can be added to a virtual machine or virtual machine scale set only if this property is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.AdditionalCapabilities.UltraSSDEnabled"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_computer_name",
				Description: "Specifies the host OS name of the virtual machine <br><br> This name cannot be updated after the VM is created <br><br> **Max-length (Windows):** 15 characters <br><br> **Max-length (Linux):** 64 characters <br><br> For naming conventions and restrictions see [Azure infrastructure services implementation guidelines](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson#1-naming-conventions)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.ComputerName"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_admin_username",
				Description: "Specifies the name of the administrator account <br><br> This property cannot be updated after the VM is created <br><br> **Windows-only restriction:** Cannot end in \"\" <br><br> **Disallowed values:** \"administrator\", \"admin\", \"user\", \"user1\", \"test\", \"user2\", \"test1\", \"user3\", \"admin1\", \"1\", \"123\", \"a\", \"actuser\", \"adm\", \"admin2\", \"aspnet\", \"backup\", \"console\", \"david\", \"guest\", \"john\", \"owner\", \"root\", \"server\", \"sql\", \"support\", \"support_388945a0\", \"sys\", \"test2\", \"test3\", \"user4\", \"user5\" <br><br> **Minimum-length (Linux):** 1  character <br><br> **Max-length (Linux):** 64 characters <br><br> **Max-length (Windows):** 20 characters  <br><br><li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)<br><li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AdminUsername"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_admin_password",
				Description: "Specifies the password of the administrator account <br><br> **Minimum-length (Windows):** 8 characters <br><br> **Minimum-length (Linux):** 6 characters <br><br> **Max-length (Windows):** 123 characters <br><br> **Max-length (Linux):** 72 characters <br><br> **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled <br> Has lower characters <br>Has upper characters <br> Has a digit <br> Has a special character (Regex match [\\W_]) <br><br> **Disallowed values:** \"abc@123\", \"P@$$w0rd\", \"P@ssw0rd\", \"P@ssword123\", \"Pa$$word\", \"pass@word1\", \"Password!\", \"Password1\", \"Password22\", \"iloveyou!\" <br><br> For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftocjson) <br><br> For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson#reset-root-password)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AdminPassword"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_custom_data",
				Description: "Specifies a base-64 encoded string of custom data The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine The maximum length of the binary array is 65535 bytes <br><br> **Note: Do not pass any secrets or passwords in customData property** <br><br> This property cannot be updated after the VM is created <br><br> customData is passed to the VM to be saved as a file, for more information see [Custom Data on Azure VMs](https://azuremicrosoftcom/en-us/blog/custom-data-and-cloud-init-on-windows-azure/) <br><br> For using cloud-init for your Linux VM, see [Using cloud-init to customize a Linux VM during creation](https://docsmicrosoftcom/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.CustomData"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_windows_configuration_provision_vm_agent",
				Description: "Indicates whether virtual machine agent should be provisioned on the virtual machine <br><br> When this property is not specified in the request body, default behavior is to set it to true  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.ProvisionVMAgent"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_windows_configuration_enable_automatic_updates",
				Description: "Indicates whether Automatic Updates is enabled for the Windows virtual machine Default value is true <br><br> For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.EnableAutomaticUpdates"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_windows_configuration_time_zone",
				Description: "Specifies the time zone of the virtual machine eg \"Pacific Standard Time\" <br><br> Possible values can be [TimeZoneInfoId](https://docsmicrosoftcom/en-us/dotnet/api/systemtimezoneinfoid?#System_TimeZoneInfo_Id) value from time zones returned by [TimeZoneInfoGetSystemTimeZones](https://docsmicrosoftcom/en-us/dotnet/api/systemtimezoneinfogetsystemtimezones)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.TimeZone"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_windows_configuration_patch_settings_patch_mode",
				Description: "the property WindowsConfigurationenableAutomaticUpdates must be false<br /><br /> **AutomaticByOS** - The virtual machine will automatically be updated by the OS The property WindowsConfigurationenableAutomaticUpdates must be true <br /><br /> **AutomaticByPlatform** - the virtual machine will automatically updated by the platform The properties provisionVMAgent and WindowsConfigurationenableAutomaticUpdates must be true Possible values include: 'WindowsVMGuestPatchModeManual', 'WindowsVMGuestPatchModeAutomaticByOS', 'WindowsVMGuestPatchModeAutomaticByPlatform'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_windows_configuration_patch_settings_enable_hotpatching",
				Description: "Enables customers to patch their Azure VMs without requiring a reboot For enableHotpatching, the 'provisionVMAgent' must be set to true and 'patchMode' must be set to 'AutomaticByPlatform'",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.WindowsConfiguration.PatchSettings.EnableHotpatching"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_linux_configuration_disable_password_authentication",
				Description: "Specifies whether password authentication should be disabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.DisablePasswordAuthentication"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_linux_configuration_provision_vm_agent",
				Description: "Indicates whether virtual machine agent should be provisioned on the virtual machine <br><br> When this property is not specified in the request body, default behavior is to set it to true  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.ProvisionVMAgent"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_linux_configuration_patch_settings_patch_mode",
				Description: "Specifies the mode of VM Guest Patching to IaaS virtual machine<br /><br /> Possible values are:<br /><br /> **ImageDefault** - The virtual machine's default patching configuration is used <br /><br /> **AutomaticByPlatform** - The virtual machine will be automatically updated by the platform The property provisionVMAgent must be true Possible values include: 'ImageDefault', 'AutomaticByPlatform'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.LinuxConfiguration.PatchSettings.PatchMode"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_allow_extension_operations",
				Description: "Specifies whether extension operations should be allowed on the virtual machine <br><br>This may only be set to False when no extensions are present on the virtual machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.AllowExtensionOperations"),
			},
			{
				Name:        "virtual_machine_properties_os_profile_require_guest_provision_signal",
				Description: "Specifies whether the guest provision signal is required to infer provision success of the virtual machine  **Note: This property is for private testing only, and all customers must not set the property to false**",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.OsProfile.RequireGuestProvisionSignal"),
			},
			{
				Name:        "virtual_machine_properties_security_profile_uefi_settings_secure_boot_enabled",
				Description: "Specifies whether secure boot should be enabled on the virtual machine <br><br>Minimum api-version: 2020-12-01",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.SecureBootEnabled"),
			},
			{
				Name:        "virtual_machine_properties_security_profile_uefi_settings_v_tpm_enabled",
				Description: "Specifies whether vTPM should be enabled on the virtual machine <br><br>Minimum api-version: 2020-12-01",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.UefiSettings.VTpmEnabled"),
			},
			{
				Name:        "virtual_machine_properties_security_profile_encryption_at_host",
				Description: "This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine or virtual machine scale set This will enable the encryption for all the disks including Resource/Temp disk at host itself <br><br> Default: The Encryption at host will be disabled unless this property is set to true for the resource",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.EncryptionAtHost"),
			},
			{
				Name:        "virtual_machine_properties_security_profile_security_type",
				Description: "Specifies the SecurityType of the virtual machine It is set as TrustedLaunch to enable UefiSettings <br><br> Default: UefiSettings will not be enabled unless this property is set as TrustedLaunch Possible values include: 'SecurityTypesTrustedLaunch'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.SecurityProfile.SecurityType"),
			},
			{
				Name:        "virtual_machine_properties_diagnostics_profile_boot_diagnostics_enabled",
				Description: "Whether boot diagnostics should be enabled on the Virtual Machine",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.Enabled"),
			},
			{
				Name:        "virtual_machine_properties_diagnostics_profile_boot_diagnostics_storage_uri",
				Description: "Uri of the storage account to use for placing the console output and screenshot <br><br>If storageUri is not specified while enabling boot diagnostics, managed storage will be used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.DiagnosticsProfile.BootDiagnostics.StorageURI"),
			},
			{
				Name:        "virtual_machine_properties_availability_set_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.AvailabilitySet.ID"),
			},
			{
				Name:        "virtual_machine_properties_virtual_machine_scale_set_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.VirtualMachineScaleSet.ID"),
			},
			{
				Name:        "virtual_machine_properties_proximity_placement_group_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ProximityPlacementGroup.ID"),
			},
			{
				Name:        "virtual_machine_properties_priority",
				Description: "Specifies the priority for the virtual machine <br><br>Minimum api-version: 2019-03-01 Possible values include: 'Regular', 'Low', 'Spot'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.Priority"),
			},
			{
				Name:        "virtual_machine_properties_eviction_policy",
				Description: "Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01 <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview Possible values include: 'Deallocate', 'Delete'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.EvictionPolicy"),
			},
			{
				Name:        "virtual_machine_properties_billing_profile_max_price",
				Description: "Specifies the maximum price you are willing to pay for a Azure Spot VM/VMSS This price is in US Dollars <br><br> This price will be compared with the current Azure Spot price for the VM size Also, the prices are compared at the time of create/update of Azure Spot VM/VMSS and the operation will only succeed if  the maxPrice is greater than the current Azure Spot price <br><br> The maxPrice will also be used for evicting a Azure Spot VM/VMSS if the current Azure Spot price goes beyond the maxPrice after creation of VM/VMSS <br><br> Possible values are: <br><br> - Any decimal value greater than zero Example: 001538 <br><br> -1 – indicates default price to be up-to on-demand <br><br> You can set the maxPrice to -1 to indicate that the Azure Spot VM/VMSS should not be evicted for price reasons Also, the default max price is -1 if it is not provided by you <br><br>Minimum api-version: 2019-03-01",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("VirtualMachineProperties.BillingProfile.MaxPrice"),
			},
			{
				Name:        "virtual_machine_properties_host_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.Host.ID"),
			},
			{
				Name:        "virtual_machine_properties_host_group_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.HostGroup.ID"),
			},
			{
				Name:        "virtual_machine_properties_provisioning_state",
				Description: "The provisioning state, which only appears in the response",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ProvisioningState"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_platform_update_domain",
				Description: "Specifies the update domain of the virtual machine",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.PlatformUpdateDomain"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_platform_fault_domain",
				Description: "Specifies the fault domain of the virtual machine",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.PlatformFaultDomain"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_computer_name",
				Description: "The computer name assigned to the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.ComputerName"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_os_name",
				Description: "The Operating System running on the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.OsName"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_os_version",
				Description: "The version of Operating System running on the virtual machine",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.OsVersion"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_hyper_v_generation",
				Description: "Specifies the HyperVGeneration Type associated with a resource Possible values include: 'HyperVGenerationTypeV1', 'HyperVGenerationTypeV2'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.HyperVGeneration"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_rdp_thumb_print",
				Description: "The Remote desktop certificate thumbprint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.RdpThumbPrint"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_vm_agent_vm_agent_version",
				Description: "The VM Agent full version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.VMAgent.VMAgentVersion"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_maintenance_redeploy_status_is_customer_initiated_maintenance_allowed",
				Description: "True, if customer is allowed to perform Maintenance",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.IsCustomerInitiatedMaintenanceAllowed"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_maintenance_redeploy_status_pre_maintenance_window_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.PreMaintenanceWindowStartTime.Time"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_maintenance_redeploy_status_pre_maintenance_window_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.PreMaintenanceWindowEndTime.Time"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_maintenance_redeploy_status_maintenance_window_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.MaintenanceWindowStartTime.Time"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_maintenance_redeploy_status_maintenance_window_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.MaintenanceWindowEndTime.Time"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_maintenance_redeploy_status_last_operation_result_code",
				Description: "The Last Maintenance Operation Result Code Possible values include: 'MaintenanceOperationResultCodeTypesNone', 'MaintenanceOperationResultCodeTypesRetryLater', 'MaintenanceOperationResultCodeTypesMaintenanceAborted', 'MaintenanceOperationResultCodeTypesMaintenanceCompleted'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.LastOperationResultCode"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_maintenance_redeploy_status_last_operation_message",
				Description: "Message returned for the last Maintenance Operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.MaintenanceRedeployStatus.LastOperationMessage"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_vm_health_status_code",
				Description: "The status code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.VMHealth.Status.Code"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_vm_health_status_level",
				Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.VMHealth.Status.Level"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_vm_health_status_display_status",
				Description: "The short localizable label for the status",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.VMHealth.Status.DisplayStatus"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_vm_health_status_message",
				Description: "The detailed status message, including for alerts and error messages",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.VMHealth.Status.Message"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_vm_health_status_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.VMHealth.Status.Time.Time"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_console_screenshot_blob_uri",
				Description: "The console screenshot blob URI <br><br>NOTE: This will **not** be set if boot diagnostics is currently enabled with managed storage",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.ConsoleScreenshotBlobURI"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_serial_console_log_blob_uri",
				Description: "The serial console log blob Uri <br><br>NOTE: This will **not** be set if boot diagnostics is currently enabled with managed storage",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.SerialConsoleLogBlobURI"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_status_code",
				Description: "The status code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.Status.Code"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_status_level",
				Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.Status.Level"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_status_display_status",
				Description: "The short localizable label for the status",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.Status.DisplayStatus"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_boot_diagnostics_status_message",
				Description: "The detailed status message, including for alerts and error messages",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.Status.Message"),
			},
			{
				Name:     "virtual_machine_properties_instance_view_boot_diagnostics_status_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("VirtualMachineProperties.InstanceView.BootDiagnostics.Status.Time.Time"),
			},
			{
				Name:        "virtual_machine_properties_instance_view_assigned_host",
				Description: "Resource id of the dedicated host, on which the virtual machine is allocated through automatic placement, when the virtual machine is associated with a dedicated host group that has automatic placement enabled <br><br>Minimum api-version: 2020-06-01",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.InstanceView.AssignedHost"),
			},
			{
				Name:        "virtual_machine_properties_license_type",
				Description: "Specifies that the image or disk that is being used was licensed on-premises <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docsmicrosoftcom/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docsmicrosoftcom/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.LicenseType"),
			},
			{
				Name:        "virtual_machine_properties_vm_id",
				Description: "Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.VMID"),
			},
			{
				Name:        "virtual_machine_properties_extensions_time_budget",
				Description: "Specifies the time alloted for all extensions to start The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format The default value is 90 minutes (PT1H30M) <br><br> Minimum api-version: 2020-06-01",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualMachineProperties.ExtensionsTimeBudget"),
			},
			{
				Name:        "virtual_machine_properties_platform_fault_domain",
				Description: "1<li>This property cannot be updated once the Virtual Machine is created<li>Fault domain assignment can be viewed in the Virtual Machine Instance View<br><br>Minimum api‐version: 2020‐12‐01",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VirtualMachineProperties.PlatformFaultDomain"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal id of virtual machine identity This property will only be provided for a system assigned identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant id associated with the virtual machine This property will only be provided for a system assigned identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the virtual machine The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities The type 'None' will remove any identities from the virtual machine Possible values include: 'ResourceIdentityTypeSystemAssigned', 'ResourceIdentityTypeUserAssigned', 'ResourceIdentityTypeSystemAssignedUserAssigned', 'ResourceIdentityTypeNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user identities associated with the Virtual Machine The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}'",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "zones",
				Description: "The virtual machine zones",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "extended_location_name",
				Description: "The name of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location Possible values include: 'EdgeZone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "resource_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_compute_virtual_machine_properties_storage_profile_data_disks",
				Description: "DataDisk describes a data disk",
				Resolver:    fetchComputeVirtualMachinePropertiesStorageProfileDataDisks,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "lun",
						Description: "Specifies the logical unit number of the data disk This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM",
						Type:        schema.TypeInt,
					},
					{
						Name:        "name",
						Description: "The disk name",
						Type:        schema.TypeString,
					},
					{
						Name:        "vhd_uri",
						Description: "Specifies the virtual hard disk's uri",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Vhd.URI"),
					},
					{
						Name:        "image_uri",
						Description: "Specifies the virtual hard disk's uri",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Image.URI"),
					},
					{
						Name:        "caching",
						Description: "Specifies the caching requirements <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage ReadOnly for Premium storage** Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'",
						Type:        schema.TypeString,
					},
					{
						Name:        "write_accelerator_enabled",
						Description: "Specifies whether writeAccelerator should be enabled or disabled on the disk",
						Type:        schema.TypeBool,
					},
					{
						Name:        "create_option",
						Description: "Specifies how the virtual machine should be created<br><br> Possible values are:<br><br> **Attach** \\u2013 This value is used when you are using a specialized disk to create the virtual machine<br><br> **FromImage** \\u2013 This value is used when you are using an image to create the virtual machine If you are using a platform image, you also use the imageReference element described above If you are using a marketplace image, you  also use the plan element previously described Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'",
						Type:        schema.TypeString,
					},
					{
						Name:        "disk_size_gb",
						Description: "Specifies the size of an empty data disk in gigabytes This element can be used to overwrite the size of the disk in a virtual machine image <br><br> This value cannot be larger than 1023 GB",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DiskSizeGB"),
					},
					{
						Name:        "managed_disk_storage_account_type",
						Description: "Specifies the storage account type for the managed disk NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk Possible values include: 'StorageAccountTypesStandardLRS', 'StorageAccountTypesPremiumLRS', 'StorageAccountTypesStandardSSDLRS', 'StorageAccountTypesUltraSSDLRS', 'StorageAccountTypesPremiumZRS', 'StorageAccountTypesStandardSSDZRS'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedDisk.StorageAccountType"),
					},
					{
						Name:        "managed_disk_disk_encryption_set_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedDisk.DiskEncryptionSet.ID"),
					},
					{
						Name:        "managed_disk_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedDisk.ID"),
					},
					{
						Name:        "to_be_detached",
						Description: "Specifies whether the data disk is in process of detachment from the VirtualMachine/VirtualMachineScaleset",
						Type:        schema.TypeBool,
					},
					{
						Name:        "disk_iops_read_write",
						Description: "Specifies the Read-Write IOPS for the managed disk when StorageAccountType is UltraSSD_LRS Returned only for VirtualMachine ScaleSet VM disks Can be updated only via updates to the VirtualMachine Scale Set",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DiskIOPSReadWrite"),
					},
					{
						Name:        "disk_mb_ps_read_write",
						Description: "Specifies the bandwidth in MB per second for the managed disk when StorageAccountType is UltraSSD_LRS Returned only for VirtualMachine ScaleSet VM disks Can be updated only via updates to the VirtualMachine Scale Set",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DiskMBpsReadWrite"),
					},
					{
						Name:        "detach_option",
						Description: "Specifies the detach behavior to be used while detaching a disk or which is already in the process of detachment from the virtual machine Supported values: **ForceDetach** <br><br> detachOption: **ForceDetach** is applicable only for managed data disks If a previous detachment attempt of the data disk did not complete due to an unexpected failure from the virtual machine and the disk is still not released then use force-detach as a last resort option to detach the disk forcibly from the VM All writes might not have been flushed when using this detach behavior <br><br> This feature is still in preview mode and is not supported for VirtualMachineScaleSet To force-detach a data disk update toBeDetached to 'true' along with setting detachOption: 'ForceDetach' Possible values include: 'ForceDetach'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_os_profile_windows_configuration_additional_unattend_contents",
				Description: "AdditionalUnattendContent specifies additional XML formatted information that can be included in the Unattendxml file, which is used by Windows Setup Contents are defined by setting name, component name, and the pass in which the content is applied",
				Resolver:    fetchComputeVirtualMachinePropertiesOsProfileWindowsConfigurationAdditionalUnattendContents,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "pass_name",
						Description: "The pass name Currently, the only allowable value is OobeSystem Possible values include: 'OobeSystem'",
						Type:        schema.TypeString,
					},
					{
						Name:        "component_name",
						Description: "The component name Currently, the only allowable value is Microsoft-Windows-Shell-Setup Possible values include: 'MicrosoftWindowsShellSetup'",
						Type:        schema.TypeString,
					},
					{
						Name:        "setting_name",
						Description: "Specifies the name of the setting to which the content applies Possible values are: FirstLogonCommands and AutoLogon Possible values include: 'AutoLogon', 'FirstLogonCommands'",
						Type:        schema.TypeString,
					},
					{
						Name:        "content",
						Description: "Specifies the XML formatted content that is added to the unattendxml file for the specified path and component The XML must be less than 4KB and must include the root element for the setting or feature that is being inserted",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_os_profile_windows_configuration_win_r_m_listeners",
				Description: "WinRMListener describes Protocol and thumbprint of Windows Remote Management listener",
				Resolver:    fetchComputeVirtualMachinePropertiesOsProfileWindowsConfigurationWinRMListeners,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "protocol",
						Description: "Specifies the protocol of WinRM listener <br><br> Possible values are: <br>**http** <br><br> **https** Possible values include: 'HTTP', 'HTTPS'",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_url",
						Description: "This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  \"data\":\"<Base64-encoded-certificate>\",<br>  \"dataType\":\"pfx\",<br>  \"password\":\"<pfx-file-password>\"<br>}",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CertificateURL"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_os_profile_linux_configuration_ssh_public_keys",
				Description: "SSHPublicKey contains information about SSH certificate public key and the path on the Linux VM where the public key is placed",
				Resolver:    fetchComputeVirtualMachinePropertiesOsProfileLinuxConfigurationSshPublicKeys,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "path",
						Description: "Specifies the full path on the created VM where ssh public key is stored If the file already exists, the specified key is appended to the file Example: /home/user/ssh/authorized_keys",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_data",
						Description: "SSH public key certificate used to authenticate with the VM through ssh The key needs to be at least 2048-bit and in ssh-rsa format <br><br> For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docsmicrosoftcom/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_os_profile_secrets",
				Description: "VaultSecretGroup describes a set of certificates which are all in the same Key Vault",
				Resolver:    fetchComputeVirtualMachinePropertiesOsProfileSecrets,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "source_vault_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceVault.ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_properties_os_profile_secret_vault_certificates",
						Description: "VaultCertificate describes a single certificate reference in a Key Vault, and where the certificate should reside on the VM",
						Resolver:    fetchComputeVirtualMachinePropertiesOsProfileSecretVaultCertificates,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_properties_os_profile_secret_id",
								Description: "Unique ID of azure_compute_virtual_machine_properties_os_profile_secrets table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "certificate_url",
								Description: "This is the URL of a certificate that has been uploaded to Key Vault as a secret For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docsmicrosoftcom/azure/key-vault/key-vault-get-started/#add) In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  \"data\":\"<Base64-encoded-certificate>\",<br>  \"dataType\":\"pfx\",<br>  \"password\":\"<pfx-file-password>\"<br>}",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("CertificateURL"),
							},
							{
								Name:        "certificate_store",
								Description: "UppercaseThumbprint&gt;crt for the X509 certificate file and &lt;UppercaseThumbprint&gt;prv for private key Both of these files are pem formatted",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_network_profile_network_interfaces",
				Description: "NetworkInterfaceReference describes a network interface reference",
				Resolver:    fetchComputeVirtualMachinePropertiesNetworkProfileNetworkInterfaces,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "network_interface_reference_properties_primary",
						Description: "Specifies the primary network interface in case the virtual machine has more than 1 network interface",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("NetworkInterfaceReferenceProperties.Primary"),
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_instance_view_vm_agent_extension_handlers",
				Description: "VirtualMachineExtensionHandlerInstanceView the instance view of a virtual machine extension handler",
				Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewVmAgentExtensionHandlers,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "an example is \"CustomScriptExtension\"",
						Type:        schema.TypeString,
					},
					{
						Name:        "type_handler_version",
						Description: "Specifies the version of the script handler",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_code",
						Description: "The status code",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.Code"),
					},
					{
						Name:        "status_level",
						Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.Level"),
					},
					{
						Name:        "status_display_status",
						Description: "The short localizable label for the status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.DisplayStatus"),
					},
					{
						Name:        "status_message",
						Description: "The detailed status message, including for alerts and error messages",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.Message"),
					},
					{
						Name:     "status_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Status.Time.Time"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_instance_view_vm_agent_statuses",
				Description: "InstanceViewStatus instance view status",
				Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewVmAgentStatuses,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "code",
						Description: "The status code",
						Type:        schema.TypeString,
					},
					{
						Name:        "level",
						Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_status",
						Description: "The short localizable label for the status",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "The detailed status message, including for alerts and error messages",
						Type:        schema.TypeString,
					},
					{
						Name:     "time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Time.Time"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_instance_view_disks",
				Description: "DiskInstanceView the instance view of the disk",
				Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewDisks,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The disk name",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_properties_instance_view_disk_encryption_settings",
						Description: "DiskEncryptionSettings describes a Encryption Settings for a Disk",
						Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewDiskEncryptionSettings,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_properties_instance_view_disk_id",
								Description: "Unique ID of azure_compute_virtual_machine_properties_instance_view_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "disk_encryption_key_secret_url",
								Description: "The URL referencing a secret in a Key Vault",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DiskEncryptionKey.SecretURL"),
							},
							{
								Name:        "disk_encryption_key_source_vault_id",
								Description: "Resource Id",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("DiskEncryptionKey.SourceVault.ID"),
							},
							{
								Name:        "key_encryption_key_key_url",
								Description: "The URL referencing a key encryption key in Key Vault",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeyEncryptionKey.KeyURL"),
							},
							{
								Name:        "key_encryption_key_source_vault_id",
								Description: "Resource Id",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("KeyEncryptionKey.SourceVault.ID"),
							},
							{
								Name:        "enabled",
								Description: "Specifies whether disk encryption should be enabled on the virtual machine",
								Type:        schema.TypeBool,
							},
						},
					},
					{
						Name:        "azure_compute_virtual_machine_properties_instance_view_disk_statuses",
						Description: "InstanceViewStatus instance view status",
						Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewDiskStatuses,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_properties_instance_view_disk_id",
								Description: "Unique ID of azure_compute_virtual_machine_properties_instance_view_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "code",
								Description: "The status code",
								Type:        schema.TypeString,
							},
							{
								Name:        "level",
								Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
								Type:        schema.TypeString,
							},
							{
								Name:        "display_status",
								Description: "The short localizable label for the status",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The detailed status message, including for alerts and error messages",
								Type:        schema.TypeString,
							},
							{
								Name:     "time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("Time.Time"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_instance_view_extensions",
				Description: "VirtualMachineExtensionInstanceView the instance view of a virtual machine extension",
				Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewExtensions,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The virtual machine extension name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "an example is \"CustomScriptExtension\"",
						Type:        schema.TypeString,
					},
					{
						Name:        "type_handler_version",
						Description: "Specifies the version of the script handler",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_properties_instance_view_extension_substatuses",
						Description: "InstanceViewStatus instance view status",
						Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewExtensionSubstatuses,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_properties_instance_view_extension_id",
								Description: "Unique ID of azure_compute_virtual_machine_properties_instance_view_extensions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "code",
								Description: "The status code",
								Type:        schema.TypeString,
							},
							{
								Name:        "level",
								Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
								Type:        schema.TypeString,
							},
							{
								Name:        "display_status",
								Description: "The short localizable label for the status",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The detailed status message, including for alerts and error messages",
								Type:        schema.TypeString,
							},
							{
								Name:     "time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("Time.Time"),
							},
						},
					},
					{
						Name:        "azure_compute_virtual_machine_properties_instance_view_extension_statuses",
						Description: "InstanceViewStatus instance view status",
						Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewExtensionStatuses,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_properties_instance_view_extension_id",
								Description: "Unique ID of azure_compute_virtual_machine_properties_instance_view_extensions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "code",
								Description: "The status code",
								Type:        schema.TypeString,
							},
							{
								Name:        "level",
								Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
								Type:        schema.TypeString,
							},
							{
								Name:        "display_status",
								Description: "The short localizable label for the status",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The detailed status message, including for alerts and error messages",
								Type:        schema.TypeString,
							},
							{
								Name:     "time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("Time.Time"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_properties_instance_view_statuses",
				Description: "InstanceViewStatus instance view status",
				Resolver:    fetchComputeVirtualMachinePropertiesInstanceViewStatuses,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "code",
						Description: "The status code",
						Type:        schema.TypeString,
					},
					{
						Name:        "level",
						Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_status",
						Description: "The short localizable label for the status",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "The detailed status message, including for alerts and error messages",
						Type:        schema.TypeString,
					},
					{
						Name:     "time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Time.Time"),
					},
				},
			},
			{
				Name:        "azure_compute_virtual_machine_resources",
				Description: "VirtualMachineExtension describes a Virtual Machine Extension",
				Resolver:    fetchComputeVirtualMachineResources,
				Columns: []schema.Column{
					{
						Name:        "virtual_machine_id",
						Description: "Unique ID of azure_compute_virtual_machines table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_machine_extension_properties_force_update_tag",
						Description: "How the extension handler should be forced to update even if the extension configuration has not changed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.ForceUpdateTag"),
					},
					{
						Name:        "virtual_machine_extension_properties_publisher",
						Description: "The name of the extension handler publisher",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.Publisher"),
					},
					{
						Name:        "virtual_machine_extension_properties_type",
						Description: "an example is \"CustomScriptExtension\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.Type"),
					},
					{
						Name:        "virtual_machine_extension_properties_type_handler_version",
						Description: "Specifies the version of the script handler",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.TypeHandlerVersion"),
					},
					{
						Name:        "virtual_machine_extension_properties_auto_upgrade_minor_version",
						Description: "Indicates whether the extension should use a newer minor version if one is available at deployment time Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.AutoUpgradeMinorVersion"),
					},
					{
						Name:        "virtual_machine_extension_properties_enable_automatic_upgrade",
						Description: "Indicates whether the extension should be automatically upgraded by the platform if there is a newer version of the extension available",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.EnableAutomaticUpgrade"),
					},
					{
						Name:        "virtual_machine_extension_properties_settings",
						Description: "Json formatted public settings for the extension",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeVirtualMachineResourceVirtualMachineExtensionPropertiesSettings,
					},
					{
						Name:        "virtual_machine_extension_properties_protected_settings",
						Description: "The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeVirtualMachineResourceVirtualMachineExtensionPropertiesProtectedSettings,
					},
					{
						Name:        "virtual_machine_extension_properties_provisioning_state",
						Description: "The provisioning state, which only appears in the response",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.ProvisioningState"),
					},
					{
						Name:        "virtual_machine_extension_properties_instance_view_name",
						Description: "The virtual machine extension name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.InstanceView.Name"),
					},
					{
						Name:        "virtual_machine_extension_properties_instance_view_type",
						Description: "an example is \"CustomScriptExtension\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.InstanceView.Type"),
					},
					{
						Name:        "virtual_machine_extension_properties_instance_view_type_handler_version",
						Description: "Specifies the version of the script handler",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualMachineExtensionProperties.InstanceView.TypeHandlerVersion"),
					},
					{
						Name:        "resource_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "Resource location",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "Resource tags",
						Type:        schema.TypeJSON,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_compute_virtual_machine_resource_virtual_machine_extension_properties_instance_view_substatuses",
						Description: "InstanceViewStatus instance view status",
						Resolver:    fetchComputeVirtualMachineResourceVirtualMachineExtensionPropertiesInstanceViewSubstatuses,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_resource_id",
								Description: "Unique ID of azure_compute_virtual_machine_resources table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "code",
								Description: "The status code",
								Type:        schema.TypeString,
							},
							{
								Name:        "level",
								Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
								Type:        schema.TypeString,
							},
							{
								Name:        "display_status",
								Description: "The short localizable label for the status",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The detailed status message, including for alerts and error messages",
								Type:        schema.TypeString,
							},
							{
								Name:     "time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("Time.Time"),
							},
						},
					},
					{
						Name:        "azure_compute_virtual_machine_resource_virtual_machine_extension_properties_instance_view_statuses",
						Description: "InstanceViewStatus instance view status",
						Resolver:    fetchComputeVirtualMachineResourceVirtualMachineExtensionPropertiesInstanceViewStatuses,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_resource_id",
								Description: "Unique ID of azure_compute_virtual_machine_resources table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "code",
								Description: "The status code",
								Type:        schema.TypeString,
							},
							{
								Name:        "level",
								Description: "The level code Possible values include: 'Info', 'Warning', 'Error'",
								Type:        schema.TypeString,
							},
							{
								Name:        "display_status",
								Description: "The short localizable label for the status",
								Type:        schema.TypeString,
							},
							{
								Name:        "message",
								Description: "The detailed status message, including for alerts and error messages",
								Type:        schema.TypeString,
							},
							{
								Name:     "time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("Time.Time"),
							},
						},
					},
					{
						Name:        "azure_compute_virtual_machine_resource_virtual_machine_properties_network_profile_network_interfaces",
						Description: "NetworkInterfaceReference describes a network interface reference",
						Resolver:    fetchComputeVirtualMachineResourceVirtualMachinePropertiesNetworkProfileNetworkInterfaces,
						Columns: []schema.Column{
							{
								Name:        "virtual_machine_resource_id",
								Description: "Unique ID of azure_compute_virtual_machine_resources table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "network_interface_reference_properties_primary",
								Description: "Specifies the primary network interface in case the virtual machine has more than 1 network interface",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("NetworkInterfaceReferenceProperties.Primary"),
							},
							{
								Name:        "resource_id",
								Description: "Resource Id",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Compute.VirtualMachines
	response, err := svc.ListAll(ctx, "false")
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func fetchComputeVirtualMachinePropertiesStorageProfileDataDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}
	if p.VirtualMachineProperties == nil || p.VirtualMachineProperties.StorageProfile == nil || p.VirtualMachineProperties.StorageProfile.DataDisks == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.StorageProfile.DataDisks
	return nil
}
func fetchComputeVirtualMachinePropertiesOsProfileWindowsConfigurationAdditionalUnattendContents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}
	if p.OsProfile == nil ||
		p.OsProfile.WindowsConfiguration == nil ||
		p.OsProfile.WindowsConfiguration.AdditionalUnattendContent == nil {
		return nil
	}

	res <- *p.OsProfile.WindowsConfiguration.AdditionalUnattendContent
	return nil
}
func fetchComputeVirtualMachinePropertiesOsProfileWindowsConfigurationWinRMListeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}
	if p.OsProfile == nil ||
		p.OsProfile.WindowsConfiguration == nil ||
		p.OsProfile.WindowsConfiguration.WinRM == nil ||
		p.OsProfile.WindowsConfiguration.WinRM.Listeners == nil {
		return nil
	}

	res <- *p.OsProfile.WindowsConfiguration.WinRM.Listeners
	return nil
}
func fetchComputeVirtualMachinePropertiesOsProfileLinuxConfigurationSshPublicKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.OsProfile == nil ||
		p.OsProfile.LinuxConfiguration == nil ||
		p.OsProfile.LinuxConfiguration.SSH == nil ||
		p.OsProfile.LinuxConfiguration.SSH.PublicKeys == nil {
		return nil
	}

	res <- *p.OsProfile.LinuxConfiguration.SSH.PublicKeys
	return nil
}
func fetchComputeVirtualMachinePropertiesOsProfileSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.OsProfile == nil || p.OsProfile.Secrets == nil {
		return nil
	}

	res <- *p.OsProfile.Secrets
	return nil
}
func fetchComputeVirtualMachinePropertiesOsProfileSecretVaultCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VaultSecretGroup)
	if !ok {
		return fmt.Errorf("expected to have compute.VaultSecretGroup but got %T", parent.Item)
	}

	if p.VaultCertificates == nil {
		return nil
	}

	res <- *p.VaultCertificates
	return nil
}
func fetchComputeVirtualMachinePropertiesNetworkProfileNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.NetworkProfile == nil ||
		p.VirtualMachineProperties.NetworkProfile.NetworkInterfaces == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.NetworkProfile.NetworkInterfaces
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewVmAgentExtensionHandlers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.InstanceView == nil ||
		p.VirtualMachineProperties.InstanceView.VMAgent == nil ||
		p.VirtualMachineProperties.InstanceView.VMAgent.ExtensionHandlers == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.InstanceView.VMAgent.ExtensionHandlers
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewVmAgentStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.InstanceView == nil ||
		p.VirtualMachineProperties.InstanceView.VMAgent == nil ||
		p.VirtualMachineProperties.InstanceView.VMAgent.Statuses == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.InstanceView.VMAgent.Statuses
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.InstanceView == nil ||
		p.VirtualMachineProperties.InstanceView.Disks == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.InstanceView.Disks
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewDiskEncryptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.DiskInstanceView)
	if !ok {
		return fmt.Errorf("expected to have compute.DiskInstanceView but got %T", parent.Item)
	}

	if p.EncryptionSettings == nil {
		return nil
	}

	res <- *p.EncryptionSettings
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewDiskStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.DiskInstanceView)
	if !ok {
		return fmt.Errorf("expected to have compute.Disk but got %T", parent.Item)
	}

	if p.Statuses == nil {
		return nil
	}

	res <- *p.Statuses
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewExtensions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.InstanceView == nil ||
		p.VirtualMachineProperties.InstanceView.Extensions == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.InstanceView.Extensions
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewExtensionSubstatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachineExtensionInstanceView)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtensionInstanceView but got %T", parent.Item)
	}
	if p.Substatuses == nil {
		return nil
	}

	res <- *p.Substatuses
	return nil
}

func fetchComputeVirtualMachinePropertiesInstanceViewExtensionStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachineExtensionInstanceView)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtensionInstanceView but got %T", parent.Item)
	}

	if p.Statuses == nil {
		return nil
	}

	res <- *p.Statuses
	return nil
}
func fetchComputeVirtualMachinePropertiesInstanceViewStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.VirtualMachineProperties == nil ||
		p.VirtualMachineProperties.InstanceView == nil ||
		p.VirtualMachineProperties.InstanceView.Statuses == nil {
		return nil
	}

	res <- *p.VirtualMachineProperties.InstanceView.Statuses
	return nil
}
func fetchComputeVirtualMachineResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachine but got %T", parent.Item)
	}

	if p.Resources == nil {
		return nil
	}

	res <- *p.Resources
	return nil
}
func resolveComputeVirtualMachineResourceVirtualMachineExtensionPropertiesSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", resource.Item)
	}

	return resource.Set(c.Name, p.Settings)
}
func resolveComputeVirtualMachineResourceVirtualMachineExtensionPropertiesProtectedSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", resource.Item)
	}
	return resource.Set(c.Name, p.ProtectedSettings)
}
func fetchComputeVirtualMachineResourceVirtualMachineExtensionPropertiesInstanceViewSubstatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", parent.Item)
	}

	if p.VirtualMachineExtensionProperties == nil ||
		p.VirtualMachineExtensionProperties.InstanceView == nil ||
		p.VirtualMachineExtensionProperties.InstanceView.Substatuses == nil {
		return nil
	}

	res <- *p.VirtualMachineExtensionProperties.InstanceView.Substatuses
	return nil
}
func fetchComputeVirtualMachineResourceVirtualMachineExtensionPropertiesInstanceViewStatuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", parent.Item)
	}

	if p.VirtualMachineExtensionProperties == nil ||
		p.VirtualMachineExtensionProperties.InstanceView == nil ||
		p.VirtualMachineExtensionProperties.InstanceView.Statuses == nil {
		return nil
	}

	res <- *p.VirtualMachineExtensionProperties.InstanceView.Statuses
	return nil
}
func fetchComputeVirtualMachineResourceVirtualMachinePropertiesNetworkProfileNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(compute.VirtualMachineExtension)
	if !ok {
		return fmt.Errorf("expected to have compute.VirtualMachineExtension but got %T", parent.Item)
	}

	if p.VirtualMachineExtensionProperties == nil ||
		p.VirtualMachineExtensionProperties.InstanceView == nil ||
		p.VirtualMachineExtensionProperties.InstanceView.Statuses == nil {
		return nil
	}

	res <- *p.VirtualMachineExtensionProperties.InstanceView.Statuses
	return nil
}
