-- Autogenerated by migration tool on 2022-02-08 00:49:00
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: network.virtual_networks
ALTER TABLE IF EXISTS "azure_network_virtual_network_subnets" DROP COLUMN IF EXISTS "ip_configurations";
ALTER TABLE IF EXISTS "azure_network_virtual_network_subnets" DROP COLUMN IF EXISTS "private_endpoints";
