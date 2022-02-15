-- Autogenerated by migration tool on 2022-02-15 17:45:26
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: eventhub.namespaces
CREATE TABLE IF NOT EXISTS "azure_eventhub_namespaces" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"subscription_id" text,
	"sku_name" text,
	"sku_tier" text,
	"sku_capacity" integer,
	"identity_principal_id" text,
	"identity_tenant_id" text,
	"identity_type" text,
	"provisioning_state" text,
	"created_at_time" timestamp without time zone,
	"updated_at_time" timestamp without time zone,
	"service_bus_endpoint" text,
	"cluster_arm_id" text,
	"metric_id" text,
	"is_auto_inflate_enabled" boolean,
	"maximum_throughput_units" integer,
	"kafka_enabled" boolean,
	"zone_redundant" boolean,
	"encryption_key_source" text,
	"location" text,
	"tags" jsonb,
	"id" text,
	"name" text,
	"type" text,
	CONSTRAINT azure_eventhub_namespaces_pk PRIMARY KEY(subscription_id,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_eventhub_namespace_encryption_key_vault_properties" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"namespace_cq_id" uuid,
	"key_name" text,
	"key_vault_uri" text,
	"key_version" text,
	CONSTRAINT azure_eventhub_namespace_encryption_key_vault_properties_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (namespace_cq_id) REFERENCES azure_eventhub_namespaces(cq_id) ON DELETE CASCADE
);
