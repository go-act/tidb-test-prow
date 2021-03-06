// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/br/pkg/lightning/backend/kv (interfaces: Encoder,Rows,Row)

// $ mockgen -package mock github.com/pingcap/tidb/br/pkg/lightning/backend/kv Encoder,Rows,Row

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kv "github.com/pingcap/tidb/br/pkg/lightning/backend/kv"
	log "github.com/pingcap/tidb/br/pkg/lightning/log"
	verification "github.com/pingcap/tidb/br/pkg/lightning/verification"
	types "github.com/pingcap/tidb/types"
)

// MockEncoder is a mock of Encoder interface.
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEncoder) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEncoderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEncoder)(nil).Close))
}

// Encode mocks base method.
func (m *MockEncoder) Encode(arg0 log.Logger, arg1 []types.Datum, arg2 int64, arg3 []int, arg4 string, arg5 int64) (kv.Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(kv.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderMockRecorder) Encode(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockRows) Clear() kv.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(kv.Rows)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockRowsMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockRows)(nil).Clear))
}

// SplitIntoChunks mocks base method.
func (m *MockRows) SplitIntoChunks(arg0 int) []kv.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SplitIntoChunks", arg0)
	ret0, _ := ret[0].([]kv.Rows)
	return ret0
}

// SplitIntoChunks indicates an expected call of SplitIntoChunks.
func (mr *MockRowsMockRecorder) SplitIntoChunks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SplitIntoChunks", reflect.TypeOf((*MockRows)(nil).SplitIntoChunks), arg0)
}

// MockRow is a mock of Row interface.
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow.
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance.
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// ClassifyAndAppend mocks base method.
func (m *MockRow) ClassifyAndAppend(arg0 *kv.Rows, arg1 *verification.KVChecksum, arg2 *kv.Rows, arg3 *verification.KVChecksum) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClassifyAndAppend", arg0, arg1, arg2, arg3)
}

// ClassifyAndAppend indicates an expected call of ClassifyAndAppend.
func (mr *MockRowMockRecorder) ClassifyAndAppend(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassifyAndAppend", reflect.TypeOf((*MockRow)(nil).ClassifyAndAppend), arg0, arg1, arg2, arg3)
}

// Size mocks base method.
func (m *MockRow) Size() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockRowMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockRow)(nil).Size))
}
