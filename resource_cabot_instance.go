package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
)

func resourceCabotInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceCabotInstanceCreate,
		Read:   resourceCabotInstanceRead,
		Update: resourceCabotInstanceUpdate,
		Delete: resourceCabotInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alerts": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"alerts_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"overall_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_checks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"users_to_notify": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func getInstanceFromResourceData(d *schema.ResourceData) *cabot.Instance {
	instanceRequest := &cabot.Instance{
		Address:       d.Get("address").(string),
		AlertsEnabled: d.Get("alerts_enabled").(bool),
		Name:          d.Get("name").(string),
	}

	alerts := d.Get("alerts").([]interface{})
	for _, alert := range alerts {
		instanceRequest.Alerts = append(instanceRequest.Alerts, alert.(int))
	}

	statusChecks := d.Get("status_checks").([]interface{})
	for _, check := range statusChecks {
		instanceRequest.StatusChecks = append(instanceRequest.StatusChecks, check.(int))
	}

	usersToNotify := d.Get("users_to_notify").([]interface{})
	for _, user := range usersToNotify {
		instanceRequest.UsersToNotify = append(instanceRequest.UsersToNotify, user.(int))
	}

	return instanceRequest
}

func resourceCabotInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	instance, err := client.Instances.Create(getInstanceFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(instance.ID))
	return resourceCabotInstanceRead(d, meta)
}

func resourceCabotInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	instance, err := client.Instances.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(instance.ID))
	d.Set("address", instance.Address)
	d.Set("alerts", instance.Alerts)
	d.Set("alerts_enabled", instance.AlertsEnabled)
	d.Set("name", instance.Name)
	d.Set("overall_status", instance.OverallStatus)
	d.Set("status_checks", instance.StatusChecks)
	d.Set("users_to_notify", instance.UsersToNotify)
	return nil
}

func resourceCabotInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	instance, err := client.Instances.Update((id), getInstanceFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(instance.ID))
	return resourceCabotInstanceRead(d, meta)
}

func resourceCabotInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.Instances.Delete(id)
}
