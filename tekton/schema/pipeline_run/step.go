package pipeline_run

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func tektonStepFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the step",
			Required:    true,
		},
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
		"workspaces": {
			Type:        schema.TypeList,
			Description: "Workspaces from the Task that this Step wants exclusive access to",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: workspaceUsageFields(),
			},
		},
		// "on_error": {
		// 	Type:        schema.TypeString,
		// 	Description: "Exiting behavior of a container on error",
		// 	Optional:    true,
		// },
		// "stdout_config": {
		// 	Type:        schema.TypeList,
		// 	Description: "Configuration for the stdout stream of the step",
		// 	Optional:    true,
		// 	Elem: &schema.Resource{
		// 		Schema: stepOutputConfigFields(),
		// 	},
		// },
		// "stderr_config": {
		// 	Type:        schema.TypeList,
		// 	Description: "Configuration for the stderr stream of the step",
		// 	Optional:    true,
		// 	Elem: &schema.Resource{
		// 		Schema: stepOutputConfigFields(),
		// 	},
		// },
	}
}

func stepOutputConfigFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource": {
			Type:        schema.TypeString,
			Description: "Name of the resource to store the stream in",
			Optional:    true,
		},
		"workspace": {
			Type:        schema.TypeString,
			Description: "Name of the workspace to store the stream in",
			Optional:    true,
		},
	}
}

func workspaceUsageFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the workspace",
			Required:    true,
		},
		"mount_path": {
			Type:        schema.TypeString,
			Description: "Path to mount the workspace at",
		},
	}
}
