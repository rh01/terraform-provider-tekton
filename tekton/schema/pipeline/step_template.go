package pipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func tektonStepTemplateFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"image": {
			Type:        schema.TypeString,
			Description: "Image of the step",
			Optional:    true,
		},
		"command": {
			Type:        schema.TypeList,
			Description: "Command to execute in the container",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"args": {
			Type:        schema.TypeList,
			Description: "Arguments to the entrypoint",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"working_dir": {
			Type:        schema.TypeString,
			Description: "Working directory to use when executing the step",
			Optional:    true,
		},
		"env": {
			Type:        schema.TypeList,
			Description: "Environment variables to set for the step",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Description: "Name of the environment variable",
						Required:    true,
					},
					"value": {
						Type:        schema.TypeString,
						Description: "Value of the environment variable",
						Required:    true,
					},
				},
			},
		},
		"volume_mounts": {
			Type:        schema.TypeList,
			Description: "Volume mounts for the step",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Description: "Name of the volume mount",
						Required:    true,
					},
					"mount_path": {
						Type:        schema.TypeString,
						Description: "Path to mount the volume at",
						Required:    true,
					},
					"read_only": {
						Type:        schema.TypeBool,
						Description: "Whether the volume should be mounted read-only",
						Optional:    true,
					},
				},
			},
		},
		"image_pull_policy": {
			Type:        schema.TypeString,
			Description: "Image pull policy for the step",
			Optional:    true,
		},

		"script": {
			Type:        schema.TypeString,
			Description: "Contents of an executable file to execute",
			Optional:    true,
		},
		"timeout": {
			Type:        schema.TypeString,
			Description: "Time after which the step times out",
			Optional:    true,
		},
	}
}

func expandTektonStepTemplate(d []interface{}) []interface{} {
	if len(d) == 0 || d[0] == nil {
		return nil
	}

	stepTemplate := d[0].(map[string]interface{})

	expanded := make([]interface{}, 1)
	expanded[0] = map[string]interface{}{
		"image":             stepTemplate["image"],
		"command":           stepTemplate["command"],
		"args":              stepTemplate["args"],
		"working_dir":       stepTemplate["working_dir"],
		"env":               stepTemplate["env"],
		"volume_mounts":     stepTemplate["volume_mounts"],
		"image_pull_policy": stepTemplate["image_pull_policy"],
		"script":            stepTemplate["script"],
		"timeout":           stepTemplate["timeout"],
	}

	return expanded
}

func flattenTektonStepTemplate(d []interface{}) []interface{} {
	if len(d) == 0 || d[0] == nil {
		return nil
	}

	stepTemplate := d[0].(map[string]interface{})

	flattened := make([]interface{}, 1)
	flattened[0] = map[string]interface{}{
		"image":             stepTemplate["image"],
		"command":           stepTemplate["command"],
		"args":              stepTemplate["args"],
		"working_dir":       stepTemplate["working_dir"],
		"env":               stepTemplate["env"],
		"volume_mounts":     stepTemplate["volume_mounts"],
		"image_pull_policy": stepTemplate["image_pull_policy"],
		"script":            stepTemplate["script"],
		"timeout":           stepTemplate["timeout"],
	}

	return flattened
}
