package resources_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cloudquery/faker/v3"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func createADGroupsTestServer(t *testing.T) (*msgraph.GraphServiceRequestBuilder, error) {
	mux := httprouter.New()
	mux.GET("/v1.0/groups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		groups := []msgraph.Group{
			*fakeGroup(t),
		}

		value, err := json.Marshal(groups)
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

func fakeGroup(t *testing.T) *msgraph.Group {
	var group msgraph.Group
	fakeSkipFields(t, &group, []string{
		"Conversations",
		"Threads",
		"CalendarView",
		"Events",
		"Calendar",
		"Drive",
		"Drives",
		"Sites",
		"Onenote",
	})

	group.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	group.Conversations = []msgraph.Conversation{fakeConversation(t)}
	group.Calendar = fakeCalendar(t)
	group.CalendarView = []msgraph.Event{fakeEvent(t)}
	group.Events = []msgraph.Event{fakeEvent(t)}
	group.Drive = fakeDrive(t)
	group.Drives = []msgraph.Drive{*fakeDrive(t)}
	group.Sites = []msgraph.Site{*fakeSite(t)}
	group.Onenote = fakeOnenote(t)
	return &group
}

func fakeOnenote(t *testing.T) *msgraph.Onenote {
	e := msgraph.Onenote{}
	fakeSkipFields(t, &e, []string{
		"Notebooks",
		"Sections",
		"SectionGroups",
		"Pages",
	})

	return &e
}
func fakeSite(t *testing.T) *msgraph.Site {
	e := msgraph.Site{}
	fakeSkipFields(t, &e, []string{
		"BaseItem",
		"Drive",
		"Drives",
		"Items",
		"Item",
		"Lists",
		"Sites",
		"Onenote",
		"Analytics",
	})
	fakeSkipFields(t, &e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	})
	return &e
}
func fakeDrive(t *testing.T) *msgraph.Drive {
	e := msgraph.Drive{}
	fakeSkipFields(t, &e, []string{
		"BaseItem",
		"Special",
		"Items",
		"List",
		"Root",
	})
	fakeSkipFields(t, &e.BaseItem, []string{
		"CreatedByUser",
		"LastModifiedByUser",
	})
	return &e
}

func fakeEvent(t *testing.T) msgraph.Event {
	e := msgraph.Event{}
	if err := faker.FakeData(&e.OutlookItem); err != nil {
		t.Fatal(err)
	}
	e.Calendar = fakeCalendar(t)
	return e
}

func fakeCalendar(t *testing.T) *msgraph.Calendar {
	e := msgraph.Calendar{}
	fakeSkipFields(t, &e, []string{
		"Events",
		"CalendarView",
	})
	return &e
}

func fakeConversationThread(t *testing.T) msgraph.ConversationThread {
	e := msgraph.ConversationThread{}
	fakeSkipFields(t, &e, []string{
		"Posts",
	})
	return e
}

func fakeConversation(t *testing.T) msgraph.Conversation {
	e := msgraph.Conversation{}
	fakeSkipFields(t, &e, []string{
		"Threads",
	})
	e.Threads = []msgraph.ConversationThread{fakeConversationThread(t)}
	return e
}

func TestADGroups(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.AdGroups(),
		Config: client.Config{
			Subscriptions: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			graph, err := createADGroupsTestServer(t)
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
