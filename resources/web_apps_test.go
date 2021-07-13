package resources_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func buildWebAppsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	apps := mocks.NewMockAppsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Apps: apps,
		},
	}

	site := web.Site{}
	err := faker.FakeData(&site)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	page := web.NewAppCollectionPage(web.AppCollection{Value: &[]web.Site{site}}, func(ctx context.Context, collection web.AppCollection) (web.AppCollection, error) {
		return web.AppCollection{}, nil
	})
	apps.EXPECT().List(gomock.Any()).Return(page, nil)

	value := ioutil.NopCloser(strings.NewReader("hello world")) // r type is io.ReadCloser

	response := web.ReadCloser{Response: autorest.Response{Response: &http.Response{Body: value}}}
	apps.EXPECT().ListPublishingProfileXMLWithSecrets(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil)

	return s
}

func TestWebApps(t *testing.T) {
	azureTestHelper(t, resources.WebApps(), buildWebAppsMock)
}
