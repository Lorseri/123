// Code generated by mockery v2.32.4. DO NOT EDIT.

package datacoord

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockCluster is an autogenerated mock type for the Cluster type
type MockCluster struct {
	mock.Mock
}

type MockCluster_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCluster) EXPECT() *MockCluster_Expecter {
	return &MockCluster_Expecter{mock: &_m.Mock}
}

// AddImportSegment provides a mock function with given fields: ctx, req
func (_m *MockCluster) AddImportSegment(ctx context.Context, req *datapb.AddImportSegmentRequest) (*datapb.AddImportSegmentResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *datapb.AddImportSegmentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.AddImportSegmentRequest) (*datapb.AddImportSegmentResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.AddImportSegmentRequest) *datapb.AddImportSegmentResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.AddImportSegmentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *datapb.AddImportSegmentRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_AddImportSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddImportSegment'
type MockCluster_AddImportSegment_Call struct {
	*mock.Call
}

// AddImportSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - req *datapb.AddImportSegmentRequest
func (_e *MockCluster_Expecter) AddImportSegment(ctx interface{}, req interface{}) *MockCluster_AddImportSegment_Call {
	return &MockCluster_AddImportSegment_Call{Call: _e.mock.On("AddImportSegment", ctx, req)}
}

func (_c *MockCluster_AddImportSegment_Call) Run(run func(ctx context.Context, req *datapb.AddImportSegmentRequest)) *MockCluster_AddImportSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.AddImportSegmentRequest))
	})
	return _c
}

func (_c *MockCluster_AddImportSegment_Call) Return(_a0 *datapb.AddImportSegmentResponse, _a1 error) *MockCluster_AddImportSegment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCluster_AddImportSegment_Call) RunAndReturn(run func(context.Context, *datapb.AddImportSegmentRequest) (*datapb.AddImportSegmentResponse, error)) *MockCluster_AddImportSegment_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockCluster) Close() {
	_m.Called()
}

// MockCluster_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockCluster_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockCluster_Expecter) Close() *MockCluster_Close_Call {
	return &MockCluster_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockCluster_Close_Call) Run(run func()) *MockCluster_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCluster_Close_Call) Return() *MockCluster_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCluster_Close_Call) RunAndReturn(run func()) *MockCluster_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields: ctx, nodeID, channel, segments
func (_m *MockCluster) Flush(ctx context.Context, nodeID int64, channel string, segments []*datapb.SegmentInfo) error {
	ret := _m.Called(ctx, nodeID, channel, segments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, []*datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, nodeID, channel, segments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type MockCluster_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - channel string
//   - segments []*datapb.SegmentInfo
func (_e *MockCluster_Expecter) Flush(ctx interface{}, nodeID interface{}, channel interface{}, segments interface{}) *MockCluster_Flush_Call {
	return &MockCluster_Flush_Call{Call: _e.mock.On("Flush", ctx, nodeID, channel, segments)}
}

func (_c *MockCluster_Flush_Call) Run(run func(ctx context.Context, nodeID int64, channel string, segments []*datapb.SegmentInfo)) *MockCluster_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].([]*datapb.SegmentInfo))
	})
	return _c
}

func (_c *MockCluster_Flush_Call) Return(_a0 error) *MockCluster_Flush_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_Flush_Call) RunAndReturn(run func(context.Context, int64, string, []*datapb.SegmentInfo) error) *MockCluster_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// FlushChannels provides a mock function with given fields: ctx, nodeID, flushTs, channels
func (_m *MockCluster) FlushChannels(ctx context.Context, nodeID int64, flushTs uint64, channels []string) error {
	ret := _m.Called(ctx, nodeID, flushTs, channels)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64, []string) error); ok {
		r0 = rf(ctx, nodeID, flushTs, channels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_FlushChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlushChannels'
type MockCluster_FlushChannels_Call struct {
	*mock.Call
}

// FlushChannels is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - flushTs uint64
//   - channels []string
func (_e *MockCluster_Expecter) FlushChannels(ctx interface{}, nodeID interface{}, flushTs interface{}, channels interface{}) *MockCluster_FlushChannels_Call {
	return &MockCluster_FlushChannels_Call{Call: _e.mock.On("FlushChannels", ctx, nodeID, flushTs, channels)}
}

func (_c *MockCluster_FlushChannels_Call) Run(run func(ctx context.Context, nodeID int64, flushTs uint64, channels []string)) *MockCluster_FlushChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(uint64), args[3].([]string))
	})
	return _c
}

func (_c *MockCluster_FlushChannels_Call) Return(_a0 error) *MockCluster_FlushChannels_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_FlushChannels_Call) RunAndReturn(run func(context.Context, int64, uint64, []string) error) *MockCluster_FlushChannels_Call {
	_c.Call.Return(run)
	return _c
}

// GetSessions provides a mock function with given fields:
func (_m *MockCluster) GetSessions() []*Session {
	ret := _m.Called()

	var r0 []*Session
	if rf, ok := ret.Get(0).(func() []*Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Session)
		}
	}

	return r0
}

// MockCluster_GetSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSessions'
type MockCluster_GetSessions_Call struct {
	*mock.Call
}

// GetSessions is a helper method to define mock.On call
func (_e *MockCluster_Expecter) GetSessions() *MockCluster_GetSessions_Call {
	return &MockCluster_GetSessions_Call{Call: _e.mock.On("GetSessions")}
}

func (_c *MockCluster_GetSessions_Call) Run(run func()) *MockCluster_GetSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCluster_GetSessions_Call) Return(_a0 []*Session) *MockCluster_GetSessions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_GetSessions_Call) RunAndReturn(run func() []*Session) *MockCluster_GetSessions_Call {
	_c.Call.Return(run)
	return _c
}

// Import provides a mock function with given fields: ctx, nodeID, it
func (_m *MockCluster) Import(ctx context.Context, nodeID int64, it *datapb.ImportTaskRequest) {
	_m.Called(ctx, nodeID, it)
}

// MockCluster_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type MockCluster_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - it *datapb.ImportTaskRequest
func (_e *MockCluster_Expecter) Import(ctx interface{}, nodeID interface{}, it interface{}) *MockCluster_Import_Call {
	return &MockCluster_Import_Call{Call: _e.mock.On("Import", ctx, nodeID, it)}
}

func (_c *MockCluster_Import_Call) Run(run func(ctx context.Context, nodeID int64, it *datapb.ImportTaskRequest)) *MockCluster_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.ImportTaskRequest))
	})
	return _c
}

func (_c *MockCluster_Import_Call) Return() *MockCluster_Import_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCluster_Import_Call) RunAndReturn(run func(context.Context, int64, *datapb.ImportTaskRequest)) *MockCluster_Import_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: node
func (_m *MockCluster) Register(node *NodeInfo) error {
	ret := _m.Called(node)

	var r0 error
	if rf, ok := ret.Get(0).(func(*NodeInfo) error); ok {
		r0 = rf(node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockCluster_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - node *NodeInfo
func (_e *MockCluster_Expecter) Register(node interface{}) *MockCluster_Register_Call {
	return &MockCluster_Register_Call{Call: _e.mock.On("Register", node)}
}

func (_c *MockCluster_Register_Call) Run(run func(node *NodeInfo)) *MockCluster_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*NodeInfo))
	})
	return _c
}

func (_c *MockCluster_Register_Call) Return(_a0 error) *MockCluster_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_Register_Call) RunAndReturn(run func(*NodeInfo) error) *MockCluster_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Startup provides a mock function with given fields: ctx, nodes
func (_m *MockCluster) Startup(ctx context.Context, nodes []*NodeInfo) error {
	ret := _m.Called(ctx, nodes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*NodeInfo) error); ok {
		r0 = rf(ctx, nodes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_Startup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Startup'
type MockCluster_Startup_Call struct {
	*mock.Call
}

// Startup is a helper method to define mock.On call
//   - ctx context.Context
//   - nodes []*NodeInfo
func (_e *MockCluster_Expecter) Startup(ctx interface{}, nodes interface{}) *MockCluster_Startup_Call {
	return &MockCluster_Startup_Call{Call: _e.mock.On("Startup", ctx, nodes)}
}

func (_c *MockCluster_Startup_Call) Run(run func(ctx context.Context, nodes []*NodeInfo)) *MockCluster_Startup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*NodeInfo))
	})
	return _c
}

func (_c *MockCluster_Startup_Call) Return(_a0 error) *MockCluster_Startup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_Startup_Call) RunAndReturn(run func(context.Context, []*NodeInfo) error) *MockCluster_Startup_Call {
	_c.Call.Return(run)
	return _c
}

// UnRegister provides a mock function with given fields: node
func (_m *MockCluster) UnRegister(node *NodeInfo) error {
	ret := _m.Called(node)

	var r0 error
	if rf, ok := ret.Get(0).(func(*NodeInfo) error); ok {
		r0 = rf(node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_UnRegister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnRegister'
type MockCluster_UnRegister_Call struct {
	*mock.Call
}

// UnRegister is a helper method to define mock.On call
//   - node *NodeInfo
func (_e *MockCluster_Expecter) UnRegister(node interface{}) *MockCluster_UnRegister_Call {
	return &MockCluster_UnRegister_Call{Call: _e.mock.On("UnRegister", node)}
}

func (_c *MockCluster_UnRegister_Call) Run(run func(node *NodeInfo)) *MockCluster_UnRegister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*NodeInfo))
	})
	return _c
}

func (_c *MockCluster_UnRegister_Call) Return(_a0 error) *MockCluster_UnRegister_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_UnRegister_Call) RunAndReturn(run func(*NodeInfo) error) *MockCluster_UnRegister_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, ch, collectionID
func (_m *MockCluster) Watch(ctx context.Context, ch string, collectionID int64) error {
	ret := _m.Called(ctx, ch, collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, ch, collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCluster_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockCluster_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - ch string
//   - collectionID int64
func (_e *MockCluster_Expecter) Watch(ctx interface{}, ch interface{}, collectionID interface{}) *MockCluster_Watch_Call {
	return &MockCluster_Watch_Call{Call: _e.mock.On("Watch", ctx, ch, collectionID)}
}

func (_c *MockCluster_Watch_Call) Run(run func(ctx context.Context, ch string, collectionID int64)) *MockCluster_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockCluster_Watch_Call) Return(_a0 error) *MockCluster_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCluster_Watch_Call) RunAndReturn(run func(context.Context, string, int64) error) *MockCluster_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCluster creates a new instance of MockCluster. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCluster(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCluster {
	mock := &MockCluster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
