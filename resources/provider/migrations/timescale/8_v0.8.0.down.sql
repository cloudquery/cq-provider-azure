-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" DROP COLUMN IF EXISTS "network_rule_set";

-- Resource: security.assessments
DROP TABLE IF EXISTS azure_security_assessments;
