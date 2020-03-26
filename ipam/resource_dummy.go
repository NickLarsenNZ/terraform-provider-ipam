package ipam

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDummy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDummyCreate,
		Read:   resourceDummyRead,
		Update: resourceDummyUpdate,
		Delete: resourceDummyDelete,
		Exists: resourceDummyExists,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path of the file we will store stuff in",
			},
			"mtime": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "File modified time (computed)",
			},
		},

		DeprecationMessage: "This is not a real resource",
	}
}

func resourceDummyCreate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Create not implemented")
}

func resourceDummyRead(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Read not implemented")
}

// https://www.terraform.io/docs/plugins/provider.html#resource-data
func resourceDummyUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Update not implemented")
}

func resourceDummyDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Delete not implemented")
}

func resourceDummyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, fmt.Errorf("Exists not implemented")
}
