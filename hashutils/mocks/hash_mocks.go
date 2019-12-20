// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/go-utils/hashutils (interfaces: Hasher,SafeHasher)

// Package mock_hashutils is a generated GoMock package.
package mock_hashutils

import (
	hash "hash"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method
func (m *MockHasher) Hash() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockHasherMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHasher)(nil).Hash))
}

// MockSafeHasher is a mock of SafeHasher interface
type MockSafeHasher struct {
	ctrl     *gomock.Controller
	recorder *MockSafeHasherMockRecorder
}

// MockSafeHasherMockRecorder is the mock recorder for MockSafeHasher
type MockSafeHasherMockRecorder struct {
	mock *MockSafeHasher
}

// NewMockSafeHasher creates a new mock instance
func NewMockSafeHasher(ctrl *gomock.Controller) *MockSafeHasher {
	mock := &MockSafeHasher{ctrl: ctrl}
	mock.recorder = &MockSafeHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSafeHasher) EXPECT() *MockSafeHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method
func (m *MockSafeHasher) Hash(arg0 hash.Hash64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (mr *MockSafeHasherMockRecorder) Hash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockSafeHasher)(nil).Hash), arg0)
}