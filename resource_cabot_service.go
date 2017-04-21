package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
)

func resourceCabotService() *schema.Resource {
	return &schema.Resource{
		Create: resourceCabotServiceCreate,
		Read:   resourceCabotServiceRead,
		Update: resourceCabotServiceUpdate,
		Delete: resourceCabotServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"alerts": &schema.Schema{
				Type:     schema.TypeList,
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
			"instances": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"users_to_notify": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func getServiceFromResourceData(d *schema.ResourceData) *cabot.Service {
	serviceRequest := &cabot.Service{
		AlertsEnabled: d.Get("alerts_enabled").(bool),
		Name:          d.Get("name").(string),
		URL:           d.Get("url").(string),
	}

	alerts := d.Get("alerts").([]interface{})
	for _, alert := range alerts {
		serviceRequest.Alerts = append(serviceRequest.Alerts, alert.(int))
	}

	instances := d.Get("instances").([]interface{})
	for _, instance := range instances {
		serviceRequest.Instances = append(serviceRequest.Instances, instance.(int))
	}

	// TODO: sort
	statusChecks := d.Get("status_checks").([]interface{})
	for _, check := range statusChecks {
		serviceRequest.StatusChecks = append(serviceRequest.StatusChecks, check.(int))
	}

	usersToNotify := d.Get("users_to_notify").([]interface{})
	for _, user := range usersToNotify {
		serviceRequest.UsersToNotify = append(serviceRequest.UsersToNotify, user.(int))
	}

	return serviceRequest
}

func resourceCabotServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	service, err := client.Services.Create(getServiceFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(service.ID))
	return resourceCabotServiceRead(d, meta)
}

func resourceCabotServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	service, err := client.Services.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(service.ID))
	d.Set("alerts", service.Alerts)
	d.Set("alerts_enabled", service.AlertsEnabled)
	d.Set("instances", service.Instances)
	d.Set("name", service.Name)
	d.Set("overall_status", service.OverallStatus)
	d.Set("status_checks", service.StatusChecks)
	d.Set("url", service.URL)
	d.Set("users_to_notify", service.UsersToNotify)
	return nil
}

func resourceCabotServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	service, err := client.Services.Update((id), getServiceFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(service.ID))
	return resourceCabotServiceRead(d, meta)
}

func resourceCabotServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.Services.Delete(id)
}
