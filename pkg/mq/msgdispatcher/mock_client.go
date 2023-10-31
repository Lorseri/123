// Code generated by mockery v2.32.4. DO NOT EDIT.

package msgdispatcher

import (
	context "context"

	mqwrapper "github.com/milvus-io/milvus/pkg/mq/msgstream/mqwrapper"
	mock "github.com/stretchr/testify/mock"

	msgpb "github.com/milvus-io/milvus-proto/go-api/v2/msgpb"

	msgstream "github.com/milvus-io/milvus/pkg/mq/msgstream"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockClient) Close() {
	_m.Called()
}

// MockClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockClient_Expecter) Close() *MockClient_Close_Call {
	return &MockClient_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockClient_Close_Call) Run(run func()) *MockClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_Close_Call) Return() *MockClient_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClient_Close_Call) RunAndReturn(run func()) *MockClient_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Deregister provides a mock function with given fields: vchannel
func (_m *MockClient) Deregister(vchannel string) {
	_m.Called(vchannel)
}

// MockClient_Deregister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deregister'
type MockClient_Deregister_Call struct {
	*mock.Call
}

// Deregister is a helper method to define mock.On call
//   - vchannel string
func (_e *MockClient_Expecter) Deregister(vchannel interface{}) *MockClient_Deregister_Call {
	return &MockClient_Deregister_Call{Call: _e.mock.On("Deregister", vchannel)}
}

func (_c *MockClient_Deregister_Call) Run(run func(vchannel string)) *MockClient_Deregister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClient_Deregister_Call) Return() *MockClient_Deregister_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClient_Deregister_Call) RunAndReturn(run func(string)) *MockClient_Deregister_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: ctx, vchannel, pos, subPos
func (_m *MockClient) Register(ctx context.Context, vchannel string, pos *msgpb.MsgPosition, subPos mqwrapper.SubscriptionInitialPosition) (<-chan *msgstream.MsgPack, error) {
	ret := _m.Called(ctx, vchannel, pos, subPos)

	var r0 <-chan *msgstream.MsgPack
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *msgpb.MsgPosition, mqwrapper.SubscriptionInitialPosition) (<-chan *msgstream.MsgPack, error)); ok {
		return rf(ctx, vchannel, pos, subPos)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *msgpb.MsgPosition, mqwrapper.SubscriptionInitialPosition) <-chan *msgstream.MsgPack); ok {
		r0 = rf(ctx, vchannel, pos, subPos)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *msgstream.MsgPack)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *msgpb.MsgPosition, mqwrapper.SubscriptionInitialPosition) error); ok {
		r1 = rf(ctx, vchannel, pos, subPos)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockClient_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - ctx context.Context
//   - vchannel string
//   - pos *msgpb.MsgPosition
//   - subPos mqwrapper.SubscriptionInitialPosition
func (_e *MockClient_Expecter) Register(ctx interface{}, vchannel interface{}, pos interface{}, subPos interface{}) *MockClient_Register_Call {
	return &MockClient_Register_Call{Call: _e.mock.On("Register", ctx, vchannel, pos, subPos)}
}

func (_c *MockClient_Register_Call) Run(run func(ctx context.Context, vchannel string, pos *msgpb.MsgPosition, subPos mqwrapper.SubscriptionInitialPosition)) *MockClient_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*msgpb.MsgPosition), args[3].(mqwrapper.SubscriptionInitialPosition))
	})
	return _c
}

func (_c *MockClient_Register_Call) Return(_a0 <-chan *msgstream.MsgPack, _a1 error) *MockClient_Register_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Register_Call) RunAndReturn(run func(context.Context, string, *msgpb.MsgPosition, mqwrapper.SubscriptionInitialPosition) (<-chan *msgstream.MsgPack, error)) *MockClient_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
