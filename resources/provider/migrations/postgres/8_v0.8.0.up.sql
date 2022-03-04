-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" ADD COLUMN IF NOT EXISTS "network_rule_set" jsonb;

-- Resource: security.assessments
CREATE TABLE IF NOT EXISTS "azure_security_assessments" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"subscription_id" text,
	"display_name" text,
	"code" text,
	"cause" text,
	"description" text,
	"additional_data" jsonb,
	"azure_portal_uri" text,
	"metadata_display_name" text,
	"metadata_policy_definition_id" text,
	"metadata_description" text,
	"metadata_remediation_description" text,
	"metadata_categories" text[],
	"metadata_severity" text,
	"metadata_user_impact" text,
	"metadata_implementation_effort" text,
	"metadata_threats" text[],
	"metadata_preview" boolean,
	"metadata_assessment_type" text,
	"metadata_partner_data_partner_name" text,
	"metadata_partner_data_product_name" text,
	"partner_name" text,
	"id" text,
	"name" text,
	"type" text,
	"resource_details" jsonb,
	CONSTRAINT azure_security_assessments_pk PRIMARY KEY(subscription_id,id),
	UNIQUE(cq_id)
);
