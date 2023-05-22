package task

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func tektonSidecarFields() map[string]*schema.Schema{
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the sidecar",
			Required:    true,
		},
		"image": {
			Type:        schema.TypeString,
			Description: "Image of the sidecar",
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
			Description: "Working directory to use when executing the sidecar",
			Optional:    true,
		},
		"env": {
			Type:        schema.TypeList,
			Description: "Environment variables to set for the sidecar",
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
			Description: "Volume mounts for the sidecar",
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
						Description: "Path to mount the volume",
						Required:    true,
					},
				},
			},
		},
	}
}