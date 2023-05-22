package pipeline_run

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/k8s"
	"github.com/rh01/terraform-provider-tekton/tekton/utils/patch"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func TektonPipelineRunFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": k8s.NamespacedMetadataSchema("PipelineRun", false),
		"spec":     tektonPipelineRunSpecSchema(),
	}
}

func ExpandTektonPipelineRun(tkpps []interface{}) (*tektonapiv1.PipelineRun, error) {
	result := &tektonapiv1.PipelineRun{}

	if len(tkpps) == 0 || tkpps[0] == nil {
		return result, nil
	}

	in := tkpps[0].(map[string]interface{})

	if v, ok := in["metadata"].([]interface{}); ok {
		result.ObjectMeta = k8s.ExpandMetadata(v)
	}
	if v, ok := in["spec"].([]interface{}); ok {
		spec, err := expandTektonPipelineRunSpec(v)
		if err != nil {
			return result, err
		}
		result.Spec = spec
	}

	return result, nil
}

func FlattenTektonPipelineRun(in tektonapiv1.PipelineRun) []interface{} {
	att := make(map[string]interface{})

	att["metadata"] = k8s.FlattenMetadata(in.ObjectMeta)
	att["spec"] = flattenTektonPipelineRunSpec(in.Spec)

	return []interface{}{att}
}

func FromResourceData(resourceData *schema.ResourceData) (*tektonapiv1.PipelineRun, error) {
	result := &tektonapiv1.PipelineRun{}

	result.ObjectMeta = k8s.ExpandMetadata(resourceData.Get("metadata").([]interface{}))
	spec, err := expandTektonPipelineRunSpec(resourceData.Get("spec").([]interface{}))
	if err != nil {
		return result, err
	}
	result.Spec = spec

	return result, nil
}

func ToResourceData(vm tektonapiv1.PipelineRun, resourceData *schema.ResourceData) error {
	if err := resourceData.Set("metadata", k8s.FlattenMetadata(vm.ObjectMeta)); err != nil {
		return err
	}
	if err := resourceData.Set("spec", flattenTektonPipelineRunSpec(vm.Spec)); err != nil {
		return err
	}

	return nil
}

func AppendPatchOps(keyPrefix, pathPrefix string, resourceData *schema.ResourceData, ops []patch.PatchOperation) patch.PatchOperations {
	return k8s.AppendPatchOps(keyPrefix+"metadata.0.", pathPrefix+"/metadata/", resourceData, ops)
}
