package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/Azure/go-autorest/autorest"
)

type MonitorClient struct {
	ActivityLogAlerts ActivityLogAlertsClient
	LogProfiles       LogProfilesClient
	ActivityLogs      ActivityLogClient
}

func NewMonitorClient(subscriptionId string, auth autorest.Authorizer) MonitorClient {
	activityLogAlerts := insights.NewActivityLogAlertsClient(subscriptionId)
	activityLogAlerts.Authorizer = auth

	logProfiles := insights.NewLogProfilesClient(subscriptionId)
	logProfiles.Authorizer = auth

	activityLogs := insights.NewActivityLogsClient(subscriptionId)
	activityLogs.Authorizer = auth
	return MonitorClient{
		ActivityLogAlerts: activityLogAlerts,
		LogProfiles:       logProfiles,
		ActivityLogs:      activityLogs,
	}
}

type ActivityLogAlertsClient interface {
	ListBySubscriptionID(ctx context.Context) (result insights.ActivityLogAlertList, err error)
}
type ActivityLogClient interface {
	List(ctx context.Context, filter string, selectParameter string) (result insights.EventDataCollectionPage, err error)
}
type LogProfilesClient interface {
	List(ctx context.Context) (result insights.LogProfileCollection, err error)
}
