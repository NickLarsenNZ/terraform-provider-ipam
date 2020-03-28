package ipam

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/nicklarsennz/terraform-provider-ipam/ipam_lib"
	"testing"
)

func TestIpamNetwork(t *testing.T) {
	var network, updated_network ipam_lib.Network

	prefix := "10.0.0.0/8"
	rn := "ipam_network.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testProviders,
		CheckDestroy: testAccIpamNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpamNetworkConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccIpamNetworkExists(rn, &network),
				),
			},
			{
				Config: testAccIpamNetworkUpdateConfig(prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccIpamNetworkExists(rn, &updated_network),
					testAccIpamNetworkIDUnchanged(&network, &updated_network),
				),
			},
		},
	})
}

func testAccIpamNetworkExists(resource_name string, _ *ipam_lib.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resource_name]
		if !ok {
			return fmt.Errorf("Not found: %s", resource_name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID set")
		}

		// Todo: get actual client to check
		// See: https://github.com/terraform-providers/terraform-provider-github/blob/master/github/resource_github_issue_label_test.go#L140-L142

		return fmt.Errorf("NOT IMPLEMENTED")
	}
}

func testAccIpamNetworkIDUnchanged(network, updated_network *ipam_lib.Network) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		if *network.ID != *updated_network.ID {
			return fmt.Errorf("network was recreated")
		}
		return nil
	}
}

func testAccIpamNetworkDestroy(s *terraform.State) error {
	// Todo: get actual client to check
	// See: https://github.com/terraform-providers/terraform-provider-github/blob/master/github/resource_github_issue_label_test.go#L182

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ipam_network" {
			continue
		}

		// Todo: Look it up and ensure it doesn't still exist
		return fmt.Errorf("NOT IMPLEMENTED") //return nil
	}

	return nil
}

func testAccIpamNetworkConfig(prefix string) string {
	return fmt.Sprintf(`
		resource "ipam_network" "test" {
			description = "foo"
			prefix = "%s"
		}
	`, prefix)
}

func testAccIpamNetworkUpdateConfig(prefix string) string {
	return fmt.Sprintf(`
		resource "ipam_network" "test" {
			description = "bar"
			prefix = "%s"
		}
	`, prefix)
}
