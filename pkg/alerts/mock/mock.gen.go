// Package mock is a generated GoMock package. DO NOT EDIT
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAlerter is a mock of Alerter interface.
type MockAlerter struct {
	ctrl     *gomock.Controller
	recorder *MockAlerterMockRecorder
}

// MockAlerterMockRecorder is the mock recorder for MockAlerter.
type MockAlerterMockRecorder struct {
	mock *MockAlerter
}

// NewMockAlerter creates a new mock instance.
func NewMockAlerter(ctrl *gomock.Controller) *MockAlerter {
	mock := &MockAlerter{ctrl: ctrl}
	mock.recorder = &MockAlerterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlerter) EXPECT() *MockAlerterMockRecorder {
	return m.recorder
}

// Alert mocks base method.
func (m *MockAlerter) Alert(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Alert", arg0, arg1)
}

// Alert indicates an expected call of Alert.
func (mr *MockAlerterMockRecorder) Alert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alert", reflect.TypeOf((*MockAlerter)(nil).Alert), arg0, arg1)
}

// Notify mocks base method.
func (m *MockAlerter) Notify(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", arg0, arg1)
}

// Notify indicates an expected call of Notify.
func (mr *MockAlerterMockRecorder) Notify(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockAlerter)(nil).Notify), arg0, arg1)
}
