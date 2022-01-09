# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- 
## Unreleased
### Added
### Changed
### Fixed
### Breaking Changes

-->

## [v0.3.9] - 2022-01-03
###### SDK Version: 0.6.1
### 🚀 Added
* Add `keyvault.vault` resource back (this result requires a special permissions) [#111](https://github.com/cloudquery/cq-provider-azure/pull/111)

## [v0.3.9] - 2022-01-03
###### SDK Version: 0.6.1
### :gear: Changed
* Update to SDK Version [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v061---2021-01-03)
* Remove ad resources (deprecated and will be migrate to msgraph)
### 💥 Breaking Changes
* Renamed table `azure_container_managed_cluster_pip_user_assigned_identity_exceptions` -> `azure_container_managed_cluster_pip_user_assigned_id_exceptions` [#97](https://github.com/cloudquery/cq-provider-azure/pull/97)
### :spider: Fixed
* Fixed disabled migrations [#104](https://github.com/cloudquery/cq-provider-azure/pull/104)

## [v0.3.8] - 2021-11-23
###### SDK Version: 0.4.9

### :rocket: Added
* Added support for ARM binary fixed [#92](https://github.com/cloudquery/cq-provider-azure/pull/92)

### :spider: Fixed
* Fixed names of `azure_network_virtual_network_subnets`, `azure_network_virtual_network_peerings`, `azure_network_virtual_network_ip_allocations` tables according to naming convention [#76](https://github.com/cloudquery/cq-provider-azure/issues/76)


## [v0.3.7] - 2021-10-07
###### SDK Version: 0.4.9

### :gear: Changed
* Upgraded to SDK Version [v0.4.9](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

## [v0.3.6] - 2021-10-03
###### SDK Version: 0.4.7

Base version at which changelog was introduced.

### Supported Resources
- ad.applications
- ad.groups
- ad.service_principals
- ad.users
- authorization.role_assignments
- authorization.role_definitions
- compute.disks
- compute.virtual_machines
- container.managed_clusters
- keyvault.vaults
- monitor.activity_log_alerts
- monitor.activity_logs
- monitor.diagnostic_settings
- monitor.log_profiles
- mysql.servers
- network.public_ip_addresses
- network.security_groups
- network.virtual_networks
- network.watchers
- postgresql.servers
- resources.groups
- resources.policy_assignments
- security.auto_provisioning_settings
- security.contacts
- security.pricings
- security.settings
- sql.servers
- storage.accounts
- subscription.subscriptions
- web.apps
