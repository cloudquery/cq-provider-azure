\echo "Segregation of Duties"
\i sql/views/nsg_rules_ports.sql
\set check_id "1229.09c1Organizational.1 - 09.c - 1"
\echo :check_id
\i sql/queries/container/aks_rbac_disabled.sql
\set check_id "1230.09c2Organizational.1 - 09.c - 1"
\echo :check_id
\i sql/queries/authorization/custom_roles.sql
\set check_id "1232.09c3Organizational.12 - 09.c - 1"
\echo :check_id
\i sql/queries/network/rdp_access_permitted.sql

