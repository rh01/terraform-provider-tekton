package pipeline_run

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/pipeline"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

/*
*

// PipelineRunSpec defines the desired state of PipelineRun

	type PipelineRunSpec struct {
		// +optional
		PipelineRef *PipelineRef `json:"pipelineRef,omitempty"`
		// +optional
		PipelineSpec *PipelineSpec `json:"pipelineSpec,omitempty"`
		// Params is a list of parameter names and values.
		// +listType=atomic
		Params Params `json:"params,omitempty"`

		// Used for cancelling a pipelinerun (and maybe more later on)
		// +optional
		Status PipelineRunSpecStatus `json:"status,omitempty"`
		// Time after which the Pipeline times out.
		// Currently three keys are accepted in the map
		// pipeline, tasks and finally
		// with Timeouts.pipeline >= Timeouts.tasks + Timeouts.finally
		// +optional
		Timeouts *TimeoutFields `json:"timeouts,omitempty"`

		// TaskRunTemplate represent template of taskrun
		// +optional
		TaskRunTemplate PipelineTaskRunTemplate `json:"taskRunTemplate,omitempty"`

		// Workspaces holds a set of workspace bindings that must match names
		// with those declared in the pipeline.
		// +optional
		// +listType=atomic
		Workspaces []WorkspaceBinding `json:"workspaces,omitempty"`
		// TaskRunSpecs holds a set of runtime specs
		// +optional
		// +listType=atomic
		TaskRunSpecs []PipelineTaskRunSpec `json:"taskRunSpecs,omitempty"`
	}
*/
func tektonPipelineRunSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pipeline_ref": {
			Type:        schema.TypeList,
			Description: "PipelineRef is a reference to a Pipeline that is used to create PipelineRuns.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonPipelineRefFields(),
			},
		},
		"pipeline_spec": {
			Type:        schema.TypeList,
			Description: "PipelineSpec is a specification of the desired behavior of a Pipeline.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: pipeline.TektonPipelineSpecFields(),
			},
		},
		"params": {
			Type:        schema.TypeList,
			Description: "Params is a list of input parameters required to run the task. Params must be supplied as inputs in PipelineRunRuns unless they declare a default value.",
			Optional:    true,
			// MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonParamSpecFields(),
			},
		},
		"status": {
			Type:        schema.TypeString,
			Description: "Status is the status of the PipelineRunSpec.",
			Optional:    true,
			Default:     "Cancelled",
			ValidateFunc: validation.StringInSlice([]string{
				"Cancelled",
				"CancelledRunFinally",
				"StoppedRunFinally",
				"PipelineRunPending",
			}, false),
		},
		"task_run_template": {
			Type:        schema.TypeList,
			Description: "StepTemplate can be used as the basis for all step containers within the PipelineRun, so that the steps inherit settings on the base container.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonPipelineTaskRunTemplateFields(),
			},
		},
	}
}

func tektonPipelineRefFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name is the name of the referenced pipeline.",
			Required:    true,
		},
		"api_version": {
			Type:        schema.TypeString,
			Description: "API version of the referent",
			Required:    true,
		},
	}
}

func tektonPipelineRunSpecSchema() *schema.Schema {
	fields := tektonPipelineRunSpecFields()

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: fmt.Sprintf("TektonPipelineRunSpec describes how the proper TektonPipelineRun should look like."),
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
			Description: "Default is the value a parameter takes if no input value is supplied. If default is set, a PipelineRun may be executed without a supplied value for the parameter.",
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

func expandTektonPipelineRunSpec(task []interface{}) (tektonapiv1.PipelineRunSpec, error) {
	result := tektonapiv1.PipelineRunSpec{}

	if len(task) == 0 || task[0] == nil {
		return result, nil
	}

	_ = task[0].(map[string]interface{})

	return result, nil
}

func flattenTektonPipelineRunSpec(in tektonapiv1.PipelineRunSpec) []interface{} {
	att := make(map[string]interface{})

	return []interface{}{att}
}
