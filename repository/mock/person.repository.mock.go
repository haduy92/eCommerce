// Code generated by MockGen. DO NOT EDIT.
// Source: repository/person.repository.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "eCommerce/model/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockPersonRepository is a mock of PersonRepository interface.
type MockPersonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPersonRepositoryMockRecorder
}

// MockPersonRepositoryMockRecorder is the mock recorder for MockPersonRepository.
type MockPersonRepositoryMockRecorder struct {
	mock *MockPersonRepository
}

// NewMockPersonRepository creates a new mock instance.
func NewMockPersonRepository(ctrl *gomock.Controller) *MockPersonRepository {
	mock := &MockPersonRepository{ctrl: ctrl}
	mock.recorder = &MockPersonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonRepository) EXPECT() *MockPersonRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPersonRepository) Create(arg0 *entity.Person) (*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPersonRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPersonRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockPersonRepository) Delete(arg0 *uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPersonRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPersonRepository)(nil).Delete), arg0)
}

// Exists mocks base method.
func (m *MockPersonRepository) Exists(arg0 *uuid.UUID) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockPersonRepositoryMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockPersonRepository)(nil).Exists), arg0)
}

// ExistsByEmail mocks base method.
func (m *MockPersonRepository) ExistsByEmail(arg0 *string) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByEmail", arg0)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByEmail indicates an expected call of ExistsByEmail.
func (mr *MockPersonRepositoryMockRecorder) ExistsByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByEmail", reflect.TypeOf((*MockPersonRepository)(nil).ExistsByEmail), arg0)
}

// ExistsByEmailExcludeID mocks base method.
func (m *MockPersonRepository) ExistsByEmailExcludeID(arg0 *uuid.UUID, arg1 *string) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByEmailExcludeID", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByEmailExcludeID indicates an expected call of ExistsByEmailExcludeID.
func (mr *MockPersonRepositoryMockRecorder) ExistsByEmailExcludeID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByEmailExcludeID", reflect.TypeOf((*MockPersonRepository)(nil).ExistsByEmailExcludeID), arg0, arg1)
}

// Get mocks base method.
func (m *MockPersonRepository) Get(arg0 *uuid.UUID) (*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPersonRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersonRepository)(nil).Get), arg0)
}

// Search mocks base method.
func (m *MockPersonRepository) Search(arg0 *string) ([]*entity.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].([]*entity.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockPersonRepositoryMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockPersonRepository)(nil).Search), arg0)
}

// Update mocks base method.
func (m *MockPersonRepository) Update(arg0 *entity.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPersonRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPersonRepository)(nil).Update), arg0)
}
