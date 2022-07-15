INSERT INTO azure_policy_results (execution_time, framework, check_id, title, subscription_id, resource_id, status)
SELECT :'execution_time'                                                          AS execution_time,
       :'framework'                                                               AS framework,
       :'check_id'                                                                AS check_id,
       'Ensure web app is using the latest version of TLS encryption (Automated)' AS title,
       awa.subscription_id                                                        AS subscription_id,
       awa.id                                                                     AS resource_id,
       CASE
           WHEN site_config -> 'minTlsVersion' IS NULL OR site_config -> 'minTlsVersion' != '1.2'
               THEN 'fail'
           ELSE 'pass'
           END                                                                    AS status
FROM azure_web_apps
