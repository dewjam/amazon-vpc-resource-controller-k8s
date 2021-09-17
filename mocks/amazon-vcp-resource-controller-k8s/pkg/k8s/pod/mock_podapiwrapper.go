// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/k8s/pod (interfaces: PodClientAPIWrapper)

// Package mock_pod is a generated GoMock package.
package mock_pod

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockPodClientAPIWrapper is a mock of PodClientAPIWrapper interface
type MockPodClientAPIWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockPodClientAPIWrapperMockRecorder
}

// MockPodClientAPIWrapperMockRecorder is the mock recorder for MockPodClientAPIWrapper
type MockPodClientAPIWrapperMockRecorder struct {
	mock *MockPodClientAPIWrapper
}

// NewMockPodClientAPIWrapper creates a new mock instance
func NewMockPodClientAPIWrapper(ctrl *gomock.Controller) *MockPodClientAPIWrapper {
	mock := &MockPodClientAPIWrapper{ctrl: ctrl}
	mock.recorder = &MockPodClientAPIWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodClientAPIWrapper) EXPECT() *MockPodClientAPIWrapperMockRecorder {
	return m.recorder
}

// AnnotatePod mocks base method
func (m *MockPodClientAPIWrapper) AnnotatePod(arg0, arg1 string, arg2 types.UID, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotatePod", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnotatePod indicates an expected call of AnnotatePod
func (mr *MockPodClientAPIWrapperMockRecorder) AnnotatePod(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotatePod", reflect.TypeOf((*MockPodClientAPIWrapper)(nil).AnnotatePod), arg0, arg1, arg2, arg3, arg4)
}

// GetPod mocks base method
func (m *MockPodClientAPIWrapper) GetPod(arg0, arg1 string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", arg0, arg1)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod
func (mr *MockPodClientAPIWrapperMockRecorder) GetPod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockPodClientAPIWrapper)(nil).GetPod), arg0, arg1)
}

// GetPodFromAPIServer mocks base method
func (m *MockPodClientAPIWrapper) GetPodFromAPIServer(arg0 context.Context, arg1, arg2 string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodFromAPIServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodFromAPIServer indicates an expected call of GetPodFromAPIServer
func (mr *MockPodClientAPIWrapperMockRecorder) GetPodFromAPIServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodFromAPIServer", reflect.TypeOf((*MockPodClientAPIWrapper)(nil).GetPodFromAPIServer), arg0, arg1, arg2)
}

// GetRunningPodsOnNode mocks base method
func (m *MockPodClientAPIWrapper) GetRunningPodsOnNode(arg0 string) ([]v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningPodsOnNode", arg0)
	ret0, _ := ret[0].([]v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunningPodsOnNode indicates an expected call of GetRunningPodsOnNode
func (mr *MockPodClientAPIWrapperMockRecorder) GetRunningPodsOnNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningPodsOnNode", reflect.TypeOf((*MockPodClientAPIWrapper)(nil).GetRunningPodsOnNode), arg0)
}

// ListPods mocks base method
func (m *MockPodClientAPIWrapper) ListPods(arg0 string) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPods", arg0)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPods indicates an expected call of ListPods
func (mr *MockPodClientAPIWrapperMockRecorder) ListPods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPods", reflect.TypeOf((*MockPodClientAPIWrapper)(nil).ListPods), arg0)
}
