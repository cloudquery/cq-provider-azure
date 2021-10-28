package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func AdUserInsightsTrending() *schema.Table {
	return &schema.Table{
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
	}
}

func AdUserInsightsShared() *schema.Table {
	return &schema.Table{
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
	}
}
func AdUserInsightsUsed() *schema.Table {
	return &schema.Table{
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
	}
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
