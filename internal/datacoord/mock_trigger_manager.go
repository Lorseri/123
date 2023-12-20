// Code generated by mockery v2.32.4. DO NOT EDIT.

package datacoord

import mock "github.com/stretchr/testify/mock"

// MockTriggerManager is an autogenerated mock type for the TriggerManager type
type MockTriggerManager struct {
	mock.Mock
}

type MockTriggerManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTriggerManager) EXPECT() *MockTriggerManager_Expecter {
	return &MockTriggerManager_Expecter{mock: &_m.Mock}
}

// Notify provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTriggerManager) Notify(_a0 int64, _a1 CompactionTriggerType, _a2 []CompactionView) {
	_m.Called(_a0, _a1, _a2)
}

// MockTriggerManager_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type MockTriggerManager_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//  - _a0 int64
//  - _a1 CompactionTriggerType
//  - _a2 []CompactionView
func (_e *MockTriggerManager_Expecter) Notify(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockTriggerManager_Notify_Call {
	return &MockTriggerManager_Notify_Call{Call: _e.mock.On("Notify", _a0, _a1, _a2)}
}

func (_c *MockTriggerManager_Notify_Call) Run(run func(_a0 int64, _a1 CompactionTriggerType, _a2 []CompactionView)) *MockTriggerManager_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(CompactionTriggerType), args[2].([]CompactionView))
	})
	return _c
}

func (_c *MockTriggerManager_Notify_Call) Return() *MockTriggerManager_Notify_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTriggerManager_Notify_Call) RunAndReturn(run func(int64, CompactionTriggerType, []CompactionView)) *MockTriggerManager_Notify_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTriggerManager creates a new instance of MockTriggerManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTriggerManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTriggerManager {
	mock := &MockTriggerManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
