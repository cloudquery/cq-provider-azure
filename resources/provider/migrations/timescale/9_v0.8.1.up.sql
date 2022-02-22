-- Autogenerated by migration tool on 2022-02-16 15:48:00
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: iothub.hubs
CREATE TABLE IF NOT EXISTS "azure_iothub_hubs" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"subscription_id" text,
	"etag" text,
	"disable_local_auth" boolean,
	"disable_device_sas" boolean,
	"disable_module_sas" boolean,
	"restrict_outbound_network_access" boolean,
	"allowed_fqdn_list" text[],
	"public_network_access" text,
	"network_rule_sets_default_action" text,
	"network_rule_sets_apply_to_built_in_event_hub_endpoint" boolean,
	"min_tls_version" text,
	"provisioning_state" text,
	"state" text,
	"host_name" text,
	"event_hub_endpoints" jsonb,
	"routing_fallback_route_name" text,
	"routing_fallback_route_source" text,
	"routing_fallback_route_condition" text,
	"routing_fallback_route_endpoint_names" text[],
	"routing_fallback_route_is_enabled" boolean,
	"routing_enrichments" jsonb,
	"storage_endpoints" jsonb,
	"messaging_endpoints" jsonb,
	"enable_file_upload_notifications" boolean,
	"cloud_to_device_max_delivery_count" integer,
	"cloud_to_device_default_ttl_as_iso8601" text,
	"cloud_to_device_feedback_lock_duration_as_iso8601" text,
	"cloud_to_device_feedback_ttl_as_iso8601" text,
	"cloud_to_device_feedback_max_delivery_count" integer,
	"comments" text,
	"features" text,
	"locations" jsonb,
	"enable_data_residency" boolean,
	"sku_name" text,
	"sku_tier" text,
	"sku_capacity" bigint,
	"identity_principal_id" text,
	"identity_tenant_id" text,
	"identity_type" text,
	"identity_user_assigned_identities" jsonb,
	"system_data_created_by" text,
	"system_data_created_by_type" text,
	"system_data_created_at_time" timestamp without time zone,
	"system_data_last_modified_by" text,
	"system_data_last_modified_by_type" text,
	"system_data_last_modified_at_time" timestamp without time zone,
	"id" text,
	"name" text,
	"type" text,
	"location" text,
	"tags" jsonb,
	CONSTRAINT azure_iothub_hubs_pk PRIMARY KEY(cq_fetch_date,subscription_id,id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('azure_iothub_hubs');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_authorization_policies" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"key_name" text,
	"primary_key" text,
	"secondary_key" text,
	"rights" text,
	CONSTRAINT azure_iothub_hub_authorization_policies_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_authorization_policies (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_authorization_policies', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_ip_filter_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"filter_name" text,
	"action" text,
	"ip_mask" text,
	CONSTRAINT azure_iothub_hub_ip_filter_rules_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_ip_filter_rules (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_ip_filter_rules', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_network_rule_sets_ip_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"filter_name" text,
	"action" text,
	"ip_mask" text,
	CONSTRAINT azure_iothub_hub_network_rule_sets_ip_rules_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_network_rule_sets_ip_rules (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_network_rule_sets_ip_rules', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_private_endpoint_connections" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"id" text,
	"name" text,
	"type" text,
	"private_endpoint_id" text,
	"status" text,
	"description" text,
	"actions_required" text,
	CONSTRAINT azure_iothub_hub_private_endpoint_connections_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_private_endpoint_connections (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_private_endpoint_connections', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_routing_endpoints_service_bus_queues" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"id" text,
	"connection_string" text,
	"endpoint_uri" text,
	"entity_path" text,
	"authentication_type" text,
	"identity_user_assigned_identity" text,
	"name" text,
	"subscription_id" text,
	"resource_group" text,
	CONSTRAINT azure_iothub_hub_routing_endpoints_service_bus_queues_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_routing_endpoints_service_bus_queues (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_routing_endpoints_service_bus_queues', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_routing_endpoints_service_bus_topics" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"id" text,
	"connection_string" text,
	"endpoint_uri" text,
	"entity_path" text,
	"authentication_type" text,
	"identity_user_assigned_identity" text,
	"name" text,
	"subscription_id" text,
	"resource_group" text,
	CONSTRAINT azure_iothub_hub_routing_endpoints_service_bus_topics_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_routing_endpoints_service_bus_topics (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_routing_endpoints_service_bus_topics', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_routing_endpoints_event_hubs" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"id" text,
	"connection_string" text,
	"endpoint_uri" text,
	"entity_path" text,
	"authentication_type" text,
	"identity_user_assigned_identity" text,
	"name" text,
	"subscription_id" text,
	"resource_group" text,
	CONSTRAINT azure_iothub_hub_routing_endpoints_event_hubs_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_routing_endpoints_event_hubs (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_routing_endpoints_event_hubs', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_routing_endpoints_storage_containers" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"id" text,
	"connection_string" text,
	"endpoint_uri" text,
	"authentication_type" text,
	"identity_user_assigned_identity" text,
	"name" text,
	"subscription_id" text,
	"resource_group" text,
	"container_name" text,
	"file_name_format" text,
	"batch_frequency_in_seconds" integer,
	"max_chunk_size_in_bytes" integer,
	"encoding" text,
	CONSTRAINT azure_iothub_hub_routing_endpoints_storage_containers_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_routing_endpoints_storage_containers (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_routing_endpoints_storage_containers', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');
CREATE TABLE IF NOT EXISTS "azure_iothub_hub_routing_routes" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"hub_cq_id" uuid,
	"name" text,
	"source" text,
	"condition" text,
	"endpoint_names" text[],
	"is_enabled" boolean,
	CONSTRAINT azure_iothub_hub_routing_routes_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON azure_iothub_hub_routing_routes (cq_fetch_date, hub_cq_id);
SELECT setup_tsdb_child('azure_iothub_hub_routing_routes', 'hub_cq_id', 'azure_iothub_hubs', 'cq_id');

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
	"events_out_of_order_max_delay_in_seconds" integer,
	"events_late_arrival_max_delay_in_seconds" integer,
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
	CONSTRAINT azure_streamanalytics_jobs_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('azure_streamanalytics_jobs');
