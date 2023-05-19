// Code generated by MockGen. DO NOT EDIT.
// Source: bookstore/bookstore.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "awesomeProject1/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBookStore is a mock of BookStore interface.
type MockBookStore struct {
	ctrl     *gomock.Controller
	recorder *MockBookStoreMockRecorder
}

// MockBookStoreMockRecorder is the mock recorder for MockBookStore.
type MockBookStoreMockRecorder struct {
	mock *MockBookStore
}

// NewMockBookStore creates a new mock instance.
func NewMockBookStore(ctrl *gomock.Controller) *MockBookStore {
	mock := &MockBookStore{ctrl: ctrl}
	mock.recorder = &MockBookStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookStore) EXPECT() *MockBookStoreMockRecorder {
	return m.recorder
}

// CreateNewBooks mocks base method.
func (m *MockBookStore) Update(books []entity.Book) ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", books)
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewBooks indicates an expected call of CreateNewBooks.
func (mr *MockBookStoreMockRecorder) CreateNewBooks(books interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookStore)(nil).Update), books)
}

// GetAllBooks mocks base method.
func (m *MockBookStore) GetAll() []entity.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Book)
	return ret0
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockBookStoreMockRecorder) GetAllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBookStore)(nil).GetAll))
}

