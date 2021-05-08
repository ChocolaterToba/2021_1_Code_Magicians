// Code generated by MockGen. DO NOT EDIT.
// Source: usage/user_app.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	io "io"
	entity "pinterest/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserAppInterface is a mock of UserAppInterface interface.
type MockUserAppInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserAppInterfaceMockRecorder
}

// MockUserAppInterfaceMockRecorder is the mock recorder for MockUserAppInterface.
type MockUserAppInterfaceMockRecorder struct {
	mock *MockUserAppInterface
}

// NewMockUserAppInterface creates a new mock instance.
func NewMockUserAppInterface(ctrl *gomock.Controller) *MockUserAppInterface {
	mock := &MockUserAppInterface{ctrl: ctrl}
	mock.recorder = &MockUserAppInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAppInterface) EXPECT() *MockUserAppInterfaceMockRecorder {
	return m.recorder
}

// CheckIfFollowed mocks base method.
func (m *MockUserAppInterface) CheckIfFollowed(arg0, arg1 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfFollowed", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfFollowed indicates an expected call of CheckIfFollowed.
func (mr *MockUserAppInterfaceMockRecorder) CheckIfFollowed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfFollowed", reflect.TypeOf((*MockUserAppInterface)(nil).CheckIfFollowed), arg0, arg1)
}

// CheckUserCredentials mocks base method.
func (m *MockUserAppInterface) CheckUserCredentials(arg0, arg1 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserCredentials", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserCredentials indicates an expected call of CheckUserCredentials.
func (mr *MockUserAppInterfaceMockRecorder) CheckUserCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserCredentials", reflect.TypeOf((*MockUserAppInterface)(nil).CheckUserCredentials), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockUserAppInterface) CreateUser(arg0 *entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserAppInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserAppInterface)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockUserAppInterface) DeleteUser(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserAppInterfaceMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserAppInterface)(nil).DeleteUser), arg0)
}

// Follow mocks base method.
func (m *MockUserAppInterface) Follow(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow.
func (mr *MockUserAppInterfaceMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserAppInterface)(nil).Follow), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserAppInterface) GetUser(arg0 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserAppInterfaceMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserAppInterface)(nil).GetUser), arg0)
}

// GetUserByUsername mocks base method.
func (m *MockUserAppInterface) GetUserByUsername(arg0 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserAppInterfaceMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserAppInterface)(nil).GetUserByUsername), arg0)
}

// GetUsers mocks base method.
func (m *MockUserAppInterface) GetUsers() ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserAppInterfaceMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserAppInterface)(nil).GetUsers))
}

// SaveUser mocks base method.
func (m *MockUserAppInterface) SaveUser(arg0 *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockUserAppInterfaceMockRecorder) SaveUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserAppInterface)(nil).SaveUser), arg0)
}

// SearchUsers mocks base method.
func (m *MockUserAppInterface) SearchUsers(arg0 string) ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", arg0)
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockUserAppInterfaceMockRecorder) SearchUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockUserAppInterface)(nil).SearchUsers), arg0)
}

// Unfollow mocks base method.
func (m *MockUserAppInterface) Unfollow(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockUserAppInterfaceMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserAppInterface)(nil).Unfollow), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockUserAppInterface) UpdateAvatar(arg0 int, arg1 io.Reader, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockUserAppInterfaceMockRecorder) UpdateAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUserAppInterface)(nil).UpdateAvatar), arg0, arg1, arg2)
}
