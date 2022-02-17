-- Autogenerated by migration tool on 2022-02-16 09:54:57
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: cosmosdb.accounts
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_accounts"
(
    "cq_id"                                        uuid NOT NULL,
    "cq_meta"                                      jsonb,
    "subscription_id"                              text,
    "provisioning_state"                           text,
    "document_endpoint"                            text,
    "database_account_offer_type"                  text,
    "ip_rules"                                     text[],
    "capabilities"                                 text[],
    "is_virtual_network_filter_enabled"            boolean,
    "enable_automatic_failover"                    boolean,
    "consistency_policy_default_consistency_level" text,
    "consistency_policy_max_staleness_prefix"      bigint,
    "consistency_policy_max_interval_in_seconds"   integer,
    "virtual_network_rules"                        jsonb,
    "enable_multiple_write_locations"              boolean,
    "enable_cassandra_connector"                   boolean,
    "connector_offer"                              text,
    "disable_key_based_metadata_write_access"      boolean,
    "key_vault_key_uri"                            text,
    "public_network_access"                        text,
    "enable_free_tier"                             boolean,
    "api_properties_server_version"                text,
    "enable_analytical_storage"                    boolean,
    "id"                                           text,
    "name"                                         text,
    "type"                                         text,
    "location"                                     text,
    "tags"                                         jsonb,
    CONSTRAINT azure_cosmosdb_accounts_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_write_locations"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "account_cq_id"      uuid,
    "id"                 text,
    "location_name"      text,
    "document_endpoint"  text,
    "provisioning_state" text,
    "failover_priority"  integer,
    "is_zone_redundant"  boolean,
    CONSTRAINT azure_cosmosdb_account_write_locations_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_read_locations"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "account_cq_id"      uuid,
    "id"                 text,
    "location_name"      text,
    "document_endpoint"  text,
    "provisioning_state" text,
    "failover_priority"  integer,
    "is_zone_redundant"  boolean,
    CONSTRAINT azure_cosmosdb_account_read_locations_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_locations"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "account_cq_id"      uuid,
    "id"                 text,
    "location_name"      text,
    "document_endpoint"  text,
    "provisioning_state" text,
    "failover_priority"  integer,
    "is_zone_redundant"  boolean,
    CONSTRAINT azure_cosmosdb_account_locations_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_failover_policies"
(
    "cq_id"             uuid NOT NULL,
    "cq_meta"           jsonb,
    "account_cq_id"     uuid,
    "id"                text,
    "location_name"     text,
    "failover_priority" integer,
    CONSTRAINT azure_cosmosdb_account_failover_policies_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_private_endpoint_connections"
(
    "cq_id"               uuid NOT NULL,
    "cq_meta"             jsonb,
    "account_cq_id"       uuid,
    "private_endpoint_id" text,
    "status"              text,
    "actions_required"    text,
    "description"         text,
    "group_id"            text,
    "provisioning_state"  text,
    "id"                  text,
    "name"                text,
    "type"                text,
    CONSTRAINT azure_cosmosdb_account_private_endpoint_connections_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_account_cors"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "account_cq_id"      uuid,
    "allowed_origins"    text,
    "allowed_methods"    text,
    "allowed_headers"    text,
    "exposed_headers"    text,
    "max_age_in_seconds" bigint,
    CONSTRAINT azure_cosmosdb_account_cors_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (account_cq_id) REFERENCES azure_cosmosdb_accounts (cq_id) ON DELETE CASCADE
);

-- Resource: cosmosdb.mongodb_databases
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_mongodb_databases"
(
    "cq_id"                             uuid NOT NULL,
    "cq_meta"                           jsonb,
    "subscription_id"                   text,
    "database_id"                       text,
    "database_rid"                      text,
    "database_ts"                       float,
    "database_etag"                     text,
    "throughput"                        integer,
    "autoscale_settings_max_throughput" integer,
    "id"                                text,
    "name"                              text,
    "type"                              text,
    "location"                          text,
    "tags"                              jsonb,
    CONSTRAINT azure_cosmosdb_mongodb_databases_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);

-- Resource: cosmosdb.sql_databases
CREATE TABLE IF NOT EXISTS "azure_cosmosdb_sql_databases"
(
    "cq_id"                                  uuid NOT NULL,
    "cq_meta"                                jsonb,
    "subscription_id"                        text,
    "database_id"                            text,
    "database_rid"                           text,
    "database_ts"                            float,
    "database_etag"                          text,
    "database_colls"                         text,
    "database_users"                         text,
    "sql_database_get_properties_throughput" integer,
    "autoscale_settings_max_throughput"      integer,
    "id"                                     text,
    "name"                                   text,
    "type"                                   text,
    "location"                               text,
    "tags"                                   jsonb,
    CONSTRAINT azure_cosmosdb_sql_databases_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);

-- Resource: eventhub.namespaces
CREATE TABLE IF NOT EXISTS "azure_eventhub_namespaces"
(
    "cq_id"                    uuid NOT NULL,
    "cq_meta"                  jsonb,
    "subscription_id"          text,
    "sku_name"                 text,
    "sku_tier"                 text,
    "sku_capacity"             integer,
    "identity_principal_id"    text,
    "identity_tenant_id"       text,
    "identity_type"            text,
    "provisioning_state"       text,
    "created_at_time"          timestamp without time zone,
    "updated_at_time"          timestamp without time zone,
    "service_bus_endpoint"     text,
    "cluster_arm_id"           text,
    "metric_id"                text,
    "is_auto_inflate_enabled"  boolean,
    "maximum_throughput_units" integer,
    "kafka_enabled"            boolean,
    "zone_redundant"           boolean,
    "encryption_key_source"    text,
    "location"                 text,
    "tags"                     jsonb,
    "id"                       text,
    "name"                     text,
    "type"                     text,
    CONSTRAINT azure_eventhub_namespaces_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_eventhub_namespace_encryption_key_vault_properties"
(
    "cq_id"           uuid NOT NULL,
    "cq_meta"         jsonb,
    "namespace_cq_id" uuid,
    "key_name"        text,
    "key_vault_uri"   text,
    "key_version"     text,
    CONSTRAINT azure_eventhub_namespace_encryption_key_vault_properties_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (namespace_cq_id) REFERENCES azure_eventhub_namespaces (cq_id) ON DELETE CASCADE
);
-- Resource: sql.managed_instances
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instances"
(
    "cq_id"                        uuid NOT NULL,
    "cq_meta"                      jsonb,
    "subscription_id"              text,
    "identity_principal_id"        uuid,
    "identity_type"                text,
    "identity_tenant_id"           uuid,
    "sku_name"                     text,
    "sku_tier"                     text,
    "sku_size"                     text,
    "sku_family"                   text,
    "sku_capacity"                 integer,
    "provisioning_state"           text,
    "managed_instance_create_mode" text,
    "fully_qualified_domain_name"  text,
    "administrator_login"          text,
    "subnet_id"                    text,
    "state"                        text,
    "license_type"                 text,
    "v_cores"                      integer,
    "storage_size_in_gb"           integer,
    "collation"                    text,
    "dns_zone"                     text,
    "dns_zone_partner"             text,
    "public_data_endpoint_enabled" boolean,
    "source_managed_instance_id"   text,
    "restore_point_in_time"        timestamp without time zone,
    "proxy_override"               text,
    "timezone_id"                  text,
    "instance_pool_id"             text,
    "maintenance_configuration_id" text,
    "minimal_tls_version"          text,
    "storage_account_type"         text,
    "zone_redundant"               boolean,
    "location"                     text,
    "tags"                         jsonb,
    "id"                           text,
    "name"                         text,
    "type"                         text,
    CONSTRAINT azure_sql_managed_instances_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_databases"
(
    "cq_id"                                  uuid NOT NULL,
    "cq_meta"                                jsonb,
    "collation"                              text,
    "status"                                 text,
    "creation_date_time"                     timestamp without time zone,
    "earliest_restore_point_time"            timestamp without time zone,
    "restore_point_in_time"                  timestamp without time zone,
    "default_secondary_location"             text,
    "catalog_collation"                      text,
    "create_mode"                            text,
    "storage_container_uri"                  text,
    "source_database_id"                     text,
    "restorable_dropped_database_id"         text,
    "storage_container_sas_token"            text,
    "failover_group_id"                      text,
    "recoverable_database_id"                text,
    "long_term_retention_backup_resource_id" text,
    "auto_complete_restore"                  boolean,
    "last_backup_name"                       text,
    "location"                               text,
    "tags"                                   jsonb,
    "id"                                     text,
    "name"                                   text,
    "type"                                   text,
    CONSTRAINT azure_sql_managed_databases_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_database_vulnerability_assessments"
(
    "cq_id"                                     uuid NOT NULL,
    "cq_meta"                                   jsonb,
    "managed_database_cq_id"                    uuid,
    "storage_container_path"                    text,
    "storage_container_sas_key"                 text,
    "storage_account_access_key"                text,
    "recurring_scans_is_enabled"                boolean,
    "recurring_scans_email_subscription_admins" boolean,
    "recurring_scans_emails"                    text[],
    "id"                                        text,
    "name"                                      text,
    "type"                                      text,
    CONSTRAINT azure_sql_managed_database_vulnerability_assessments_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (managed_database_cq_id) REFERENCES azure_sql_managed_databases (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_database_vulnerability_assessment_scans"
(
    "cq_id"                            uuid NOT NULL,
    "cq_meta"                          jsonb,
    "managed_database_cq_id"           uuid,
    "scan_id"                          text,
    "trigger_type"                     text,
    "state"                            text,
    "start_time"                       timestamp without time zone,
    "end_time"                         timestamp without time zone,
    "errors"                           jsonb,
    "storage_container_path"           text,
    "number_of_failed_security_checks" integer,
    "id"                               text,
    "name"                             text,
    "type"                             text,
    CONSTRAINT azure_sql_managed_database_vulnerability_assessment_scans_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (managed_database_cq_id) REFERENCES azure_sql_managed_databases (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_private_endpoint_connections"
(
    "cq_id"                                                  uuid NOT NULL,
    "cq_meta"                                                jsonb,
    "managed_instance_cq_id"                                 uuid,
    "id"                                                     text,
    "private_endpoint_id"                                    text,
    "private_link_service_connection_state_status"           text,
    "private_link_service_connection_state_description"      text,
    "private_link_service_connection_state_actions_required" text,
    "provisioning_state"                                     text,
    CONSTRAINT azure_sql_managed_instance_private_endpoint_connections_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (managed_instance_cq_id) REFERENCES azure_sql_managed_instances (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_vulnerability_assessments"
(
    "cq_id"                                     uuid NOT NULL,
    "cq_meta"                                   jsonb,
    "managed_instance_cq_id"                    uuid,
    "storage_container_path"                    text,
    "storage_container_sas_key"                 text,
    "storage_account_access_key"                text,
    "recurring_scans_is_enabled"                boolean,
    "recurring_scans_email_subscription_admins" boolean,
    "recurring_scans_emails"                    text[],
    "id"                                        text,
    "name"                                      text,
    "type"                                      text,
    CONSTRAINT azure_sql_managed_instance_vulnerability_assessments_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (managed_instance_cq_id) REFERENCES azure_sql_managed_instances (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_encryption_protectors"
(
    "cq_id"                  uuid NOT NULL,
    "cq_meta"                jsonb,
    "managed_instance_cq_id" uuid,
    "kind"                   text,
    "server_key_name"        text,
    "server_key_type"        text,
    "uri"                    text,
    "thumbprint"             text,
    "id"                     text,
    "name"                   text,
    "type"                   text,
    CONSTRAINT azure_sql_managed_instance_encryption_protectors_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (managed_instance_cq_id) REFERENCES azure_sql_managed_instances (cq_id) ON DELETE CASCADE
);

-- Resource: datalake.analytics_accounts
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_accounts"
(
    "cq_id"                             uuid NOT NULL,
    "cq_meta"                           jsonb,
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
    CONSTRAINT azure_datalake_analytics_accounts_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_data_lake_store_accounts"
(
    "cq_id"                   uuid NOT NULL,
    "cq_meta"                 jsonb,
    "analytics_account_cq_id" uuid,
    "suffix"                  text,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_data_lake_store_accounts_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analytics_account_cq_id) REFERENCES azure_datalake_analytics_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_storage_accounts"
(
    "cq_id"                   uuid NOT NULL,
    "cq_meta"                 jsonb,
    "analytics_account_cq_id" uuid,
    "suffix"                  text,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_storage_accounts_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analytics_account_cq_id) REFERENCES azure_datalake_analytics_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_compute_policies"
(
    "cq_id"                             uuid NOT NULL,
    "cq_meta"                           jsonb,
    "analytics_account_cq_id"           uuid,
    "object_id"                         uuid,
    "object_type"                       text,
    "max_degree_of_parallelism_per_job" integer,
    "min_priority_per_job"              integer,
    "id"                                text,
    "name"                              text,
    "type"                              text,
    CONSTRAINT azure_datalake_analytics_account_compute_policies_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analytics_account_cq_id) REFERENCES azure_datalake_analytics_accounts (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_datalake_analytics_account_firewall_rules"
(
    "cq_id"                   uuid NOT NULL,
    "cq_meta"                 jsonb,
    "analytics_account_cq_id" uuid,
    "start_ip_address"        inet,
    "end_ip_address"          inet,
    "id"                      text,
    "name"                    text,
    "type"                    text,
    CONSTRAINT azure_datalake_analytics_account_firewall_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analytics_account_cq_id) REFERENCES azure_datalake_analytics_accounts (cq_id) ON DELETE CASCADE
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
    "start_ip_address"      inet,
    "end_ip_address"        inet,
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


-- Resource: compute.virtual_machine_scale_sets
CREATE TABLE IF NOT EXISTS "azure_compute_virtual_machine_scale_sets"
(
    "cq_id"                                         uuid NOT NULL,
    "cq_meta"                                       jsonb,
    "subscription_id"                               text,
    "sku_name"                                      text,
    "sku_tier"                                      text,
    "sku_capacity"                                  bigint,
    "plan_name"                                     text,
    "plan_publisher"                                text,
    "plan_product"                                  text,
    "plan_promotion_code"                           text,
    "upgrade_policy"                                jsonb,
    "automatic_repairs_policy_enabled"              boolean,
    "automatic_repairs_policy_grace_period"         text,
    "os_profile_computer_name_prefix"               text,
    "os_profile_admin_username"                     text,
    "os_profile_admin_password"                     text,
    "os_profile_custom_data"                        text,
    "os_profile_windows_configuration"              jsonb,
    "os_profile_linux_configuration"                jsonb,
    "storage_profile"                               jsonb,
    "network_profile"                               jsonb,
    "security_profile"                              jsonb,
    "diagnostics_profile"                           jsonb,
    "extension_profile_extensions_time_budget"      text,
    "license_type"                                  text,
    "priority"                                      text,
    "eviction_policy"                               text,
    "billing_profile_max_price"                     float,
    "scheduled_events_profile"                      jsonb,
    "user_data"                                     text,
    "provisioning_state"                            text,
    "overprovision"                                 boolean,
    "do_not_run_extensions_on_overprovisioned_vms" boolean,
    "unique_id"                                     text,
    "single_placement_group"                        boolean,
    "zone_balance"                                  boolean,
    "platform_fault_domain_count"                   integer,
    "proximity_placement_group_id"                  text,
    "host_group_id"                                 text,
    "additional_capabilities_ultra_ssd_enabled"   boolean,
    "scale_in_policy_rules"                         text[],
    "orchestration_mode"                            text,
    "identity_principal_id"                         text,
    "identity_tenant_id"                            text,
    "identity_type"                                 text,
    "identity_user_assigned_identities"             jsonb,
    "zones"                                         text[],
    "extended_location_name"                        text,
    "extended_location_type"                        text,
    "id"                                            text,
    "name"                                          text,
    "type"                                          text,
    "location"                                      text,
    "tags"                                          jsonb,
    CONSTRAINT azure_compute_virtual_machine_scale_sets_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_compute_virtual_machine_scale_set_os_profile_secrets"
(
    "cq_id"                           uuid NOT NULL,
    "cq_meta"                         jsonb,
    "virtual_machine_scale_set_cq_id" uuid,
    "source_vault_id"                 text,
    "vault_certificates"              jsonb,
    CONSTRAINT azure_compute_virtual_machine_scale_set_os_profile_secrets_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (virtual_machine_scale_set_cq_id) REFERENCES azure_compute_virtual_machine_scale_sets (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_compute_virtual_machine_scale_set_extensions"
(
    "cq_id"                           uuid NOT NULL,
    "cq_meta"                         jsonb,
    "virtual_machine_scale_set_cq_id" uuid,
    "type"                            text,
    "extension_type"                  text,
    "name"                            text,
    "force_update_tag"                text,
    "publisher"                       text,
    "type_handler_version"            text,
    "auto_upgrade_minor_version"      boolean,
    "enable_automatic_upgrade"        boolean,
    "settings"                        jsonb,
    "protected_settings"              jsonb,
    "provisioning_state"              text,
    "provision_after_extensions"      text[],
    "id"                              text,
    CONSTRAINT azure_compute_virtual_machine_scale_set_extensions_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (virtual_machine_scale_set_cq_id) REFERENCES azure_compute_virtual_machine_scale_sets (cq_id) ON DELETE CASCADE
);
