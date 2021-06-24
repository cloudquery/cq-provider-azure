package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

type RBACUsersClient interface {
	List(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultPage, err error)
}

type RBAC struct {
	Users RBACUsersClient
}

func NewRBACClient(subscriptionId string, auth autorest.Authorizer) RBAC {
	client := graphrbac.NewUsersClient(subscriptionId)
	client.Authorizer = auth
	return RBAC{Users: client}
}
