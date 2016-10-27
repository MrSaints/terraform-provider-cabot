package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
	"strings"
)

func dataSourceCabotPlugin() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCabotPluginRead,

		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCabotPluginRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	plugins, err := client.Plugins.List()
	if err != nil {
		return err
	}

	target := d.Get("title").(string)
	for _, plugin := range plugins {
		if !strings.EqualFold(plugin.Title, target) {
			continue
		}

		d.SetId(strconv.Itoa(plugin.ID))
		d.Set("title", plugin.Title)
		return nil
	}

	return fmt.Errorf("no matching plugin found: %v", target)
}
