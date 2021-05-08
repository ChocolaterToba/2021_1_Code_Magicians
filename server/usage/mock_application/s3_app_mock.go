// Code generated by MockGen. DO NOT EDIT.
// Source: usage/s3_app.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockS3AppInterface is a mock of S3AppInterface interface.
type MockS3AppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockS3AppInterfaceMockRecorder
}

// MockS3AppInterfaceMockRecorder is the mock recorder for MockS3AppInterface.
type MockS3AppInterfaceMockRecorder struct {
	mock *MockS3AppInterface
}

// NewMockS3AppInterface creates a new mock instance.
func NewMockS3AppInterface(ctrl *gomock.Controller) *MockS3AppInterface {
	mock := &MockS3AppInterface{ctrl: ctrl}
	mock.recorder = &MockS3AppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3AppInterface) EXPECT() *MockS3AppInterfaceMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockS3AppInterface) DeleteFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockS3AppInterfaceMockRecorder) DeleteFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockS3AppInterface)(nil).DeleteFile), arg0)
}

// UploadFile mocks base method.
func (m *MockS3AppInterface) UploadFile(arg0 io.Reader, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockS3AppInterfaceMockRecorder) UploadFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockS3AppInterface)(nil).UploadFile), arg0, arg1)
}
