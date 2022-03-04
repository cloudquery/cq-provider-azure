-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" ADD COLUMN IF NOT EXISTS "network_rule_set" jsonb;

-- Resource: streamanalytics.jobs
CREATE TABLE IF NOT EXISTS "azure_streamanalytics_jobs" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"subscription_id" text,
	"sku_name" text,
	"job_id" text,
	"provisioning_state" text,
	"job_state" text,
	"job_type" text,
	"output_start_mode" text,
	"output_start_time" timestamp without time zone,
	"last_output_event_time" timestamp without time zone,
	"events_out_of_order_policy" text,
	"output_error_policy" text,
	"events_out_of_order_max_delay" integer,
	"events_late_arrival_max_delay" integer,
	"data_locale" text,
	"compatibility_level" text,
	"created_date" timestamp without time zone,
	"transformation_properties_streaming_units" integer,
	"transformation_properties_valid_streaming_units" integer[],
	"transformation_properties_query" text,
	"transformation_properties_etag" text,
	"transformation_id" text,
	"transformation_name" text,
	"transformation_type" text,
	"etag" text,
	"job_storage_account_authentication_mode" text,
	"job_storage_account_name" text,
	"job_storage_account_key" text,
	"content_storage_policy" text,
	"cluster_id" text,
	"identity_tenant_id" text,
	"identity_principal_id" text,
	"identity_type" text,
	"tags" jsonb,
	"location" text,
	"id" text,
	"name" text,
	"type" text,
	CONSTRAINT azure_streamanalytics_jobs_pk PRIMARY KEY(cq_fetch_date,subscription_id,id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('azure_streamanalytics_jobs');
