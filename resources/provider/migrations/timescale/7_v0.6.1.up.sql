-- Resource: datalake.analytics_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_accounts"
(
    "cq_id"                             uuid                        NOT NULL,
    "cq_meta"                           jsonb,
    "cq_fetch_date"                     timestamp without time zone NOT NULL,
    "subscription_id"                   text,
    "default_data_lake_store_account"   text,
    "firewall_state"                    text,
    "firewall_allow_azure_ips"          text,
    "new_tier"                          text,
    "current_tier"                      text,
    "max_job_count"                     integer,
    "system_max_job_count"              integer,
    "max_degree_of_parallelism"         integer,
    "system_max_degree_of_parallelism"  integer,
    "max_degree_of_parallelism_per_job" integer,
    "min_priority_per_job"              integer,
    "query_store_retention"             integer,
    "account_id"                        uuid,
    "provisioning_state"                text,
    "state"                             text,
    "creation_time"                     timestamp without time zone,
    "last_modified_time"                timestamp without time zone,
    "endpoint"                          text,
    "id"                                text,
    "name"                              text,
    "type"                              text,
    "location"                          text,
    "tags"                              jsonb,
    CONSTRAINT azure_datalake_analytics_accounts_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_datalake_analytics_accounts');
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_data_lake_store_accounts"
(
    "cq_id"                   uuid                        NOT NULL,
    "cq_meta"                 jsonb,
    "cq_fetch_date"           timestamp without time zone NOT NULL,
    "analytics_account_cq_id" uuid,
    "suffix"                  text,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_data_lake_store_accounts_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_analytics_account_data_lake_store_accounts (cq_fetch_date, analytics_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_analytics_account_data_lake_store_accounts', 'analytics_account_cq_id',
                        'azure_datalake_analytics_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_storage_accounts"
(
    "cq_id"                   uuid                        NOT NULL,
    "cq_meta"                 jsonb,
    "cq_fetch_date"           timestamp without time zone NOT NULL,
    "analytics_account_cq_id" uuid,
    "suffix"                  text,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_storage_accounts_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_analytics_account_storage_accounts (cq_fetch_date, analytics_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_analytics_account_storage_accounts', 'analytics_account_cq_id',
                        'azure_datalake_analytics_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_compute_policies"
(
    "cq_id"                             uuid                        NOT NULL,
    "cq_meta"                           jsonb,
    "cq_fetch_date"                     timestamp without time zone NOT NULL,
    "analytics_account_cq_id"           uuid,
    "object_id"                         uuid,
    "object_type"                       text,
    "max_degree_of_parallelism_per_job" integer,
    "min_priority_per_job"              integer,
    "id"                                text,
    "name"                              text,
    "type"                              text,
    CONSTRAINT azure_datalake_analytics_account_compute_policies_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_analytics_account_compute_policies (cq_fetch_date, analytics_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_analytics_account_compute_policies', 'analytics_account_cq_id',
                        'azure_datalake_analytics_accounts', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_firewall_rules"
(
    "cq_id"                   uuid                        NOT NULL,
    "cq_meta"                 jsonb,
    "cq_fetch_date"           timestamp without time zone NOT NULL,
    "analytics_account_cq_id" uuid,
    "start_ip_address"        inet,
    "end_ip_address"          inet,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_firewall_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_datalake_analytics_account_firewall_rules (cq_fetch_date, analytics_account_cq_id);
SELECT setup_tsdb_child('azure_datalake_analytics_account_firewall_rules', 'analytics_account_cq_id',
                        'azure_datalake_analytics_accounts', 'cq_id');

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
    "start_ip_address"      inet,
    "end_ip_address"        inet,
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
