ALTER TABLE IF EXISTS azure_networks_virtual_network_subnets
    RENAME TO azure_network_virtual_network_subnets;
ALTER TABLE IF EXISTS azure_networks_virtual_network_peerings
    RENAME TO azure_network_virtual_network_peerings;
ALTER TABLE IF EXISTS azure_networks_virtual_network_ip_allocations
    RENAME TO azure_network_virtual_network_ip_allocations;

--make azure_network_public_ip_address_ip_tags table a json column of azure_network_public_ip_addresses
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN ip_tags JSON;

UPDATE azure_network_public_ip_addresses ips
SET ip_tags = (SELECT json_object_agg(ip_tag_type, tag)
               FROM azure_network_public_ip_address_ip_tags
               WHERE public_ip_address_cq_id = ips.cq_id);

DROP TABLE IF EXISTS azure_network_public_ip_address_ip_tags;

--ip configuration columns of azure_network_public_ip_addresses
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN private_ip_address;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN private_ip_allocation_method;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN subnet;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN public_ip_address;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN ip_configuration JSON;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN service_public_ip_address JSON;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN nat_gateway JSON;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN linked_public_ip_address JSON;



