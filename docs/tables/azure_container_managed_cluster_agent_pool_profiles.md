
# Table: azure_container_managed_cluster_agent_pool_profiles
ManagedClusterAgentPoolProfile profile for the container service agent pool
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_cluster_cq_id|uuid|Unique ID of azure_container_managed_clusters table (FK)|
|name|text|Unique name of the agent pool profile in the context of the subscription and resource group|
|count|integer|Number of agents (VMs) to host docker containers Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools The default value is 1|
|vm_size|text|Size of agent VMs|
|os_disk_size_gb|integer|OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool If you specify 0, it will apply the default osDisk size according to the vmSize specified|
|os_disk_type|text|OS disk type to be used for machines in a given agent pool Allowed values are 'Ephemeral' and 'Managed' If unspecified, defaults to 'Ephemeral' when the VM supports ephemeral OS and has a cache disk larger than the requested OSDiskSizeGB Otherwise, defaults to 'Managed' May not be changed after creation Possible values include: 'Managed', 'Ephemeral'|
|kubelet_disk_type|text|KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage Currently allows one value, OS, resulting in Kubelet using the OS disk for data Possible values include: 'OS', 'Temporary'|
|vnet_subnet_id|text|VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods|
|pod_subnet_id|text|Pod SubnetID specifies the VNet's subnet identifier for pods|
|max_pods|integer|Maximum number of pods that can run on a node|
|os_type|text|OsType to be used to specify os type Choose from Linux and Windows Default to Linux Possible values include: 'Linux', 'Windows'|
|os_sku|text|OsSKU to be used to specify os sku Choose from Ubuntu(default) and CBLMariner for Linux OSType Not applicable to Windows OSType Possible values include: 'Ubuntu', 'CBLMariner'|
|max_count|integer|Maximum number of nodes for auto-scaling|
|min_count|integer|Minimum number of nodes for auto-scaling|
|enable_auto_scaling|boolean|Whether to enable auto-scaler|
|type|text|AgentPoolType represents types of an agent pool Possible values include: 'VirtualMachineScaleSets', 'AvailabilitySet'|
|mode|text|AgentPoolMode represents mode of an agent pool Possible values include: 'System', 'User'|
|orchestrator_version|text|Version of orchestrator specified when creating the managed cluster|
|node_image_version|text|Version of node image|
|upgrade_settings_max_surge|text|Count or percentage of additional nodes to be added during upgrade If empty uses AKS default|
|provisioning_state|text|The current deployment or provisioning state, which only appears in the response|
|power_state_code|text|Tells whether the cluster is Running or Stopped Possible values include: 'Running', 'Stopped'|
|availability_zones|text[]|Availability zones for nodes Must use VirtualMachineScaleSets AgentPoolType|
|enable_node_public_ip|boolean|Enable public IP for nodes|
|node_public_ip_prefix_id|text|Public IP Prefix ID VM nodes use IPs assigned from this Public IP Prefix|
|scale_set_priority|text|ScaleSetPriority to be used to specify virtual machine scale set priority Default to regular Possible values include: 'Spot', 'Regular'|
|scale_set_eviction_policy|text|ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set Default to Delete Possible values include: 'Delete', 'Deallocate'|
|spot_max_price|float|SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand|
|tags|jsonb|Agent pool tags to be persisted on the agent pool virtual machine scale set|
|node_labels|jsonb|Agent pool node labels to be persisted across all nodes in agent pool|
|node_taints|text[]|Taints added to new nodes during node pool create and scale For example, key=value:NoSchedule|
|proximity_placement_group_id|text|The ID for Proximity Placement Group|
|kubelet_config_cpu_manager_policy|text|CPU Manager policy to use|
|kubelet_config_cpu_cfs_quota|boolean|Enable CPU CFS quota enforcement for containers that specify CPU limits|
|kubelet_config_cpu_cfs_quota_period|text|Sets CPU CFS quota period value|
|kubelet_config_image_gc_high_threshold|integer|The percent of disk usage after which image garbage collection is always run|
|kubelet_config_image_gc_low_threshold|integer|The percent of disk usage before which image garbage collection is never run|
|kubelet_config_topology_manager_policy|text|Topology Manager policy to use|
|kubelet_config_allowed_unsafe_sysctls|text[]|Allowlist of unsafe sysctls or unsafe sysctl patterns (ending in `*`)|
|kubelet_config_fail_swap_on|boolean|If set to true it will make the Kubelet fail to start if swap is enabled on the node|
|kubelet_config_container_log_max_size_mb|integer|The maximum size (eg 10Mi) of container log file before it is rotated|
|kubelet_config_container_log_max_files|integer|The maximum number of container log files that can be present for a container The number must be ≥ 2|
|kubelet_config_pod_max_pids|integer|The maximum number of processes per pod|
|linux_os_config|jsonb|LinuxOSConfig specifies the OS configuration of linux agent nodes|
|enable_encryption_at_host|boolean|Whether to enable EncryptionAtHost|
|enable_fips|boolean|Whether to use FIPS enabled OS|
|gpu_instance_profile|text|GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU Supported values are MIG1g, MIG2g, MIG3g, MIG4g and MIG7g Possible values include: 'MIG1g', 'MIG2g', 'MIG3g', 'MIG4g', 'MIG7g'|
