
# Table: azure_ad_user_managed_devices

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of azure_ad_users table (FK)|
|id|text||
|user_id|text||
|device_name|text||
|managed_device_owner_type|text||
|enrolled_date_time|timestamp without time zone||
|last_sync_date_time|timestamp without time zone||
|operating_system|text||
|compliance_state|text||
|jail_broken|text||
|management_agent|text||
|os_version|text||
|eas_activated|boolean||
|eas_device_id|text||
|eas_activation_date_time|timestamp without time zone||
|azure_a_d_registered|boolean||
|device_enrollment_type|text||
|activation_lock_bypass_code|text||
|email_address|text||
|azure_a_d_device_id|text||
|device_registration_state|text||
|device_category_display_name|text||
|is_supervised|boolean||
|exchange_last_successful_sync_date_time|timestamp without time zone||
|exchange_access_state|text||
|exchange_access_state_reason|text||
|remote_assistance_session_url|text||
|remote_assistance_session_error_details|text||
|is_encrypted|boolean||
|user_principal_name|text||
|model|text||
|manufacturer|text||
|imei|text||
|compliance_grace_period_expiration_date_time|timestamp without time zone||
|serial_number|text||
|phone_number|text||
|android_security_patch_level|text||
|user_display_name|text||
|configuration_manager_client_enabled_features_inventory|boolean||
|configuration_manager_client_enabled_features_modern_apps|boolean||
|configuration_manager_client_enabled_features_resource_access|boolean||
|configuration_manager_client_enabled_features_device_configuration|boolean||
|configuration_manager_client_enabled_features_compliance_policy|boolean||
|configuration_manager_client_enabled_features_windows_update_for_business|boolean||
|wi_fi_mac_address|text||
|device_health_attestation_state_last_update_date_time|text||
|device_health_attestation_state_content_namespace_url|text||
|device_health_attestation_state_device_health_attestation_status|text||
|device_health_attestation_state_content_version|text||
|device_health_attestation_state_issued_date_time|timestamp without time zone||
|device_health_attestation_state_attestation_identity_key|text||
|device_health_attestation_state_reset_count|bigint||
|device_health_attestation_state_restart_count|bigint||
|device_health_attestation_state_data_excution_policy|text||
|device_health_attestation_state_bit_locker_status|text||
|device_health_attestation_state_boot_manager_version|text||
|device_health_attestation_state_code_integrity_check_version|text||
|device_health_attestation_state_secure_boot|text||
|device_health_attestation_state_boot_debugging|text||
|device_health_attestation_state_operating_system_kernel_debugging|text||
|device_health_attestation_state_code_integrity|text||
|device_health_attestation_state_test_signing|text||
|device_health_attestation_state_safe_mode|text||
|device_health_attestation_state_windows_p_e|text||
|device_health_attestation_state_early_launch_anti_malware_driver_protection|text||
|device_health_attestation_state_virtual_secure_mode|text||
|device_health_attestation_state_pcr_hash_algorithm|text||
|device_health_attestation_state_boot_app_security_version|text||
|device_health_attestation_state_boot_manager_security_version|text||
|device_health_attestation_state_tpm_version|text||
|device_health_attestation_state_pcr0|text||
|device_health_attestation_state_secure_boot_configuration_policy_finger_print|text||
|device_health_attestation_state_code_integrity_policy|text||
|device_health_attestation_state_boot_revision_list_info|text||
|device_health_attestation_state_operating_system_rev_list_info|text||
|device_health_attestation_state_health_status_mismatch_info|text||
|device_health_attestation_state_health_attestation_supported_status|text||
|subscriber_carrier|text||
|meid|text||
|total_storage_space_in_bytes|bigint||
|free_storage_space_in_bytes|bigint||
|managed_device_name|text||
|partner_reported_threat_state|text||
