\set ON_ERROR_STOP on
\set framework 'hipaa_hitrust_v9.2'
\set execution_time `date '+%Y-%m-%d %H:%M:%S'`
\echo "Creating azure_policy_results table if not exist"
\i sql/create_azure_policy_results.sql
-- \i sql/hipaa_hitrust_v9.2/privilege_management.sql
\i sql/hipaa_hitrust_v9.2/user_authentication_for_external_connections.sql
\i sql/hipaa_hitrust_v9.2/segregation_in_networks.sql
\i sql/hipaa_hitrust_v9.2/network_connection_control.sql
\i sql/hipaa_hitrust_v9.2/user_identification_and_authentication.sql
\i sql/hipaa_hitrust_v9.2/identification_of_risks_related_to_external_parties.sql
\i sql/hipaa_hitrust_v9.2/audit_logging.sql
\i sql/hipaa_hitrust_v9.2/monitoring_system_use.sql
\i sql/hipaa_hitrust_v9.2/administrator_and_operator_logs.sql
\i sql/hipaa_hitrust_v9.2/segregation_of_duties.sql
\i sql/hipaa_hitrust_v9.2/controls_against_malicious_code.sql
\i sql/hipaa_hitrust_v9.2/back_up.sql
\i sql/hipaa_hitrust_v9.2/network_controls.sql
\i sql/hipaa_hitrust_v9.2/security_of_network_services.sql
\i sql/hipaa_hitrust_v9.2/management_of_removable_media.sql
\i sql/hipaa_hitrust_v9.2/information_exchange_policies_and_procedures.sql
\i sql/hipaa_hitrust_v9.2/on_line_transactions.sql
-- \i sql/hipaa_hitrust_v9.2/control_of_operational_software.sql
-- \i sql/hipaa_hitrust_v9.2/change_control_procedures.sql
\i sql/hipaa_hitrust_v9.2/control_of_technical_vulnerabilities.sql
\i sql/hipaa_hitrust_v9.2/business_continuity_and_risk_assessment.sql