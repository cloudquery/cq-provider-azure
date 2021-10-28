package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdUsers() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_users",
		Resolver:     fetchAdUsers,
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
				Name: "account_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "age_group",
				Type: schema.TypeString,
			},
			{
				Name: "business_phones",
				Type: schema.TypeStringArray,
			},
			{
				Name: "city",
				Type: schema.TypeString,
			},
			{
				Name: "company_name",
				Type: schema.TypeString,
			},
			{
				Name: "consent_provided_for_minor",
				Type: schema.TypeString,
			},
			{
				Name: "country",
				Type: schema.TypeString,
			},
			{
				Name: "creation_type",
				Type: schema.TypeString,
			},
			{
				Name: "department",
				Type: schema.TypeString,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name:     "employee_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmployeeID"),
			},
			{
				Name: "fax_number",
				Type: schema.TypeString,
			},
			{
				Name: "given_name",
				Type: schema.TypeString,
			},
			{
				Name: "im_addresses",
				Type: schema.TypeStringArray,
			},
			{
				Name: "is_resource_account",
				Type: schema.TypeBool,
			},
			{
				Name: "job_title",
				Type: schema.TypeString,
			},
			{
				Name: "last_password_change_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "legal_age_group_classification",
				Type: schema.TypeString,
			},
			{
				Name: "mail",
				Type: schema.TypeString,
			},
			{
				Name: "mail_nickname",
				Type: schema.TypeString,
			},
			{
				Name: "mobile_phone",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_distinguished_name",
				Type: schema.TypeString,
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute1",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute1"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute2",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute2"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute3",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute3"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute4",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute4"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute5",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute5"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute6",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute6"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute7",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute7"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute8",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute8"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute9",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute9"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute10",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute10"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute11",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute11"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute12",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute12"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute13",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute13"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute14",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute14"),
			},
			{
				Name:     "on_premises_extension_attributes_extension_attribute15",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesExtensionAttributes.ExtensionAttribute15"),
			},
			{
				Name:     "on_premises_immutable_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OnPremisesImmutableID"),
			},
			{
				Name: "on_premises_last_sync_date_time",
				Type: schema.TypeTimestamp,
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
				Name: "on_premises_domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_sam_account_name",
				Type: schema.TypeString,
			},
			{
				Name: "on_premises_user_principal_name",
				Type: schema.TypeString,
			},
			{
				Name: "other_mails",
				Type: schema.TypeStringArray,
			},
			{
				Name: "password_policies",
				Type: schema.TypeString,
			},
			{
				Name:     "password_profile_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PasswordProfile.Password"),
			},
			{
				Name:     "password_profile_force_change_password_next_sign_in",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PasswordProfile.ForceChangePasswordNextSignIn"),
			},
			{
				Name:     "password_profile_force_change_password_next_sign_in_with_mfa",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PasswordProfile.ForceChangePasswordNextSignInWithMFA"),
			},
			{
				Name: "office_location",
				Type: schema.TypeString,
			},
			{
				Name: "postal_code",
				Type: schema.TypeString,
			},
			{
				Name: "preferred_language",
				Type: schema.TypeString,
			},
			{
				Name: "proxy_addresses",
				Type: schema.TypeStringArray,
			},
			{
				Name: "show_in_address_list",
				Type: schema.TypeBool,
			},
			{
				Name: "sign_in_sessions_valid_from_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "street_address",
				Type: schema.TypeString,
			},
			{
				Name: "surname",
				Type: schema.TypeString,
			},
			{
				Name: "usage_location",
				Type: schema.TypeString,
			},
			{
				Name: "user_principal_name",
				Type: schema.TypeString,
			},
			{
				Name: "user_type",
				Type: schema.TypeString,
			},
			{
				Name:     "automatic_replies_setting_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.Status"),
			},
			{
				Name:     "automatic_replies_setting_external_audience",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ExternalAudience"),
			},
			{
				Name:     "automatic_replies_setting_scheduled_start_date_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ScheduledStartDateTime.DateTime"),
			},
			{
				Name:     "automatic_replies_setting_scheduled_start_date_time_time_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ScheduledStartDateTime.TimeZone"),
			},
			{
				Name:     "automatic_replies_setting_scheduled_end_date_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ScheduledEndDateTime.DateTime"),
			},
			{
				Name:     "automatic_replies_setting_scheduled_end_date_time_time_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ScheduledEndDateTime.TimeZone"),
			},
			{
				Name:     "automatic_replies_setting_internal_reply_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.InternalReplyMessage"),
			},
			{
				Name:     "automatic_replies_setting_external_reply_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.AutomaticRepliesSetting.ExternalReplyMessage"),
			},
			{
				Name:     "archive_folder",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.ArchiveFolder"),
			},
			{
				Name:     "time_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.TimeZone"),
			},
			{
				Name:     "language_locale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.Language.Locale"),
			},
			{
				Name:     "language_display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.Language.DisplayName"),
			},
			{
				Name:     "working_hours_days_of_week",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MailboxSettings.WorkingHours.DaysOfWeek"),
			},
			{
				Name:     "working_hours_start_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.WorkingHours.StartTime"),
			},
			{
				Name:     "working_hours_end_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.WorkingHours.EndTime"),
			},
			{
				Name:     "working_hours_time_zone_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.WorkingHours.TimeZone.Name"),
			},
			{
				Name:     "date_format",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.DateFormat"),
			},
			{
				Name:     "time_format",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MailboxSettings.TimeFormat"),
			},
			{
				Name: "device_enrollment_limit",
				Type: schema.TypeBigInt,
			},
			{
				Name: "about_me",
				Type: schema.TypeString,
			},
			{
				Name: "birthday",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "hire_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "interests",
				Type: schema.TypeStringArray,
			},
			{
				Name: "my_site",
				Type: schema.TypeString,
			},
			{
				Name: "past_projects",
				Type: schema.TypeStringArray,
			},
			{
				Name: "preferred_name",
				Type: schema.TypeString,
			},
			{
				Name: "responsibilities",
				Type: schema.TypeStringArray,
			},
			{
				Name: "schools",
				Type: schema.TypeStringArray,
			},
			{
				Name: "skills",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "manager_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Manager.Entity.ID"),
			},
			{
				Name:     "manager_deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Manager.DeletedDateTime"),
			},
			{
				Name:     "outlook",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersOutlook,
			},
			{
				Name:     "messages",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersMessages,
			},
			{
				Name:     "mail_folders",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersMailFolders,
			},
			{
				Name:     "calendar",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersCalendar,
			},
			{
				Name:     "calendars",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersCalendars,
			},
			{
				Name:     "calendar_groups",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersCalendarGroups,
			},
			{
				Name:     "calendar_view",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersCalendarView,
			},
			{
				Name:     "events",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersEvents,
			},
			{
				Name:     "contacts",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersContacts,
			},
			{
				Name:     "contact_folders",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersContactFolders,
			},
			{
				Name:     "inference_classification_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InferenceClassification.Entity.ID"),
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
				Resolver: resolveAdUsersDrive,
			},
			{
				Name:     "drives",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersDrives,
			},
			{
				Name:     "planner",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersPlanner,
			},
			{
				Name:     "insights_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Insights.Entity.ID"),
			},
			{
				Name:     "contribution_to_content_discovery_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.ContributionToContentDiscoveryDisabled"),
			},
			{
				Name:     "contribution_to_content_discovery_as_organization_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.ContributionToContentDiscoveryAsOrganizationDisabled"),
			},
			{
				Name:     "onenote",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersOnenote,
			},
			{
				Name:     "activities",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersActivities,
			},
			{
				Name:     "online_meetings",
				Type:     schema.TypeJSON,
				Resolver: resolveAdUsersOnlineMeetings,
			},
			{
				Name:     "joined_teams",
				Type:     schema.TypeStringArray,
				Resolver: resolveAdUsersJoinedTeams,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_ad_user_assigned_licenses",
				Resolver: fetchAdUserAssignedLicenses,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
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
				Name:     "azure_ad_user_assigned_plans",
				Resolver: fetchAdUserAssignedPlans,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "assigned_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "capability_status",
						Type: schema.TypeString,
					},
					{
						Name: "service",
						Type: schema.TypeString,
					},
					{
						Name:     "service_plan_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ServicePlanID"),
					},
				},
			},
			{
				Name:     "azure_ad_user_license_assignment_states",
				Resolver: fetchAdUserLicenseAssignmentStates,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "sku_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SKUID"),
					},
					{
						Name: "disabled_plans",
						Type: schema.TypeStringArray,
					},
					{
						Name: "assigned_by_group",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "error",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_user_on_premises_provisioning_errors",
				Resolver: fetchAdUserOnPremisesProvisioningErrors,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
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
				Name:     "azure_ad_user_provisioned_plans",
				Resolver: fetchAdUserProvisionedPlans,
				Columns: []schema.Column{
					{
						Name:        "user_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_users table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "capability_status",
						Type: schema.TypeString,
					},
					{
						Name: "provisioning_status",
						Type: schema.TypeString,
					},
					{
						Name: "service",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_user_owned_devices",
				Resolver: fetchAdUserOwnedDevices,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_registered_devices",
				Resolver: fetchAdUserRegisteredDevices,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_direct_reports",
				Resolver: fetchAdUserDirectReports,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_member_of",
				Resolver: fetchAdUserMemberOfs,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_created_objects",
				Resolver: fetchAdUserCreatedObjects,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_owned_objects",
				Resolver: fetchAdUserOwnedObjects,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_license_details",
				Resolver: fetchAdUserLicenseDetails,
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
						Name:     "sku_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SKUID"),
					},
					{
						Name:     "sku_part_number",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SKUPartNumber"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_user_license_detail_service_plans",
						Resolver: fetchAdUserLicenseDetailServicePlans,
						Columns: []schema.Column{
							{
								Name:        "user_license_detail_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_license_details table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "service_plan_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ServicePlanID"),
							},
							{
								Name: "service_plan_name",
								Type: schema.TypeString,
							},
							{
								Name: "provisioning_status",
								Type: schema.TypeString,
							},
							{
								Name: "applies_to",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_user_transitive_member_of",
				Resolver: fetchAdUserTransitiveMemberOfs,
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
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "azure_ad_user_people",
				Resolver: fetchAdUserPeoples,
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
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "given_name",
						Type: schema.TypeString,
					},
					{
						Name: "surname",
						Type: schema.TypeString,
					},
					{
						Name: "birthday",
						Type: schema.TypeString,
					},
					{
						Name: "person_notes",
						Type: schema.TypeString,
					},
					{
						Name: "is_favorite",
						Type: schema.TypeBool,
					},
					{
						Name: "job_title",
						Type: schema.TypeString,
					},
					{
						Name: "company_name",
						Type: schema.TypeString,
					},
					{
						Name: "yomi_company",
						Type: schema.TypeString,
					},
					{
						Name: "department",
						Type: schema.TypeString,
					},
					{
						Name: "office_location",
						Type: schema.TypeString,
					},
					{
						Name: "profession",
						Type: schema.TypeString,
					},
					{
						Name:     "person_type_class",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PersonType.Class"),
					},
					{
						Name:     "person_type_subclass",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PersonType.Subclass"),
					},
					{
						Name: "user_principal_name",
						Type: schema.TypeString,
					},
					{
						Name: "im_address",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_user_person_scored_email_addresses",
						Resolver: fetchAdUserPeopleScoredEmailAddresses,
						Columns: []schema.Column{
							{
								Name:        "user_people_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_people table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "address",
								Type: schema.TypeString,
							},
							{
								Name: "relevance_score",
								Type: schema.TypeFloat,
							},
							{
								Name: "selection_likelihood",
								Type: schema.TypeString,
							},
							{
								Name:     "item_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ItemID"),
							},
						},
					},
					{
						Name:     "azure_ad_user_person_phones",
						Resolver: fetchAdUserPeoplePhones,
						Columns: []schema.Column{
							{
								Name:        "user_people_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_people table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
							{
								Name: "number",
								Type: schema.TypeString,
							},
							{
								Name: "region",
								Type: schema.TypeString,
							},
							{
								Name: "language",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "azure_ad_user_person_postal_addresses",
						Resolver: fetchAdUserPeoplePostalAddresses,
						Columns: []schema.Column{
							{
								Name:        "user_people_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_people table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "display_name",
								Type: schema.TypeString,
							},
							{
								Name: "location_email_address",
								Type: schema.TypeString,
							},
							{
								Name:     "address_street",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Address.Street"),
							},
							{
								Name:     "address_city",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Address.City"),
							},
							{
								Name:     "address_state",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Address.State"),
							},
							{
								Name:     "address_country_or_region",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Address.CountryOrRegion"),
							},
							{
								Name:     "address_postal_code",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Address.PostalCode"),
							},
							{
								Name:     "location_uri",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("LocationURI"),
							},
							{
								Name:     "coordinates_latitude",
								Type:     schema.TypeFloat,
								Resolver: schema.PathResolver("Coordinates.Latitude"),
							},
							{
								Name:     "coordinates_longitude",
								Type:     schema.TypeFloat,
								Resolver: schema.PathResolver("Coordinates.Longitude"),
							},
							{
								Name:     "coordinates_accuracy",
								Type:     schema.TypeFloat,
								Resolver: schema.PathResolver("Coordinates.Accuracy"),
							},
							{
								Name:     "coordinates_altitude",
								Type:     schema.TypeFloat,
								Resolver: schema.PathResolver("Coordinates.Altitude"),
							},
							{
								Name:     "coordinates_altitude_accuracy",
								Type:     schema.TypeFloat,
								Resolver: schema.PathResolver("Coordinates.AltitudeAccuracy"),
							},
							{
								Name: "location_type",
								Type: schema.TypeString,
							},
							{
								Name:     "unique_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("UniqueID"),
							},
							{
								Name:     "unique_id_type",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("UniqueIDType"),
							},
						},
					},
					{
						Name:     "azure_ad_user_person_websites",
						Resolver: fetchAdUserPeopleWebsites,
						Columns: []schema.Column{
							{
								Name:        "user_people_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_people table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
							{
								Name: "address",
								Type: schema.TypeString,
							},
							{
								Name: "display_name",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_user_inference_classification_overrides",
				Resolver: fetchAdUserInferenceClassificationOverrides,
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
						Name: "classify_as",
						Type: schema.TypeString,
					},
					{
						Name:     "sender_email_address_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SenderEmailAddress.Name"),
					},
					{
						Name:     "sender_email_address",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SenderEmailAddress.Address"),
					},
				},
			},
			{
				Name:     "azure_ad_user_photos",
				Resolver: fetchAdUserPhotos,
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
				Name:     "azure_ad_user_extensions",
				Resolver: fetchAdUserExtensions,
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
				},
			},
			AdUserManagedDevices(),
			AdUserManagedAppsRegistrations(),
			{
				Name:     "azure_ad_user_device_management_troubleshooting_events",
				Resolver: fetchAdUserDeviceManagementTroubleshootingEvents,
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
						Name: "event_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "correlation_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CorrelationID"),
					},
				},
			},
			AdUserInsightsTrending(),
			AdUserInsightsShared(),
			AdUserInsightsUsed(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAdUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Graph
	response, err := svc.Users().Request().Get(ctx)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
func resolveAdUsersOutlook(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	if p.Outlook == nil {
		return nil
	}
	j, err := json.Marshal(p.Outlook)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersMessages(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Messages)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersMailFolders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.MailFolders)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersCalendar(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
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
func resolveAdUsersCalendars(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Calendars)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersCalendarGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.CalendarGroups)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersCalendarView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.CalendarView)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersEvents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Events)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersContacts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Contacts)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersContactFolders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.ContactFolders)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersDrive(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
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
func resolveAdUsersDrives(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Drives)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersPlanner(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Planner)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersOnenote(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Onenote)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersActivities(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	j, err := json.Marshal(p.Activities)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersOnlineMeetings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}
	if p.Outlook == nil {
		return nil
	}
	j, err := json.Marshal(p.OnlineMeetings)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdUsersJoinedTeams(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", resource.Item)
	}

	joinedTeams := make([]string, 0, len(p.JoinedTeams))
	for _, t := range p.JoinedTeams {
		joinedTeams = append(joinedTeams, *t.DirectoryObject.Entity.ID)
	}

	return resource.Set(c.Name, joinedTeams)
}
func fetchAdUserAssignedLicenses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.AssignedLicenses
	return nil
}
func fetchAdUserAssignedPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.AssignedPlans
	return nil
}
func fetchAdUserLicenseAssignmentStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.LicenseAssignmentStates
	return nil
}
func fetchAdUserOnPremisesProvisioningErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.OnPremisesProvisioningErrors
	return nil
}
func fetchAdUserProvisionedPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.ProvisionedPlans
	return nil
}
func fetchAdUserOwnedDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.OwnedDevices
	return nil
}
func fetchAdUserRegisteredDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.RegisteredDevices
	return nil
}
func fetchAdUserDirectReports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.DirectReports
	return nil
}
func fetchAdUserMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.MemberOf
	return nil
}
func fetchAdUserCreatedObjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.CreatedObjects
	return nil
}
func fetchAdUserOwnedObjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.OwnedObjects
	return nil
}
func fetchAdUserLicenseDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.LicenseDetails
	return nil
}
func fetchAdUserLicenseDetailServicePlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.LicenseDetails)
	if !ok {
		return fmt.Errorf("expected to have msgraph.LicenseDetails but got %T", parent.Item)
	}
	res <- p.ServicePlans
	return nil
}
func fetchAdUserTransitiveMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.TransitiveMemberOf
	return nil
}
func fetchAdUserPeoples(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.People
	return nil
}
func fetchAdUserPeopleScoredEmailAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Person)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Person but got %T", parent.Item)
	}
	res <- p.ScoredEmailAddresses
	return nil
}
func fetchAdUserPeoplePhones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Person)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Person but got %T", parent.Item)
	}
	res <- p.Phones
	return nil
}
func fetchAdUserPeoplePostalAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Person)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Person but got %T", parent.Item)
	}
	res <- p.PostalAddresses
	return nil
}
func fetchAdUserPeopleWebsites(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Person)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Person but got %T", parent.Item)
	}
	res <- p.Websites
	return nil
}
func fetchAdUserInferenceClassificationOverrides(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.InferenceClassification.Overrides
	return nil
}
func fetchAdUserPhotos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.Photos
	return nil
}
func fetchAdUserExtensions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.Extensions
	return nil
}
func fetchAdUserDeviceManagementTroubleshootingEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.DeviceManagementTroubleshootingEvents
	return nil
}
