package ipam

import (
	"fmt"
	"os"
	"testing"

	"errors"
	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"io/ioutil"
)

const (
	path            = "/tmp/something.txt"
	firstline       = "alpha"
	secondline      = "bravo"
	ipamDummyConfig = `
		resource "ipam_dummy" "d" {
			path       = "%s"
			firstline  = "%s"
			secondline = "%s"
		}
	`
)

func TestIpamResourceDummy(t *testing.T) {
	expectedContent := fmt.Sprintf("%s\n%s\n", firstline, secondline)
	renderedConfig := fmt.Sprintf(ipamDummyConfig, path, firstline, secondline)

	r.UnitTest(t, r.TestCase{
		Providers: testProviders,
		Steps: []r.TestStep{
			{
				Config: renderedConfig,
				Check: func(s *terraform.State) error {
					content, err := ioutil.ReadFile(path)
					if err != nil {
						return err
						//return fmt.Errorf("was an error\ntemplate:\n%s\ngot:\n%s\nexpected:\n%s\n", renderedConfig, content, expectedContent)
					}
					if string(content) != expectedContent {
						return fmt.Errorf("template:\n%s\ngot:\n%s\nexpected:\n%s\n", renderedConfig, content, expectedContent)
					}
					return nil
				},
			},
		},
		CheckDestroy: func(*terraform.State) error {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				return nil
			}
			return errors.New("resource_dummy did not get destroyed")
		},
	})

}
