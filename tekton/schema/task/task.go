package task

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rh01/terraform-provider-tekton/tekton/schema/k8s"
	"github.com/rh01/terraform-provider-tekton/tekton/utils/patch"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func TektonTaskFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": k8s.NamespacedMetadataSchema("Task", false),
		"spec":     tektonTaskSpecSchema(),
	}
}

func ExpandTektonTask(virtualMachine []interface{}) (*tektonapiv1.Task, error) {
	result := &tektonapiv1.Task{}

	if len(virtualMachine) == 0 || virtualMachine[0] == nil {
		return result, nil
	}

	in := virtualMachine[0].(map[string]interface{})

	if v, ok := in["metadata"].([]interface{}); ok {
		result.ObjectMeta = k8s.ExpandMetadata(v)
	}
	if v, ok := in["spec"].([]interface{}); ok {
		spec, err := expandTektonTaskSpec(v)
		if err != nil {
			return result, err
		}
		result.Spec = spec
	}

	return result, nil
}

func FlattenTektonTask(in tektonapiv1.Task) []interface{} {
	att := make(map[string]interface{})

	att["metadata"] = k8s.FlattenMetadata(in.ObjectMeta)
	att["spec"] = flattenTektonTaskSpec(in.Spec)

	return []interface{}{att}
}

func FromResourceData(resourceData *schema.ResourceData) (*tektonapiv1.Task, error) {
	result := &tektonapiv1.Task{}

	result.ObjectMeta = k8s.ExpandMetadata(resourceData.Get("metadata").([]interface{}))
	spec, err := expandTektonTaskSpec(resourceData.Get("spec").([]interface{}))
	if err != nil {
		return result, err
	}
	result.Spec = spec

	return result, nil
}

func ToResourceData(vm tektonapiv1.Task, resourceData *schema.ResourceData) error {
	if err := resourceData.Set("metadata", k8s.FlattenMetadata(vm.ObjectMeta)); err != nil {
		return err
	}
	if err := resourceData.Set("spec", flattenTektonTaskSpec(vm.Spec)); err != nil {
		return err
	}

	return nil
}

func AppendPatchOps(keyPrefix, pathPrefix string, resourceData *schema.ResourceData, ops []patch.PatchOperation) patch.PatchOperations {
	return k8s.AppendPatchOps(keyPrefix+"metadata.0.", pathPrefix+"/metadata/", resourceData, ops)
}
