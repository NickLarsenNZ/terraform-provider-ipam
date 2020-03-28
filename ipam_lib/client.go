package ipam_lib

import (
	"fmt"
	"github.com/nicklarsennz/terraform-provider-ipam/ipam_lib/backend"
)

type Config struct {
	Backend string // DynamoDB

	*AWS_Config
}

type AWS_Config struct {
	Username   string
	Password   string
	Key_ID     string
	Access_Key string
	Profile    string
}

type Client interface {
	// Todo: define Client interface
}

func NewClient(c *Config) (*Client, error) {
	switch c.Backend {
	case backend.DYNAMO_DB:
		return nil, fmt.Errorf("Still working on the DynamoDB client")
	default:
		return nil, fmt.Errorf("Unsupported IPAM backend '%s'", c.Backend)
	}
}
