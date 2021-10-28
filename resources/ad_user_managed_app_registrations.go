package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdUserManagedAppsRegistrations() *schema.Table {
	return &schema.Table{
		Name:     "azure_ad_user_managed_app_registrations",
		Resolver: fetchAdUserManagedAppRegistrations,
		Columns: []schema.Column{
			{
				Name:        "user_cq_id",
				Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Entity.ID"),
			},
			{
				Name: "created_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "last_sync_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "application_version",
				Type: schema.TypeString,
			},
			{
				Name: "management_sdk_version",
				Type: schema.TypeString,
			},
			{
				Name: "platform_version",
				Type: schema.TypeString,
			},
			{
				Name: "device_type",
				Type: schema.TypeString,
			},
			{
				Name: "device_tag",
				Type: schema.TypeString,
			},
			{
				Name: "device_name",
				Type: schema.TypeString,
			},
			{
				Name: "flagged_reasons",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserID"),
			},
			{
				Name: "version",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_ad_user_managed_app_registration_applied_policies",
				Resolver: fetchAdUserManagedAppRegistrationAppliedPolicies,
				Columns: []schema.Column{
					{
						Name:        "user_managed_app_registration_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_app_registrations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "created_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_modified_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "version",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_user_managed_app_registration_intended_policies",
				Resolver: fetchAdUserManagedAppRegistrationIntendedPolicies,
				Columns: []schema.Column{
					{
						Name:        "user_managed_app_registration_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_app_registrations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "created_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_modified_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "version",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_user_managed_app_registration_operations",
				Resolver: fetchAdUserManagedAppRegistrationOperations,
				Columns: []schema.Column{
					{
						Name:        "user_managed_app_registration_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_user_managed_app_registrations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "last_modified_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func fetchAdUserManagedAppRegistrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.ManagedAppRegistrations
	return nil
}
func fetchAdUserManagedAppRegistrationAppliedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedAppRegistration)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedAppRegistration but got %T", parent.Item)
	}
	res <- p.AppliedPolicies
	return nil
}
func fetchAdUserManagedAppRegistrationIntendedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedAppRegistration)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedAppRegistration but got %T", parent.Item)
	}
	res <- p.IntendedPolicies
	return nil
}
func fetchAdUserManagedAppRegistrationOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedAppRegistration)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedAppRegistration but got %T", parent.Item)
	}
	res <- p.Operations
	return nil
}
