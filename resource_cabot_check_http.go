package main

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
	"time"
)

func resourceCabotCheckHTTP() *schema.Resource {
	s := map[string]*schema.Schema{
		"endpoint": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"username": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"text_match": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"status_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  200,
		},
		"timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"verify_ssl": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}

	return &schema.Resource{
		Create: resourceCabotCheckHTTPCreate,
		Read:   resourceCabotCheckHTTPRead,
		Update: resourceCabotCheckHTTPUpdate,
		Delete: resourceCabotCheckHTTPDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: CombineWithBaseCheckSchema(s),
	}
}

func getHTTPCheckFromResourceData(d *schema.ResourceData) *cabot.HTTPCheck {
	checkRequest := &cabot.HTTPCheck{
		Endpoint:   d.Get("endpoint").(string),
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		TextMatch:  d.Get("text_match").(string),
		StatusCode: d.Get("status_code").(string),
		Timeout:    d.Get("timeout").(int),
		VerifySSL:  d.Get("verify_ssl").(bool),
	}
	checkRequest.StatusCheck = getStatusCheckFromResourceData(d)
	return checkRequest
}

func resourceCabotCheckHTTPCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	check, err := client.HTTPChecks.Create(getHTTPCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckHTTPRead(d, meta)
}

func resourceCabotCheckHTTPRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())

	return resource.Retry(DEFAULT_RETRY_TIMEOUT, func() *resource.RetryError {
		check, err := client.HTTPChecks.Get(id)
		if err != nil {
			time.Sleep(DEFAULT_GRACE_PERIOD)
			return resource.RetryableError(err)
		}

		d.SetId(strconv.Itoa(check.ID))
		setResourceDataForStatusCheck(d, check.StatusCheck)
		d.Set("endpoint", check.Endpoint)
		d.Set("username", check.Username)
		d.Set("password", check.Password)
		d.Set("text_match", check.TextMatch)
		d.Set("status_code", check.StatusCode)
		d.Set("timeout", check.Timeout)
		d.Set("verify_ssl", check.VerifySSL)

		return nil
	})
}

func resourceCabotCheckHTTPUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.HTTPChecks.Update((id), getHTTPCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckHTTPRead(d, meta)
}

func resourceCabotCheckHTTPDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.HTTPChecks.Delete(id)
}
