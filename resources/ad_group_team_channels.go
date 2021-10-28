package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdGroupsTeamChannels() *schema.Table {
	return &schema.Table{
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
	}
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
