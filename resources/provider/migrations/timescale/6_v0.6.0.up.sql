-- Autogenerated by migration tool on 2022-02-08 13:28:41
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: resources.links
CREATE TABLE IF NOT EXISTS "azure_resources_links" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"subscription_id" text,
	"id" text,
	"name" text,
	"type" text,
	"source_id" text,
	"target_id" text,
	"notes" text,
	CONSTRAINT azure_resources_links_pk PRIMARY KEY(cq_fetch_date,subscription_id,id),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('azure_resources_links');
