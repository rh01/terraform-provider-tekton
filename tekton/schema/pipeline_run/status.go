package pipeline_run

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/pipeline"
)

func tektonPipelineRunStatusSchema() *schema.Schema {
	fields := tektonPipelineRunStatusFields()

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Status is the current status of the PipelineRun",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: fields,
		},
	}
}

func tektonPipelineRunStatusFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start_time": {
			Type:        schema.TypeString,
			Description: "StartTime is the time the PipelineRun is actually started.",
			Optional:    true,
		},
		"completion_time": {
			Type:        schema.TypeString,
			Description: "CompletionTime is the time the PipelineRun completed.",
			Optional:    true,
		},
		"results": {
			Type:        schema.TypeList,
			Description: "Results are the list of results written out by the pipeline task's containers",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonPipelineRunResultSchema(),
			},
		},
		"pipeline_spec": {
			Type:        schema.TypeList,
			Description: "PipelineSpec contains the exact spec used to instantiate the run",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: pipeline.TektonPipelineSpecFields(),
			},
		},
		"skipped_tasks": {
			Type:        schema.TypeList,
			Description: "list of tasks that were skipped due to when expressions evaluating to false",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonSkippedTaskSchema(),
			},
		},
		"child_references": {
			Type:        schema.TypeList,
			Description: "list of TaskRun and Run names, PipelineTask names, and API versions/kinds for children of this PipelineRun.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonChildStatusReferenceSchema(),
			},
		},
		"finally_start_time": {
			Type:        schema.TypeString,
			Description: "FinallyStartTime is when all non-finally tasks have been completed and only finally tasks are being executed.",
			Optional:    true,
		},

		"span_context": {
			Type:        schema.TypeMap,
			Description: "SpanContext contains tracing span context fields",
			Optional:    true,
		},
	}
}

func tektonSkippedTaskSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name is the name of the task that was skipped",
			Optional:    true,
		},
		"reason": {
			Type:        schema.TypeString,
			Description: "Reason is the reason the task was skipped",
			Optional:    true,
		},
		"when_expression": {
			Type:        schema.TypeString,
			Description: "WhenExpression is the when expression that evaluated to false",
			Optional:    true,
		},
	}
}

func tektonPipelineRunResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name is the result's name as declared by the Pipeline",
			Optional:    true,
		},
		"value": {
			Type:        schema.TypeList,
			Description: "Value is the result returned from the execution of this PipelineRun",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonParamValueFields(),
			},
		},
	}
}

/*
*
// ChildStatusReference is used to point to the statuses of individual TaskRuns and Runs within this PipelineRun.

	type ChildStatusReference struct {
		runtime.TypeMeta `json:",inline"`
		// Name is the name of the TaskRun or Run this is referencing.
		Name string `json:"name,omitempty"`
		// PipelineTaskName is the name of the PipelineTask this is referencing.
		PipelineTaskName string `json:"pipelineTaskName,omitempty"`

		// WhenExpressions is the list of checks guarding the execution of the PipelineTask
		// +optional
		// +listType=atomic
		WhenExpressions []WhenExpression `json:"whenExpressions,omitempty"`
	}
*/
func tektonChildStatusReferenceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name is the name of the child",
			Optional:    true,
		},
		"pipeline_task_name": {
			Type:        schema.TypeString,
			Description: "PipelineTaskName is the name of the PipelineTask this is referencing",
			Optional:    true,
		},
		"when_expressions": {
			Type:        schema.TypeList,
			Description: "WhenExpressions is the list of checks guarding the execution of the PipelineTask",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonWhenExpressionSchema(),
			},
		},
	}
}

func tektonWhenExpressionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"input": {
			Type:        schema.TypeString,
			Description: "Input is the string for guard checking which can be a static input or an output from a parent Task",
			Optional:    true,
		},
		"operator": {
			Type:        schema.TypeString,
			Description: "Operator that represents an Input's relationship to the values",
			Optional:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"!",
				"=",
				"==",
				"in",
				"!=",
				"notin",
				"exists",
				"gt",
				"lt",
			}, false),
		},
		"values": {
			Type:        schema.TypeList,
			Description: "Values is an array of strings, which is compared against the input, for guard checking",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
