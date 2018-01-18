package main

import (
	"encoding/json"
	"testing"

	"github.com/bitwurx/jrpc2"
	"github.com/stretchr/testify/mock"
)

func TestAp1V1AddResource(t *testing.T) {
	var table = []struct {
		Body    []byte
		Name    string
		CallErr error
		Result  int
		ErrCode jrpc2.ErrorCode
		ErrMsg  jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"name": "test"}`),
			"test",
			nil,
			0,
			-1,
			"",
		},
		{
			[]byte(`["test2"]`),
			"test2",
			nil,
			0,
			-1,
			"",
		},
		{
			[]byte(`["test"]`),
			"test",
			ResourceExistsError,
			-1,
			AddResourceErrorCode,
			AddResourceErrorMsg,
		},
	}

	for _, tt := range table {
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		models := map[string]Model{"resources": rescModel}
		ctrl := &MockController{}
		ctrl.On("AddResource", tt.Name, rescModel).Return(tt.CallErr)
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.AddResource(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && result != tt.Result {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}

func TestAp1V1AddTask(t *testing.T) {
	var table = []struct {
		Body    []byte
		CallErr error
		Result  int
		ErrCode jrpc2.ErrorCode
		ErrMsg  jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"key": "test", "priority": 2.1}`),
			nil,
			0,
			-1,
			"",
		},
		{
			[]byte(`{"key": "test", "priority": 2.1}`),
			TaskAddFailedError,
			-1,
			AddTaskErrorCode,
			AddTaskErrorMsg,
		},
		{
			[]byte(`["test", {}, 2.1, "2017-01-01T12:00:00Z"]`),
			nil,
			0,
			-1,
			"",
		},
		{
			[]byte(`["test", {}, 2.1, "2017-01-01T12:00:00Z", 8]`),
			nil,
			0,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"key": "test"}`),
			nil,
			0,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{}`),
			nil,
			0,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
	}

	for _, tt := range table {
		taskModel := &MockModel{}
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		models := map[string]Model{"tasks": taskModel, "resources": rescModel}
		ctrl := &MockController{}
		ctrl.On("AddTask", mock.AnythingOfType("*main.Task"), taskModel).Return(tt.CallErr).Once()
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.AddTask(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && result != tt.Result {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}

func TestAp1V1CompleteTask(t *testing.T) {
	var table = []struct {
		Body      []byte
		Key       string
		CallErr   error
		Result    int
		Notify    bool
		NotifyErr error
		ErrCode   jrpc2.ErrorCode
		ErrMsg    jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"key": "test"}`),
			"test",
			nil,
			0,
			true,
			nil,
			-1,
			"",
		},
		{
			[]byte(`["test2"]`),
			"test2",
			nil,
			0,
			true,
			nil,
			-1,
			"",
		},
		{
			[]byte(`[]`),
			"",
			nil,
			0,
			true,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"keys": "test"}`),
			"",
			nil,
			0,
			true,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`[1, 2]`),
			"",
			nil,
			0,
			true,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`["test"]`),
			"test",
			TaskNotStartedError,
			-1,
			false,
			nil,
			CompleteTaskErrorCode,
			CompleteTaskErrorMsg,
		},
		{
			[]byte(`{"key": "test"}`),
			"test",
			nil,
			0,
			true,
			NotificationFailedError,
			NotificationFailedErrorCode,
			NotificationFailedErrorMsg,
		},
	}

	for _, tt := range table {
		taskModel := &MockModel{}
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		ctrl := &MockController{}
		ctrl.On("CompleteTask", tt.Key, taskModel).Return(tt.CallErr)
		if tt.Notify == true {
			ctrl.On("Notify", mock.MatchedBy(func(e *Event) bool { return e.Kind == EventTaskStatusChanged })).Return(tt.NotifyErr)
		}
		models := map[string]Model{"resources": rescModel, "tasks": taskModel}
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.CompleteTask(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && result != tt.Result {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}

func TestAp1V1GetTask(t *testing.T) {
	var table = []struct {
		Body    []byte
		TaskId  string
		Err     error
		Result  *Task
		ErrCode jrpc2.ErrorCode
		ErrMsg  jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"id": "abc123"}`),
			"abc123",
			nil,
			NewTask([]byte(`{"key": "test", "id": "abc123", "priority": 2.3}`)),
			-1,
			"",
		},
		{
			[]byte(`{"id": "123xyz"}`),
			"123xyz",
			TaskNotFoundError,
			nil,
			GetTaskErrorCode,
			GetTaskErrorMsg,
		},
	}

	for _, tt := range table {
		taskModel := &MockModel{}
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		models := map[string]Model{"resources": rescModel, "tasks": taskModel}
		ctrl := &MockController{}
		ctrl.On("GetTask", tt.TaskId, taskModel).Return(tt.Result, tt.Err)
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.GetTask(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && result != tt.Result {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}

func TestApiV1ListPriorityQueue(t *testing.T) {
	var table = []struct {
		Body    []byte
		Key     string
		Err     error
		Result  json.RawMessage
		ErrCode jrpc2.ErrorCode
		ErrMsg  jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"key": "test"}`),
			"test",
			nil,
			[]byte(`[{"_key": "test"}]`),
			-1,
			"",
		},
		{
			[]byte(`["test"]`),
			"test",
			nil,
			[]byte(`[{"_key": "test"}]`),
			-1,
			"",
		},
		{
			[]byte(`["test", 25]`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"keys", 25}`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"key", 25}`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"key": "test"}`),
			"test",
			QueueNotFoundError,
			nil,
			ListPriorityQueueErrorCode,
			ListPriorityQueueErrorMsg,
		},
	}

	for _, tt := range table {
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		models := map[string]Model{"resources": rescModel}
		ctrl := &MockController{}
		ctrl.On("ListPriorityQueue", tt.Key).Return(tt.Result, tt.Err)
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.ListPriorityQueue(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && string(result.(json.RawMessage)) != string(tt.Result) {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}

func TestApiV1ListTimetable(t *testing.T) {
	var table = []struct {
		Body    []byte
		Key     string
		Err     error
		Result  json.RawMessage
		ErrCode jrpc2.ErrorCode
		ErrMsg  jrpc2.ErrorMsg
	}{
		{
			[]byte(`{"key": "test"}`),
			"test",
			nil,
			[]byte(`[{"_key": "test"}]`),
			-1,
			"",
		},
		{
			[]byte(`["test"]`),
			"test",
			nil,
			[]byte(`[{"_key": "test"}]`),
			-1,
			"",
		},
		{
			[]byte(`["test", 25]`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"keys", 25}`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"key", 25}`),
			"test",
			nil,
			nil,
			jrpc2.InvalidParamsCode,
			jrpc2.InvalidParamsMsg,
		},
		{
			[]byte(`{"key": "test"}`),
			"test",
			QueueNotFoundError,
			nil,
			ListPriorityQueueErrorCode,
			ListPriorityQueueErrorMsg,
		},
	}

	for _, tt := range table {
		rescModel := &MockModel{}
		rescModel.On("FetchAll").Return(make([]interface{}, 0), nil)
		models := map[string]Model{"resources": rescModel}
		ctrl := &MockController{}
		ctrl.On("ListTimetable", tt.Key).Return(tt.Result, tt.Err)
		api := NewApiV1(models, ctrl, jrpc2.NewServer("", ""))
		result, errObj := api.ListTimetable(tt.Body)
		if errObj != nil && errObj.Code != tt.ErrCode && errObj.Message != tt.ErrMsg {
			t.Fatal(errObj.Message)
		}
		if result != nil && string(result.(json.RawMessage)) != string(tt.Result) {
			t.Fatalf("expected result to be %d, go %d", tt.Result, result)
		}
		if errObj == nil || errObj.Code != jrpc2.InvalidParamsCode {
			ctrl.AssertExpectations(t)
		}
	}
}