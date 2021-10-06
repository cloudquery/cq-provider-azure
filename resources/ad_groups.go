package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdGroups() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_groups",
		Description:  "Group undocumented",
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
				Name:        "id",
				Description: "ID undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryObject.Entity.ID"),
			},
			{
				Name:        "deleted_date_time",
				Description: "DeletedDateTime undocumented",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DirectoryObject.DeletedDateTime"),
			},
			{
				Name:        "classification",
				Description: "Classification undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_date_time",
				Description: "CreatedDateTime undocumented",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "Description undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_name",
				Description: "DisplayName undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "has_members_with_license_errors",
				Description: "HasMembersWithLicenseErrors undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "group_types",
				Description: "GroupTypes undocumented",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "license_processing_state",
				Description: "State undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LicenseProcessingState.State"),
			},
			{
				Name:        "mail",
				Description: "Mail undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "mail_enabled",
				Description: "MailEnabled undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "mail_nickname",
				Description: "MailNickname undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "on_premises_domain_name",
				Description: "OnPremisesDomainName undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "on_premises_last_sync_date_time",
				Description: "OnPremisesLastSyncDateTime undocumented",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "on_premises_net_bios_name",
				Description: "OnPremisesNetBiosName undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "on_premises_sam_account_name",
				Description: "OnPremisesSamAccountName undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "on_premises_security_identifier",
				Description: "OnPremisesSecurityIdentifier undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "on_premises_sync_enabled",
				Description: "OnPremisesSyncEnabled undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "preferred_data_location",
				Description: "PreferredDataLocation undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "proxy_addresses",
				Description: "ProxyAddresses undocumented",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "renewed_date_time",
				Description: "RenewedDateTime undocumented",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "security_enabled",
				Description: "SecurityEnabled undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "security_identifier",
				Description: "SecurityIdentifier undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "visibility",
				Description: "Visibility undocumented",
				Type:        schema.TypeString,
			},
			{
				Name:        "allow_external_senders",
				Description: "AllowExternalSenders undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "auto_subscribe_new_members",
				Description: "AutoSubscribeNewMembers undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "is_subscribed_by_mail",
				Description: "IsSubscribedByMail undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "unseen_count",
				Description: "UnseenCount undocumented",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "is_archived",
				Description: "IsArchived undocumented",
				Type:        schema.TypeBool,
			},
			{
				Name:        "created_on_behalf_of_id",
				Description: "ID undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CreatedOnBehalfOf.Entity.ID"),
			},
			{
				Name:        "created_on_behalf_of_deleted_date_time",
				Description: "DeletedDateTime undocumented",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CreatedOnBehalfOf.DeletedDateTime"),
			},
			{
				Name:        "photo_id",
				Description: "ID undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Photo.Entity.ID"),
			},
			{
				Name:        "photo_height",
				Description: "Height undocumented",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Photo.Height"),
			},
			{
				Name:        "photo_width",
				Description: "Width undocumented",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Photo.Width"),
			},
			{
				Name:        "team_id",
				Description: "ID undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Team.Entity.ID"),
			},
			{
				Name:        "team_web_url",
				Description: "WebURL undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Team.WebURL"),
			},
			{
				Name:        "team_member_settings_allow_create_update_channels",
				Description: "AllowCreateUpdateChannels undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MemberSettings.AllowCreateUpdateChannels"),
			},
			{
				Name:        "team_member_settings_allow_delete_channels",
				Description: "AllowDeleteChannels undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MemberSettings.AllowDeleteChannels"),
			},
			{
				Name:        "team_member_settings_allow_add_remove_apps",
				Description: "AllowAddRemoveApps undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MemberSettings.AllowAddRemoveApps"),
			},
			{
				Name:        "team_member_settings_allow_create_update_remove_tabs",
				Description: "AllowCreateUpdateRemoveTabs undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MemberSettings.AllowCreateUpdateRemoveTabs"),
			},
			{
				Name:        "team_member_settings_allow_create_update_remove_connectors",
				Description: "AllowCreateUpdateRemoveConnectors undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MemberSettings.AllowCreateUpdateRemoveConnectors"),
			},
			{
				Name:        "team_guest_settings_allow_create_update_channels",
				Description: "AllowCreateUpdateChannels undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.GuestSettings.AllowCreateUpdateChannels"),
			},
			{
				Name:        "team_guest_settings_allow_delete_channels",
				Description: "AllowDeleteChannels undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.GuestSettings.AllowDeleteChannels"),
			},
			{
				Name:        "team_messaging_settings_allow_user_edit_messages",
				Description: "AllowUserEditMessages undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MessagingSettings.AllowUserEditMessages"),
			},
			{
				Name:        "team_messaging_settings_allow_user_delete_messages",
				Description: "AllowUserDeleteMessages undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MessagingSettings.AllowUserDeleteMessages"),
			},
			{
				Name:        "team_messaging_settings_allow_owner_delete_messages",
				Description: "AllowOwnerDeleteMessages undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MessagingSettings.AllowOwnerDeleteMessages"),
			},
			{
				Name:        "team_messaging_settings_allow_team_mentions",
				Description: "AllowTeamMentions undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MessagingSettings.AllowTeamMentions"),
			},
			{
				Name:        "team_messaging_settings_allow_channel_mentions",
				Description: "AllowChannelMentions undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.MessagingSettings.AllowChannelMentions"),
			},
			{
				Name:        "team_fun_settings_allow_giphy",
				Description: "AllowGiphy undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.FunSettings.AllowGiphy"),
			},
			{
				Name:        "team_fun_settings_giphy_content_rating",
				Description: "GiphyContentRating undocumented",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Team.FunSettings.GiphyContentRating"),
			},
			{
				Name:        "team_fun_settings_allow_stickers_and_memes",
				Description: "AllowStickersAndMemes undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.FunSettings.AllowStickersAndMemes"),
			},
			{
				Name:        "team_fun_settings_allow_custom_memes",
				Description: "AllowCustomMemes undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.FunSettings.AllowCustomMemes"),
			},
			{
				Name:        "team_is_archived",
				Description: "IsArchived undocumented",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Team.IsArchived"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_ad_group_assigned_licenses",
				Description: "AssignedLicense undocumented",
				Resolver:    fetchAdGroupAssignedLicenses,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "disabled_plans",
						Description: "DisabledPlans undocumented",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "s_k_uid",
						Description: "SKUID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SKUID"),
					},
				},
			},
			{
				Name:        "azure_ad_group_on_premises_provisioning_errors",
				Description: "OnPremisesProvisioningError undocumented",
				Resolver:    fetchAdGroupOnPremisesProvisioningErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "value",
						Description: "Value undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "category",
						Description: "Category undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "property_causing_error",
						Description: "PropertyCausingError undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "occurred_date_time",
						Description: "OccurredDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_members",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_settings",
				Description: "GroupSetting undocumented",
				Resolver:    fetchAdGroupSettings,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "display_name",
						Description: "DisplayName undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "template_id",
						Description: "TemplateID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TemplateID"),
					},
					{
						Name:        "values",
						Description: "Values undocumented",
						Type:        schema.TypeJSON,
						Resolver:    resolveAdGroupSettingValues,
					},
				},
			},
			{
				Name:        "azure_ad_group_photos",
				Description: "ProfilePhoto undocumented",
				Resolver:    fetchAdGroupPhotos,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "height",
						Description: "Height undocumented",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "width",
						Description: "Width undocumented",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "azure_ad_group_lifecycle_policies",
				Description: "GroupLifecyclePolicy undocumented",
				Resolver:    fetchAdGroupLifecyclePolicies,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "group_lifetime_in_days",
						Description: "GroupLifetimeInDays undocumented",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "managed_group_types",
						Description: "ManagedGroupTypes undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "alternate_notification_emails",
						Description: "AlternateNotificationEmails undocumented",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_ad_group_team_channels",
				Description: "Channel undocumented",
				Resolver:    fetchAdGroupTeamChannels,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "display_name",
						Description: "DisplayName undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Description undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "email",
						Description: "Email undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "web_url",
						Description: "WebURL undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("WebURL"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_ad_group_team_channel_tabs",
						Description: "TeamsTab undocumented",
						Resolver:    fetchAdGroupTeamChannelTabs,
						Columns: []schema.Column{
							{
								Name:        "group_team_channel_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_group_team_channels table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "id",
								Description: "ID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Entity.ID"),
							},
							{
								Name:        "display_name",
								Description: "DisplayName undocumented",
								Type:        schema.TypeString,
							},
							{
								Name:        "web_url",
								Description: "WebURL undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("WebURL"),
							},
							{
								Name:        "configuration_id",
								Description: "EntityID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Configuration.EntityID"),
							},
							{
								Name:        "configuration_content_url",
								Description: "ContentURL undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Configuration.ContentURL"),
							},
							{
								Name:        "configuration_remove_url",
								Description: "RemoveURL undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Configuration.RemoveURL"),
							},
							{
								Name:        "configuration_website_url",
								Description: "WebsiteURL undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Configuration.WebsiteURL"),
							},
							{
								Name:        "teams_app_id",
								Description: "ID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TeamsApp.Entity.ID"),
							},
							{
								Name:        "teams_app_external_id",
								Description: "ExternalID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TeamsApp.ExternalID"),
							},
							{
								Name:        "teams_app_display_name",
								Description: "DisplayName undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TeamsApp.DisplayName"),
							},
							{
								Name:        "teams_app_distribution_method",
								Description: "DistributionMethod undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TeamsApp.DistributionMethod"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "azure_ad_group_team_channel_tab_teams_app_app_definitions",
								Description: "TeamsAppDefinition undocumented",
								Resolver:    fetchAdGroupTeamChannelTabTeamsAppAppDefinitions,
								Columns: []schema.Column{
									{
										Name:        "group_team_channel_tab_cq_id",
										Description: "Unique CloudQuery ID of azure_ad_group_team_channel_tabs table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "id",
										Description: "ID undocumented",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("Entity.ID"),
									},
									{
										Name:        "teams_app_id",
										Description: "TeamsAppID undocumented",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("TeamsAppID"),
									},
									{
										Name:        "display_name",
										Description: "DisplayName undocumented",
										Type:        schema.TypeString,
									},
									{
										Name:        "version",
										Description: "Version undocumented",
										Type:        schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "azure_ad_group_team_installed_apps",
				Description: "TeamsAppInstallation undocumented",
				Resolver:    fetchAdGroupTeamInstalledApps,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "teams_app_id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsApp.Entity.ID"),
					},
					{
						Name:        "teams_app_external_id",
						Description: "ExternalID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsApp.ExternalID"),
					},
					{
						Name:        "teams_app_display_name",
						Description: "DisplayName undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsApp.DisplayName"),
					},
					{
						Name:        "teams_app_distribution_method",
						Description: "DistributionMethod undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsApp.DistributionMethod"),
					},
					{
						Name:        "teams_app_definition_id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsAppDefinition.Entity.ID"),
					},
					{
						Name:        "teams_app_definition_teams_app_id",
						Description: "TeamsAppID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsAppDefinition.TeamsAppID"),
					},
					{
						Name:        "teams_app_definition_display_name",
						Description: "DisplayName undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsAppDefinition.DisplayName"),
					},
					{
						Name:        "teams_app_definition_version",
						Description: "Version undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TeamsAppDefinition.Version"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_ad_group_team_installed_app_teams_app_app_definitions",
						Description: "TeamsAppDefinition undocumented",
						Resolver:    fetchAdGroupTeamInstalledAppTeamsAppAppDefinitions,
						Columns: []schema.Column{
							{
								Name:        "group_team_installed_app_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_group_team_installed_apps table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "id",
								Description: "ID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Entity.ID"),
							},
							{
								Name:        "teams_app_id",
								Description: "TeamsAppID undocumented",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("TeamsAppID"),
							},
							{
								Name:        "display_name",
								Description: "DisplayName undocumented",
								Type:        schema.TypeString,
							},
							{
								Name:        "version",
								Description: "Version undocumented",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_ad_group_team_operations",
				Description: "TeamsAsyncOperation undocumented",
				Resolver:    fetchAdGroupTeamOperations,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "operation_type",
						Description: "OperationType undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_date_time",
						Description: "CreatedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "status",
						Description: "Status undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_action_date_time",
						Description: "LastActionDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "attempts_count",
						Description: "AttemptsCount undocumented",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "target_resource_id",
						Description: "TargetResourceID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetResourceID"),
					},
					{
						Name:        "target_resource_location",
						Description: "TargetResourceLocation undocumented",
						Type:        schema.TypeString,
					},
					{
						Name:        "error_code",
						Description: "Code undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Error.Code"),
					},
					{
						Name:        "error_message",
						Description: "Message undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Error.Message"),
					},
				},
			},
			{
				Name:        "azure_ad_group_member_ofs",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_members_with_license_errors",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupMembersWithLicenseErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_transitive_members",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupTransitiveMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_transitive_member_ofs",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupTransitiveMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_owners",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupOwners,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_accepted_senders",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupAcceptedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "azure_ad_group_rejected_senders",
				Description: "DirectoryObject Represents an Azure Active Directory object",
				Resolver:    fetchAdGroupRejectedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID undocumented",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.ID"),
					},
					{
						Name:        "deleted_date_time",
						Description: "DeletedDateTime undocumented",
						Type:        schema.TypeTimestamp,
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
func fetchAdGroupSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.Settings
	return nil
}
func resolveAdGroupSettingValues(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchAdGroupLifecyclePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Group but got %T", parent.Item)
	}
	res <- p.GroupLifecyclePolicies
	return nil
}
func fetchAdGroupTeamChannels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	if p.Team == nil || p.Team.Channels == nil {
		return nil
	}
	res <- p.Team.Channels
	return nil
}
func fetchAdGroupTeamChannelTabs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Channel)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Channel but got %T", parent.Item)
	}
	res <- p.Tabs
	return nil
}
func fetchAdGroupTeamChannelTabTeamsAppAppDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.TeamsTab)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Channel but got %T", parent.Item)
	}
	res <- p.TeamsApp.AppDefinitions
	return nil
}
func fetchAdGroupTeamInstalledApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	if p.Team == nil || p.Team.InstalledApps == nil {
		return nil
	}
	res <- p.Team.InstalledApps
	return nil
}
func fetchAdGroupTeamInstalledAppTeamsAppAppDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.TeamsAppInstallation)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	res <- p.TeamsAppDefinition
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
