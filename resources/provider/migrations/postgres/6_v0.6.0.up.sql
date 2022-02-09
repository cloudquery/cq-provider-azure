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
DROP TABLE IF EXISTS "azure_compute_virtual_machine_network_interfaces"

-- Resource: security.jit_network_access_policies
CREATE TABLE IF NOT EXISTS "azure_security_jit_network_access_policies"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "subscription_id"    text,
    "id"                 text,
    "name"               text,
    "type"               text,
    "kind"               text,
    "location"           text,
    "provisioning_state" text,
    CONSTRAINT azure_security_jit_network_access_policies_pk PRIMARY KEY (subscription_id, id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_security_jit_network_access_policy_virtual_machines"
(
    "cq_id"                           uuid NOT NULL,
    "cq_meta"                         jsonb,
    "jit_network_access_policy_cq_id" uuid,
    "id"                              text,
    "ports"                           jsonb,
    "public_ip_address"               inet,
    CONSTRAINT azure_security_jit_network_access_policy_virtual_machines_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (jit_network_access_policy_cq_id) REFERENCES azure_security_jit_network_access_policies (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "azure_security_jit_network_access_policy_requests"
(
    "cq_id"                           uuid NOT NULL,
    "cq_meta"                         jsonb,
    "jit_network_access_policy_cq_id" uuid,
    "virtual_machines"                text[],
    "start_time_utc"                  timestamp without time zone,
    "requestor"                       text,
    "justification"                   text,
    CONSTRAINT azure_security_jit_network_access_policy_requests_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (jit_network_access_policy_cq_id) REFERENCES azure_security_jit_network_access_policies (cq_id) ON DELETE CASCADE
);
