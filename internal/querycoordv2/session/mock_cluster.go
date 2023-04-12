// Code generated by mockery v2.16.0. DO NOT EDIT.

package session

import (
	context "context"

	commonpb "github.com/milvus-io/milvus-proto/go-api/commonpb"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/milvuspb"

	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
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

// GetComponentStates provides a mock function with given fields: ctx, nodeID
func (_m *MockCluster) GetComponentStates(ctx context.Context, nodeID int64) (*milvuspb.ComponentStates, error) {
	ret := _m.Called(ctx, nodeID)

	var r0 *milvuspb.ComponentStates
	if rf, ok := ret.Get(0).(func(context.Context, int64) *milvuspb.ComponentStates); ok {
		r0 = rf(ctx, nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.ComponentStates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, nodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_GetComponentStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComponentStates'
type MockCluster_GetComponentStates_Call struct {
	*mock.Call
}

// GetComponentStates is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
func (_e *MockCluster_Expecter) GetComponentStates(ctx interface{}, nodeID interface{}) *MockCluster_GetComponentStates_Call {
	return &MockCluster_GetComponentStates_Call{Call: _e.mock.On("GetComponentStates", ctx, nodeID)}
}

func (_c *MockCluster_GetComponentStates_Call) Run(run func(ctx context.Context, nodeID int64)) *MockCluster_GetComponentStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockCluster_GetComponentStates_Call) Return(_a0 *milvuspb.ComponentStates, _a1 error) *MockCluster_GetComponentStates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDataDistribution provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) GetDataDistribution(ctx context.Context, nodeID int64, req *querypb.GetDataDistributionRequest) (*querypb.GetDataDistributionResponse, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *querypb.GetDataDistributionResponse
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.GetDataDistributionRequest) *querypb.GetDataDistributionResponse); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querypb.GetDataDistributionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.GetDataDistributionRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_GetDataDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataDistribution'
type MockCluster_GetDataDistribution_Call struct {
	*mock.Call
}

// GetDataDistribution is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.GetDataDistributionRequest
func (_e *MockCluster_Expecter) GetDataDistribution(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_GetDataDistribution_Call {
	return &MockCluster_GetDataDistribution_Call{Call: _e.mock.On("GetDataDistribution", ctx, nodeID, req)}
}

func (_c *MockCluster_GetDataDistribution_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.GetDataDistributionRequest)) *MockCluster_GetDataDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.GetDataDistributionRequest))
	})
	return _c
}

func (_c *MockCluster_GetDataDistribution_Call) Return(_a0 *querypb.GetDataDistributionResponse, _a1 error) *MockCluster_GetDataDistribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMetrics provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) GetMetrics(ctx context.Context, nodeID int64, req *milvuspb.GetMetricsRequest) (*milvuspb.GetMetricsResponse, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *milvuspb.GetMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int64, *milvuspb.GetMetricsRequest) *milvuspb.GetMetricsResponse); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.GetMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *milvuspb.GetMetricsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_GetMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetrics'
type MockCluster_GetMetrics_Call struct {
	*mock.Call
}

// GetMetrics is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *milvuspb.GetMetricsRequest
func (_e *MockCluster_Expecter) GetMetrics(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_GetMetrics_Call {
	return &MockCluster_GetMetrics_Call{Call: _e.mock.On("GetMetrics", ctx, nodeID, req)}
}

func (_c *MockCluster_GetMetrics_Call) Run(run func(ctx context.Context, nodeID int64, req *milvuspb.GetMetricsRequest)) *MockCluster_GetMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*milvuspb.GetMetricsRequest))
	})
	return _c
}

func (_c *MockCluster_GetMetrics_Call) Return(_a0 *milvuspb.GetMetricsResponse, _a1 error) *MockCluster_GetMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadPartitions provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) LoadPartitions(ctx context.Context, nodeID int64, req *querypb.LoadPartitionsRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.LoadPartitionsRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.LoadPartitionsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_LoadPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadPartitions'
type MockCluster_LoadPartitions_Call struct {
	*mock.Call
}

// LoadPartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.LoadPartitionsRequest
func (_e *MockCluster_Expecter) LoadPartitions(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_LoadPartitions_Call {
	return &MockCluster_LoadPartitions_Call{Call: _e.mock.On("LoadPartitions", ctx, nodeID, req)}
}

func (_c *MockCluster_LoadPartitions_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.LoadPartitionsRequest)) *MockCluster_LoadPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.LoadPartitionsRequest))
	})
	return _c
}

func (_c *MockCluster_LoadPartitions_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_LoadPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadSegments provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) LoadSegments(ctx context.Context, nodeID int64, req *querypb.LoadSegmentsRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.LoadSegmentsRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.LoadSegmentsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_LoadSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSegments'
type MockCluster_LoadSegments_Call struct {
	*mock.Call
}

// LoadSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.LoadSegmentsRequest
func (_e *MockCluster_Expecter) LoadSegments(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_LoadSegments_Call {
	return &MockCluster_LoadSegments_Call{Call: _e.mock.On("LoadSegments", ctx, nodeID, req)}
}

func (_c *MockCluster_LoadSegments_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.LoadSegmentsRequest)) *MockCluster_LoadSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.LoadSegmentsRequest))
	})
	return _c
}

func (_c *MockCluster_LoadSegments_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_LoadSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleasePartitions provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) ReleasePartitions(ctx context.Context, nodeID int64, req *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.ReleasePartitionsRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.ReleasePartitionsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_ReleasePartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleasePartitions'
type MockCluster_ReleasePartitions_Call struct {
	*mock.Call
}

// ReleasePartitions is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.ReleasePartitionsRequest
func (_e *MockCluster_Expecter) ReleasePartitions(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_ReleasePartitions_Call {
	return &MockCluster_ReleasePartitions_Call{Call: _e.mock.On("ReleasePartitions", ctx, nodeID, req)}
}

func (_c *MockCluster_ReleasePartitions_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.ReleasePartitionsRequest)) *MockCluster_ReleasePartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.ReleasePartitionsRequest))
	})
	return _c
}

func (_c *MockCluster_ReleasePartitions_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_ReleasePartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleaseSegments provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) ReleaseSegments(ctx context.Context, nodeID int64, req *querypb.ReleaseSegmentsRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.ReleaseSegmentsRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.ReleaseSegmentsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_ReleaseSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseSegments'
type MockCluster_ReleaseSegments_Call struct {
	*mock.Call
}

// ReleaseSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.ReleaseSegmentsRequest
func (_e *MockCluster_Expecter) ReleaseSegments(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_ReleaseSegments_Call {
	return &MockCluster_ReleaseSegments_Call{Call: _e.mock.On("ReleaseSegments", ctx, nodeID, req)}
}

func (_c *MockCluster_ReleaseSegments_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.ReleaseSegmentsRequest)) *MockCluster_ReleaseSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.ReleaseSegmentsRequest))
	})
	return _c
}

func (_c *MockCluster_ReleaseSegments_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_ReleaseSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *MockCluster) Start(ctx context.Context) {
	_m.Called(ctx)
}

// MockCluster_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockCluster_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCluster_Expecter) Start(ctx interface{}) *MockCluster_Start_Call {
	return &MockCluster_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *MockCluster_Start_Call) Run(run func(ctx context.Context)) *MockCluster_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCluster_Start_Call) Return() *MockCluster_Start_Call {
	_c.Call.Return()
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockCluster) Stop() {
	_m.Called()
}

// MockCluster_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockCluster_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockCluster_Expecter) Stop() *MockCluster_Stop_Call {
	return &MockCluster_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockCluster_Stop_Call) Run(run func()) *MockCluster_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCluster_Stop_Call) Return() *MockCluster_Stop_Call {
	_c.Call.Return()
	return _c
}

// SyncDistribution provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) SyncDistribution(ctx context.Context, nodeID int64, req *querypb.SyncDistributionRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.SyncDistributionRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.SyncDistributionRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_SyncDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncDistribution'
type MockCluster_SyncDistribution_Call struct {
	*mock.Call
}

// SyncDistribution is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.SyncDistributionRequest
func (_e *MockCluster_Expecter) SyncDistribution(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_SyncDistribution_Call {
	return &MockCluster_SyncDistribution_Call{Call: _e.mock.On("SyncDistribution", ctx, nodeID, req)}
}

func (_c *MockCluster_SyncDistribution_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.SyncDistributionRequest)) *MockCluster_SyncDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.SyncDistributionRequest))
	})
	return _c
}

func (_c *MockCluster_SyncDistribution_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_SyncDistribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UnsubDmChannel provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) UnsubDmChannel(ctx context.Context, nodeID int64, req *querypb.UnsubDmChannelRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.UnsubDmChannelRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.UnsubDmChannelRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_UnsubDmChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsubDmChannel'
type MockCluster_UnsubDmChannel_Call struct {
	*mock.Call
}

// UnsubDmChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.UnsubDmChannelRequest
func (_e *MockCluster_Expecter) UnsubDmChannel(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_UnsubDmChannel_Call {
	return &MockCluster_UnsubDmChannel_Call{Call: _e.mock.On("UnsubDmChannel", ctx, nodeID, req)}
}

func (_c *MockCluster_UnsubDmChannel_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.UnsubDmChannelRequest)) *MockCluster_UnsubDmChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.UnsubDmChannelRequest))
	})
	return _c
}

func (_c *MockCluster_UnsubDmChannel_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_UnsubDmChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// WatchDmChannels provides a mock function with given fields: ctx, nodeID, req
func (_m *MockCluster) WatchDmChannels(ctx context.Context, nodeID int64, req *querypb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	ret := _m.Called(ctx, nodeID, req)

	var r0 *commonpb.Status
	if rf, ok := ret.Get(0).(func(context.Context, int64, *querypb.WatchDmChannelsRequest) *commonpb.Status); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *querypb.WatchDmChannelsRequest) error); ok {
		r1 = rf(ctx, nodeID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCluster_WatchDmChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchDmChannels'
type MockCluster_WatchDmChannels_Call struct {
	*mock.Call
}

// WatchDmChannels is a helper method to define mock.On call
//   - ctx context.Context
//   - nodeID int64
//   - req *querypb.WatchDmChannelsRequest
func (_e *MockCluster_Expecter) WatchDmChannels(ctx interface{}, nodeID interface{}, req interface{}) *MockCluster_WatchDmChannels_Call {
	return &MockCluster_WatchDmChannels_Call{Call: _e.mock.On("WatchDmChannels", ctx, nodeID, req)}
}

func (_c *MockCluster_WatchDmChannels_Call) Run(run func(ctx context.Context, nodeID int64, req *querypb.WatchDmChannelsRequest)) *MockCluster_WatchDmChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*querypb.WatchDmChannelsRequest))
	})
	return _c
}

func (_c *MockCluster_WatchDmChannels_Call) Return(_a0 *commonpb.Status, _a1 error) *MockCluster_WatchDmChannels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockCluster interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCluster creates a new instance of MockCluster. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCluster(t mockConstructorTestingTNewMockCluster) *MockCluster {
	mock := &MockCluster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
