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
				Name:     "password_profile_force_change_password_next_sign_in_with_m_f_a",
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
				Name:     "insights_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Insights.Entity.ID"),
			},
			{
				Name:     "settings_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Settings.Entity.ID"),
			},
			{
				Name:     "settings_contribution_to_content_discovery_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.ContributionToContentDiscoveryDisabled"),
			},
			{
				Name:     "settings_contribution_to_content_discovery_as_organization_disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Settings.ContributionToContentDiscoveryAsOrganizationDisabled"),
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
			{
				Name:     "azure_ad_user_managed_devices",
				Resolver: fetchAdUserManagedDevices,
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
						Name:     "user_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("UserID"),
					},
					{
						Name: "device_name",
						Type: schema.TypeString,
					},
					{
						Name: "managed_device_owner_type",
						Type: schema.TypeString,
					},
					{
						Name: "enrolled_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_sync_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "operating_system",
						Type: schema.TypeString,
					},
					{
						Name: "compliance_state",
						Type: schema.TypeString,
					},
					{
						Name: "jail_broken",
						Type: schema.TypeString,
					},
					{
						Name: "management_agent",
						Type: schema.TypeString,
					},
					{
						Name: "os_version",
						Type: schema.TypeString,
					},
					{
						Name: "eas_activated",
						Type: schema.TypeBool,
					},
					{
						Name:     "eas_device_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("EasDeviceID"),
					},
					{
						Name: "eas_activation_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "azure_a_d_registered",
						Type: schema.TypeBool,
					},
					{
						Name: "device_enrollment_type",
						Type: schema.TypeString,
					},
					{
						Name: "activation_lock_bypass_code",
						Type: schema.TypeString,
					},
					{
						Name: "email_address",
						Type: schema.TypeString,
					},
					{
						Name:     "azure_a_d_device_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AzureADDeviceID"),
					},
					{
						Name: "device_registration_state",
						Type: schema.TypeString,
					},
					{
						Name: "device_category_display_name",
						Type: schema.TypeString,
					},
					{
						Name: "is_supervised",
						Type: schema.TypeBool,
					},
					{
						Name: "exchange_last_successful_sync_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "exchange_access_state",
						Type: schema.TypeString,
					},
					{
						Name: "exchange_access_state_reason",
						Type: schema.TypeString,
					},
					{
						Name:     "remote_assistance_session_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("RemoteAssistanceSessionURL"),
					},
					{
						Name: "remote_assistance_session_error_details",
						Type: schema.TypeString,
					},
					{
						Name: "is_encrypted",
						Type: schema.TypeBool,
					},
					{
						Name: "user_principal_name",
						Type: schema.TypeString,
					},
					{
						Name: "model",
						Type: schema.TypeString,
					},
					{
						Name: "manufacturer",
						Type: schema.TypeString,
					},
					{
						Name: "imei",
						Type: schema.TypeString,
					},
					{
						Name: "compliance_grace_period_expiration_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "serial_number",
						Type: schema.TypeString,
					},
					{
						Name: "phone_number",
						Type: schema.TypeString,
					},
					{
						Name: "android_security_patch_level",
						Type: schema.TypeString,
					},
					{
						Name: "user_display_name",
						Type: schema.TypeString,
					},
					{
						Name:     "configuration_manager_client_enabled_features_inventory",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.Inventory"),
					},
					{
						Name:     "configuration_manager_client_enabled_features_modern_apps",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.ModernApps"),
					},
					{
						Name:     "configuration_manager_client_enabled_features_resource_access",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.ResourceAccess"),
					},
					{
						Name:     "configuration_manager_client_enabled_features_device_configuration",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.DeviceConfiguration"),
					},
					{
						Name:     "configuration_manager_client_enabled_features_compliance_policy",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.CompliancePolicy"),
					},
					{
						Name:     "configuration_manager_client_enabled_features_windows_update_for_business",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("ConfigurationManagerClientEnabledFeatures.WindowsUpdateForBusiness"),
					},
					{
						Name: "wi_fi_mac_address",
						Type: schema.TypeString,
					},
					{
						Name:     "device_health_attestation_state_last_update_date_time",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.LastUpdateDateTime"),
					},
					{
						Name:     "device_health_attestation_state_content_namespace_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.ContentNamespaceURL"),
					},
					{
						Name:     "device_health_attestation_state_device_health_attestation_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.DeviceHealthAttestationStatus"),
					},
					{
						Name:     "device_health_attestation_state_content_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.ContentVersion"),
					},
					{
						Name:     "device_health_attestation_state_issued_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.IssuedDateTime"),
					},
					{
						Name:     "device_health_attestation_state_attestation_identity_key",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.AttestationIdentityKey"),
					},
					{
						Name:     "device_health_attestation_state_reset_count",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.ResetCount"),
					},
					{
						Name:     "device_health_attestation_state_restart_count",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.RestartCount"),
					},
					{
						Name:     "device_health_attestation_state_data_excution_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.DataExcutionPolicy"),
					},
					{
						Name:     "device_health_attestation_state_bit_locker_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BitLockerStatus"),
					},
					{
						Name:     "device_health_attestation_state_boot_manager_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BootManagerVersion"),
					},
					{
						Name:     "device_health_attestation_state_code_integrity_check_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrityCheckVersion"),
					},
					{
						Name:     "device_health_attestation_state_secure_boot",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.SecureBoot"),
					},
					{
						Name:     "device_health_attestation_state_boot_debugging",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BootDebugging"),
					},
					{
						Name:     "device_health_attestation_state_operating_system_kernel_debugging",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.OperatingSystemKernelDebugging"),
					},
					{
						Name:     "device_health_attestation_state_code_integrity",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrity"),
					},
					{
						Name:     "device_health_attestation_state_test_signing",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.TestSigning"),
					},
					{
						Name:     "device_health_attestation_state_safe_mode",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.SafeMode"),
					},
					{
						Name:     "device_health_attestation_state_windows_p_e",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.WindowsPE"),
					},
					{
						Name:     "device_health_attestation_state_early_launch_anti_malware_driver_protection",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.EarlyLaunchAntiMalwareDriverProtection"),
					},
					{
						Name:     "device_health_attestation_state_virtual_secure_mode",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.VirtualSecureMode"),
					},
					{
						Name:     "device_health_attestation_state_pcr_hash_algorithm",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.PcrHashAlgorithm"),
					},
					{
						Name:     "device_health_attestation_state_boot_app_security_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BootAppSecurityVersion"),
					},
					{
						Name:     "device_health_attestation_state_boot_manager_security_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BootManagerSecurityVersion"),
					},
					{
						Name:     "device_health_attestation_state_tpm_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.TpmVersion"),
					},
					{
						Name:     "device_health_attestation_state_pcr0",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.Pcr0"),
					},
					{
						Name:     "device_health_attestation_state_secure_boot_configuration_policy_finger_print",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.SecureBootConfigurationPolicyFingerPrint"),
					},
					{
						Name:     "device_health_attestation_state_code_integrity_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.CodeIntegrityPolicy"),
					},
					{
						Name:     "device_health_attestation_state_boot_revision_list_info",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.BootRevisionListInfo"),
					},
					{
						Name:     "device_health_attestation_state_operating_system_rev_list_info",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.OperatingSystemRevListInfo"),
					},
					{
						Name:     "device_health_attestation_state_health_status_mismatch_info",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.HealthStatusMismatchInfo"),
					},
					{
						Name:     "device_health_attestation_state_health_attestation_supported_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeviceHealthAttestationState.HealthAttestationSupportedStatus"),
					},
					{
						Name: "subscriber_carrier",
						Type: schema.TypeString,
					},
					{
						Name: "meid",
						Type: schema.TypeString,
					},
					{
						Name: "total_storage_space_in_bytes",
						Type: schema.TypeBigInt,
					},
					{
						Name: "free_storage_space_in_bytes",
						Type: schema.TypeBigInt,
					},
					{
						Name: "managed_device_name",
						Type: schema.TypeString,
					},
					{
						Name: "partner_reported_threat_state",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_user_managed_device_device_action_results",
						Resolver: fetchAdUserManagedDeviceDeviceActionResults,
						Columns: []schema.Column{
							{
								Name:        "user_managed_device_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "action_name",
								Type: schema.TypeString,
							},
							{
								Name: "action_state",
								Type: schema.TypeString,
							},
							{
								Name: "start_date_time",
								Type: schema.TypeTimestamp,
							},
							{
								Name: "last_updated_date_time",
								Type: schema.TypeTimestamp,
							},
						},
					},
					{
						Name:     "azure_ad_user_managed_device_device_configuration_states",
						Resolver: fetchAdUserManagedDeviceDeviceConfigurationStates,
						Columns: []schema.Column{
							{
								Name:        "user_managed_device_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Entity.ID"),
							},
							{
								Name:     "setting_states",
								Type:     schema.TypeJSON,
								Resolver: resolveAdUserManagedDeviceDeviceConfigurationStatesSettingStates,
							},
							{
								Name: "display_name",
								Type: schema.TypeString,
							},
							{
								Name: "version",
								Type: schema.TypeBigInt,
							},
							{
								Name: "platform_type",
								Type: schema.TypeString,
							},
							{
								Name: "state",
								Type: schema.TypeString,
							},
							{
								Name: "setting_count",
								Type: schema.TypeBigInt,
							},
						},
					},
					{
						Name:     "azure_ad_user_managed_device_device_compliance_policy_states",
						Resolver: fetchAdUserManagedDeviceDeviceCompliancePolicyStates,
						Columns: []schema.Column{
							{
								Name:        "user_managed_device_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_managed_devices table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Entity.ID"),
							},
							{
								Name:     "setting_states",
								Type:     schema.TypeJSON,
								Resolver: resolveAdUserManagedDeviceDeviceCompliancePolicyStatesSettingStates,
							},
							{
								Name: "display_name",
								Type: schema.TypeString,
							},
							{
								Name: "version",
								Type: schema.TypeBigInt,
							},
							{
								Name: "platform_type",
								Type: schema.TypeString,
							},
							{
								Name: "state",
								Type: schema.TypeString,
							},
							{
								Name: "setting_count",
								Type: schema.TypeBigInt,
							},
						},
					},
				},
			},
			{
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
			},
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
			{
				Name:     "azure_ad_user_insights_trending",
				Resolver: fetchAdUserInsightsTrendings,
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
						Name: "weight",
						Type: schema.TypeFloat,
					},
					{
						Name:     "resource_visualization_title",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Title"),
					},
					{
						Name:     "resource_visualization_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Type"),
					},
					{
						Name:     "resource_visualization_media_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.MediaType"),
					},
					{
						Name:     "resource_visualization_preview_image_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewImageURL"),
					},
					{
						Name:     "resource_visualization_preview_text",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewText"),
					},
					{
						Name:     "resource_visualization_container_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerWebURL"),
					},
					{
						Name:     "resource_visualization_container_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerDisplayName"),
					},
					{
						Name:     "resource_visualization_container_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerType"),
					},
					{
						Name:     "resource_reference_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.WebURL"),
					},
					{
						Name:     "resource_reference_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.ID"),
					},
					{
						Name:     "resource_reference_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.Type"),
					},
					{
						Name: "last_modified_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Resource.ID"),
					},
				},
			},
			{
				Name:     "azure_ad_user_insights_shared",
				Resolver: fetchAdUserInsightsShareds,
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
						Name:     "last_shared_shared_by_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharedBy.DisplayName"),
					},
					{
						Name:     "last_shared_shared_by_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharedBy.ID"),
					},
					{
						Name:     "last_shared_shared_by_address",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharedBy.Address"),
					},
					{
						Name:     "last_shared_shared_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("LastShared.SharedDateTime"),
					},
					{
						Name:     "last_shared_sharing_subject",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharingSubject"),
					},
					{
						Name:     "last_shared_sharing_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharingType"),
					},
					{
						Name:     "last_shared_sharing_reference_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharingReference.WebURL"),
					},
					{
						Name:     "last_shared_sharing_reference_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharingReference.ID"),
					},
					{
						Name:     "last_shared_sharing_reference_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastShared.SharingReference.Type"),
					},
					{
						Name:     "resource_visualization_title",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Title"),
					},
					{
						Name:     "resource_visualization_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Type"),
					},
					{
						Name:     "resource_visualization_media_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.MediaType"),
					},
					{
						Name:     "resource_visualization_preview_image_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewImageURL"),
					},
					{
						Name:     "resource_visualization_preview_text",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewText"),
					},
					{
						Name:     "resource_visualization_container_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerWebURL"),
					},
					{
						Name:     "resource_visualization_container_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerDisplayName"),
					},
					{
						Name:     "resource_visualization_container_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerType"),
					},
					{
						Name:     "resource_reference_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.WebURL"),
					},
					{
						Name:     "resource_reference_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.ID"),
					},
					{
						Name:     "resource_reference_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.Type"),
					},
					{
						Name:     "last_shared_method_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LastSharedMethod.ID"),
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Resource.ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_user_insights_shared_sharing_history",
						Resolver: fetchAdUserInsightsSharedSharingHistories,
						Columns: []schema.Column{
							{
								Name:        "user_insights_shared_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_user_insights_shared table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "shared_by_display_name",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharedBy.DisplayName"),
							},
							{
								Name:     "shared_by_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharedBy.ID"),
							},
							{
								Name:     "shared_by_address",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharedBy.Address"),
							},
							{
								Name: "shared_date_time",
								Type: schema.TypeTimestamp,
							},
							{
								Name: "sharing_subject",
								Type: schema.TypeString,
							},
							{
								Name: "sharing_type",
								Type: schema.TypeString,
							},
							{
								Name:     "sharing_reference_web_url",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharingReference.WebURL"),
							},
							{
								Name:     "sharing_reference_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharingReference.ID"),
							},
							{
								Name:     "sharing_reference_type",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("SharingReference.Type"),
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_user_insights_used",
				Resolver: fetchAdUserInsightsUseds,
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
						Name:     "last_used_last_accessed_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("LastUsed.LastAccessedDateTime"),
					},
					{
						Name:     "last_used_last_modified_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("LastUsed.LastModifiedDateTime"),
					},
					{
						Name:     "resource_visualization_title",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Title"),
					},
					{
						Name:     "resource_visualization_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.Type"),
					},
					{
						Name:     "resource_visualization_media_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.MediaType"),
					},
					{
						Name:     "resource_visualization_preview_image_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewImageURL"),
					},
					{
						Name:     "resource_visualization_preview_text",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.PreviewText"),
					},
					{
						Name:     "resource_visualization_container_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerWebURL"),
					},
					{
						Name:     "resource_visualization_container_display_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerDisplayName"),
					},
					{
						Name:     "resource_visualization_container_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceVisualization.ContainerType"),
					},
					{
						Name:     "resource_reference_web_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.WebURL"),
					},
					{
						Name:     "resource_reference_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.ID"),
					},
					{
						Name:     "resource_reference_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceReference.Type"),
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Resource.ID"),
					},
				},
			},
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
func fetchAdUserManagedDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.ManagedDevices
	return nil
}
func fetchAdUserManagedDeviceDeviceActionResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceActionResults
	return nil
}
func fetchAdUserManagedDeviceDeviceConfigurationStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceConfigurationStates
	return nil
}
func resolveAdUserManagedDeviceDeviceConfigurationStatesSettingStates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.DeviceConfigurationState)
	if !ok {
		return fmt.Errorf("expected to have msgraph.DeviceConfigurationState but got %T", resource.Item)
	}
	out, err := json.Marshal(p.SettingStates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func fetchAdUserManagedDeviceDeviceCompliancePolicyStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.ManagedDevice)
	if !ok {
		return fmt.Errorf("expected to have msgraph.ManagedDevice but got %T", parent.Item)
	}
	res <- p.DeviceCompliancePolicyStates
	return nil
}
func resolveAdUserManagedDeviceDeviceCompliancePolicyStatesSettingStates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.DeviceCompliancePolicyState)
	if !ok {
		return fmt.Errorf("expected to have msgraph.DeviceCompliancePolicyState but got %T", resource.Item)
	}
	out, err := json.Marshal(p.SettingStates)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
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
func fetchAdUserDeviceManagementTroubleshootingEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.DeviceManagementTroubleshootingEvents
	return nil
}
func fetchAdUserInsightsTrendings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.Insights.Trending
	return nil
}
func fetchAdUserInsightsShareds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.Insights.Shared
	return nil
}
func fetchAdUserInsightsSharedSharingHistories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.SharedInsight)
	if !ok {
		return fmt.Errorf("expected to have msgraph.SharedInsight but got %T", parent.Item)
	}
	res <- p.SharingHistory
	return nil
}
func fetchAdUserInsightsUseds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.User)
	if !ok {
		return fmt.Errorf("expected to have msgraph.User but got %T", parent.Item)
	}
	res <- p.Insights.Used
	return nil
}
