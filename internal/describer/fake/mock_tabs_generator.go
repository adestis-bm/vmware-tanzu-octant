// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/describer (interfaces: TabsGenerator)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	describer "github.com/vmware-tanzu/octant/internal/describer"
	component "github.com/vmware-tanzu/octant/pkg/view/component"
	reflect "reflect"
)

// MockTabsGenerator is a mock of TabsGenerator interface
type MockTabsGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTabsGeneratorMockRecorder
}

// MockTabsGeneratorMockRecorder is the mock recorder for MockTabsGenerator
type MockTabsGeneratorMockRecorder struct {
	mock *MockTabsGenerator
}

// NewMockTabsGenerator creates a new mock instance
func NewMockTabsGenerator(ctrl *gomock.Controller) *MockTabsGenerator {
	mock := &MockTabsGenerator{ctrl: ctrl}
	mock.recorder = &MockTabsGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTabsGenerator) EXPECT() *MockTabsGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockTabsGenerator) Generate(arg0 context.Context, arg1 describer.TabsGeneratorConfig) ([]component.Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0, arg1)
	ret0, _ := ret[0].([]component.Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate
func (mr *MockTabsGeneratorMockRecorder) Generate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTabsGenerator)(nil).Generate), arg0, arg1)
}
