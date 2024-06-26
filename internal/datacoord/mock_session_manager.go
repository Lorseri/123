// Code generated by mockery v2.30.1. DO NOT EDIT.

package datacoord

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"

	typeutil "github.com/milvus-io/milvus/pkg/util/typeutil"
)

// MockSessionManager is an autogenerated mock type for the SessionManager type
type MockSessionManager struct {
	mock.Mock
}

type MockSessionManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSessionManager) EXPECT() *MockSessionManager_Expecter {
	return &MockSessionManager_Expecter{mock: &_m.Mock}
}

// AddSession provides a mock function with given fields: node
func (_m *MockSessionManager) AddSession(node *NodeInfo) {
	_m.Called(node)
}

// MockSessionManager_AddSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSession'
type MockSessionManager_AddSession_Call struct {
	*mock.Call
}

// AddSession is a helper method to define mock.On call
//  - node *NodeInfo
func (_e *MockSessionManager_Expecter) AddSession(node interface{}) *MockSessionManager_AddSession_Call {
	return &MockSessionManager_AddSession_Call{Call: _e.mock.On("AddSession", node)}
}

func (_c *MockSessionManager_AddSession_Call) Run(run func(node *NodeInfo)) *MockSessionManager_AddSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*NodeInfo))
	})
	return _c
}

func (_c *MockSessionManager_AddSession_Call) Return() *MockSessionManager_AddSession_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSessionManager_AddSession_Call) RunAndReturn(run func(*NodeInfo)) *MockSessionManager_AddSession_Call {
	_c.Call.Return(run)
	return _c
}

// CheckChannelOperationProgress provides a mock function with given fields: ctx, nodeID, info
func (_m *MockSessionManager) CheckChannelOperationProgress(ctx context.Context, nodeID int64, info *datapb.ChannelWatchInfo) (*datapb.ChannelOperationProgressResponse, error) {
	ret := _m.Called(ctx, nodeID, info)

	var r0 *datapb.ChannelOperationProgressResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *datapb.ChannelWatchInfo) (*datapb.ChannelOperationProgressResponse, error)); ok {
		return rf(ctx, nodeID, info)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *datapb.ChannelWatchInfo) *datapb.ChannelOperationProgressResponse); ok {
		r0 = rf(ctx, nodeID, info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.ChannelOperationProgressResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *datapb.ChannelWatchInfo) error); ok {
		r1 = rf(ctx, nodeID, info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSessionManager_CheckChannelOperationProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckChannelOperationProgress'
type MockSessionManager_CheckChannelOperationProgress_Call struct {
	*mock.Call
}

// CheckChannelOperationProgress is a helper method to define mock.On call
//  - ctx context.Context
//  - nodeID int64
//  - info *datapb.ChannelWatchInfo
func (_e *MockSessionManager_Expecter) CheckChannelOperationProgress(ctx interface{}, nodeID interface{}, info interface{}) *MockSessionManager_CheckChannelOperationProgress_Call {
	return &MockSessionManager_CheckChannelOperationProgress_Call{Call: _e.mock.On("CheckChannelOperationProgress", ctx, nodeID, info)}
}

func (_c *MockSessionManager_CheckChannelOperationProgress_Call) Run(run func(ctx context.Context, nodeID int64, info *datapb.ChannelWatchInfo)) *MockSessionManager_CheckChannelOperationProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.ChannelWatchInfo))
	})
	return _c
}

func (_c *MockSessionManager_CheckChannelOperationProgress_Call) Return(_a0 *datapb.ChannelOperationProgressResponse, _a1 error) *MockSessionManager_CheckChannelOperationProgress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSessionManager_CheckChannelOperationProgress_Call) RunAndReturn(run func(context.Context, int64, *datapb.ChannelWatchInfo) (*datapb.ChannelOperationProgressResponse, error)) *MockSessionManager_CheckChannelOperationProgress_Call {
	_c.Call.Return(run)
	return _c
}

// CheckHealth provides a mock function with given fields: ctx
func (_m *MockSessionManager) CheckHealth(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_CheckHealth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckHealth'
type MockSessionManager_CheckHealth_Call struct {
	*mock.Call
}

// CheckHealth is a helper method to define mock.On call
//  - ctx context.Context
func (_e *MockSessionManager_Expecter) CheckHealth(ctx interface{}) *MockSessionManager_CheckHealth_Call {
	return &MockSessionManager_CheckHealth_Call{Call: _e.mock.On("CheckHealth", ctx)}
}

func (_c *MockSessionManager_CheckHealth_Call) Run(run func(ctx context.Context)) *MockSessionManager_CheckHealth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSessionManager_CheckHealth_Call) Return(_a0 error) *MockSessionManager_CheckHealth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_CheckHealth_Call) RunAndReturn(run func(context.Context) error) *MockSessionManager_CheckHealth_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockSessionManager) Close() {
	_m.Called()
}

// MockSessionManager_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSessionManager_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSessionManager_Expecter) Close() *MockSessionManager_Close_Call {
	return &MockSessionManager_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSessionManager_Close_Call) Run(run func()) *MockSessionManager_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionManager_Close_Call) Return() *MockSessionManager_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSessionManager_Close_Call) RunAndReturn(run func()) *MockSessionManager_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Compaction provides a mock function with given fields: ctx, nodeID, plan
func (_m *MockSessionManager) Compaction(ctx context.Context, nodeID int64, plan *datapb.CompactionPlan) error {
	ret := _m.Called(ctx, nodeID, plan)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *datapb.CompactionPlan) error); ok {
		r0 = rf(ctx, nodeID, plan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_Compaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compaction'
type MockSessionManager_Compaction_Call struct {
	*mock.Call
}

// Compaction is a helper method to define mock.On call
//  - ctx context.Context
//  - nodeID int64
//  - plan *datapb.CompactionPlan
func (_e *MockSessionManager_Expecter) Compaction(ctx interface{}, nodeID interface{}, plan interface{}) *MockSessionManager_Compaction_Call {
	return &MockSessionManager_Compaction_Call{Call: _e.mock.On("Compaction", ctx, nodeID, plan)}
}

func (_c *MockSessionManager_Compaction_Call) Run(run func(ctx context.Context, nodeID int64, plan *datapb.CompactionPlan)) *MockSessionManager_Compaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.CompactionPlan))
	})
	return _c
}

func (_c *MockSessionManager_Compaction_Call) Return(_a0 error) *MockSessionManager_Compaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_Compaction_Call) RunAndReturn(run func(context.Context, int64, *datapb.CompactionPlan) error) *MockSessionManager_Compaction_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSession provides a mock function with given fields: node
func (_m *MockSessionManager) DeleteSession(node *NodeInfo) {
	_m.Called(node)
}

// MockSessionManager_DeleteSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSession'
type MockSessionManager_DeleteSession_Call struct {
	*mock.Call
}

// DeleteSession is a helper method to define mock.On call
//  - node *NodeInfo
func (_e *MockSessionManager_Expecter) DeleteSession(node interface{}) *MockSessionManager_DeleteSession_Call {
	return &MockSessionManager_DeleteSession_Call{Call: _e.mock.On("DeleteSession", node)}
}

func (_c *MockSessionManager_DeleteSession_Call) Run(run func(node *NodeInfo)) *MockSessionManager_DeleteSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*NodeInfo))
	})
	return _c
}

func (_c *MockSessionManager_DeleteSession_Call) Return() *MockSessionManager_DeleteSession_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSessionManager_DeleteSession_Call) RunAndReturn(run func(*NodeInfo)) *MockSessionManager_DeleteSession_Call {
	_c.Call.Return(run)
	return _c
}

// DropImport provides a mock function with given fields: nodeID, in
func (_m *MockSessionManager) DropImport(nodeID int64, in *datapb.DropImportRequest) error {
	ret := _m.Called(nodeID, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.DropImportRequest) error); ok {
		r0 = rf(nodeID, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_DropImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropImport'
type MockSessionManager_DropImport_Call struct {
	*mock.Call
}

// DropImport is a helper method to define mock.On call
//  - nodeID int64
//  - in *datapb.DropImportRequest
func (_e *MockSessionManager_Expecter) DropImport(nodeID interface{}, in interface{}) *MockSessionManager_DropImport_Call {
	return &MockSessionManager_DropImport_Call{Call: _e.mock.On("DropImport", nodeID, in)}
}

func (_c *MockSessionManager_DropImport_Call) Run(run func(nodeID int64, in *datapb.DropImportRequest)) *MockSessionManager_DropImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.DropImportRequest))
	})
	return _c
}

func (_c *MockSessionManager_DropImport_Call) Return(_a0 error) *MockSessionManager_DropImport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_DropImport_Call) RunAndReturn(run func(int64, *datapb.DropImportRequest) error) *MockSessionManager_DropImport_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields: ctx, nodeID, req
func (_m *MockSessionManager) Flush(ctx context.Context, nodeID int64, req *datapb.FlushSegmentsRequest) {
	_m.Called(ctx, nodeID, req)
}

// MockSessionManager_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type MockSessionManager_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//  - ctx context.Context
//  - nodeID int64
//  - req *datapb.FlushSegmentsRequest
func (_e *MockSessionManager_Expecter) Flush(ctx interface{}, nodeID interface{}, req interface{}) *MockSessionManager_Flush_Call {
	return &MockSessionManager_Flush_Call{Call: _e.mock.On("Flush", ctx, nodeID, req)}
}

func (_c *MockSessionManager_Flush_Call) Run(run func(ctx context.Context, nodeID int64, req *datapb.FlushSegmentsRequest)) *MockSessionManager_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.FlushSegmentsRequest))
	})
	return _c
}

func (_c *MockSessionManager_Flush_Call) Return() *MockSessionManager_Flush_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSessionManager_Flush_Call) RunAndReturn(run func(context.Context, int64, *datapb.FlushSegmentsRequest)) *MockSessionManager_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// FlushChannels provides a mock function with given fields: ctx, nodeID, req
func (_m *MockSessionManager) FlushChannels(ctx context.Context, nodeID int64, req *datapb.FlushChannelsRequest) error {
	ret := _m.Called(ctx, nodeID, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *datapb.FlushChannelsRequest) error); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_FlushChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlushChannels'
type MockSessionManager_FlushChannels_Call struct {
	*mock.Call
}

// FlushChannels is a helper method to define mock.On call
//  - ctx context.Context
//  - nodeID int64
//  - req *datapb.FlushChannelsRequest
func (_e *MockSessionManager_Expecter) FlushChannels(ctx interface{}, nodeID interface{}, req interface{}) *MockSessionManager_FlushChannels_Call {
	return &MockSessionManager_FlushChannels_Call{Call: _e.mock.On("FlushChannels", ctx, nodeID, req)}
}

func (_c *MockSessionManager_FlushChannels_Call) Run(run func(ctx context.Context, nodeID int64, req *datapb.FlushChannelsRequest)) *MockSessionManager_FlushChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.FlushChannelsRequest))
	})
	return _c
}

func (_c *MockSessionManager_FlushChannels_Call) Return(_a0 error) *MockSessionManager_FlushChannels_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_FlushChannels_Call) RunAndReturn(run func(context.Context, int64, *datapb.FlushChannelsRequest) error) *MockSessionManager_FlushChannels_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompactionPlansResults provides a mock function with given fields:
func (_m *MockSessionManager) GetCompactionPlansResults() (map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult], error) {
	ret := _m.Called()

	var r0 map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult]
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult], error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult])
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSessionManager_GetCompactionPlansResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompactionPlansResults'
type MockSessionManager_GetCompactionPlansResults_Call struct {
	*mock.Call
}

// GetCompactionPlansResults is a helper method to define mock.On call
func (_e *MockSessionManager_Expecter) GetCompactionPlansResults() *MockSessionManager_GetCompactionPlansResults_Call {
	return &MockSessionManager_GetCompactionPlansResults_Call{Call: _e.mock.On("GetCompactionPlansResults")}
}

func (_c *MockSessionManager_GetCompactionPlansResults_Call) Run(run func()) *MockSessionManager_GetCompactionPlansResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionManager_GetCompactionPlansResults_Call) Return(_a0 map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult], _a1 error) *MockSessionManager_GetCompactionPlansResults_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSessionManager_GetCompactionPlansResults_Call) RunAndReturn(run func() (map[int64]*typeutil.Pair[int64, *datapb.CompactionPlanResult], error)) *MockSessionManager_GetCompactionPlansResults_Call {
	_c.Call.Return(run)
	return _c
}

// GetSessionIDs provides a mock function with given fields:
func (_m *MockSessionManager) GetSessionIDs() []int64 {
	ret := _m.Called()

	var r0 []int64
	if rf, ok := ret.Get(0).(func() []int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockSessionManager_GetSessionIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSessionIDs'
type MockSessionManager_GetSessionIDs_Call struct {
	*mock.Call
}

// GetSessionIDs is a helper method to define mock.On call
func (_e *MockSessionManager_Expecter) GetSessionIDs() *MockSessionManager_GetSessionIDs_Call {
	return &MockSessionManager_GetSessionIDs_Call{Call: _e.mock.On("GetSessionIDs")}
}

func (_c *MockSessionManager_GetSessionIDs_Call) Run(run func()) *MockSessionManager_GetSessionIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionManager_GetSessionIDs_Call) Return(_a0 []int64) *MockSessionManager_GetSessionIDs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_GetSessionIDs_Call) RunAndReturn(run func() []int64) *MockSessionManager_GetSessionIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetSessions provides a mock function with given fields:
func (_m *MockSessionManager) GetSessions() []*Session {
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

// MockSessionManager_GetSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSessions'
type MockSessionManager_GetSessions_Call struct {
	*mock.Call
}

// GetSessions is a helper method to define mock.On call
func (_e *MockSessionManager_Expecter) GetSessions() *MockSessionManager_GetSessions_Call {
	return &MockSessionManager_GetSessions_Call{Call: _e.mock.On("GetSessions")}
}

func (_c *MockSessionManager_GetSessions_Call) Run(run func()) *MockSessionManager_GetSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionManager_GetSessions_Call) Return(_a0 []*Session) *MockSessionManager_GetSessions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_GetSessions_Call) RunAndReturn(run func() []*Session) *MockSessionManager_GetSessions_Call {
	_c.Call.Return(run)
	return _c
}

// ImportV2 provides a mock function with given fields: nodeID, in
func (_m *MockSessionManager) ImportV2(nodeID int64, in *datapb.ImportRequest) error {
	ret := _m.Called(nodeID, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.ImportRequest) error); ok {
		r0 = rf(nodeID, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_ImportV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ImportV2'
type MockSessionManager_ImportV2_Call struct {
	*mock.Call
}

// ImportV2 is a helper method to define mock.On call
//  - nodeID int64
//  - in *datapb.ImportRequest
func (_e *MockSessionManager_Expecter) ImportV2(nodeID interface{}, in interface{}) *MockSessionManager_ImportV2_Call {
	return &MockSessionManager_ImportV2_Call{Call: _e.mock.On("ImportV2", nodeID, in)}
}

func (_c *MockSessionManager_ImportV2_Call) Run(run func(nodeID int64, in *datapb.ImportRequest)) *MockSessionManager_ImportV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.ImportRequest))
	})
	return _c
}

func (_c *MockSessionManager_ImportV2_Call) Return(_a0 error) *MockSessionManager_ImportV2_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_ImportV2_Call) RunAndReturn(run func(int64, *datapb.ImportRequest) error) *MockSessionManager_ImportV2_Call {
	_c.Call.Return(run)
	return _c
}

// NotifyChannelOperation provides a mock function with given fields: ctx, nodeID, req
func (_m *MockSessionManager) NotifyChannelOperation(ctx context.Context, nodeID int64, req *datapb.ChannelOperationsRequest) error {
	ret := _m.Called(ctx, nodeID, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *datapb.ChannelOperationsRequest) error); ok {
		r0 = rf(ctx, nodeID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_NotifyChannelOperation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NotifyChannelOperation'
type MockSessionManager_NotifyChannelOperation_Call struct {
	*mock.Call
}

// NotifyChannelOperation is a helper method to define mock.On call
//  - ctx context.Context
//  - nodeID int64
//  - req *datapb.ChannelOperationsRequest
func (_e *MockSessionManager_Expecter) NotifyChannelOperation(ctx interface{}, nodeID interface{}, req interface{}) *MockSessionManager_NotifyChannelOperation_Call {
	return &MockSessionManager_NotifyChannelOperation_Call{Call: _e.mock.On("NotifyChannelOperation", ctx, nodeID, req)}
}

func (_c *MockSessionManager_NotifyChannelOperation_Call) Run(run func(ctx context.Context, nodeID int64, req *datapb.ChannelOperationsRequest)) *MockSessionManager_NotifyChannelOperation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(*datapb.ChannelOperationsRequest))
	})
	return _c
}

func (_c *MockSessionManager_NotifyChannelOperation_Call) Return(_a0 error) *MockSessionManager_NotifyChannelOperation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_NotifyChannelOperation_Call) RunAndReturn(run func(context.Context, int64, *datapb.ChannelOperationsRequest) error) *MockSessionManager_NotifyChannelOperation_Call {
	_c.Call.Return(run)
	return _c
}

// PreImport provides a mock function with given fields: nodeID, in
func (_m *MockSessionManager) PreImport(nodeID int64, in *datapb.PreImportRequest) error {
	ret := _m.Called(nodeID, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.PreImportRequest) error); ok {
		r0 = rf(nodeID, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_PreImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PreImport'
type MockSessionManager_PreImport_Call struct {
	*mock.Call
}

// PreImport is a helper method to define mock.On call
//  - nodeID int64
//  - in *datapb.PreImportRequest
func (_e *MockSessionManager_Expecter) PreImport(nodeID interface{}, in interface{}) *MockSessionManager_PreImport_Call {
	return &MockSessionManager_PreImport_Call{Call: _e.mock.On("PreImport", nodeID, in)}
}

func (_c *MockSessionManager_PreImport_Call) Run(run func(nodeID int64, in *datapb.PreImportRequest)) *MockSessionManager_PreImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.PreImportRequest))
	})
	return _c
}

func (_c *MockSessionManager_PreImport_Call) Return(_a0 error) *MockSessionManager_PreImport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_PreImport_Call) RunAndReturn(run func(int64, *datapb.PreImportRequest) error) *MockSessionManager_PreImport_Call {
	_c.Call.Return(run)
	return _c
}

// QueryImport provides a mock function with given fields: nodeID, in
func (_m *MockSessionManager) QueryImport(nodeID int64, in *datapb.QueryImportRequest) (*datapb.QueryImportResponse, error) {
	ret := _m.Called(nodeID, in)

	var r0 *datapb.QueryImportResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.QueryImportRequest) (*datapb.QueryImportResponse, error)); ok {
		return rf(nodeID, in)
	}
	if rf, ok := ret.Get(0).(func(int64, *datapb.QueryImportRequest) *datapb.QueryImportResponse); ok {
		r0 = rf(nodeID, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.QueryImportResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, *datapb.QueryImportRequest) error); ok {
		r1 = rf(nodeID, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSessionManager_QueryImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryImport'
type MockSessionManager_QueryImport_Call struct {
	*mock.Call
}

// QueryImport is a helper method to define mock.On call
//  - nodeID int64
//  - in *datapb.QueryImportRequest
func (_e *MockSessionManager_Expecter) QueryImport(nodeID interface{}, in interface{}) *MockSessionManager_QueryImport_Call {
	return &MockSessionManager_QueryImport_Call{Call: _e.mock.On("QueryImport", nodeID, in)}
}

func (_c *MockSessionManager_QueryImport_Call) Run(run func(nodeID int64, in *datapb.QueryImportRequest)) *MockSessionManager_QueryImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.QueryImportRequest))
	})
	return _c
}

func (_c *MockSessionManager_QueryImport_Call) Return(_a0 *datapb.QueryImportResponse, _a1 error) *MockSessionManager_QueryImport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSessionManager_QueryImport_Call) RunAndReturn(run func(int64, *datapb.QueryImportRequest) (*datapb.QueryImportResponse, error)) *MockSessionManager_QueryImport_Call {
	_c.Call.Return(run)
	return _c
}

// QueryPreImport provides a mock function with given fields: nodeID, in
func (_m *MockSessionManager) QueryPreImport(nodeID int64, in *datapb.QueryPreImportRequest) (*datapb.QueryPreImportResponse, error) {
	ret := _m.Called(nodeID, in)

	var r0 *datapb.QueryPreImportResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.QueryPreImportRequest) (*datapb.QueryPreImportResponse, error)); ok {
		return rf(nodeID, in)
	}
	if rf, ok := ret.Get(0).(func(int64, *datapb.QueryPreImportRequest) *datapb.QueryPreImportResponse); ok {
		r0 = rf(nodeID, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.QueryPreImportResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, *datapb.QueryPreImportRequest) error); ok {
		r1 = rf(nodeID, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSessionManager_QueryPreImport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryPreImport'
type MockSessionManager_QueryPreImport_Call struct {
	*mock.Call
}

// QueryPreImport is a helper method to define mock.On call
//  - nodeID int64
//  - in *datapb.QueryPreImportRequest
func (_e *MockSessionManager_Expecter) QueryPreImport(nodeID interface{}, in interface{}) *MockSessionManager_QueryPreImport_Call {
	return &MockSessionManager_QueryPreImport_Call{Call: _e.mock.On("QueryPreImport", nodeID, in)}
}

func (_c *MockSessionManager_QueryPreImport_Call) Run(run func(nodeID int64, in *datapb.QueryPreImportRequest)) *MockSessionManager_QueryPreImport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.QueryPreImportRequest))
	})
	return _c
}

func (_c *MockSessionManager_QueryPreImport_Call) Return(_a0 *datapb.QueryPreImportResponse, _a1 error) *MockSessionManager_QueryPreImport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSessionManager_QueryPreImport_Call) RunAndReturn(run func(int64, *datapb.QueryPreImportRequest) (*datapb.QueryPreImportResponse, error)) *MockSessionManager_QueryPreImport_Call {
	_c.Call.Return(run)
	return _c
}

// QuerySlot provides a mock function with given fields: nodeID
func (_m *MockSessionManager) QuerySlot(nodeID int64) (*datapb.QuerySlotResponse, error) {
	ret := _m.Called(nodeID)

	var r0 *datapb.QuerySlotResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*datapb.QuerySlotResponse, error)); ok {
		return rf(nodeID)
	}
	if rf, ok := ret.Get(0).(func(int64) *datapb.QuerySlotResponse); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.QuerySlotResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSessionManager_QuerySlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QuerySlot'
type MockSessionManager_QuerySlot_Call struct {
	*mock.Call
}

// QuerySlot is a helper method to define mock.On call
//  - nodeID int64
func (_e *MockSessionManager_Expecter) QuerySlot(nodeID interface{}) *MockSessionManager_QuerySlot_Call {
	return &MockSessionManager_QuerySlot_Call{Call: _e.mock.On("QuerySlot", nodeID)}
}

func (_c *MockSessionManager_QuerySlot_Call) Run(run func(nodeID int64)) *MockSessionManager_QuerySlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockSessionManager_QuerySlot_Call) Return(_a0 *datapb.QuerySlotResponse, _a1 error) *MockSessionManager_QuerySlot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSessionManager_QuerySlot_Call) RunAndReturn(run func(int64) (*datapb.QuerySlotResponse, error)) *MockSessionManager_QuerySlot_Call {
	_c.Call.Return(run)
	return _c
}

// SyncSegments provides a mock function with given fields: nodeID, req
func (_m *MockSessionManager) SyncSegments(nodeID int64, req *datapb.SyncSegmentsRequest) error {
	ret := _m.Called(nodeID, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, *datapb.SyncSegmentsRequest) error); ok {
		r0 = rf(nodeID, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSessionManager_SyncSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncSegments'
type MockSessionManager_SyncSegments_Call struct {
	*mock.Call
}

// SyncSegments is a helper method to define mock.On call
//  - nodeID int64
//  - req *datapb.SyncSegmentsRequest
func (_e *MockSessionManager_Expecter) SyncSegments(nodeID interface{}, req interface{}) *MockSessionManager_SyncSegments_Call {
	return &MockSessionManager_SyncSegments_Call{Call: _e.mock.On("SyncSegments", nodeID, req)}
}

func (_c *MockSessionManager_SyncSegments_Call) Run(run func(nodeID int64, req *datapb.SyncSegmentsRequest)) *MockSessionManager_SyncSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(*datapb.SyncSegmentsRequest))
	})
	return _c
}

func (_c *MockSessionManager_SyncSegments_Call) Return(_a0 error) *MockSessionManager_SyncSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionManager_SyncSegments_Call) RunAndReturn(run func(int64, *datapb.SyncSegmentsRequest) error) *MockSessionManager_SyncSegments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSessionManager creates a new instance of MockSessionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSessionManager(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MockSessionManager {
	mock := &MockSessionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
