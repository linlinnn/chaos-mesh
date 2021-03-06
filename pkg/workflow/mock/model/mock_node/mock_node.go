// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/node/node.go

// Package mock_node is a generated GoMock package.
package mock_node

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	node "github.com/chaos-mesh/chaos-mesh/pkg/workflow/model/node"
)

// MockNode is a mock of Node interface.
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode.
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance.
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockNode) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNodeMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNode)(nil).Name))
}

// NodePhase mocks base method.
func (m *MockNode) NodePhase() node.NodePhase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodePhase")
	ret0, _ := ret[0].(node.NodePhase)
	return ret0
}

// NodePhase indicates an expected call of NodePhase.
func (mr *MockNodeMockRecorder) NodePhase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodePhase", reflect.TypeOf((*MockNode)(nil).NodePhase))
}

// ParentNodeName mocks base method.
func (m *MockNode) ParentNodeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentNodeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ParentNodeName indicates an expected call of ParentNodeName.
func (mr *MockNodeMockRecorder) ParentNodeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentNodeName", reflect.TypeOf((*MockNode)(nil).ParentNodeName))
}

// TemplateName mocks base method.
func (m *MockNode) TemplateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TemplateName indicates an expected call of TemplateName.
func (mr *MockNodeMockRecorder) TemplateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateName", reflect.TypeOf((*MockNode)(nil).TemplateName))
}

// MockNodeTreeNode is a mock of NodeTreeNode interface.
type MockNodeTreeNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeTreeNodeMockRecorder
}

// MockNodeTreeNodeMockRecorder is the mock recorder for MockNodeTreeNode.
type MockNodeTreeNodeMockRecorder struct {
	mock *MockNodeTreeNode
}

// NewMockNodeTreeNode creates a new mock instance.
func NewMockNodeTreeNode(ctrl *gomock.Controller) *MockNodeTreeNode {
	mock := &MockNodeTreeNode{ctrl: ctrl}
	mock.recorder = &MockNodeTreeNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeTreeNode) EXPECT() *MockNodeTreeNodeMockRecorder {
	return m.recorder
}

// Children mocks base method.
func (m *MockNodeTreeNode) Children() node.NodeTreeChildren {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Children")
	ret0, _ := ret[0].(node.NodeTreeChildren)
	return ret0
}

// Children indicates an expected call of Children.
func (mr *MockNodeTreeNodeMockRecorder) Children() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Children", reflect.TypeOf((*MockNodeTreeNode)(nil).Children))
}

// FetchNodeByName mocks base method.
func (m *MockNodeTreeNode) FetchNodeByName(nodeName string) (node.NodeTreeNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNodeByName", nodeName)
	ret0, _ := ret[0].(node.NodeTreeNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchNodeByName indicates an expected call of FetchNodeByName.
func (mr *MockNodeTreeNodeMockRecorder) FetchNodeByName(nodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNodeByName", reflect.TypeOf((*MockNodeTreeNode)(nil).FetchNodeByName), nodeName)
}

// Name mocks base method.
func (m *MockNodeTreeNode) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNodeTreeNodeMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNodeTreeNode)(nil).Name))
}

// TemplateName mocks base method.
func (m *MockNodeTreeNode) TemplateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TemplateName indicates an expected call of TemplateName.
func (mr *MockNodeTreeNodeMockRecorder) TemplateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateName", reflect.TypeOf((*MockNodeTreeNode)(nil).TemplateName))
}

// MockNodeTreeChildren is a mock of NodeTreeChildren interface.
type MockNodeTreeChildren struct {
	ctrl     *gomock.Controller
	recorder *MockNodeTreeChildrenMockRecorder
}

// MockNodeTreeChildrenMockRecorder is the mock recorder for MockNodeTreeChildren.
type MockNodeTreeChildrenMockRecorder struct {
	mock *MockNodeTreeChildren
}

// NewMockNodeTreeChildren creates a new mock instance.
func NewMockNodeTreeChildren(ctrl *gomock.Controller) *MockNodeTreeChildren {
	mock := &MockNodeTreeChildren{ctrl: ctrl}
	mock.recorder = &MockNodeTreeChildrenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeTreeChildren) EXPECT() *MockNodeTreeChildrenMockRecorder {
	return m.recorder
}

// GetAllChildrenNode mocks base method.
func (m *MockNodeTreeChildren) GetAllChildrenNode() []node.NodeTreeNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChildrenNode")
	ret0, _ := ret[0].([]node.NodeTreeNode)
	return ret0
}

// GetAllChildrenNode indicates an expected call of GetAllChildrenNode.
func (mr *MockNodeTreeChildrenMockRecorder) GetAllChildrenNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChildrenNode", reflect.TypeOf((*MockNodeTreeChildren)(nil).GetAllChildrenNode))
}

// Length mocks base method.
func (m *MockNodeTreeChildren) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockNodeTreeChildrenMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockNodeTreeChildren)(nil).Length))
}
