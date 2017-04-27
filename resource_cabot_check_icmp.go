package main

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
	"time"
)

func resourceCabotCheckICMP() *schema.Resource {
	return &schema.Resource{
		Create: resourceCabotCheckICMPCreate,
		Read:   resourceCabotCheckICMPRead,
		Update: resourceCabotCheckICMPUpdate,
		Delete: resourceCabotCheckICMPDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: baseCheckSchema,
	}
}

func getICMPCheckFromResourceData(d *schema.ResourceData) *cabot.ICMPCheck {
	checkRequest := &cabot.ICMPCheck{}
	checkRequest.StatusCheck = getStatusCheckFromResourceData(d)
	return checkRequest
}

func resourceCabotCheckICMPCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	check, err := client.ICMPChecks.Create(getICMPCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckICMPRead(d, meta)
}

func resourceCabotCheckICMPRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())

	return resource.Retry(DEFAULT_RETRY_TIMEOUT, func() *resource.RetryError {
		check, err := client.ICMPChecks.Get(id)
		if err != nil {
			time.Sleep(DEFAULT_GRACE_PERIOD)
			return resource.RetryableError(err)
		}

		d.SetId(strconv.Itoa(check.ID))
		setResourceDataForStatusCheck(d, check.StatusCheck)

		return nil
	})
}

func resourceCabotCheckICMPUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.ICMPChecks.Update((id), getICMPCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckICMPRead(d, meta)
}

func resourceCabotCheckICMPDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO: handle deletion when a linked instance is deleted
	// That is, the ICMP check will be removed automatically resulting in a
	// 404 not found error

	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.ICMPChecks.Delete(id)
}
