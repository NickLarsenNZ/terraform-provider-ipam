package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/nicklarsennz/terraform-provider-ipam/ipam"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ipam.Provider,
	})
}
