-- Resource: eventhub.namespaces
ALTER TABLE IF EXISTS "azure_eventhub_namespaces" DROP COLUMN IF EXISTS "network_rule_set";

-- Resource: sql.servers
ALTER TABLE IF EXISTS "azure_sql_databases" DROP COLUMN IF EXISTS "backup_long_term_retention_policy";
