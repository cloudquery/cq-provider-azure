CREATE TABLE public.azure_network_virtual_network_ip_allocations
(
    cq_id                 uuid NULL,
    meta                  jsonb NULL,
    virtual_network_cq_id uuid NOT NULL,
    id                    text NOT NULL,
    CONSTRAINT azure_network_virtual_network_ip_allocations_cq_id_key UNIQUE (cq_id),
    CONSTRAINT azure_network_virtual_network_ip_allocations_pk PRIMARY KEY (virtual_network_cq_id, id)
);

-- public.azure_network_virtual_network_ip_allocations foreign keys

ALTER TABLE public.azure_network_virtual_network_ip_allocations
    ADD CONSTRAINT azure_network_virtual_network_ip_all_virtual_network_cq_id_fkey FOREIGN KEY (virtual_network_cq_id) REFERENCES public.azure_network_virtual_networks (cq_id) ON DELETE CASCADE;

INSERT INTO azure_network_virtual_network_ip_allocations(cq_id, virtual_network_cq_id, id)
SELECT gen_random_uuid(),
       cq_id,
       unnest(ip_allocations)
FROM azure_network_virtual_networks;

alter table if exists azure_network_virtual_networks drop column ip_allocations;

--todo drop ip_allocation column


ALTER TABLE IF EXISTS azure_network_virtual_network_subnets
    RENAME TO azure_networks_virtual_network_subnets;
ALTER TABLE IF EXISTS azure_network_virtual_network_peerings
    RENAME TO azure_networks_virtual_network_peerings;
ALTER TABLE IF EXISTS azure_network_virtual_network_ip_allocations
    RENAME TO azure_networks_virtual_network_ip_allocations;

