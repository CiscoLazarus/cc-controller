// Code generated by mockery v1.0.0
package main

import mock "github.com/stretchr/testify/mock"

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// AddResource provides a mock function with given fields: _a0, _a1
func (_m *MockController) AddResource(_a0 string, _a1 Model) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, Model) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddTask provides a mock function with given fields: _a0, _a1
func (_m *MockController) AddTask(_a0 *Task, _a1 Model) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Task, Model) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteTask provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockController) CompleteTask(_a0 string, _a1 int, _a2 Model) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, Model) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTask provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetTask(_a0 string, _a1 Model) (*Task, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Task
	if rf, ok := ret.Get(0).(func(string, Model) *Task); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, Model) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPriorityQueue provides a mock function with given fields: _a0
func (_m *MockController) ListPriorityQueue(_a0 string) (map[string]interface{}, error) {
	ret := _m.Called(_a0)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTimetable provides a mock function with given fields: _a0
func (_m *MockController) ListTimetable(_a0 string) (map[string]interface{}, error) {
	ret := _m.Called(_a0)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Notify provides a mock function with given fields: _a0
func (_m *MockController) Notify(_a0 *Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StageTask provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockController) StageTask(_a0 *Task, _a1 Model, _a2 bool) {
	_m.Called(_a0, _a1, _a2)
}

// StartTask provides a mock function with given fields: _a0, _a1
func (_m *MockController) StartTask(_a0 string, _a1 Model) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, Model) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
