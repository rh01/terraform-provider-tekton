// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method.
func (m *MockClient) CreatePipeline(obj *v1.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockClientMockRecorder) CreatePipeline(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockClient)(nil).CreatePipeline), obj)
}

// CreatePipelineRun mocks base method.
func (m *MockClient) CreatePipelineRun(obj *v1.PipelineRun) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipelineRun", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipelineRun indicates an expected call of CreatePipelineRun.
func (mr *MockClientMockRecorder) CreatePipelineRun(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipelineRun", reflect.TypeOf((*MockClient)(nil).CreatePipelineRun), obj)
}

// CreateTask mocks base method.
func (m *MockClient) CreateTask(obj *v1.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockClientMockRecorder) CreateTask(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockClient)(nil).CreateTask), obj)
}

// CreateTaskRun mocks base method.
func (m *MockClient) CreateTaskRun(obj *v1.TaskRun) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTaskRun", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTaskRun indicates an expected call of CreateTaskRun.
func (mr *MockClientMockRecorder) CreateTaskRun(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTaskRun", reflect.TypeOf((*MockClient)(nil).CreateTaskRun), obj)
}

// DeletePipeline mocks base method.
func (m *MockClient) DeletePipeline(namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockClientMockRecorder) DeletePipeline(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockClient)(nil).DeletePipeline), namespace, name)
}

// DeletePipelineRun mocks base method.
func (m *MockClient) DeletePipelineRun(namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipelineRun", namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipelineRun indicates an expected call of DeletePipelineRun.
func (mr *MockClientMockRecorder) DeletePipelineRun(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipelineRun", reflect.TypeOf((*MockClient)(nil).DeletePipelineRun), namespace, name)
}

// DeleteTask mocks base method.
func (m *MockClient) DeleteTask(namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockClientMockRecorder) DeleteTask(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockClient)(nil).DeleteTask), namespace, name)
}

// DeleteTaskRun mocks base method.
func (m *MockClient) DeleteTaskRun(namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTaskRun", namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTaskRun indicates an expected call of DeleteTaskRun.
func (mr *MockClientMockRecorder) DeleteTaskRun(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTaskRun", reflect.TypeOf((*MockClient)(nil).DeleteTaskRun), namespace, name)
}

// GetPipeline mocks base method.
func (m *MockClient) GetPipeline(namespace, name string) (*v1.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipeline", namespace, name)
	ret0, _ := ret[0].(*v1.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipeline indicates an expected call of GetPipeline.
func (mr *MockClientMockRecorder) GetPipeline(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockClient)(nil).GetPipeline), namespace, name)
}

// GetPipelineRun mocks base method.
func (m *MockClient) GetPipelineRun(namespace, name string) (*v1.PipelineRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineRun", namespace, name)
	ret0, _ := ret[0].(*v1.PipelineRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineRun indicates an expected call of GetPipelineRun.
func (mr *MockClientMockRecorder) GetPipelineRun(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineRun", reflect.TypeOf((*MockClient)(nil).GetPipelineRun), namespace, name)
}

// GetTask mocks base method.
func (m *MockClient) GetTask(namespace, name string) (*v1.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", namespace, name)
	ret0, _ := ret[0].(*v1.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockClientMockRecorder) GetTask(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockClient)(nil).GetTask), namespace, name)
}

// GetTaskRun mocks base method.
func (m *MockClient) GetTaskRun(namespace, name string) (*v1.TaskRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskRun", namespace, name)
	ret0, _ := ret[0].(*v1.TaskRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskRun indicates an expected call of GetTaskRun.
func (mr *MockClientMockRecorder) GetTaskRun(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskRun", reflect.TypeOf((*MockClient)(nil).GetTaskRun), namespace, name)
}

// UpdatePipeline mocks base method.
func (m *MockClient) UpdatePipeline(namespace, name string, obj *v1.Pipeline, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", namespace, name, obj, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockClientMockRecorder) UpdatePipeline(namespace, name, obj, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockClient)(nil).UpdatePipeline), namespace, name, obj, data)
}

// UpdatePipelineRun mocks base method.
func (m *MockClient) UpdatePipelineRun(namespace, name string, obj *v1.PipelineRun, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineRun", namespace, name, obj, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipelineRun indicates an expected call of UpdatePipelineRun.
func (mr *MockClientMockRecorder) UpdatePipelineRun(namespace, name, obj, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineRun", reflect.TypeOf((*MockClient)(nil).UpdatePipelineRun), namespace, name, obj, data)
}

// UpdateTask mocks base method.
func (m *MockClient) UpdateTask(namespace, name string, obj *v1.Task, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", namespace, name, obj, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockClientMockRecorder) UpdateTask(namespace, name, obj, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockClient)(nil).UpdateTask), namespace, name, obj, data)
}

// UpdateTaskRun mocks base method.
func (m *MockClient) UpdateTaskRun(namespace, name string, obj *v1.TaskRun, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskRun", namespace, name, obj, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTaskRun indicates an expected call of UpdateTaskRun.
func (mr *MockClientMockRecorder) UpdateTaskRun(namespace, name, obj, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskRun", reflect.TypeOf((*MockClient)(nil).UpdateTaskRun), namespace, name, obj, data)
}
