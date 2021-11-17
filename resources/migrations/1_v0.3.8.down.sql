ALTER TABLE IF EXISTS azure_network_virtual_network_subnets
    RENAME TO azure_networks_virtual_network_subnets;
ALTER TABLE IF EXISTS azure_network_virtual_network_peerings
    RENAME TO azure_networks_virtual_network_peerings;
ALTER TABLE IF EXISTS azure_network_virtual_network_ip_allocations
    RENAME TO azure_networks_virtual_network_ip_allocations;

--make azure_network_public_ip_address_ip_tags table a json column of azure_network_public_ip_addresses
CREATE TABLE IF NOT EXISTS public.azure_network_public_ip_address_ip_tags
(
    cq_id
    uuid
    NOT
    NULL,
    meta
    jsonb
    NULL,
    public_ip_address_cq_id
    uuid
    NULL,
    ip_tag_type
    text
    NULL,
    tag
    text
    NULL,
    CONSTRAINT
    azure_network_public_ip_address_ip_tags_pk
    null
);
ALTER TABLE public.azure_network_public_ip_address_ip_tags
    ADD CONSTRAINT azure_network_public_ip_address_ip_public_ip_address_cq_id_fkey FOREIGN KEY (public_ip_address_cq_id) REFERENCES public.azure_network_public_ip_addresses (cq_id) ON DELETE CASCADE;


INSERT INTO azure_network_public_ip_address_ip_tags(cq_id, public_ip_address_cq_id, ip_tag_type, tag)
SELECT gen_random_uuid(),
       cq_id,
       json_data.key,
       json_data.value
FROM azure_network_public_ip_addresses,
     json_each_text(ip_tags) AS json_data

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN ip_tags;

--ip configuration columns of azure_network_public_ip_addresses
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD
    COLUMN private_ip_address TEXT;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD
    COLUMN private_ip_allocation_method TEXT;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD
    COLUMN subnet JSON;
ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD
    COLUMN public_ip_address JSON;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN ip_configuration;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN service_public_ip_address;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN nat_gateway;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
DROP
COLUMN linked_public_ip_address;

