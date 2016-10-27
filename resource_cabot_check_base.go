package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
)

// TODO: Implement `CalculatedStatus` and `Importance` enums
var baseCheckSchema = map[string]*schema.Schema{
	"active": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	// "calculated_status": &schema.Schema{
	// 	Type:     schema.TypeString,
	// 	Computed: true,
	// },
	"debounce": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  0,
	},
	"frequency": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  5,
	},
	// "importance": &schema.Schema{
	// 	Type:     schema.TypeString,
	// 	Optional: true,
	// },
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
}

func CombineWithBaseCheckSchema(s map[string]*schema.Schema) map[string]*schema.Schema {
	for k, v := range baseCheckSchema {
		s[k] = v
	}
	return s
}

func getStatusCheckFromResourceData(d *schema.ResourceData) cabot.StatusCheck {
	return cabot.StatusCheck{
		Active:    d.Get("active").(bool),
		Debounce:  d.Get("debounce").(int),
		Frequency: d.Get("frequency").(int),
		Name:      d.Get("name").(string),
	}
}

func setResourceDataForStatusCheck(d *schema.ResourceData, c cabot.StatusCheck) {
	d.Set("active", c.Active)
	d.Set("debounce", c.Debounce)
	d.Set("frequency", c.Frequency)
	d.Set("name", c.Name)
}
