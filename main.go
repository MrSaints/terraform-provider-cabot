package main

import (
	"github.com/hashicorp/terraform/plugin"
	"time"
)

const (
	DEFAULT_RETRY_TIMEOUT = 1 * time.Minute
	DEFAULT_GRACE_PERIOD  = 2 * time.Second
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
