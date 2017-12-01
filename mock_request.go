// Code generated by MockGen. DO NOT EDIT.
// Source: request.go

// Package ripo is a generated GoMock package.
package ripo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockRequest is a mock of Request interface
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// RemoteIP mocks base method
func (m *MockRequest) RemoteIP() (string, error) {
	ret := m.ctrl.Call(m, "RemoteIP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteIP indicates an expected call of RemoteIP
func (mr *MockRequestMockRecorder) RemoteIP() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteIP", reflect.TypeOf((*MockRequest)(nil).RemoteIP))
}

// URL mocks base method
func (m *MockRequest) URL() *url.URL {
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockRequestMockRecorder) URL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockRequest)(nil).URL))
}

// Host mocks base method
func (m *MockRequest) Host() string {
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockRequestMockRecorder) Host() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockRequest)(nil).Host))
}

// HandlerName mocks base method
func (m *MockRequest) HandlerName() string {
	ret := m.ctrl.Call(m, "HandlerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HandlerName indicates an expected call of HandlerName
func (mr *MockRequestMockRecorder) HandlerName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlerName", reflect.TypeOf((*MockRequest)(nil).HandlerName))
}

// Body mocks base method
func (m *MockRequest) Body() ([]byte, error) {
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body
func (mr *MockRequestMockRecorder) Body() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockRequest)(nil).Body))
}

// BodyMap mocks base method
func (m *MockRequest) BodyMap() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "BodyMap")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BodyMap indicates an expected call of BodyMap
func (mr *MockRequestMockRecorder) BodyMap() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyMap", reflect.TypeOf((*MockRequest)(nil).BodyMap))
}

// BodyTo mocks base method
func (m *MockRequest) BodyTo(model interface{}) error {
	ret := m.ctrl.Call(m, "BodyTo", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// BodyTo indicates an expected call of BodyTo
func (mr *MockRequestMockRecorder) BodyTo(model interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyTo", reflect.TypeOf((*MockRequest)(nil).BodyTo), model)
}

// GetHeader mocks base method
func (m *MockRequest) GetHeader(arg0 string) string {
	ret := m.ctrl.Call(m, "GetHeader", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockRequestMockRecorder) GetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockRequest)(nil).GetHeader), arg0)
}

// HeaderKeys mocks base method
func (m *MockRequest) HeaderKeys() []string {
	ret := m.ctrl.Call(m, "HeaderKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// HeaderKeys indicates an expected call of HeaderKeys
func (mr *MockRequestMockRecorder) HeaderKeys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderKeys", reflect.TypeOf((*MockRequest)(nil).HeaderKeys))
}

// GetFormValue mocks base method
func (m *MockRequest) GetFormValue(key string) string {
	ret := m.ctrl.Call(m, "GetFormValue", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFormValue indicates an expected call of GetFormValue
func (mr *MockRequestMockRecorder) GetFormValue(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFormValue", reflect.TypeOf((*MockRequest)(nil).GetFormValue), key)
}

// Context mocks base method
func (m *MockRequest) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockRequestMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRequest)(nil).Context))
}

// GetString mocks base method
func (m *MockRequest) GetString(key string, sources ...FromX) (*string, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetString", varargs...)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetString indicates an expected call of GetString
func (mr *MockRequestMockRecorder) GetString(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockRequest)(nil).GetString), varargs...)
}

// GetStringList mocks base method
func (m *MockRequest) GetStringList(key string, sources ...FromX) ([]string, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStringList", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringList indicates an expected call of GetStringList
func (mr *MockRequestMockRecorder) GetStringList(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringList", reflect.TypeOf((*MockRequest)(nil).GetStringList), varargs...)
}

// GetInt mocks base method
func (m *MockRequest) GetInt(key string, sources ...FromX) (*int, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInt", varargs...)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInt indicates an expected call of GetInt
func (mr *MockRequestMockRecorder) GetInt(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockRequest)(nil).GetInt), varargs...)
}

// GetIntDefault mocks base method
func (m *MockRequest) GetIntDefault(key string, defaultValue int, sources ...FromX) (int, error) {
	varargs := []interface{}{key, defaultValue}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntDefault", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntDefault indicates an expected call of GetIntDefault
func (mr *MockRequestMockRecorder) GetIntDefault(key, defaultValue interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key, defaultValue}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntDefault", reflect.TypeOf((*MockRequest)(nil).GetIntDefault), varargs...)
}

// GetFloat mocks base method
func (m *MockRequest) GetFloat(key string, sources ...FromX) (*float64, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFloat", varargs...)
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloat indicates an expected call of GetFloat
func (mr *MockRequestMockRecorder) GetFloat(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat", reflect.TypeOf((*MockRequest)(nil).GetFloat), varargs...)
}

// GetBool mocks base method
func (m *MockRequest) GetBool(key string, sources ...FromX) (*bool, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBool", varargs...)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBool indicates an expected call of GetBool
func (mr *MockRequestMockRecorder) GetBool(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockRequest)(nil).GetBool), varargs...)
}

// GetTime mocks base method
func (m *MockRequest) GetTime(key string, sources ...FromX) (*time.Time, error) {
	varargs := []interface{}{key}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTime", varargs...)
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTime indicates an expected call of GetTime
func (mr *MockRequestMockRecorder) GetTime(key interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockRequest)(nil).GetTime), varargs...)
}

// GetObject mocks base method
func (m *MockRequest) GetObject(key string, _type reflect.Type, sources ...FromX) (interface{}, error) {
	varargs := []interface{}{key, _type}
	for _, a := range sources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockRequestMockRecorder) GetObject(key, _type interface{}, sources ...interface{}) *gomock.Call {
	varargs := append([]interface{}{key, _type}, sources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockRequest)(nil).GetObject), varargs...)
}

// FullMap mocks base method
func (m *MockRequest) FullMap() map[string]interface{} {
	ret := m.ctrl.Call(m, "FullMap")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// FullMap indicates an expected call of FullMap
func (mr *MockRequestMockRecorder) FullMap() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullMap", reflect.TypeOf((*MockRequest)(nil).FullMap))
}
