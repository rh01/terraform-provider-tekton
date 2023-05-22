package task_run

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func tektonTaskRunSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"params": {
			Type:        schema.TypeList,
			Description: "Params is a list of input parameters required to run the task. Params must be supplied as inputs in TaskRunRuns unless they declare a default value.",
			Optional:    true,
			// MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonParamSpecFields(),
			},
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "DisplayName is a user-facing name of the task that may be used to populate a UI.",
			Optional:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description is a user-facing description of the task that may be used to populate a UI.",
			Optional:    true,
		},
		"steps": {
			Type:        schema.TypeList,
			Description: "Steps are the steps of the build; each step is run sequentially with the source mounted into /workspace.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonStepFields(),
			},
		},
		// "volumes": {
		// 	Type:        schema.TypeList,
		// 	Description: "Volumes is a collection of volumes that are available to mount into the steps of the build.",
		// 	Optional:    true,
		// 	Elem: &schema.Resource{
		// 		Schema: tektonVolumeFields(),
		// 	},
		// },
		"step_template": {
			Type:        schema.TypeList,
			Description: "StepTemplate can be used as the basis for all step containers within the TaskRun, so that the steps inherit settings on the base container.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonStepTemplateFields(),
			},
		},
		"sidecars": {
			Type:        schema.TypeList,
			Description: "Sidecars are run alongside the TaskRun's step containers. They begin before the steps start and end after the steps complete.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonSidecarFields(),
			},
		},
		"workspaces": {
			Type:        schema.TypeList,
			Description: "Workspaces are the volumes that this TaskRun requires.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonWorkspaceDeclarationFields(),
			},
		},
		// "results": {
		// 	Type:        schema.TypeList,
		// 	Description: "Results are values that this TaskRun can output",
		// 	Optional:    true,
		// 	Elem: &schema.Resource{
		// 		Schema: tektonTaskRunResultFields(),
		// 	},
		// },
	}
}

func tektonTaskRunSpecSchema() *schema.Schema {
	fields := tektonTaskRunSpecFields()

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: fmt.Sprintf("TektonTaskRunSpec describes how the proper TektonTaskRun should look like."),
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: fields,
		},
	}

}

func tektonParamSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name declares the name by which a parameter is referenced.",
			Required:    true,
		},
		"type": {
			Type:         schema.TypeString,
			Description:  "Type is the user-specified type of the parameter. The possible types are currently string, array and object, and string is the default.",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"string", "array", "object"}, false),
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description is a user-facing description of the parameter that may be used to populate a UI.",
			Optional:    true,
		},
		"properties": {
			Type:        schema.TypeMap,
			Description: "Properties is the JSON Schema properties to support key-value pairs parameter.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonPropertySpecFields(),
			},
		},
		"default": {
			Type:        schema.TypeList,
			Description: "Default is the value a parameter takes if no input value is supplied. If default is set, a TaskRun may be executed without a supplied value for the parameter.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonParamValueFields(),
			},
		},
	}
}

func tektonPropertySpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Description:  "Type is the user-specified type of the parameter. The possible types are currently string, array and object, and string is the default.",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"string", "array", "object"}, false),
		},
	}
}

func tektonParamValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Description:  "Type is the user-specified type of the parameter. The possible types are currently string, array and object, and string is the default.",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"string", "array", "object"}, false),
		},
		"string_val": {
			Type:        schema.TypeString,
			Description: "StringVal is a string value.",
			Optional:    true,
		},
		"array_val": {
			Type:        schema.TypeList,
			Description: "ArrayVal is an array of strings.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"object_val": {
			Type:        schema.TypeMap,
			Description: "ObjectVal is a map of strings to strings.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func expandTektonTaskRunSpec(task []interface{}) (tektonapiv1.TaskRunSpec, error) {
	result := tektonapiv1.TaskRunSpec{}

	if len(task) == 0 || task[0] == nil {
		return result, nil
	}

	_ = task[0].(map[string]interface{})

	return result, nil
}

func flattenTektonTaskRunSpec(in tektonapiv1.TaskRunSpec) []interface{} {
	att := make(map[string]interface{})

	return []interface{}{att}
}
