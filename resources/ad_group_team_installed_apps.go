package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdGroupsTeamInstalledApps() *schema.Table {
	return &schema.Table{
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
	}
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
