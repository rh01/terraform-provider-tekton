package pipeline

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/task"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func tektonPipelineSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"params": {
			Type:        schema.TypeList,
			Description: "Params is a list of input parameters required to run the task. Params must be supplied as inputs in PipelineRuns unless they declare a default value.",
			Optional:    true,
			MaxItems:    1,
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
		"tasks": {
			Type:        schema.TypeList,
			Description: "Tasks are the tasks of the build; each task is run sequentially with the source mounted into /workspace.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: task.TektonTaskFields(),
			},
		},
		"workspaces": {
			Type:        schema.TypeList,
			Description: "Workspaces are the volumes that this Pipeline requires.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonWorkspaceDeclarationFields(),
			},
		},
		"results": {
			Type:        schema.TypeList,
			Description: "Results are values that this Pipeline can output",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: tektonPipelineResultFields(),
			},
		},
	}
}

func tektonPipelineResultFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name declares the name by which a result is referenced.",
			Required:    true,
		},
		"type": {
			Type: schema.TypeString,

			Description:  "Type is the user-specified type of the result. The possible types are currently string, array and object, and string is the default.",
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"string", "array", "object"}, false),
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description is a user-facing description of the result that may be used to populate a UI.",
			Optional:    true,
		},
		"value": {
			Type:        schema.TypeList,
			Description: "Value is the expression used to retrieve the value.",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: tektonParamValueFields(),
			},
		},
	}
}

func tektonPipelineSpecSchema() *schema.Schema {
	fields := tektonPipelineSpecFields()

	return &schema.Schema{
		Type:        schema.TypeList,
		Description: fmt.Sprintf("TektonPipelineSpec describes how the proper TektonPipeline should look like."),
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
			Description: "Default is the value a parameter takes if no input value is supplied. If default is set, a Pipeline may be executed without a supplied value for the parameter.",
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

func expandTektonPipelineSpec(task []interface{}) (tektonapiv1.PipelineSpec, error) {
	ppSpec := tektonapiv1.PipelineSpec{}

	if len(task) == 0 || task[0] == nil {
		return ppSpec, nil
	}

	tktask := task[0].(map[string]interface{})

	// display_name is a user-facing name of the pipeline
	if v, ok := tktask["display_name"]; ok {
		ppSpec.DisplayName = v.(string)
	}

	// description is a user-facing description of the pipeline
	if v, ok := tktask["description"]; ok {
		ppSpec.Description = v.(string)
	}

	// params
	if v, ok := tktask["params"]; ok {
		params := v.([]interface{})
		for _, param := range params {
			p := param.(map[string]interface{})
			ppSpec.Params = append(ppSpec.Params, tektonapiv1.ParamSpec{
				Name:        p["name"].(string),
				Type:        tektonapiv1.ParamType(p["type"].(string)),
				Description: p["description"].(string),
				Default:     expandTektonParamValue(p["default"].([]interface{})),
			})
		}
	}

	//results
	if v, ok := tktask["results"]; ok {
		results := v.([]interface{})
		for _, res := range results {
			r := res.(map[string]interface{})
			ppSpec.Results = append(ppSpec.Results, tektonapiv1.PipelineResult{
				Name:        r["name"].(string),
				Type:        tektonapiv1.ResultsType(r["type"].(string)),
				Description: r["description"].(string),
				Value:       tektonapiv1.ResultValue(*expandTektonParamValue(r["value"].([]interface{}))),
			})
		}
	}

	// workspaces
	if v, ok := tktask["workspaces"]; ok {
		workspaces := v.([]interface{})
		for _, ws := range workspaces {
			w := ws.(map[string]interface{})
			ppSpec.Workspaces = append(ppSpec.Workspaces, tektonapiv1.PipelineWorkspaceDeclaration{
				Name:        w["name"].(string),
				Description: w["description"].(string),
				Optional:    w["optional"].(bool),
			})
		}
	}

	return ppSpec, nil
}

func expandTektonParamValue(value []interface{}) *tektonapiv1.ParamValue {
	if len(value) == 0 || value[0] == nil {
		return nil
	}

	v := value[0].(map[string]interface{})

	return &tektonapiv1.ParamValue{
		Type:      tektonapiv1.ParamType(v["type"].(string)),
		StringVal: v["string_val"].(string),
		ArrayVal:  expandTektonArrayValue(v["array_val"].([]interface{})),
	}
}

func expandTektonArrayValue(value []interface{}) []string {
	var result []string

	for _, v := range value {
		result = append(result, v.(string))
	}

	return result
}

func flattenTektonPipelineSpec(in tektonapiv1.PipelineSpec) []interface{} {
	att := make(map[string]interface{})
	att["display_name"] = in.DisplayName
	att["description"] = in.Description
	att["params"] = flattenTektonParamSpec(in.Params)
	att["results"] = flattenTektonPipelineResult(in.Results)
	att["workspaces"] = flattenTektonPipelineWorkspaceDeclaration(in.Workspaces)

	return []interface{}{att}
}

func flattenTektonParamSpec(in []tektonapiv1.ParamSpec) []interface{} {

	var result []interface{}

	for _, v := range in {
		att := make(map[string]interface{})
		att["name"] = v.Name
		att["type"] = v.Type
		att["description"] = v.Description
		att["default"] = flattenTektonParamValue(v.Default)

		result = append(result, att)
	}

	return result
}

func flattenTektonPipelineResult(in []tektonapiv1.PipelineResult) []interface{} {

	var result []interface{}

	for _, v := range in {
		att := make(map[string]interface{})
		att["name"] = v.Name
		att["type"] = v.Type
		att["description"] = v.Description
		att["value"] = flattenTektonParamValue(&v.Value)

		result = append(result, att)
	}

	return result
}

func flattenTektonPipelineWorkspaceDeclaration(in []tektonapiv1.PipelineWorkspaceDeclaration) []interface{} {

	var result []interface{}

	for _, v := range in {
		att := make(map[string]interface{})
		att["name"] = v.Name
		att["description"] = v.Description
		att["optional"] = v.Optional

		result = append(result, att)
	}

	return result
}

func flattenTektonParamValue(in *tektonapiv1.ParamValue) []interface{} {

	var result []interface{}

	if in == nil {
		return result
	}

	att := make(map[string]interface{})
	att["type"] = in.Type
	att["string_val"] = in.StringVal
	att["array_val"] = in.ArrayVal

	result = append(result, att)

	return result
}
