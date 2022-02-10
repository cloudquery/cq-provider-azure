//go:generate mockgen -destination=./mocks/datalake.go -package=mocks . DataLakeStorageAccountsClient,
package services

import (
	"context"
	storeAccount "github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
	"github.com/Azure/go-autorest/autorest"
)

type DataLakeClient struct {
	DataLakeStorageAccounts DataLakeStorageAccountsClient
}

func NewDataLakeClient(subscriptionId string, auth autorest.Authorizer) DataLakeClient {
	storeAccounts := storeAccount.NewAccountsClient(subscriptionId)
	storeAccounts.Authorizer = auth
	return DataLakeClient{
		DataLakeStorageAccounts: storeAccounts,
	}
}

type DataLakeStorageAccountsClient interface {
	List(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result storeAccount.DataLakeStoreAccountListResultPage, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result storeAccount.DataLakeStoreAccount, err error)
}
