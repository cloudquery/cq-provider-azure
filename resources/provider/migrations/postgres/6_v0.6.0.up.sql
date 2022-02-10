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

--sql.servers
CREATE TABLE IF NOT EXISTS "azure_sql_database_db_vulnerability_assessment_scans"
(
    "cq_id"                            uuid NOT NULL,
    "cq_meta"                          jsonb,
    "database_cq_id"                   uuid,
    "scan_id"                          text,
    "trigger_type"                     text,
    "state"                            text,
    "start_time"                       timestamp without time zone,
    "end_time"                         timestamp without time zone,
    "storage_container_path"           text,
    "number_of_failed_security_checks" integer,
    "id"                               text,
    "errors"                           jsonb,
    "name"                             text,
    "type"                             text,
    CONSTRAINT azure_sql_database_db_vulnerability_assessment_scans_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (database_cq_id) REFERENCES azure_sql_databases (cq_id) ON DELETE CASCADE
);
