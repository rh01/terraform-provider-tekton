package tekton

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rh01/terraform-provider-tekton/tekton/client"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/task"
	"github.com/rh01/terraform-provider-tekton/tekton/utils"
	"github.com/rh01/terraform-provider-tekton/tekton/utils/patch"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

func resourceTektonTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTektonTaskCreate,
		Read:   resourceTektonTaskRead,
		Update: resourceTektonTaskUpdate,
		Delete: resourceTektonTaskDelete,
		Exists: resourceTektonTaskExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: task.TektonTaskFields(),
	}
}

func resourceTektonTaskCreate(resourceData *schema.ResourceData, meta interface{}) error {
	cli := (meta).(client.Client)

	dv, err := task.FromResourceData(resourceData)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating new tekton task: %#v", dv)
	if err := cli.CreateTask(dv); err != nil {
		return err
	}
	log.Printf("[INFO] Submitted new tekton task: %#v", dv)
	if err := task.ToResourceData(*dv, resourceData); err != nil {
		return err
	}
	resourceData.SetId(utils.BuildId(dv.ObjectMeta))

	// Wait for tekton task instance's status phase to be succeeded:
	name := dv.ObjectMeta.Name
	namespace := dv.ObjectMeta.Namespace

	stateConf := &resource.StateChangeConf{
		Pending: []string{"Creating"},
		Target:  []string{"Succeeded"},
		Timeout: resourceData.Timeout(schema.TimeoutCreate),
		Refresh: func() (interface{}, string, error) {
			var err error
			dv, err = cli.GetTask(namespace, name)
			if err != nil {
				if errors.IsNotFound(err) {
					log.Printf("[DEBUG] tekton task %s is not created yet", name)
					return dv, "Creating", nil
				}
				return dv, "", err
			}

			log.Printf("[DEBUG] tekton task %s is being created", name)
			return dv, "Creating", nil
		},
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("%s", err)
	}
	return task.ToResourceData(*dv, resourceData)
}

func resourceTektonTaskRead(resourceData *schema.ResourceData, meta interface{}) error {
	cli := (meta).(client.Client)

	namespace, name, err := utils.IdParts(resourceData.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Reading tekton task %s", name)

	dv, err := cli.GetTask(namespace, name)
	if err != nil {
		log.Printf("[DEBUG] Received error: %#v", err)
		return err
	}
	log.Printf("[INFO] Received tekton task: %#v", dv)

	return task.ToResourceData(*dv, resourceData)
}

func resourceTektonTaskUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	cli := (meta).(client.Client)

	namespace, name, err := utils.IdParts(resourceData.Id())
	if err != nil {
		return err
	}

	ops := task.AppendPatchOps("", "", resourceData, make([]patch.PatchOperation, 0, 0))
	data, err := ops.MarshalJSON()
	if err != nil {
		return fmt.Errorf("[DEBUG] Failed to marshal update operations: %s", err)
	}

	log.Printf("[INFO] Updating tekton task: %s", ops)
	out := &tektonapiv1.Task{}
	if err := cli.UpdateTask(namespace, name, out, data); err != nil {
		return err
	}

	log.Printf("[INFO] Submitted updated tekton task: %#v", out)

	return resourceTektonTaskRead(resourceData, meta)
}

func resourceTektonTaskDelete(resourceData *schema.ResourceData, meta interface{}) error {
	cli := (meta).(client.Client)

	namespace, name, err := utils.IdParts(resourceData.Id())
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting tekton task: %#v", name)
	if err := cli.DeleteTask(namespace, name); err != nil {
		return err
	}

	// Wait for tekton task instance to be removed:
	stateConf := &resource.StateChangeConf{
		Pending: []string{"Deleting"},
		Timeout: resourceData.Timeout(schema.TimeoutDelete),
		Refresh: func() (interface{}, string, error) {
			dv, err := cli.GetTask(namespace, name)
			if err != nil {
				if errors.IsNotFound(err) {
					return nil, "", nil
				}
				return dv, "", err
			}

			log.Printf("[DEBUG] tekton task %s is being deleted", dv.GetName())
			return dv, "Deleting", nil
		},
	}

	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("%s", err)
	}

	log.Printf("[INFO] tekton task %s deleted", name)

	resourceData.SetId("")
	return nil
}

func resourceTektonTaskExists(resourceData *schema.ResourceData, meta interface{}) (bool, error) {
	cli := (meta).(client.Client)

	namespace, name, err := utils.IdParts(resourceData.Id())
	if err != nil {
		return false, err
	}

	log.Printf("[INFO] Checking tekton task %s", name)
	if _, err := cli.GetTask(namespace, name); err != nil {
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			return false, nil
		}
		log.Printf("[DEBUG] Received error: %#v", err)
		return true, err
	}
	return true, nil
}
