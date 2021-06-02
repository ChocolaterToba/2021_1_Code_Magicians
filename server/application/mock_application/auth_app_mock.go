// Code generated by MockGen. DO NOT EDIT.
// Source: application/auth_app.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	http "net/http"
	entity "pinterest/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthAppInterface is a mock of AuthAppInterface interface.
type MockAuthAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthAppInterfaceMockRecorder
}

// MockAuthAppInterfaceMockRecorder is the mock recorder for MockAuthAppInterface.
type MockAuthAppInterfaceMockRecorder struct {
	mock *MockAuthAppInterface
}

// NewMockAuthAppInterface creates a new mock instance.
func NewMockAuthAppInterface(ctrl *gomock.Controller) *MockAuthAppInterface {
	mock := &MockAuthAppInterface{ctrl: ctrl}
	mock.recorder = &MockAuthAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthAppInterface) EXPECT() *MockAuthAppInterfaceMockRecorder {
	return m.recorder
}

// AddVkCode mocks base method.
func (m *MockAuthAppInterface) AddVkCode(userID int, code, redirectURI string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVkCode", userID, code, redirectURI)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVkCode indicates an expected call of AddVkCode.
func (mr *MockAuthAppInterfaceMockRecorder) AddVkCode(userID, code, redirectURI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVkCode", reflect.TypeOf((*MockAuthAppInterface)(nil).AddVkCode), userID, code, redirectURI)
}

// AddVkToken mocks base method.
func (m *MockAuthAppInterface) AddVkToken(userID int, tokenInput *entity.UserVkTokenInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVkToken", userID, tokenInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVkToken indicates an expected call of AddVkToken.
func (mr *MockAuthAppInterfaceMockRecorder) AddVkToken(userID, tokenInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVkToken", reflect.TypeOf((*MockAuthAppInterface)(nil).AddVkToken), userID, tokenInput)
}

// CheckCookie mocks base method.
func (m *MockAuthAppInterface) CheckCookie(cookie *http.Cookie) (*entity.CookieInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCookie", cookie)
	ret0, _ := ret[0].(*entity.CookieInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckCookie indicates an expected call of CheckCookie.
func (mr *MockAuthAppInterfaceMockRecorder) CheckCookie(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCookie", reflect.TypeOf((*MockAuthAppInterface)(nil).CheckCookie), cookie)
}

// CheckUserCredentials mocks base method.
func (m *MockAuthAppInterface) CheckUserCredentials(username, password string) (*entity.CookieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserCredentials", username, password)
	ret0, _ := ret[0].(*entity.CookieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserCredentials indicates an expected call of CheckUserCredentials.
func (mr *MockAuthAppInterfaceMockRecorder) CheckUserCredentials(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserCredentials", reflect.TypeOf((*MockAuthAppInterface)(nil).CheckUserCredentials), username, password)
}

// CheckVkCode mocks base method.
func (m *MockAuthAppInterface) CheckVkCode(code, redirectURI string) (*entity.CookieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVkCode", code, redirectURI)
	ret0, _ := ret[0].(*entity.CookieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVkCode indicates an expected call of CheckVkCode.
func (mr *MockAuthAppInterfaceMockRecorder) CheckVkCode(code, redirectURI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVkCode", reflect.TypeOf((*MockAuthAppInterface)(nil).CheckVkCode), code, redirectURI)
}

// LogoutUser mocks base method.
func (m *MockAuthAppInterface) LogoutUser(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogoutUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogoutUser indicates an expected call of LogoutUser.
func (mr *MockAuthAppInterfaceMockRecorder) LogoutUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutUser", reflect.TypeOf((*MockAuthAppInterface)(nil).LogoutUser), userID)
}

// VkCodeToToken mocks base method.
func (m *MockAuthAppInterface) VkCodeToToken(code, redirectURI string) (*entity.UserVkTokenInput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkCodeToToken", code, redirectURI)
	ret0, _ := ret[0].(*entity.UserVkTokenInput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkCodeToToken indicates an expected call of VkCodeToToken.
func (mr *MockAuthAppInterfaceMockRecorder) VkCodeToToken(code, redirectURI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCodeToToken", reflect.TypeOf((*MockAuthAppInterface)(nil).VkCodeToToken), code, redirectURI)
}
