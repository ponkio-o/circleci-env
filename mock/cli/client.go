// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_cli is a generated GoMock package.
package mock_cli

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUI is a mock of UI interface.
type MockUI struct {
	ctrl     *gomock.Controller
	recorder *MockUIMockRecorder
}

// MockUIMockRecorder is the mock recorder for MockUI.
type MockUIMockRecorder struct {
	mock *MockUI
}

// NewMockUI creates a new mock instance.
func NewMockUI(ctrl *gomock.Controller) *MockUI {
	mock := &MockUI{ctrl: ctrl}
	mock.recorder = &MockUIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUI) EXPECT() *MockUIMockRecorder {
	return m.recorder
}

// ReadInput mocks base method.
func (m *MockUI) ReadInput(msg string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInput", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInput indicates an expected call of ReadInput.
func (mr *MockUIMockRecorder) ReadInput(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInput", reflect.TypeOf((*MockUI)(nil).ReadInput), msg)
}

// ReadSecret mocks base method.
func (m *MockUI) ReadSecret(msg string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadSecret", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSecret indicates an expected call of ReadSecret.
func (mr *MockUIMockRecorder) ReadSecret(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSecret", reflect.TypeOf((*MockUI)(nil).ReadSecret), msg)
}

// SelectFromList mocks base method.
func (m *MockUI) SelectFromList(msg string, ls []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFromList", msg, ls)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFromList indicates an expected call of SelectFromList.
func (mr *MockUIMockRecorder) SelectFromList(msg, ls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFromList", reflect.TypeOf((*MockUI)(nil).SelectFromList), msg, ls)
}

// YesNo mocks base method.
func (m *MockUI) YesNo(msg string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "YesNo", msg)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// YesNo indicates an expected call of YesNo.
func (mr *MockUIMockRecorder) YesNo(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YesNo", reflect.TypeOf((*MockUI)(nil).YesNo), msg)
}
