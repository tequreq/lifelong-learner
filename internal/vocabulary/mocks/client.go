// Code generated by MockGen. DO NOT EDIT.
// Source: internal/vocabulary/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	vocabulary "github.com/Abdulsametileri/lifelong-learner/internal/vocabulary"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// FindMeaningByWord mocks base method.
func (m *MockClient) FindMeaningByWord(ctx context.Context, word string) (*vocabulary.Vocabulary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMeaningByWord", ctx, word)
	ret0, _ := ret[0].(*vocabulary.Vocabulary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMeaningByWord indicates an expected call of FindMeaningByWord.
func (mr *MockClientMockRecorder) FindMeaningByWord(ctx, word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMeaningByWord", reflect.TypeOf((*MockClient)(nil).FindMeaningByWord), ctx, word)
}