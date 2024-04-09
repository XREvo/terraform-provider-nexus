package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ResourceName = &schema.Schema{
		Description: "A unique identifier for this repository",
		Required:    true,
		Type:        schema.TypeString,
	}
	DataSourceName   = ResourceName
	IgnoreIfNotFound = &schema.Schema{
		Description: "Return a null object if not found on Nexus",
		Optional:    true,
		Type:        schema.TypeBool,
		Default:     false,
	}
)
