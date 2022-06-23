package repository

import (
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	repositorySchema "github.com/datadrivers/terraform-provider-nexus/internal/schema/repository"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRepositoryCocoapodsProxy() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get an existing cocoapods proxy repository.",

		Read: dataSourceRepositoryCocoapodsProxyRead,

		Schema: map[string]*schema.Schema{
			// Common schemas
			"id":     common.DataSourceID,
			"name":   repositorySchema.DataSourceName,
			"online": repositorySchema.DataSourceOnline,
			// Proxy schemas
			"cleanup":        repositorySchema.DataSourceCleanup,
			"http_client":    repositorySchema.DataSourceHTTPClient,
			"negative_cache": repositorySchema.DataSourceNegativeCache,
			"proxy":          repositorySchema.DataSourceProxy,
			"routing_rule":   repositorySchema.DataSourceRoutingRule,
			"storage":        repositorySchema.DataSourceStorage,
		},
	}
}

func dataSourceRepositoryCocoapodsProxyRead(resourceData *schema.ResourceData, m interface{}) error {
	resourceData.SetId(resourceData.Get("name").(string))

	return resourceCocoapodsProxyRepositoryRead(resourceData, m)
}
