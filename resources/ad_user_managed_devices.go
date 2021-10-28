package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdUserManagedDevices() *schema.Table {
	return &schema.Table{
		Name:     "azure_ad_user_managed_devices",
		Resolver: fetchAdUserManagedDevices,
		Columns: []schema.Column{
			{
				Name:        "user_cq_id",
				Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Entity.ID"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserID"),
			},
			{
				Name: "device_name",
				Type: schema.TypeString,
			},
			{
				Name: "managed_device_owner_type",
				Type: schema.TypeString,
			},
			{
				Name: "enrolled_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "last_sync_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "operating_system",
				Type: schema.TypeString,
			},
			{
				Name: "compliance_state",
				Type: schema.TypeString,
			},
			{
				Name: "jail_broken",
				Type: schema.TypeString,
			},
			{
				Name: "management_agent",
				Type: schema.TypeString,
			},
			{
				Name: "os_version",
				Type: schema.TypeString,
			},
			{
				Name: "eas_activated",
				Type: schema.TypeBool,
			},
			{
				Name:     "eas_device_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EasDeviceID"),
			},
			{
				Name: "eas_activation_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "azure_ad_registered",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AzureADRegistered"),
			},
			{
				Name: "device_enrollment_type",
				Type: schema.TypeString,
			},
			{
				Name: "activation_lock_bypass_code",
				Type: schema.TypeString,
			},
			{
				Name: "email_address",
				Type: schema.TypeString,
			},
			{
				Name:     "azure_ad_device_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AzureADDeviceID"),
			},
			{
				Name: "device_registration_state",
				Type: schema.TypeString,
			},
			{
				Name: "device_category_display_name",
				Type: schema.TypeString,
			},
			{
				Name: "is_supervised",
				Type: schema.TypeBool,
			},
			{
				Name: "exchange_last_successful_sync_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "exchange_access_state",
				Type: schema.TypeString,
			},
			{
				Name: "exchange_access_state_reason",
				Type: schema.TypeString,
			},
			{
				Name:     "remote_assistance_session_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RemoteAssistanceSessionURL"),
			},
			{
				Name: "remote_assistance_session_error_details",
				Type: schema.TypeString,
			},
			{
				Name: "is_encrypted",
				Type: schema.TypeBool,
			},
			{
				Name: "user_principal_name",
				Type: schema.TypeString,
			},
			{
				Name: "model",
				Type: schema.TypeString,
			},
			{
				Name: "manufacturer",
				Type: schema.TypeString,
			},
			{
				Name: "imei",
				Type: schema.TypeString,
			},
			{
				Name: "compliance_grace_period_expiration_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "serial_number",
				Type: schema.TypeString,
			},
			{
				Name: "phone_number",
				Type: schema.TypeString,
			},
			{
				Name: "android_security_patch_level",
				Type: schema.TypeString,
			},
			{
				Name: "user_display_name",
				Type: schema.TypeString,
			},
			{
				Name:     "enabled_features_inventory",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.Inventory"),
			},
			{
				Name:     "enabled_features_modern_apps",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.ModernApps"),
			},
			{
				Name:     "enabled_features_resource_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.ResourceAccess"),
			},
			{
				Name:     "enabled_features_device_configuration",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.DeviceConfiguration"),
			},
			{
				Name:     "enabled_features_compliance_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.CompliancePolicy"),
			},
			{
				Name:     "enabled_features_windows_update_for_business",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.WindowsUpdateForBusiness"),
			},
			{
				Name: "wi_fi_mac_address",
				Type: schema.TypeString,
			},
			{
				Name:     "health_state_last_update_date_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.LastUpdateDateTime"),
			},
			{
				Name:     "health_state_content_namespace_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.ContentNamespaceURL"),
			},
			{
				Name:     "health_state_device_health_attestation_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.DeviceHealthAttestationStatus"),
			},
			{
				Name:     "health_state_content_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.ContentVersion"),
			},
			{
				Name:     "health_state_issued_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.IssuedDateTime"),
			},
			{
				Name:     "health_state_attestation_identity_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.AttestationIdentityKey"),
			},
			{
				Name:     "health_state_reset_count",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.ResetCount"),
			},
			{
				Name:     "health_state_restart_count",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.RestartCount"),
			},
			{
				Name:     "health_state_data_excution_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.DataExcutionPolicy"),
			},
			{
				Name:     "health_state_bit_locker_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BitLockerStatus"),
			},
			{
				Name:     "health_state_boot_manager_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BootManagerVersion"),
			},
			{
				Name:     "health_state_code_integrity_check_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrityCheckVersion"),
			},
			{
				Name:     "health_state_secure_boot",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.SecureBoot"),
			},
			{
				Name:     "health_state_boot_debugging",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BootDebugging"),
			},
			{
				Name:     "health_state_operating_system_kernel_debugging",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.OperatingSystemKernelDebugging"),
			},
			{
				Name:     "health_state_code_integrity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrity"),
			},
			{
				Name:     "health_state_test_signing",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.TestSigning"),
			},
			{
				Name:     "health_state_safe_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.SafeMode"),
			},
			{
				Name:     "health_state_windows_p_e",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.WindowsPE"),
			},
			{
				Name:     "health_state_early_launch_anti_malware_driver_protection",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.EarlyLaunchAntiMalwareDriverProtection"),
			},
			{
				Name:     "health_state_virtual_secure_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.VirtualSecureMode"),
			},
			{
				Name:     "health_state_pcr_hash_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.PcrHashAlgorithm"),
			},
			{
				Name:     "health_state_boot_app_security_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BootAppSecurityVersion"),
			},
			{
				Name:     "health_state_boot_manager_security_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BootManagerSecurityVersion"),
			},
			{
				Name:     "health_state_tpm_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.TpmVersion"),
			},
			{
				Name:     "health_state_pcr0",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.Pcr0"),
			},
			{
				Name:     "health_state_secure_boot_configuration_policy_finger_print",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.SecureBootConfigurationPolicyFingerPrint"),
			},
			{
				Name:     "health_state_code_integrity_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrityPolicy"),
			},
			{
				Name:     "health_state_boot_revision_list_info",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.BootRevisionListInfo"),
			},
			{
				Name:     "health_state_operating_system_rev_list_info",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.OperatingSystemRevListInfo"),
			},
			{
				Name:     "health_state_health_status_mismatch_info",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.HealthStatusMismatchInfo"),
			},
			{
				Name:     "health_state_health_attestation_supported_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceHealthAttestationState.HealthAttestationSupportedStatus"),
			},
			{
				Name: "subscriber_carrier",
				Type: schema.TypeString,
			},
			{
				Name: "meid",
				Type: schema.TypeString,
			},
			{
				Name: "total_storage_space_in_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "free_storage_space_in_bytes",
				Type: schema.TypeBigInt,
			},
			{
				Name: "managed_device_name",
				Type: schema.TypeString,
			},
			{
				Name: "partner_reported_threat_state",
				Type: schema.TypeString,
			},
			{
				Name: "device_category",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_ad_user_managed_device_device_action_results",
				Resolver: fetchAdUserManagedDeviceDeviceActionResults,
				Columns: []schema.Column{
					{
						Name:        "user_managed_device_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "action_name",
						Type: schema.TypeString,
					},
					{
						Name: "action_state",
						Type: schema.TypeString,
					},
					{
						Name: "start_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_updated_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_managed_device_device_configuration_states",
				Resolver: fetchAdUserManagedDeviceDeviceConfigurationStates,
				Columns: []schema.Column{
					{
						Name:        "user_managed_device_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name:     "setting_states",
						Type:     schema.TypeJSON,
						Resolver: resolveAdUserManagedDeviceDeviceConfigurationStatesSettingStates,
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeBigInt,
					},
					{
						Name: "platform_type",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "setting_count",
						Type: schema.TypeBigInt,
					},
				},
			},
			{
				Name:     "azure_ad_user_managed_device_device_compliance_policy_states",
				Resolver: fetchAdUserManagedDeviceDeviceCompliancePolicyStates,
				Columns: []schema.Column{
					{
						Name:        "user_managed_device_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name:     "setting_states",
						Type:     schema.TypeJSON,
						Resolver: resolveAdUserManagedDeviceDeviceCompliancePolicyStatesSettingStates,
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeBigInt,
					},
					{
						Name: "platform_type",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "setting_count",
						Type: schema.TypeBigInt,
					},
				},
			},
		},
	}
}

func fetchAdUserManagedDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.ManagedDevices
	return nil
}
func fetchAdUserManagedDeviceDeviceActionResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceActionResults
	return nil
}
func fetchAdUserManagedDeviceDeviceConfigurationStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceConfigurationStates
	return nil
}
func resolveAdUserManagedDeviceDeviceConfigurationStatesSettingStates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.DeviceConfigurationState)
	if !ok {
		return fmt.Errorf("expected to have msgraph.DeviceConfigurationState but got %T", resource.Item)
	}
	out, err := json.Marshal(p.SettingStates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func fetchAdUserManagedDeviceDeviceCompliancePolicyStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceCompliancePolicyStates
	return nil
}
func resolveAdUserManagedDeviceDeviceCompliancePolicyStatesSettingStates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.DeviceCompliancePolicyState)
	if !ok {
		return fmt.Errorf("expected to have msgraph.DeviceCompliancePolicyState but got %T", resource.Item)
	}
	out, err := json.Marshal(p.SettingStates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
