package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cabot": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CABOT_BASE_URL"); v == "" {
		t.Fatal("CABOT_BASE_URL must be set for acceptance tests")
	}
	if v := os.Getenv("CABOT_USERNAME"); v == "" {
		t.Fatal("CABOT_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("CABOT_PASSWORD"); v == "" {
		t.Fatal("CABOT_PASSWORD must be set for acceptance tests")
	}
}
