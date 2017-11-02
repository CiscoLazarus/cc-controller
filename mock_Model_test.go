// Code generated by mockery v1.0.0
package main

import mock "github.com/stretchr/testify/mock"

// MockModel is an autogenerated mock type for the Model type
type MockModel struct {
	mock.Mock
}

// Create provides a mock function with given fields:
func (_m *MockModel) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: _a0, _a1
func (_m *MockModel) Query(_a0 string, _a1 interface{}) ([]interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}) []interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *MockModel) Save(_a0 interface{}) (DocumentMeta, error) {
	ret := _m.Called(_a0)

	var r0 DocumentMeta
	if rf, ok := ret.Get(0).(func(interface{}) DocumentMeta); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}