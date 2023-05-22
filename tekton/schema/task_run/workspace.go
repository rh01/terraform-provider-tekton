package task_run

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func tektonWorkspaceDeclarationFields() map[string]*schema.Schema {
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
		"mount_path": {
			Type:        schema.TypeString,
			Description: "MountPath overrides the directory that the volume will be made available at",
			Optional:    true,
		},
		"read_only": {
			Type:        schema.TypeBool,
			Description: "ReadOnly dictates whether a mounted volume is writable.",
			Optional:    true,
		},
		"optional": {
			Type:        schema.TypeBool,
			Description: "Optional marks a Workspace as not being required in TaskRunRuns. By default  this field is false and so declared workspaces are required.",
			Optional:    true,
		},
	}
}
