-- Resource: sql.managed_instances
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instances"
(
    "cq_id"                        uuid                        NOT NULL,
    "cq_meta"                      jsonb,
    "cq_fetch_date"                timestamp without time zone NOT NULL,
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
    "administrator_login_password" text,
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
    CONSTRAINT azure_sql_managed_instances_pk PRIMARY KEY (cq_fetch_date, subscription_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_sql_managed_instances');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_databases"
(
    "cq_id"                                  uuid                        NOT NULL,
    "cq_meta"                                jsonb,
    "cq_fetch_date"                          timestamp without time zone NOT NULL,
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
    CONSTRAINT azure_sql_managed_databases_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('azure_sql_managed_databases');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_database_vulnerability_assessments"
(
    "cq_id"                                     uuid                        NOT NULL,
    "cq_meta"                                   jsonb,
    "cq_fetch_date"                             timestamp without time zone NOT NULL,
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
    CONSTRAINT azure_sql_managed_database_vulnerability_assessments_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_managed_database_vulnerability_assessments (cq_fetch_date, managed_database_cq_id);
SELECT setup_tsdb_child('azure_sql_managed_database_vulnerability_assessments', 'managed_database_cq_id',
                        'azure_sql_managed_databases', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_database_vulnerability_assessment_scans"
(
    "cq_id"                            uuid                        NOT NULL,
    "cq_meta"                          jsonb,
    "cq_fetch_date"                    timestamp without time zone NOT NULL,
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
    CONSTRAINT azure_sql_managed_database_vulnerability_assessment_scans_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_managed_database_vulnerability_assessment_scans (cq_fetch_date, managed_database_cq_id);
SELECT setup_tsdb_child('azure_sql_managed_database_vulnerability_assessment_scans', 'managed_database_cq_id',
                        'azure_sql_managed_databases', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_private_endpoint_connections"
(
    "cq_id"                                                  uuid                        NOT NULL,
    "cq_meta"                                                jsonb,
    "cq_fetch_date"                                          timestamp without time zone NOT NULL,
    "managed_instance_cq_id"                                 uuid,
    "id"                                                     text,
    "private_endpoint_id"                                    text,
    "private_link_service_connection_state_status"           text,
    "private_link_service_connection_state_description"      text,
    "private_link_service_connection_state_actions_required" text,
    "provisioning_state"                                     text,
    CONSTRAINT azure_sql_managed_instance_private_endpoint_connections_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_managed_instance_private_endpoint_connections (cq_fetch_date, managed_instance_cq_id);
SELECT setup_tsdb_child('azure_sql_managed_instance_private_endpoint_connections', 'managed_instance_cq_id',
                        'azure_sql_managed_instances', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_vulnerability_assessments"
(
    "cq_id"                                     uuid                        NOT NULL,
    "cq_meta"                                   jsonb,
    "cq_fetch_date"                             timestamp without time zone NOT NULL,
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
    CONSTRAINT azure_sql_managed_instance_vulnerability_assessments_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_managed_instance_vulnerability_assessments (cq_fetch_date, managed_instance_cq_id);
SELECT setup_tsdb_child('azure_sql_managed_instance_vulnerability_assessments', 'managed_instance_cq_id',
                        'azure_sql_managed_instances', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_sql_managed_instance_encryption_protectors"
(
    "cq_id"                  uuid                        NOT NULL,
    "cq_meta"                jsonb,
    "cq_fetch_date"          timestamp without time zone NOT NULL,
    "managed_instance_cq_id" uuid,
    "kind"                   text,
    "server_key_name"        text,
    "server_key_type"        text,
    "uri"                    text,
    "thumbprint"             text,
    "id"                     text,
    "name"                   text,
    "type"                   text,
    CONSTRAINT azure_sql_managed_instance_encryption_protectors_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_sql_managed_instance_encryption_protectors (cq_fetch_date, managed_instance_cq_id);
SELECT setup_tsdb_child('azure_sql_managed_instance_encryption_protectors', 'managed_instance_cq_id',
                        'azure_sql_managed_instances', 'cq_id');
