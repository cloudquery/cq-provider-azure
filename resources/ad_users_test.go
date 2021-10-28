package resources_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"

	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
)

func createADUsersTestServer(t *testing.T) (*msgraph.GraphServiceRequestBuilder, error) {
	var user msgraph.User
	fakeSkipFields(t, &user, []string{
		"Drive",
		"Drives",
		"MailFolders",
		"Calendar",
		"Calendars",
		"CalendarGroups",
		"CalendarView",
		"Events",
		"Onenote",
		"ContactFolders",
		"Activities",
		"JoinedTeams",
	})
	user.JoinedTeams = []msgraph.Group{
		*fakeGroup(t),
	}
	user.MailFolders = []msgraph.MailFolder{
		*fakeMailFolder(t),
	}
	user.ContactFolders = []msgraph.ContactFolder{
		*fakeContactFolder(t),
	}
	user.Calendar = fakeCalendar(t)
	user.Calendars = []msgraph.Calendar{*fakeCalendar(t)}
	user.CalendarGroups = []msgraph.CalendarGroup{*fakeCalendarGroup(t)}
	user.CalendarView = []msgraph.Event{fakeEvent(t)}
	user.Events = []msgraph.Event{fakeEvent(t)}
	user.Drive = fakeDrive(t)
	user.Drives = []msgraph.Drive{*fakeDrive(t)}
	user.Onenote = fakeOnenote(t)
	user.Activities = []msgraph.UserActivity{*fakeUserActivity(t)}
	mux := httprouter.New()
	mux.GET("/v1.0/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users := []msgraph.User{
			user,
		}

		value, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		resp := msgraph.Paging{
			NextLink: "",
			Value:    value,
		}

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	ts := httptest.NewTLSServer(mux)
	u, _ := url.Parse(ts.URL)
	client := createTestClient(u.Host)
	svc := msgraph.NewClient(&client)
	return svc, nil
}

func fakeContactFolder(t *testing.T) *msgraph.ContactFolder {
	e := msgraph.ContactFolder{}
	fakeSkipFields(t, &e, []string{
		"ChildFolders",
	})
	return &e
}

func fakeCalendarGroup(t *testing.T) *msgraph.CalendarGroup {
	e := msgraph.CalendarGroup{}
	fakeSkipFields(t, &e, []string{
		"Calendars",
	})
	e.Calendars = []msgraph.Calendar{*fakeCalendar(t)}
	return &e
}

func fakeUserActivity(t *testing.T) *msgraph.UserActivity {
	e := msgraph.UserActivity{}
	fakeSkipFields(t, &e, []string{
		"HistoryItems",
	})
	j := "{\"test\":1}"
	e.ContentInfo = []byte(j)
	e.VisualElements.Content = []byte(j)

	return &e
}

func fakeMailFolder(t *testing.T) *msgraph.MailFolder {
	e := msgraph.MailFolder{}
	fakeSkipFields(t, &e, []string{
		"ChildFolders",
	})

	return &e
}

func TestADUsers(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.AdUsers(),
		Config: client.Config{
			Subscriptions: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			graph, err := createADUsersTestServer(t)
			if err != nil {
				return nil, err
			}
			c := client.NewAzureClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"})
			c.SetSubscriptionServices("testProject", services.Services{
				Graph: graph,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}
