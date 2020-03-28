package ipam

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/nicklarsennz/terraform-provider-ipam/ipam_lib"
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
		ConfigureFunc: providerConfigure,
	}
}

type ipamClient struct {
	client ipam_lib.Client
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &ipam_lib.Config{}

	client, err := ipam_lib.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to configure IPAM Client %s", err)
	}

	return &ipamClient{
		client: *client,
	}, nil
}
