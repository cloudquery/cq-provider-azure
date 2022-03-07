-- Autogenerated by migration tool on 2022-02-22 11:00:53

-- Resource: compute.virtual_machine_scale_sets
DROP TABLE IF EXISTS azure_compute_virtual_machine_scale_set_os_profile_secrets;
DROP TABLE IF EXISTS azure_compute_virtual_machine_scale_set_extensions;
DROP TABLE IF EXISTS azure_compute_virtual_machine_scale_sets;

-- Resource: container.registies
DROP TABLE IF EXISTS azure_container_registry_network_rule_set_virtual_network_rules;
DROP TABLE IF EXISTS azure_container_registry_network_rule_set_ip_rules;
DROP TABLE IF EXISTS azure_container_registry_replications;
DROP TABLE IF EXISTS azure_container_registries;

-- Resource: cosmosdb.accounts
DROP TABLE IF EXISTS azure_cosmosdb_account_write_locations;
DROP TABLE IF EXISTS azure_cosmosdb_account_read_locations;
DROP TABLE IF EXISTS azure_cosmosdb_account_locations;
DROP TABLE IF EXISTS azure_cosmosdb_account_failover_policies;
DROP TABLE IF EXISTS azure_cosmosdb_account_private_endpoint_connections;
DROP TABLE IF EXISTS azure_cosmosdb_account_cors;
DROP TABLE IF EXISTS azure_cosmosdb_accounts;

-- Resource: cosmosdb.mongodb_databases
DROP TABLE IF EXISTS azure_cosmosdb_mongodb_databases;

-- Resource: cosmosdb.sql_databases
DROP TABLE IF EXISTS azure_cosmosdb_sql_databases;

-- Resource: datalake.analytics_accounts
DROP TABLE IF EXISTS azure_datalake_analytics_account_data_lake_store_accounts;
DROP TABLE IF EXISTS azure_datalake_analytics_account_storage_accounts;
DROP TABLE IF EXISTS azure_datalake_analytics_account_compute_policies;
DROP TABLE IF EXISTS azure_datalake_analytics_account_firewall_rules;
DROP TABLE IF EXISTS azure_datalake_analytics_accounts;

-- Resource: datalake.storage_accounts
DROP TABLE IF EXISTS azure_datalake_storage_account_firewall_rules;
DROP TABLE IF EXISTS azure_datalake_storage_account_virtual_network_rules;
DROP TABLE IF EXISTS azure_datalake_storage_account_trusted_id_providers;
DROP TABLE IF EXISTS azure_datalake_storage_accounts;

-- Resource: eventhub.namespaces
DROP TABLE IF EXISTS azure_eventhub_namespace_encryption_key_vault_properties;
DROP TABLE IF EXISTS azure_eventhub_namespaces;

-- Resource: iothub.hubs
DROP TABLE IF EXISTS azure_iothub_hub_authorization_policies;
DROP TABLE IF EXISTS azure_iothub_hub_ip_filter_rules;
DROP TABLE IF EXISTS azure_iothub_hub_network_rule_sets_ip_rules;
DROP TABLE IF EXISTS azure_iothub_hub_private_endpoint_connections;
DROP TABLE IF EXISTS azure_iothub_hub_routing_endpoints_service_bus_queues;
DROP TABLE IF EXISTS azure_iothub_hub_routing_endpoints_service_bus_topics;
DROP TABLE IF EXISTS azure_iothub_hub_routing_endpoints_event_hubs;
DROP TABLE IF EXISTS azure_iothub_hub_routing_endpoints_storage_containers;
DROP TABLE IF EXISTS azure_iothub_hub_routing_routes;
DROP TABLE IF EXISTS azure_iothub_hubs;

-- Resource: mariadb.servers
DROP TABLE IF EXISTS azure_mariadb_server_private_endpoint_connections;
DROP TABLE IF EXISTS azure_mariadb_server_configurations;
DROP TABLE IF EXISTS azure_mariadb_servers;

-- Resource: network.interfaces
DROP TABLE IF EXISTS azure_network_interface_ip_configurations;
DROP TABLE IF EXISTS azure_network_interfaces;

-- Resource: servicebus.namespaces
DROP TABLE IF EXISTS azure_servicebus_namespace_private_endpoint_connections;
DROP TABLE IF EXISTS azure_servicebus_namespaces;

-- Resource: sql.managed_instances
DROP TABLE IF EXISTS azure_sql_managed_database_vulnerability_assessments;
DROP TABLE IF EXISTS azure_sql_managed_database_vulnerability_assessment_scans;
DROP TABLE IF EXISTS azure_sql_managed_databases;
DROP TABLE IF EXISTS azure_sql_managed_instance_private_endpoint_connections;
DROP TABLE IF EXISTS azure_sql_managed_instance_vulnerability_assessments;
DROP TABLE IF EXISTS azure_sql_managed_instance_encryption_protectors;
DROP TABLE IF EXISTS azure_sql_managed_instances;

-- Resource: sql.servers
DROP TABLE IF EXISTS azure_sql_database_db_vulnerability_assessment_scans;
