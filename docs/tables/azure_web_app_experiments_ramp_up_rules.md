
# Table: azure_web_app_experiments_ramp_up_rules
RampUpRule routing rules for ramp up testing This rule allows to redirect static traffic % to a slot or to gradually change routing % based on performance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|action_host_name|text|Hostname of a slot to which the traffic will be redirected if decided to Eg myapp-stageazurewebsitesnet|
|reroute_percentage|float|Percentage of the traffic which will be redirected to <code>ActionHostName</code>|
|change_step|float|In auto ramp up scenario this is the step to add/remove from <code>ReroutePercentage</code> until it reaches \n<code>MinReroutePercentage</code> or <code>MaxReroutePercentage</code> Site metrics are checked every N minutes specified in <code>ChangeIntervalInMinutes</code>\nCustom decision algorithm can be provided in TiPCallback site extension which URL can be specified in <code>ChangeDecisionCallbackUrl</code>|
|change_interval_in_minutes|integer|Specifies interval in minutes to reevaluate ReroutePercentage|
|min_reroute_percentage|float|Specifies lower boundary above which ReroutePercentage will stay|
|max_reroute_percentage|float|Specifies upper boundary below which ReroutePercentage will stay|
|change_decision_callback_url|text|Custom decision algorithm can be provided in TiPCallback site extension which URL can be specified See TiPCallback site extension for the scaffold and contracts https://wwwsiteextensionsnet/packages/TiPCallback/|
|name|text|Name of the routing rule The recommended name would be to point to the slot which will receive the traffic in the experiment|
