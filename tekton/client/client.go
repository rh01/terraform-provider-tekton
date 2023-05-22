/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	tektonapiv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	pkgApi "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
)

//go:generate mockgen -source=./client.go -destination=./mock/client_generated.go -package=mock

type Client interface {
	// Task CRUD operations
	CreateTask(obj *tektonapiv1.Task) error
	GetTask(namespace string, name string) (*tektonapiv1.Task, error)
	UpdateTask(namespace string, name string, obj *tektonapiv1.Task, data []byte) error
	DeleteTask(namespace string, name string) error

	// TaskRun CRUD operations
	CreateTaskRun(obj *tektonapiv1.TaskRun) error
	GetTaskRun(namespace string, name string) (*tektonapiv1.TaskRun, error)
	UpdateTaskRun(namespace string, name string, obj *tektonapiv1.TaskRun, data []byte) error
	DeleteTaskRun(namespace string, name string) error

	// Pipeline CRUD operations
	CreatePipeline(obj *tektonapiv1.Pipeline) error
	GetPipeline(namespace string, name string) (*tektonapiv1.Pipeline, error)
	UpdatePipeline(namespace string, name string, obj *tektonapiv1.Pipeline, data []byte) error
	DeletePipeline(namespace string, name string) error

	// PipelineRun CRUD operations
	CreatePipelineRun(obj *tektonapiv1.PipelineRun) error
	GetPipelineRun(namespace string, name string) (*tektonapiv1.PipelineRun, error)
	UpdatePipelineRun(namespace string, name string, obj *tektonapiv1.PipelineRun, data []byte) error
	DeletePipelineRun(namespace string, name string) error
}

type client struct {
	dynamicClient dynamic.Interface
}

// CreatePipeline implements Client
func (c *client) CreatePipeline(obj *tektonapiv1.Pipeline) error {
	pipelineUpdateTypeMeta(obj)
	return c.createResource(obj, obj.Namespace, pipelineRes())
}

// CreateTaskRun implements Client
func (c *client) CreateTaskRun(obj *tektonapiv1.TaskRun) error {
	taskrunUpdateTypeMeta(obj)
	return c.createResource(obj, obj.Namespace, taskRunRes())
}

// DeletePipeline implements Client
func (c *client) DeletePipeline(namespace string, name string) error {
	return c.deleteResource(namespace, name, pipelineRes())
}

// DeleteTaskRun implements Client
func (c *client) DeleteTaskRun(namespace string, name string) error {
	return c.deleteResource(namespace, name, taskRunRes())
}

// GetPipeline implements Client
func (c *client) GetPipeline(namespace string, name string) (*tektonapiv1.Pipeline, error) {
	obj, err := c.getResource(namespace, name, pipelineRes())
	if err != nil {
		return nil, err
	}
	pipeline := &tektonapiv1.Pipeline{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, pipeline)
	if err != nil {
		return nil, err
	}
	return pipeline, nil
}

// GetTaskRun implements Client
func (c *client) GetTaskRun(namespace string, name string) (*tektonapiv1.TaskRun, error) {
	obj, err := c.getResource(namespace, name, taskRunRes())
	if err != nil {
		return nil, err
	}
	taskrun := &tektonapiv1.TaskRun{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, taskrun)
	if err != nil {
		return nil, err
	}
	return taskrun, nil
}

// UpdatePipeline implements Client
func (c *client) UpdatePipeline(namespace string, name string, obj *tektonapiv1.Pipeline, data []byte) error {
	pipelineUpdateTypeMeta(obj)
	return c.updateResource(namespace, name, pipelineRes(), obj, data)
}

// UpdateTaskRun implements Client
func (c *client) UpdateTaskRun(namespace string, name string, obj *tektonapiv1.TaskRun, data []byte) error {
	taskrunUpdateTypeMeta(obj)
	return c.updateResource(namespace, name, taskRunRes(), obj, data)
}

// New creates our client wrapper object for the actual kubeVirt and kubernetes clients we use.
func NewClient(cfg *restclient.Config) (Client, diag.Diagnostics) {
	var diags diag.Diagnostics

	result := &client{}
	c, err := dynamic.NewForConfig(cfg)
	if err != nil {
		msg := fmt.Sprintf("Failed to create client, with error: %v", err)
		log.Printf("[Error] %s", msg)
		return nil, diag.FromErr(fmt.Errorf(msg))
	}
	result.dynamicClient = c
	return result, diags
}

func taskrunUpdateTypeMeta(obj *tektonapiv1.TaskRun) {
	obj.APIVersion = "tekton.dev/v1alpha1"
	obj.Kind = "TaskRun"
}

func taskRunRes() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  "v1alpha1",
		Resource: "taskruns",
	}
}

func pipelineUpdateTypeMeta(obj *tektonapiv1.Pipeline) {
	obj.APIVersion = "tekton.dev/v1alpha1"
	obj.Kind = "Pipeline"
}

func pipelineRes() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  "v1alpha1",
		Resource: "pipelines",
	}
}

// PipelineRun CRUD operations

// CreatePipelineRun implements Client
func (c *client) CreatePipelineRun(obj *tektonapiv1.PipelineRun) error {
	pipelineRunUpdateTypeMeta(obj)
	return c.createResource(obj, obj.Namespace, pipelineRunRes())
}

// DeletePipelineRun implements Client
func (c *client) DeletePipelineRun(namespace string, name string) error {
	return c.deleteResource(namespace, name, pipelineRunRes())
}

// GetPipelineRun implements Client
func (c *client) GetPipelineRun(namespace string, name string) (*tektonapiv1.PipelineRun, error) {
	var obj tektonapiv1.PipelineRun
	resp, err := c.getResource(namespace, name, pipelineRunRes())
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("[Warning] PipelineRun %s not found (namespace=%s)", name, namespace)
			return nil, err
		}
		msg := fmt.Sprintf("Failed to get pipelinerun, with error: %v", err)
		log.Printf("[Error] %s", msg)
		return nil, fmt.Errorf(msg)
	}
	unstructured := resp.UnstructuredContent()
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &obj); err != nil {
		msg := fmt.Sprintf("Failed to translate unstructed to PipelineRun, with error: %v", err)
		log.Printf("[Error] %s", msg)
		return nil, fmt.Errorf(msg)
	}
	return &obj, nil
}

// UpdatePipelineRun implements Client
func (c *client) UpdatePipelineRun(namespace string, name string, obj *tektonapiv1.PipelineRun, data []byte) error {
	pipelineRunUpdateTypeMeta(obj)
	return c.updateResource(namespace, name, pipelineRunRes(), obj, data)
}

func pipelineRunRes() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  "v1alpha1",
		Resource: "pipelineruns",
	}
}

func pipelineRunUpdateTypeMeta(obj *tektonapiv1.PipelineRun) {
	obj.APIVersion = "tekton.dev/v1alpha1"
	obj.Kind = "PipelineRun"
}

// Pipeline CRUD operations

// Task CRUD operations

func (c *client) CreateTask(obj *tektonapiv1.Task) error {
	taskUpdateTypeMeta(obj)
	return c.createResource(obj, obj.Namespace, taskRes())
}

func (c *client) GetTask(namespace string, name string) (*tektonapiv1.Task, error) {
	var obj tektonapiv1.Task
	resp, err := c.getResource(namespace, name, taskRes())
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("[Warning] Task %s not found (namespace=%s)", name, namespace)
			return nil, err
		}
		msg := fmt.Sprintf("Failed to get Task, with error: %v", err)
		log.Printf("[Error] %s", msg)
		return nil, fmt.Errorf(msg)
	}
	unstructured := resp.UnstructuredContent()
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &obj); err != nil {
		msg := fmt.Sprintf("Failed to translate unstructed to Task, with error: %v", err)
		log.Printf("[Error] %s", msg)
		return nil, fmt.Errorf(msg)
	}
	return &obj, nil
}

func (c *client) UpdateTask(namespace string, name string, obj *tektonapiv1.Task, data []byte) error {
	taskUpdateTypeMeta(obj)
	return c.updateResource(namespace, name, taskRes(), obj, data)
}

func (c *client) DeleteTask(namespace string, name string) error {
	return c.deleteResource(namespace, name, taskRes())
}

func taskUpdateTypeMeta(obj *tektonapiv1.Task) {
	obj.TypeMeta = metav1.TypeMeta{
		Kind:       "Task",
		APIVersion: tektonapiv1.SchemeGroupVersion.String(),
	}
}

func taskRes() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    tektonapiv1.SchemeGroupVersion.Group,
		Version:  tektonapiv1.SchemeGroupVersion.Version,
		Resource: "task",
	}

}

// Generic Resource CRUD operations

func (c *client) createResource(obj interface{}, namespace string, resource schema.GroupVersionResource) error {
	resultMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		msg := fmt.Sprintf("Failed to translate %s to Unstructed (for create operation), with error: %v", resource.Resource, err)
		log.Printf("[Error] %s", msg)
		return fmt.Errorf(msg)
	}
	input := unstructured.Unstructured{}
	input.SetUnstructuredContent(resultMap)
	resp, err := c.dynamicClient.Resource(resource).Namespace(namespace).Create(context.Background(), &input, metav1.CreateOptions{})
	if err != nil {
		msg := fmt.Sprintf("Failed to create %s, with error: %v", resource.Resource, err)
		log.Printf("[Error] %s", msg)
		return fmt.Errorf(msg)
	}
	unstructured := resp.UnstructuredContent()
	return runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, obj)
}

func (c *client) getResource(namespace string, name string, resource schema.GroupVersionResource) (*unstructured.Unstructured, error) {
	return c.dynamicClient.Resource(resource).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
}

func (c *client) updateResource(namespace string, name string, resource schema.GroupVersionResource, obj interface{}, data []byte) error {
	// patch, merge
	resp, err := c.dynamicClient.Resource(resource).Namespace(namespace).Patch(
		context.Background(),
		name,
		pkgApi.JSONPatchType,
		data,
		metav1.PatchOptions{})
	if err != nil {
		msg := fmt.Sprintf("Failed to update %s, with error: %v", resource.Resource, err)
		log.Printf("[Error] %s", msg)
		return fmt.Errorf(msg)
	}
	unstructured := resp.UnstructuredContent()
	return runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, obj)
}

func (c *client) deleteResource(namespace string, name string, resource schema.GroupVersionResource) error {
	return c.dynamicClient.Resource(resource).Namespace(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
