ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "windows_configuration_patch_settings_assessment_mode" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "linux_configuration_patch_settings_assessment_mode" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "network_profile_network_api_version" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "network_profile_network_interface_configurations" jsonb;
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "scheduled_events_profile" jsonb;
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    ADD COLUMN "user_data" text;


ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "virtual_machine_extension_properties";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "force_update_tag" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "publisher" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "type_handler_version" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "auto_upgrade_minor_version" boolean;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "enable_automatic_upgrade" boolean;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "settings" jsonb;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "protected_settings" jsonb;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "provisioning_state" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "extension_type" text;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "instance_view" jsonb;

--it was duplicated as a json column of virtual machine
DROP TABLE IF EXISTS "azure_compute_virtual_machine_network_interfaces";

CREATE TABLE IF NOT EXISTS "azure_resources_links"
(
    "cq_id"           uuid                        NOT NULL,
    "cq_meta"         jsonb,
    "cq_fetch_date"   timestamp without time zone NOT NULL,
    "subscription_id" text,
    "id"              text,
    "name"            text,
    "type"            text,
    "source_id"       text,
    "target_id"       text,
    "notes"           text,
    CONSTRAINT azure_resources_links_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_resources_links');


-- Resource: datalake.analytics_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_accounts"
(
    "cq_id"                                                        uuid                        NOT NULL,
    "cq_meta"                                                      jsonb,
    "cq_fetch_date"                                                timestamp without time zone NOT NULL,
    "subscription_id"                                              text,
    "identity_type"                                                text,
    "identity_principal_id"                                        uuid,
    "identity_tenant_id"                                           uuid,
    "default_group"                                                text,
    "encryption_config_type"                                       text,
    "encryption_config_key_vault_meta_info_key_vault_resource_id"  text,
    "encryption_config_key_vault_meta_info_encryption_key_name"    text,
    "encryption_config_key_vault_meta_info_encryption_key_version" text,
    "encryption_state"                                             text,
    "encryption_provisioning_state"                                text,
    "firewall_state"                                               text,
    "firewall_allow_azure_ips"                                     text,
    "trusted_id_provider_state"                                    text,
    "new_tier"                                                     text,
    "current_tier"                                                 text,
    "account_id"                                                   uuid,
    "provisioning_state"                                           text,
    "state"                                                        text,
    "creation_time"                                                timestamp without time zone,
    "last_modified_time"                                           timestamp without time zone,
    "endpoint"                                                     text,
    "id"                                                           text,
    "name"                                                         text,
    "type"                                                         text,
    "location"                                                     text,
    "tags"                                                         jsonb,
    CONSTRAINT azure_datalake_storage_accounts_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_datalake_storage_accounts');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_firewall_rules"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "start_ip_address"      text,
    "end_ip_address"        text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_firewall_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_firewall_rules (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_firewall_rules', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_virtual_network_rules"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "subnet_id"             text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_virtual_network_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_virtual_network_rules (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_virtual_network_rules', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_trusted_id_providers"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "id_provider"           text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_trusted_id_providers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_trusted_id_providers (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_trusted_id_providers', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');

-- Resource: datalake.storage_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_accounts"
(
    "cq_id"                                                        uuid                        NOT NULL,
    "cq_meta"                                                      jsonb,
    "cq_fetch_date"                                                timestamp without time zone NOT NULL,
    "subscription_id"                                              text,
    "identity_type"                                                text,
    "identity_principal_id"                                        uuid,
    "identity_tenant_id"                                           uuid,
    "default_group"                                                text,
    "encryption_config_type"                                       text,
    "encryption_config_key_vault_meta_info_key_vault_resource_id"  text,
    "encryption_config_key_vault_meta_info_encryption_key_name"    text,
    "encryption_config_key_vault_meta_info_encryption_key_version" text,
    "encryption_state"                                             text,
    "encryption_provisioning_state"                                text,
    "firewall_state"                                               text,
    "firewall_allow_azure_ips"                                     text,
    "trusted_id_provider_state"                                    text,
    "new_tier"                                                     text,
    "current_tier"                                                 text,
    "account_id"                                                   uuid,
    "provisioning_state"                                           text,
    "state"                                                        text,
    "creation_time"                                                timestamp without time zone,
    "last_modified_time"                                           timestamp without time zone,
    "endpoint"                                                     text,
    "id"                                                           text,
    "name"                                                         text,
    "type"                                                         text,
    "location"                                                     text,
    "tags"                                                         jsonb,
    CONSTRAINT azure_datalake_storage_accounts_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_datalake_storage_accounts');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_firewall_rules"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "start_ip_address"      text,
    "end_ip_address"        text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_firewall_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_firewall_rules (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_firewall_rules', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_virtual_network_rules"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "subnet_id"             text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_virtual_network_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_virtual_network_rules (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_virtual_network_rules', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_trusted_id_providers"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "storage_account_cq_id" uuid,
    "id_provider"           text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_trusted_id_providers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_storage_account_trusted_id_providers (cq_fetch_date, storage_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_storage_account_trusted_id_providers', 'storage_account_cq_id',
                        'azure_datalake_storage_accounts', 'cq_id');
