package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdGroups() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_groups",
		Resolver:     fetchAdGroups,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryObject.Entity.ID"),
			},
			{
				Name:     "deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DirectoryObject.DeletedDateTime"),
			},
			{
				Name: "classification",
				Type: schema.TypeString,
			},
			{
				Name: "created_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "has_members_with_license_errors",
				Type: schema.TypeBool,
			},
			{
				Name: "group_types",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "license_processing_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseProcessingState.State"),
			},
			{
				Name: "mail",
				Type: schema.TypeString,
			},
			{
				Name: "mail_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "mail_nickname",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_last_sync_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "on_premises_net_bios_name",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_sam_account_name",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_security_identifier",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_sync_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "preferred_data_location",
				Type: schema.TypeString,
			},
			{
				Name: "proxy_addresses",
				Type: schema.TypeStringArray,
			},
			{
				Name: "renewed_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "security_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "security_identifier",
				Type: schema.TypeString,
			},
			{
				Name: "visibility",
				Type: schema.TypeString,
			},
			{
				Name: "allow_external_senders",
				Type: schema.TypeBool,
			},
			{
				Name: "auto_subscribe_new_members",
				Type: schema.TypeBool,
			},
			{
				Name: "is_subscribed_by_mail",
				Type: schema.TypeBool,
			},
			{
				Name: "unseen_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "is_archived",
				Type: schema.TypeBool,
			},
			{
				Name:     "created_on_behalf_of_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedOnBehalfOf.Entity.ID"),
			},
			{
				Name:     "created_on_behalf_of_deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOnBehalfOf.DeletedDateTime"),
			},
			{
				Name:     "conversations",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsConversations,
			},
			{
				Name:     "threads",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsThreads,
			},
			{
				Name:     "calendar",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsCalendar,
			},
			{
				Name:     "calendar_view",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsCalendarView,
			},
			{
				Name:     "events",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsEvents,
			},
			{
				Name:     "photo_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Photo.Entity.ID"),
			},
			{
				Name:     "photo_height",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Photo.Height"),
			},
			{
				Name:     "photo_width",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Photo.Width"),
			},
			{
				Name:     "drive",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsDrive,
			},
			{
				Name:     "drives",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsDrives,
			},
			{
				Name:     "sites",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsSites,
			},
			{
				Name:     "extensions",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsExtensions,
			},
			{
				Name:     "planner",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsPlanner,
			},
			{
				Name:     "onenote",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsOnenote,
			},
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Team.Entity.ID"),
			},
			{
				Name:     "team_web_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Team.WebURL"),
			},
			{
				Name:     "team_member_settings_allow_create_update_channels",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.MemberSettings.AllowCreateUpdateChannels"),
			},
			{
				Name:     "team_member_settings_allow_delete_channels",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.MemberSettings.AllowDeleteChannels"),
			},
			{
				Name:     "team_member_settings_allow_add_remove_apps",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.MemberSettings.AllowAddRemoveApps"),
			},
			{
				Name:     "team_member_settings_allow_create_update_remove_tabs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.MemberSettings.AllowCreateUpdateRemoveTabs"),
			},
			{
				Name:     "team_member_settings_allow_create_update_remove_connectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.MemberSettings.AllowCreateUpdateRemoveConnectors"),
			},
			{
				Name:     "team_guest_settings_allow_create_update_channels",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.GuestSettings.AllowCreateUpdateChannels"),
			},
			{
				Name:     "team_guest_settings_allow_delete_channels",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.GuestSettings.AllowDeleteChannels"),
			},
			{
				Name:     "team_messaging_settings",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsTeamMessagingSettings,
			},
			{
				Name:     "team_fun_settings_allow_giphy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.FunSettings.AllowGiphy"),
			},
			{
				Name:     "team_fun_settings_giphy_content_rating",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Team.FunSettings.GiphyContentRating"),
			},
			{
				Name:     "team_fun_settings_allow_stickers_and_memes",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.FunSettings.AllowStickersAndMemes"),
			},
			{
				Name:     "team_fun_settings_allow_custom_memes",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.FunSettings.AllowCustomMemes"),
			},
			{
				Name:     "team_is_archived",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Team.IsArchived"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_ad_group_assigned_licenses",
				Resolver: fetchAdGroupAssignedLicenses,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "disabled_plans",
						Type: schema.TypeStringArray,
					},
					{
						Name:     "sku_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SKUID"),
					},
				},
			},
			{
				Name:     "azure_ad_group_on_premises_provisioning_errors",
				Resolver: fetchAdGroupOnPremisesProvisioningErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "value",
						Type: schema.TypeString,
					},
					{
						Name: "category",
						Type: schema.TypeString,
					},
					{
						Name: "property_causing_error",
						Type: schema.TypeString,
					},
					{
						Name: "occurred_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_members",
				Resolver: fetchAdGroupMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_member_of",
				Resolver: fetchAdGroupMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_members_with_license_errors",
				Resolver: fetchAdGroupMembersWithLicenseErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_transitive_members",
				Resolver: fetchAdGroupTransitiveMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_transitive_member_of",
				Resolver: fetchAdGroupTransitiveMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_owners",
				Resolver: fetchAdGroupOwners,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_settings",
				Resolver: fetchAdGroupSettings,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
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
						Name:     "template_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TemplateID"),
					},
					{
						Name:     "values",
						Type:     schema.TypeJSON,
						Resolver: resolveAdGroupSettingsValues,
					},
				},
			},
			{
				Name:     "azure_ad_group_photos",
				Resolver: fetchAdGroupPhotos,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "height",
						Type: schema.TypeBigInt,
					},
					{
						Name: "width",
						Type: schema.TypeBigInt,
					},
				},
			},
			{
				Name:     "azure_ad_group_accepted_senders",
				Resolver: fetchAdGroupAcceptedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_rejected_senders",
				Resolver: fetchAdGroupRejectedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_group_lifecycle_policies",
				Resolver: fetchAdGroupLifecyclePolicies,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "group_lifetime_in_days",
						Type: schema.TypeBigInt,
					},
					{
						Name: "managed_group_types",
						Type: schema.TypeString,
					},
					{
						Name: "alternate_notification_emails",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_group_team_channels",
				Resolver: fetchAdGroupTeamChannels,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
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
						Name: "email",
						Type: schema.TypeString,
					},
					{
						Name:     "web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("WebURL"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_group_team_channel_tabs",
						Resolver: fetchAdGroupTeamChannelTabs,
						Columns: []schema.Column{
							{
								Name:        "group_team_channel_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_group_team_channels table (FK)",
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
								Name:     "web_url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("WebURL"),
							},
							{
								Name:     "configuration_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Configuration.EntityID"),
							},
							{
								Name:     "configuration_content_url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Configuration.ContentURL"),
							},
							{
								Name:     "configuration_remove_url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Configuration.RemoveURL"),
							},
							{
								Name:     "configuration_website_url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Configuration.WebsiteURL"),
							},
							{
								Name:     "teams_app_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("TeamsApp.Entity.ID"),
							},
							{
								Name:     "teams_app_external_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("TeamsApp.ExternalID"),
							},
							{
								Name:     "teams_app_display_name",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("TeamsApp.DisplayName"),
							},
							{
								Name:     "teams_app_distribution_method",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("TeamsApp.DistributionMethod"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:     "azure_ad_group_team_channel_tab_teams_app_app_definitions",
								Resolver: fetchAdGroupTeamChannelTabTeamsAppAppDefinitions,
								Columns: []schema.Column{
									{
										Name:        "group_team_channel_tab_cq_id",
										Description: "Unique CloudQuery ID of azure_ad_group_team_channel_tabs table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:     "id",
										Type:     schema.TypeString,
										Resolver: schema.PathResolver("Entity.ID"),
									},
									{
										Name:     "teams_app_id",
										Type:     schema.TypeString,
										Resolver: schema.PathResolver("TeamsAppID"),
									},
									{
										Name: "display_name",
										Type: schema.TypeString,
									},
									{
										Name: "version",
										Type: schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_group_team_installed_apps",
				Resolver: fetchAdGroupTeamInstalledApps,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name:     "teams_app_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsApp.Entity.ID"),
					},
					{
						Name:     "teams_app_external_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsApp.ExternalID"),
					},
					{
						Name:     "teams_app_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsApp.DisplayName"),
					},
					{
						Name:     "teams_app_distribution_method",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsApp.DistributionMethod"),
					},
					{
						Name:     "teams_app_definition_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsAppDefinition.Entity.ID"),
					},
					{
						Name:     "teams_app_definition_teams_app_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsAppDefinition.TeamsAppID"),
					},
					{
						Name:     "teams_app_definition_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsAppDefinition.DisplayName"),
					},
					{
						Name:     "teams_app_definition_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TeamsAppDefinition.Version"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_group_team_installed_app_teams_app_app_definitions",
						Resolver: fetchAdGroupTeamInstalledAppTeamsAppAppDefinitions,
						Columns: []schema.Column{
							{
								Name:        "group_team_installed_app_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_group_team_installed_apps table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Entity.ID"),
							},
							{
								Name:     "teams_app_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("TeamsAppID"),
							},
							{
								Name: "display_name",
								Type: schema.TypeString,
							},
							{
								Name: "version",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_group_team_operations",
				Resolver: fetchAdGroupTeamOperations,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "operation_type",
						Type: schema.TypeString,
					},
					{
						Name: "created_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "last_action_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "attempts_count",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "target_resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TargetResourceID"),
					},
					{
						Name: "target_resource_location",
						Type: schema.TypeString,
					},
					{
						Name:     "error_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Error.Code"),
					},
					{
						Name:     "error_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Error.Message"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAdGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Graph
	response, err := svc.Groups().Request().Get(ctx)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
func resolveAdGroupsConversations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Conversations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsThreads(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Threads)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsCalendar(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	if p.Calendar == nil {
		return nil
	}
	j, err := json.Marshal(p.Calendar)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsCalendarView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.CalendarView)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsEvents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Events)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsDrive(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	if p.Drive == nil {
		return nil
	}
	j, err := json.Marshal(p.Drive)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsDrives(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Drives)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsSites(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Sites)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsExtensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Extensions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsPlanner(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	if p.Planner == nil {
		return nil
	}
	j, err := json.Marshal(p.Planner)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsOnenote(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	if p.Onenote == nil {
		return nil
	}
	j, err := json.Marshal(p.Onenote)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsTeamMessagingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", resource.Item)
	}
	if p.Team == nil || p.Team.MessagingSettings == nil {
		return nil
	}
	j, err := json.Marshal(p.Team.MessagingSettings)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func fetchAdGroupAssignedLicenses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.AssignedLicenses
	return nil
}
func fetchAdGroupOnPremisesProvisioningErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.OnPremisesProvisioningErrors
	return nil
}
func fetchAdGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.Members
	return nil
}
func fetchAdGroupMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.MemberOf
	return nil
}
func fetchAdGroupMembersWithLicenseErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.MembersWithLicenseErrors
	return nil
}
func fetchAdGroupTransitiveMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.TransitiveMembers
	return nil
}
func fetchAdGroupTransitiveMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.TransitiveMemberOf
	return nil
}
func fetchAdGroupOwners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.Owners
	return nil
}
func fetchAdGroupSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.Settings
	return nil
}
func resolveAdGroupSettingsValues(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.GroupSetting)
	if !ok {
		return fmt.Errorf("expected to have msgraph.GroupSetting but got %T", resource.Item)
	}

	j := map[string]interface{}{}
	for _, v := range p.Values {
		j[*v.Name] = *v.Value
	}

	return resource.Set(c.Name, j)
}
func fetchAdGroupPhotos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.Photos
	return nil
}
func fetchAdGroupAcceptedSenders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.AcceptedSenders
	return nil
}
func fetchAdGroupRejectedSenders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.RejectedSenders
	return nil
}
func fetchAdGroupLifecyclePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.GroupLifecyclePolicies
	return nil
}
func fetchAdGroupTeamOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	if p.Team == nil || p.Team.Operations == nil {
		return nil
	}
	res <- p.Team.Operations
	return nil
}
