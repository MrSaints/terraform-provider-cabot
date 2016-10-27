package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CABOT_BASE_URL", nil),
				Description: "The URL of the root of the target Cabot server.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CABOT_USERNAME", nil),
				Description: "The basic auth username for accessing Cabot API.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CABOT_PASSWORD", nil),
				Description: "The basic auth password for accessing Cabot API.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cabot_plugin": dataSourceCabotPlugin(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cabot_check_graphite": resourceCabotCheckGraphite(),
			"cabot_check_http":     resourceCabotCheckHTTP(),
			"cabot_check_icmp":     resourceCabotCheckICMP(),
			"cabot_check_jenkins":  resourceCabotCheckJenkins(),
			"cabot_instance":       resourceCabotInstance(),
			"cabot_service":        resourceCabotService(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		BaseURL:  d.Get("base_url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}

	return config.Client()
}
