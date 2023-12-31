// Code generated by MockGen. DO NOT EDIT.
// Source: controller/enricher.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	enricher "github.com/wtkeqrf0/restService/internal/enricher"
	gomock "go.uber.org/mock/gomock"
)

// MockEnricher is a mock of Enricher interface.
type MockEnricher struct {
	ctrl     *gomock.Controller
	recorder *MockEnricherMockRecorder
}

// MockEnricherMockRecorder is the mock recorder for MockEnricher.
type MockEnricherMockRecorder struct {
	mock *MockEnricher
}

// NewMockEnricher creates a new mock instance.
func NewMockEnricher(ctrl *gomock.Controller) *MockEnricher {
	mock := &MockEnricher{ctrl: ctrl}
	mock.recorder = &MockEnricherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnricher) EXPECT() *MockEnricherMockRecorder {
	return m.recorder
}

// EnrichFIO mocks base method.
func (m *MockEnricher) EnrichFIO(ctx context.Context, fio enricher.FIO) (enricher.EnrichedFIO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrichFIO", ctx, fio)
	ret0, _ := ret[0].(enricher.EnrichedFIO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnrichFIO indicates an expected call of EnrichFIO.
func (mr *MockEnricherMockRecorder) EnrichFIO(ctx, fio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrichFIO", reflect.TypeOf((*MockEnricher)(nil).EnrichFIO), ctx, fio)
}

// ValidateName mocks base method.
func (m *MockEnricher) ValidateName(name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateName", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateName indicates an expected call of ValidateName.
func (mr *MockEnricherMockRecorder) ValidateName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateName", reflect.TypeOf((*MockEnricher)(nil).ValidateName), name)
}
