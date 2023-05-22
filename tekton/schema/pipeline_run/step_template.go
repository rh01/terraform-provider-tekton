package pipeline_run

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/k8s"
)

func tektonPipelineTaskRunTemplateFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pod_template": {
			Type:        schema.TypeList,
			Description: "PodTemplate is used to specify run specifications for all Task in pipelinerun",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: k8s.PodTemplateFields("pipelinerun"),
			},
		},
		"service_account_name": {
			Type:        schema.TypeString,
			Description: "ServiceAccountName is the name of the ServiceAccount to use to run this PipelineRun's Pods",
			Optional:    true,
		},
	}
}
