package frontdoor

import (
	"math/rand"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/cloudquery/faker/v3"
)

func fakeBasicRouteConfiguration(t *testing.T) frontdoor.BasicRouteConfiguration {
	switch rand.Int() % 3 {
	case 1:
		var config frontdoor.ForwardingConfiguration
		if err := faker.FakeData(&config); err != nil {
			t.Fatal(err)
		}
		return config
	case 2:
		var config frontdoor.RedirectConfiguration
		if err := faker.FakeData(&config); err != nil {
			t.Fatal(err)
		}
		return config
	default:
		var config frontdoor.RouteConfiguration
		if err := faker.FakeData(&config); err != nil {
			t.Fatal(err)
		}
		return config
	}
}

func fakeHeaderAction(t *testing.T) frontdoor.HeaderAction {
	var action frontdoor.HeaderAction
	if err := faker.FakeDataSkipFields(&action, []string{"HeaderActionType"}); err != nil {
		t.Fatal(err)
	}
	action.HeaderActionType = frontdoor.HeaderActionTypeAppend
	return action
}

func fakeRulesEngineMatchCondition(t *testing.T) frontdoor.RulesEngineMatchCondition {
	var condition frontdoor.RulesEngineMatchCondition
	if err := faker.FakeDataSkipFields(&condition,
		[]string{"RulesEngineMatchVariable", "RulesEngineOperator", "Transforms"}); err != nil {
		t.Fatal(err)
	}
	condition.RulesEngineMatchVariable = frontdoor.RulesEngineMatchVariableQueryString
	condition.RulesEngineOperator = frontdoor.RulesEngineOperatorContains
	condition.Transforms = &[]frontdoor.Transform{
		frontdoor.TransformLowercase,
		frontdoor.TransformTrim,
	}
	return condition
}

func fakeRulesEngineRule(t *testing.T) frontdoor.RulesEngineRule {
	var rule frontdoor.RulesEngineRule
	if err := faker.FakeDataSkipFields(&rule,
		[]string{"Action", "MatchConditions", "MatchProcessingBehavior"}); err != nil {
		t.Fatal(err)
	}
	rule.Action = &frontdoor.RulesEngineAction{
		RequestHeaderActions: &[]frontdoor.HeaderAction{
			fakeHeaderAction(t),
			fakeHeaderAction(t),
			fakeHeaderAction(t),
		},
		ResponseHeaderActions: &[]frontdoor.HeaderAction{
			fakeHeaderAction(t),
			fakeHeaderAction(t),
			fakeHeaderAction(t),
		},
		RouteConfigurationOverride: fakeBasicRouteConfiguration(t),
	}
	rule.MatchConditions = &[]frontdoor.RulesEngineMatchCondition{
		fakeRulesEngineMatchCondition(t),
		fakeRulesEngineMatchCondition(t),
		fakeRulesEngineMatchCondition(t),
	}
	rule.MatchProcessingBehavior = frontdoor.MatchProcessingBehaviorContinue

	return rule
}

func fakeRulesEngine(t *testing.T) frontdoor.RulesEngine {
	var engine frontdoor.RulesEngine
	if err := faker.FakeDataSkipFields(&engine, []string{"RulesEngineProperties"}); err != nil {
		t.Fatal(err)
	}
	engine.RulesEngineProperties = &frontdoor.RulesEngineProperties{
		ResourceState: frontdoor.ResourceStateEnabled,
		Rules: &[]frontdoor.RulesEngineRule{
			fakeRulesEngineRule(t),
			fakeRulesEngineRule(t),
			fakeRulesEngineRule(t),
		},
	}
	return engine
}

func fakeRoutingRule(t *testing.T) frontdoor.RoutingRule {
	var rule frontdoor.RoutingRule
	if err := faker.FakeDataSkipFields(&rule, []string{"RoutingRuleProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.RoutingRuleProperties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{
			"ResourceState",
			"AcceptedProtocols",
			"EnabledState",
			"RouteConfiguration",
		}); err != nil {
		t.Fatal(err)
	}

	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.AcceptedProtocols = &[]frontdoor.Protocol{frontdoor.ProtocolHTTPS}
	properties.EnabledState = frontdoor.RoutingRuleEnabledStateEnabled
	properties.RouteConfiguration = fakeBasicRouteConfiguration(t)

	rule.RoutingRuleProperties = &properties
	return rule
}

func fakeLoadBalancingSettingsModel(t *testing.T) frontdoor.LoadBalancingSettingsModel {
	var model frontdoor.LoadBalancingSettingsModel
	if err := faker.FakeDataSkipFields(&model, []string{"LoadBalancingSettingsProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.LoadBalancingSettingsProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	model.LoadBalancingSettingsProperties = &properties
	return model
}

func fakeHealthProbeSettings(t *testing.T) frontdoor.HealthProbeSettingsModel {
	var model frontdoor.HealthProbeSettingsModel
	if err := faker.FakeDataSkipFields(&model, []string{"HealthProbeSettingsProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.HealthProbeSettingsProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState", "Protocol", "HealthProbeMethod", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.Protocol = frontdoor.ProtocolHTTPS
	properties.HealthProbeMethod = frontdoor.HealthProbeMethodGET
	properties.EnabledState = frontdoor.HealthProbeEnabledEnabled

	model.HealthProbeSettingsProperties = &properties
	return model
}

func fakeBackend(t *testing.T) frontdoor.Backend {
	var backend frontdoor.Backend
	if err := faker.FakeDataSkipFields(&backend, []string{"PrivateEndpointStatus", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	backend.PrivateEndpointStatus = frontdoor.PrivateEndpointStatusApproved
	backend.EnabledState = frontdoor.BackendEnabledStateEnabled

	return backend
}

func fakeBackendPools(t *testing.T) frontdoor.BackendPool {
	var pool frontdoor.BackendPool
	if err := faker.FakeDataSkipFields(&pool, []string{"BackendPoolProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.BackendPoolProperties
	if err := faker.FakeDataSkipFields(&properties, []string{"ResourceState", "Backends"}); err != nil {
		t.Fatal(err)
	}

	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.Backends = &[]frontdoor.Backend{
		fakeBackend(t),
		fakeBackend(t),
		fakeBackend(t),
	}
	pool.BackendPoolProperties = &properties
	return pool
}

func fakeCustomHTTPSConfiguration(t *testing.T) *frontdoor.CustomHTTPSConfiguration {
	var config frontdoor.CustomHTTPSConfiguration
	if err := faker.FakeDataSkipFields(&config,
		[]string{"CertificateSource", "MinimumTLSVersion", "CertificateSourceParameters"}); err != nil {
		t.Fatal(err)
	}
	config.CertificateSource = frontdoor.CertificateSourceAzureKeyVault
	config.MinimumTLSVersion = frontdoor.MinimumTLSVersionOneFullStopTwo
	config.CertificateSourceParameters = &frontdoor.CertificateSourceParameters{
		CertificateType: frontdoor.CertificateTypeDedicated,
	}

	return &config
}

func fakeFrontendEndpoint(t *testing.T) frontdoor.FrontendEndpoint {
	var endpoint frontdoor.FrontendEndpoint
	if err := faker.FakeDataSkipFields(&endpoint, []string{"FrontendEndpointProperties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.FrontendEndpointProperties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{
			"ResourceState",
			"CustomHTTPSProvisioningState",
			"CustomHTTPSProvisioningSubstate",
			"CustomHTTPSConfiguration",
			"SessionAffinityEnabledState",
		}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabling
	properties.CustomHTTPSProvisioningState = frontdoor.CustomHTTPSProvisioningStateEnabling
	properties.CustomHTTPSProvisioningSubstate = frontdoor.CustomHTTPSProvisioningSubstateCertificateDeployed
	properties.CustomHTTPSConfiguration = fakeCustomHTTPSConfiguration(t)
	properties.SessionAffinityEnabledState = frontdoor.SessionAffinityEnabledStateDisabled

	endpoint.FrontendEndpointProperties = &properties
	return endpoint
}

func fakeFrontDoor(t *testing.T) frontdoor.FrontDoor {
	var frontDoor frontdoor.FrontDoor
	if err := faker.FakeDataSkipFields(&frontDoor, []string{"Properties"}); err != nil {
		t.Fatal(err)
	}

	var properties frontdoor.Properties
	if err := faker.FakeDataSkipFields(&properties,
		[]string{"ResourceState", "RulesEngines", "RoutingRules", "LoadBalancingSettings", "HealthProbeSettings",
			"BackendPools", "FrontendEndpoints", "BackendPoolsSettings", "EnabledState"}); err != nil {
		t.Fatal(err)
	}
	properties.ResourceState = frontdoor.ResourceStateEnabled
	properties.RulesEngines = &[]frontdoor.RulesEngine{
		fakeRulesEngine(t),
		fakeRulesEngine(t),
		fakeRulesEngine(t),
	}
	properties.RoutingRules = &[]frontdoor.RoutingRule{
		fakeRoutingRule(t),
		fakeRoutingRule(t),
		fakeRoutingRule(t),
	}
	properties.LoadBalancingSettings = &[]frontdoor.LoadBalancingSettingsModel{
		fakeLoadBalancingSettingsModel(t),
		fakeLoadBalancingSettingsModel(t),
	}
	properties.HealthProbeSettings = &[]frontdoor.HealthProbeSettingsModel{
		fakeHealthProbeSettings(t),
		fakeHealthProbeSettings(t),
	}
	properties.BackendPools = &[]frontdoor.BackendPool{
		fakeBackendPools(t),
		fakeBackendPools(t),
		fakeBackendPools(t),
	}
	properties.FrontendEndpoints = &[]frontdoor.FrontendEndpoint{
		fakeFrontendEndpoint(t),
		fakeFrontendEndpoint(t),
		fakeFrontendEndpoint(t),
	}

	var backendPoolsSettings frontdoor.BackendPoolsSettings
	if err := faker.FakeDataSkipFields(&backendPoolsSettings, []string{"EnforceCertificateNameCheck"}); err != nil {
		t.Fatal(err)
	}
	backendPoolsSettings.EnforceCertificateNameCheck = frontdoor.EnforceCertificateNameCheckEnabledStateEnabled
	properties.BackendPoolsSettings = &backendPoolsSettings
	properties.EnabledState = frontdoor.EnabledStateEnabled

	frontDoor.Properties = &properties
	return frontDoor
}
