// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_lazygrpc

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService[T interface{}] struct {
	mock.Mock
}

type MockService_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockService[T]) EXPECT() *MockService_Expecter[T] {
	return &MockService_Expecter[T]{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockService[T]) Close() {
	_m.Called()
}

// MockService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockService_Close_Call[T interface{}] struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockService_Expecter[T]) Close() *MockService_Close_Call[T] {
	return &MockService_Close_Call[T]{Call: _e.mock.On("Close")}
}

func (_c *MockService_Close_Call[T]) Run(run func()) *MockService_Close_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockService_Close_Call[T]) Return() *MockService_Close_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockService_Close_Call[T]) RunAndReturn(run func()) *MockService_Close_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetConn provides a mock function with given fields: ctx
func (_m *MockService[T]) GetConn(ctx context.Context) (*grpc.ClientConn, error) {
	ret := _m.Called(ctx)

	var r0 *grpc.ClientConn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*grpc.ClientConn, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *grpc.ClientConn); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetConn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConn'
type MockService_GetConn_Call[T interface{}] struct {
	*mock.Call
}

// GetConn is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter[T]) GetConn(ctx interface{}) *MockService_GetConn_Call[T] {
	return &MockService_GetConn_Call[T]{Call: _e.mock.On("GetConn", ctx)}
}

func (_c *MockService_GetConn_Call[T]) Run(run func(ctx context.Context)) *MockService_GetConn_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_GetConn_Call[T]) Return(_a0 *grpc.ClientConn, _a1 error) *MockService_GetConn_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetConn_Call[T]) RunAndReturn(run func(context.Context) (*grpc.ClientConn, error)) *MockService_GetConn_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetService provides a mock function with given fields: ctx
func (_m *MockService[T]) GetService(ctx context.Context) (T, error) {
	ret := _m.Called(ctx)

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (T, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) T); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockService_GetService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetService'
type MockService_GetService_Call[T interface{}] struct {
	*mock.Call
}

// GetService is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockService_Expecter[T]) GetService(ctx interface{}) *MockService_GetService_Call[T] {
	return &MockService_GetService_Call[T]{Call: _e.mock.On("GetService", ctx)}
}

func (_c *MockService_GetService_Call[T]) Run(run func(ctx context.Context)) *MockService_GetService_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockService_GetService_Call[T]) Return(_a0 T, _a1 error) *MockService_GetService_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockService_GetService_Call[T]) RunAndReturn(run func(context.Context) (T, error)) *MockService_GetService_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService[T] {
	mock := &MockService[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
