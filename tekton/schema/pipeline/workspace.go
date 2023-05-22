package pipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func tektonPipelineWorkspaceDeclarationFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the workspace",
			Required:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "Description is an optional human readable description of this volume.",
			Optional:    true,
		},
		"optional": {
			Type:        schema.TypeBool,
			Description: "Optional marks a Workspace as not being required in PipelineRuns. By default  this field is false and so declared workspaces are required.",
			Optional:    true,
		},
	}
}

func expandPipelineWorkspaceDeclaration(in []interface{}) []tektonapiv1.WorkspaceDeclaration {
	if in == nil {
		return nil
	}
	out := make([]tektonapiv1.WorkspaceDeclaration, 0)
	for _, v := range in {
		m := v.(map[string]interface{})
		out = append(out, tektonapiv1.WorkspaceDeclaration{
			Name:        m["name"].(string),
			Description: m["description"].(string),
			MountPath:   m["mount_path"].(string),
			ReadOnly:    m["read_only"].(bool),
			Optional:    m["optional"].(bool),
		})
	}
	return out
}

func flattenPipelineWorkspaceDeclaration(in []tektonapiv1.WorkspaceDeclaration) []interface{} {
	if in == nil {
		return nil
	}
	out := make([]interface{}, 0)
	for _, v := range in {
		m := make(map[string]interface{})
		m["name"] = v.Name
		m["description"] = v.Description
		m["mount_path"] = v.MountPath
		m["read_only"] = v.ReadOnly
		m["optional"] = v.Optional
		out = append(out, m)
	}
	return out
}
