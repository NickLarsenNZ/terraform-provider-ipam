package ipam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"

	"testing"
)

var testBackend = "DynamoDB"
var testBackendDynamoDbTable = os.Getenv("IPAM_BACKEND_DYNAMODB_TABLE")

var testAccProvider = Provider().(*schema.Provider)
var testProviders = map[string]terraform.ResourceProvider{
	"ipam": testAccProvider,
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// Ensure Environment variables are set for Acceptance Testing
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IPAM_BACKEND_DYNAMODB_TABLE"); v == "" {
		t.Fatal("IPAM_BACKEND_DYNAMODB_TABLE must be set for acceptance tests")
	}
}
