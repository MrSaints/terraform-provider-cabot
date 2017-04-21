package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
)

func resourceCabotCheckGraphite() *schema.Resource {
	s := map[string]*schema.Schema{
		"metric": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"check_type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"value": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"expected_num_hosts": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"allowed_num_failures": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}

	return &schema.Resource{
		Create: resourceCabotCheckGraphiteCreate,
		Read:   resourceCabotCheckGraphiteRead,
		Update: resourceCabotCheckGraphiteUpdate,
		Delete: resourceCabotCheckGraphiteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: CombineWithBaseCheckSchema(s),
	}
}

func getGraphiteCheckFromResourceData(d *schema.ResourceData) *cabot.GraphiteCheck {
	checkRequest := &cabot.GraphiteCheck{
		Metric:             d.Get("metric").(string),
		CheckType:          d.Get("check_type").(string),
		Value:              d.Get("value").(string),
		ExpectedNumHosts:   d.Get("expected_num_hosts").(int),
		AllowedNumFailures: d.Get("allowed_num_failures").(int),
	}
	checkRequest.StatusCheck = getStatusCheckFromResourceData(d)
	return checkRequest
}

func resourceCabotCheckGraphiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	check, err := client.GraphiteChecks.Create(getGraphiteCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckGraphiteRead(d, meta)
}

func resourceCabotCheckGraphiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.GraphiteChecks.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	setResourceDataForStatusCheck(d, check.StatusCheck)
	d.Set("metric", check.Metric)
	d.Set("check_type", check.CheckType)
	d.Set("value", check.Value)
	d.Set("expected_num_hosts", check.ExpectedNumHosts)
	d.Set("allowed_num_failures", check.AllowedNumFailures)
	return nil
}

func resourceCabotCheckGraphiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.GraphiteChecks.Update((id), getGraphiteCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckGraphiteRead(d, meta)
}

func resourceCabotCheckGraphiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.GraphiteChecks.Delete(id)
}
