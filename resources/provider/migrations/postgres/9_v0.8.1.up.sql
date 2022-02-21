-- Autogenerated by migration tool on 2022-02-21 18:25:49
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: batch.accounts
CREATE TABLE IF NOT EXISTS "azure_batch_accounts" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"subscription_id" text,
	"account_endpoint" text,
	"provisioning_state" text,
	"pool_allocation_mode" text,
	"key_vault_reference_id" text,
	"key_vault_reference_url" text,
	"public_network_access" text,
	"auto_storage_last_key_sync_time" timestamp without time zone,
	"auto_storage_storage_account_id" text,
	"auto_storage_authentication_mode" text,
	"auto_storage_node_identity_reference_resource_id" text,
	"encryption_key_source" text,
	"encryption_key_vault_properties_key_identifier" text,
	"dedicated_core_quota" integer,
	"low_priority_core_quota" integer,
	"dedicated_core_quota_per_vm_family" jsonb,
	"dedicated_core_quota_per_vm_family_enforced" boolean,
	"pool_quota" integer,
	"active_job_and_job_schedule_quota" integer,
	"allowed_authentication_modes" text[],
	"identity_principal_id" text,
	"identity_tenant_id" text,
	"identity_type" text,
	"identity_user_assigned_identities" jsonb,
	"id" text,
	"name" text,
	"type" text,
	"location" text,
	"tags" jsonb,
	CONSTRAINT azure_batch_accounts_pk PRIMARY KEY(subscription_id,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_batch_account_private_endpoint_connections" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_cq_id" uuid,
	"provisioning_state" text,
	"private_endpoint_id" text,
	"private_link_service_connection_state_status" text,
	"private_link_service_connection_state_description" text,
	"private_link_service_connection_state_action_required" text,
	"id" text,
	"name" text,
	"type" text,
	"etag" text,
	CONSTRAINT azure_batch_account_private_endpoint_connections_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (account_cq_id) REFERENCES azure_batch_accounts(cq_id) ON DELETE CASCADE
);
