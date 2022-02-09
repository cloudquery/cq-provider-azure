-- Autogenerated by migration tool on 2022-02-09 11:25:41
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: keyvault.managed_hsm
CREATE TABLE IF NOT EXISTS "azure_keyvault_managed_hsm" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"subscription_id" text,
	"tenant_id" uuid,
	"initial_admin_object_ids" text[],
	"hsm_uri" text,
	"enable_soft_delete" boolean,
	"soft_delete_retention_in_days" integer,
	"enable_purge_protection" boolean,
	"create_mode" text,
	"status_message" text,
	"provisioning_state" text,
	"id" text,
	"name" text,
	"type" text,
	"location" text,
	"sku_family" text,
	"sku_name" text,
	"tags" jsonb,
	CONSTRAINT azure_keyvault_managed_hsm_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
