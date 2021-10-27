package services

import (
	"github.com/Azure/go-autorest/autorest"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

type Services struct {
	AD            AD
	Authorization AuthorizationClient
	Compute       ComputeClient
	Container     ContainerServiceClient
	KeyVault      KeyVaultClient
	Monitor       MonitorClient
	MySQL         MySQL
	Network       NetworksClient
	PostgreSQL    PostgreSQL
	Resources     ResourcesClient
	Security      SecurityClient
	SQL           SQLClient
	Storage       StorageClient
	Subscriptions SubscriptionsClient
	Web           WebClient
	Graph         *msgraph.GraphServiceRequestBuilder
}

func InitServices(subscriptionId string, auth autorest.Authorizer) (*Services, error) {
	graph, err := NewGraphClient(subscriptionId, auth)
	if err != nil {
		return nil, err
	}
	return &Services{
		AD:            NewADClient(subscriptionId, auth),
		Authorization: NewAuthorizationClient(subscriptionId, auth),
		Compute:       NewComputeClient(subscriptionId, auth),
		Container:     NewContainerServiceClient(subscriptionId, auth),
		KeyVault:      NewKeyVaultClient(subscriptionId, auth),
		Monitor:       NewMonitorClient(subscriptionId, auth),
		MySQL:         NewMySQLClient(subscriptionId, auth),
		Network:       NewNetworksClient(subscriptionId, auth),
		PostgreSQL:    NewPostgresClient(subscriptionId, auth),
		Resources:     NewResourcesClient(subscriptionId, auth),
		Security:      NewSecurityClient(subscriptionId, auth),
		SQL:           NewSQLClient(subscriptionId, auth),
		Storage:       NewStorageClient(subscriptionId, auth),
		Subscriptions: NewSubscriptionsClient(subscriptionId, auth),
		Web:           NewWebClient(subscriptionId, auth),
		Graph:         graph,
	}, nil
}
