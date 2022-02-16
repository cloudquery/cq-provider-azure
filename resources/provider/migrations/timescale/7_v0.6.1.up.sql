
-- Resource: datalake.analytics_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_accounts"
(
    "cq_id"                                                        uuid NOT NULL,
    "cq_meta"                                                      jsonb,
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
    CONSTRAINT azure_datalake_storage_accounts_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_firewall_rules"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "start_ip_address"      text,
    "end_ip_address"        text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_firewall_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_virtual_network_rules"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "subnet_id"             text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_virtual_network_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_trusted_id_providers"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "id_provider"           text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_trusted_id_providers_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);

-- Resource: datalake.storage_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_accounts"
(
    "cq_id"                                                        uuid NOT NULL,
    "cq_meta"                                                      jsonb,
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
    CONSTRAINT azure_datalake_storage_accounts_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_firewall_rules"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "start_ip_address"      text,
    "end_ip_address"        text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_firewall_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_virtual_network_rules"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "subnet_id"             text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_virtual_network_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_storage_account_trusted_id_providers"
(
    "cq_id"                 uuid NOT NULL,
    "cq_meta"               jsonb,
    "storage_account_cq_id" uuid,
    "id_provider"           text,
    "id"                    text,
    "name"                  text,
    "type"                  text,
    CONSTRAINT azure_datalake_storage_account_trusted_id_providers_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (storage_account_cq_id) REFERENCES azure_datalake_storage_accounts (cq_id) ON DELETE CASCADE
);