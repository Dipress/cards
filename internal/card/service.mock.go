// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package card is a generated GoMock package.
package card

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepository) Create(arg0 context.Context, arg1 *NewCard, arg2 *Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1, arg2)
}

// MockValidater is a mock of Validater interface
type MockValidater struct {
	ctrl     *gomock.Controller
	recorder *MockValidaterMockRecorder
}

// MockValidaterMockRecorder is the mock recorder for MockValidater
type MockValidaterMockRecorder struct {
	mock *MockValidater
}

// NewMockValidater creates a new mock instance
func NewMockValidater(ctrl *gomock.Controller) *MockValidater {
	mock := &MockValidater{ctrl: ctrl}
	mock.recorder = &MockValidaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidater) EXPECT() *MockValidaterMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockValidater) Validate(arg0 context.Context, arg1 *Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockValidaterMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidater)(nil).Validate), arg0, arg1)
}
