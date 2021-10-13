package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdApplications() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_applications",
		Resolver:     fetchAdApplications,
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
				Name:     "api_accept_mapped_claims",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("API.AcceptMappedClaims"),
			},
			{
				Name:     "api_known_client_applications",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("API.KnownClientApplications"),
			},
			{
				Name:     "api_requested_access_token_version",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("API.RequestedAccessTokenVersion"),
			},
			{
				Name:     "app_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppID"),
			},
			{
				Name:     "application_template_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationTemplateID"),
			},
			{
				Name: "is_fallback_public_client",
				Type: schema.TypeBool,
			},
			{
				Name: "identifier_uris",
				Type: schema.TypeStringArray,
			},
			{
				Name: "created_date_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "public_client_redirect_uris",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PublicClient.RedirectUris"),
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "group_membership_claims",
				Type: schema.TypeString,
			},
			{
				Name:     "info_logo_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Info.LogoURL"),
			},
			{
				Name:     "info_marketing_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Info.MarketingURL"),
			},
			{
				Name:     "info_privacy_statement_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Info.PrivacyStatementURL"),
			},
			{
				Name:     "info_support_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Info.SupportURL"),
			},
			{
				Name:     "info_terms_of_service_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Info.TermsOfServiceURL"),
			},
			{
				Name: "is_device_only_auth_supported",
				Type: schema.TypeBool,
			},
			{
				Name:     "oauth_2_require_post_response",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("OAuth2RequirePostResponse"),
			},
			{
				Name:     "parental_control_settings_countries_blocked_for_minors",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ParentalControlSettings.CountriesBlockedForMinors"),
			},
			{
				Name:     "parental_control_settings_legal_age_group_rule",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParentalControlSettings.LegalAgeGroupRule"),
			},
			{
				Name: "publisher_domain",
				Type: schema.TypeString,
			},
			{
				Name: "sign_in_audience",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "token_encryption_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TokenEncryptionKeyID"),
			},
			{
				Name:     "web_home_page_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Web.HomePageURL"),
			},
			{
				Name:     "web_redirect_uris",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Web.RedirectUris"),
			},
			{
				Name:     "web_logout_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Web.LogoutURL"),
			},
			{
				Name:     "web_implicit_grant_settings_enable_id_token_issuance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Web.ImplicitGrantSettings.EnableIDTokenIssuance"),
			},
			{
				Name:     "web_implicit_grant_settings_enable_access_token_issuance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Web.ImplicitGrantSettings.EnableAccessTokenIssuance"),
			},
			{
				Name:     "created_on_behalf_of_entity_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedOnBehalfOf.Entity.ID"),
			},
			{
				Name:     "created_on_behalf_of_deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOnBehalfOf.DeletedDateTime"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_ad_application_add_ins",
				Resolver: fetchAdApplicationAddIns,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_application_add_in_properties",
						Resolver: fetchAdApplicationAddInProperties,
						Columns: []schema.Column{
							{
								Name:        "application_add_in_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_application_add_ins table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name: "key",
								Type: schema.TypeString,
							},
							{
								Name: "value",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_application_api_pre_authorized_applications",
				Resolver: fetchAdApplicationApiPreAuthorizedApplications,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "app_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AppID"),
					},
					{
						Name:     "delegated_permission_ids",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("DelegatedPermissionIDs"),
					},
				},
			},
			{
				Name:     "azure_ad_application_api_oauth_2_permission_scopes",
				Resolver: fetchAdApplicationApiOauth2PermissionScopes,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "admin_consent_description",
						Type: schema.TypeString,
					},
					{
						Name: "admin_consent_display_name",
						Type: schema.TypeString,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "is_enabled",
						Type: schema.TypeBool,
					},
					{
						Name: "origin",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "user_consent_description",
						Type: schema.TypeString,
					},
					{
						Name: "user_consent_display_name",
						Type: schema.TypeString,
					},
					{
						Name: "value",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_application_app_roles",
				Resolver: fetchAdApplicationAppRoles,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "allowed_member_types",
						Type: schema.TypeStringArray,
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
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "is_enabled",
						Type: schema.TypeBool,
					},
					{
						Name: "origin",
						Type: schema.TypeString,
					},
					{
						Name: "value",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_application_key_credentials",
				Resolver: fetchAdApplicationKeyCredentials,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "custom_key_identifier",
						Type:     schema.TypeByteArray,
						Resolver: resolveAdApplicationKeyCredentialsCustomKeyIdentifier,
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "end_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyID"),
					},
					{
						Name: "start_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "usage",
						Type: schema.TypeString,
					},
					{
						Name:     "key",
						Type:     schema.TypeByteArray,
						Resolver: resolveAdApplicationKeyCredentialsKey,
					},
				},
			},
			{
				Name:     "azure_ad_application_optional_claims_id_token",
				Resolver: fetchAdApplicationOptionalClaimsIdTokens,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "source",
						Type: schema.TypeString,
					},
					{
						Name: "essential",
						Type: schema.TypeBool,
					},
					{
						Name: "additional_properties",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "azure_ad_application_optional_claims_access_token",
				Resolver: fetchAdApplicationOptionalClaimsAccessTokens,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "source",
						Type: schema.TypeString,
					},
					{
						Name: "essential",
						Type: schema.TypeBool,
					},
					{
						Name: "additional_properties",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "azure_ad_application_optional_claims_saml2_token",
				Resolver: fetchAdApplicationOptionalClaimsSaml2Tokens,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "source",
						Type: schema.TypeString,
					},
					{
						Name: "essential",
						Type: schema.TypeBool,
					},
					{
						Name: "additional_properties",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "azure_ad_application_password_credentials",
				Resolver: fetchAdApplicationPasswordCredentials,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "custom_key_identifier",
						Type:     schema.TypeByteArray,
						Resolver: resolveAdApplicationPasswordCredentialsCustomKeyIdentifier,
					},
					{
						Name: "display_name",
						Type: schema.TypeString,
					},
					{
						Name: "end_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyID"),
					},
					{
						Name: "start_date_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "secret_text",
						Type: schema.TypeString,
					},
					{
						Name: "hint",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_ad_application_required_resource_access",
				Resolver: fetchAdApplicationRequiredResourceAccesses,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "resource_app_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ResourceAppID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "azure_ad_application_required_resource_access_resource_access",
						Resolver: fetchAdApplicationRequiredResourceAccessResourceAccesses,
						Columns: []schema.Column{
							{
								Name:        "application_required_resource_access_cq_id",
								Description: "Unique CloudQuery ID of azure_ad_application_required_resource_access table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("ID"),
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "azure_ad_application_extension_properties",
				Resolver: fetchAdApplicationExtensionProperties,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "directory_object_entity_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DirectoryObject.Entity.ID"),
					},
					{
						Name:     "directory_object_deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("DirectoryObject.DeletedDateTime"),
					},
					{
						Name: "app_display_name",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "data_type",
						Type: schema.TypeString,
					},
					{
						Name: "is_synced_from_on_premises",
						Type: schema.TypeBool,
					},
					{
						Name: "target_objects",
						Type: schema.TypeStringArray,
					},
				},
			},
			{
				Name:     "azure_ad_application_owners",
				Resolver: fetchAdApplicationOwners,
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique CloudQuery ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "entity_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Entity.ID"),
					},
					{
						Name: "deleted_date_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAdApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Graph
	response, err := svc.Applications().Request().Get(ctx)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
func fetchAdApplicationAddIns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.AddIns
	return nil
}
func fetchAdApplicationAddInProperties(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.AddIn)
	if !ok {
		return fmt.Errorf("expected to have msgraph.AddIn but got %T", parent.Item)
	}
	res <- p.Properties
	return nil
}
func fetchAdApplicationApiPreAuthorizedApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.API.PreAuthorizedApplications
	return nil
}
func fetchAdApplicationApiOauth2PermissionScopes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.API.OAuth2PermissionScopes
	return nil
}
func fetchAdApplicationAppRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.AppRoles
	return nil
}
func fetchAdApplicationKeyCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.KeyCredentials
	return nil
}
func resolveAdApplicationKeyCredentialsCustomKeyIdentifier(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.KeyCredential)
	if !ok {
		return fmt.Errorf("expected to have msgraph.KeyCredential but got %T", resource.Item)
	}

	return resource.Set(c.Name, []byte(*p.CustomKeyIdentifier))
}
func resolveAdApplicationKeyCredentialsKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.KeyCredential)
	if !ok {
		return fmt.Errorf("expected to have msgraph.KeyCredential but got %T", resource.Item)
	}

	return resource.Set(c.Name, []byte(*p.Key))
}
func fetchAdApplicationOptionalClaimsIdTokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.OptionalClaims.IDToken
	return nil
}
func fetchAdApplicationOptionalClaimsAccessTokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.OptionalClaims.AccessToken
	return nil
}
func fetchAdApplicationOptionalClaimsSaml2Tokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.OptionalClaims.Saml2Token
	return nil
}
func fetchAdApplicationPasswordCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.PasswordCredentials
	return nil
}
func resolveAdApplicationPasswordCredentialsCustomKeyIdentifier(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(msgraph.PasswordCredential)
	if !ok {
		return fmt.Errorf("expected to have msgraph.PasswordCredential but got %T", resource.Item)
	}

	return resource.Set(c.Name, []byte(*p.CustomKeyIdentifier))
}
func fetchAdApplicationRequiredResourceAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.RequiredResourceAccess
	return nil
}
func fetchAdApplicationRequiredResourceAccessResourceAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.RequiredResourceAccess)
	if !ok {
		return fmt.Errorf("expected to have msgraph.RequiredResourceAccess but got %T", parent.Item)
	}
	res <- p.ResourceAccess
	return nil
}
func fetchAdApplicationExtensionProperties(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.ExtensionProperties
	return nil
}
func fetchAdApplicationOwners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(msgraph.Application)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Application but got %T", parent.Item)
	}
	res <- p.Owners
	return nil
}
