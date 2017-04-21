package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mrsaints/go-cabot/cabot"
	"strconv"
)

func resourceCabotCheckJenkins() *schema.Resource {
	s := map[string]*schema.Schema{
		"max_queued_build_time": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
	}

	return &schema.Resource{
		Create: resourceCabotCheckJenkinsCreate,
		Read:   resourceCabotCheckJenkinsRead,
		Update: resourceCabotCheckJenkinsUpdate,
		Delete: resourceCabotCheckJenkinsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: CombineWithBaseCheckSchema(s),
	}
}

func getJenkinsCheckFromResourceData(d *schema.ResourceData) *cabot.JenkinsCheck {
	checkRequest := &cabot.JenkinsCheck{
		MaxQueuedBuildTime: d.Get("max_queued_build_time").(int),
	}
	checkRequest.StatusCheck = getStatusCheckFromResourceData(d)
	return checkRequest
}

func resourceCabotCheckJenkinsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	check, err := client.JenkinsChecks.Create(getJenkinsCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckJenkinsRead(d, meta)
}

func resourceCabotCheckJenkinsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.JenkinsChecks.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	setResourceDataForStatusCheck(d, check.StatusCheck)
	d.Set("max_queued_build_time", check.MaxQueuedBuildTime)
	return nil
}

func resourceCabotCheckJenkinsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)

	id, _ := strconv.Atoi(d.Id())
	check, err := client.JenkinsChecks.Update((id), getJenkinsCheckFromResourceData(d))
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(check.ID))
	return resourceCabotCheckJenkinsRead(d, meta)
}

func resourceCabotCheckJenkinsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cabot.Client)
	id, _ := strconv.Atoi(d.Id())
	return client.JenkinsChecks.Delete(id)
}
