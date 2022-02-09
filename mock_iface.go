// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIface is a mock of Iface interface
type MockIface struct {
	ctrl     *gomock.Controller
	recorder *MockIfaceMockRecorder
}

// MockIfaceMockRecorder is the mock recorder for MockIface
type MockIfaceMockRecorder struct {
	mock *MockIface
}

// NewMockIface creates a new mock instance
func NewMockIface(ctrl *gomock.Controller) *MockIface {
	mock := &MockIface{ctrl: ctrl}
	mock.recorder = &MockIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIface) EXPECT() *MockIfaceMockRecorder {
	return m.recorder
}

// Foo mocks base method
func (m *MockIface) Foo(a, b, c, d int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Foo", a, b, c, d)
}

// Foo indicates an expected call of Foo
func (mr *MockIfaceMockRecorder) Foo(a, b, c, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockIface)(nil).Foo), a, b, c, d)
}
