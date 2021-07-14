package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/go-autorest/autorest"
)

type AuthorizationClient struct {
	RoleAssignments RoleAssignmentsClient
}

func NewAuthorizationClient(subscriptionId string, auth autorest.Authorizer) AuthorizationClient {
	assignments := authorization.NewRoleAssignmentsClient(subscriptionId)
	assignments.Authorizer = auth
	return AuthorizationClient{
		RoleAssignments: assignments,
	}
}

type RoleAssignmentsClient interface {
	List(ctx context.Context, filter string) (result authorization.RoleAssignmentListResultPage, err error)
}
