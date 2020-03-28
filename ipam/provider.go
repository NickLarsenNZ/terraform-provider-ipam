package ipam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		//DataSourcesMap: map[string]*schema.Resource{
		//	"template_file":             dataSourceFile(),
		//	"template_cloudinit_config": dataSourceCloudinitConfig(),
		//},
		ResourcesMap: map[string]*schema.Resource{
			"ipam_dummy": resourceDummy(),
		},

		// Later, specify a provider config schema
		// Schema: nil,
		// Later: pass in a configured API client
		// ConfigureFunc: nil,
	}
}
