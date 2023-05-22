package pipeline

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func tektonSidecarFields() map[string]*schema.Schema {
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

func expandTektonSidecar(d []interface{}) []interface{} {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	sidecars := make([]interface{}, len(d))
	for i, sidecar := range d {
		sidecars[i] = expandTektonSidecarElement(sidecar)
	}
	return sidecars
}

func expandTektonSidecarElement(d interface{}) interface{} {
	sidecar := d.(map[string]interface{})
	return map[string]interface{}{
		"name":         sidecar["name"].(string),
		"image":        sidecar["image"].(string),
		"command":      sidecar["command"].([]interface{}),
		"args":         sidecar["args"].([]interface{}),
		"working_dir":  sidecar["working_dir"].(string),
		"env":          sidecar["env"].([]interface{}),
		"volume_mount": sidecar["volume_mounts"].([]interface{}),
	}
}

func flattenTektonSidecar(d []interface{}) []interface{} {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	sidecars := make([]interface{}, len(d))
	for i, sidecar := range d {
		sidecars[i] = flattenTektonSidecarElement(sidecar)
	}
	return sidecars
}

func flattenTektonSidecarElement(d interface{}) interface{} {
	sidecar := d.(map[string]interface{})
	return map[string]interface{}{
		"name":          sidecar["name"].(string),
		"image":         sidecar["image"].(string),
		"command":       sidecar["command"].([]interface{}),
		"args":          sidecar["args"].([]interface{}),
		"working_dir":   sidecar["working_dir"].(string),
		"env":           sidecar["env"].([]interface{}),
		"volume_mounts": sidecar["volume_mount"].([]interface{}),
	}
}
