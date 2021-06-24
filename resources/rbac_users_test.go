package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRBACUsers(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockRBACUsersClient(ctrl)
	var user graphrbac.User
	faker.SetIgnoreInterface(true)
	defer faker.SetIgnoreInterface(false)
	if err := faker.FakeData(&user); err != nil {
		t.Fatal(err)
	}
	var signInName graphrbac.SignInName
	if err := faker.FakeData(&signInName); err != nil {
		t.Fatal(err)
	}
	signInName.AdditionalProperties = map[string]interface{}{"test": "value"}
	user.SignInNames = &[]graphrbac.SignInName{signInName}

	userListPage := graphrbac.NewUserListResultPage(
		graphrbac.UserListResult{Value: &[]graphrbac.User{user}},
		func(ctx context.Context, list graphrbac.UserListResult) (graphrbac.UserListResult, error) {
			return graphrbac.UserListResult{}, nil
		},
	)
	m.EXPECT().List(gomock.Any(), "", "").Return(userListPage, nil)
	return services.Services{
		RBAC: services.RBAC{Users: m},
	}
}

func TestRBACUsers(t *testing.T) {
	azureTestHelper(t, resources.RbacUsers(), buildRBACUsers)
}
