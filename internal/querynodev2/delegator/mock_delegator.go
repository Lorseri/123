// Code generated by mockery v2.16.0. DO NOT EDIT.

package delegator

import (
	context "context"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"
	mock "github.com/stretchr/testify/mock"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
)

// MockShardDelegator is an autogenerated mock type for the ShardDelegator type
type MockShardDelegator struct {
	mock.Mock
}

type MockShardDelegator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShardDelegator) EXPECT() *MockShardDelegator_Expecter {
	return &MockShardDelegator_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockShardDelegator) Close() {
	_m.Called()
}

// MockShardDelegator_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockShardDelegator_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) Close() *MockShardDelegator_Close_Call {
	return &MockShardDelegator_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockShardDelegator_Close_Call) Run(run func()) *MockShardDelegator_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_Close_Call) Return() *MockShardDelegator_Close_Call {
	_c.Call.Return()
	return _c
}

// Collection provides a mock function with given fields:
func (_m *MockShardDelegator) Collection() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockShardDelegator_Collection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collection'
type MockShardDelegator_Collection_Call struct {
	*mock.Call
}

// Collection is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) Collection() *MockShardDelegator_Collection_Call {
	return &MockShardDelegator_Collection_Call{Call: _e.mock.On("Collection")}
}

func (_c *MockShardDelegator_Collection_Call) Run(run func()) *MockShardDelegator_Collection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_Collection_Call) Return(_a0 int64) *MockShardDelegator_Collection_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetDistribution provides a mock function with given fields:
func (_m *MockShardDelegator) GetDistribution() *distribution {
	ret := _m.Called()

	var r0 *distribution
	if rf, ok := ret.Get(0).(func() *distribution); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*distribution)
		}
	}

	return r0
}

// MockShardDelegator_GetDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDistribution'
type MockShardDelegator_GetDistribution_Call struct {
	*mock.Call
}

// GetDistribution is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) GetDistribution() *MockShardDelegator_GetDistribution_Call {
	return &MockShardDelegator_GetDistribution_Call{Call: _e.mock.On("GetDistribution")}
}

func (_c *MockShardDelegator_GetDistribution_Call) Run(run func()) *MockShardDelegator_GetDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_GetDistribution_Call) Return(_a0 *distribution) *MockShardDelegator_GetDistribution_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetStatistics provides a mock function with given fields: ctx, req
func (_m *MockShardDelegator) GetStatistics(ctx context.Context, req *querypb.GetStatisticsRequest) ([]*internalpb.GetStatisticsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 []*internalpb.GetStatisticsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.GetStatisticsRequest) []*internalpb.GetStatisticsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internalpb.GetStatisticsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.GetStatisticsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShardDelegator_GetStatistics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatistics'
type MockShardDelegator_GetStatistics_Call struct {
	*mock.Call
}

// GetStatistics is a helper method to define mock.On call
//  - ctx context.Context
//  - req *querypb.GetStatisticsRequest
func (_e *MockShardDelegator_Expecter) GetStatistics(ctx interface{}, req interface{}) *MockShardDelegator_GetStatistics_Call {
	return &MockShardDelegator_GetStatistics_Call{Call: _e.mock.On("GetStatistics", ctx, req)}
}

func (_c *MockShardDelegator_GetStatistics_Call) Run(run func(ctx context.Context, req *querypb.GetStatisticsRequest)) *MockShardDelegator_GetStatistics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.GetStatisticsRequest))
	})
	return _c
}

func (_c *MockShardDelegator_GetStatistics_Call) Return(_a0 []*internalpb.GetStatisticsResponse, _a1 error) *MockShardDelegator_GetStatistics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadGrowing provides a mock function with given fields: ctx, infos, version
func (_m *MockShardDelegator) LoadGrowing(ctx context.Context, infos []*querypb.SegmentLoadInfo, version int64) error {
	ret := _m.Called(ctx, infos, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*querypb.SegmentLoadInfo, int64) error); ok {
		r0 = rf(ctx, infos, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockShardDelegator_LoadGrowing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadGrowing'
type MockShardDelegator_LoadGrowing_Call struct {
	*mock.Call
}

// LoadGrowing is a helper method to define mock.On call
//  - ctx context.Context
//  - infos []*querypb.SegmentLoadInfo
//  - version int64
func (_e *MockShardDelegator_Expecter) LoadGrowing(ctx interface{}, infos interface{}, version interface{}) *MockShardDelegator_LoadGrowing_Call {
	return &MockShardDelegator_LoadGrowing_Call{Call: _e.mock.On("LoadGrowing", ctx, infos, version)}
}

func (_c *MockShardDelegator_LoadGrowing_Call) Run(run func(ctx context.Context, infos []*querypb.SegmentLoadInfo, version int64)) *MockShardDelegator_LoadGrowing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*querypb.SegmentLoadInfo), args[2].(int64))
	})
	return _c
}

func (_c *MockShardDelegator_LoadGrowing_Call) Return(_a0 error) *MockShardDelegator_LoadGrowing_Call {
	_c.Call.Return(_a0)
	return _c
}

// LoadSegments provides a mock function with given fields: ctx, req, force
func (_m *MockShardDelegator) LoadSegments(ctx context.Context, req *querypb.LoadSegmentsRequest, force bool) error {
	ret := _m.Called(ctx, req, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.LoadSegmentsRequest, bool) error); ok {
		r0 = rf(ctx, req, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockShardDelegator_LoadSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSegments'
type MockShardDelegator_LoadSegments_Call struct {
	*mock.Call
}

// LoadSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - req *querypb.LoadSegmentsRequest
//   - force bool
func (_e *MockShardDelegator_Expecter) LoadSegments(ctx interface{}, req interface{}, force interface{}) *MockShardDelegator_LoadSegments_Call {
	return &MockShardDelegator_LoadSegments_Call{Call: _e.mock.On("LoadSegments", ctx, req, force)}
}

func (_c *MockShardDelegator_LoadSegments_Call) Run(run func(ctx context.Context, req *querypb.LoadSegmentsRequest, force bool)) *MockShardDelegator_LoadSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.LoadSegmentsRequest), args[2].(bool))
	})
	return _c
}

func (_c *MockShardDelegator_LoadSegments_Call) Return(_a0 error) *MockShardDelegator_LoadSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

// ProcessDelete provides a mock function with given fields: deleteData, ts
func (_m *MockShardDelegator) ProcessDelete(deleteData []*DeleteData, ts uint64) {
	_m.Called(deleteData, ts)
}

// MockShardDelegator_ProcessDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessDelete'
type MockShardDelegator_ProcessDelete_Call struct {
	*mock.Call
}

// ProcessDelete is a helper method to define mock.On call
//  - deleteData []*DeleteData
//  - ts uint64
func (_e *MockShardDelegator_Expecter) ProcessDelete(deleteData interface{}, ts interface{}) *MockShardDelegator_ProcessDelete_Call {
	return &MockShardDelegator_ProcessDelete_Call{Call: _e.mock.On("ProcessDelete", deleteData, ts)}
}

func (_c *MockShardDelegator_ProcessDelete_Call) Run(run func(deleteData []*DeleteData, ts uint64)) *MockShardDelegator_ProcessDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*DeleteData), args[1].(uint64))
	})
	return _c
}

func (_c *MockShardDelegator_ProcessDelete_Call) Return() *MockShardDelegator_ProcessDelete_Call {
	_c.Call.Return()
	return _c
}

// ProcessInsert provides a mock function with given fields: insertRecords
func (_m *MockShardDelegator) ProcessInsert(insertRecords map[int64]*InsertData) {
	_m.Called(insertRecords)
}

// MockShardDelegator_ProcessInsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessInsert'
type MockShardDelegator_ProcessInsert_Call struct {
	*mock.Call
}

// ProcessInsert is a helper method to define mock.On call
//  - insertRecords map[int64]*InsertData
func (_e *MockShardDelegator_Expecter) ProcessInsert(insertRecords interface{}) *MockShardDelegator_ProcessInsert_Call {
	return &MockShardDelegator_ProcessInsert_Call{Call: _e.mock.On("ProcessInsert", insertRecords)}
}

func (_c *MockShardDelegator_ProcessInsert_Call) Run(run func(insertRecords map[int64]*InsertData)) *MockShardDelegator_ProcessInsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[int64]*InsertData))
	})
	return _c
}

func (_c *MockShardDelegator_ProcessInsert_Call) Return() *MockShardDelegator_ProcessInsert_Call {
	_c.Call.Return()
	return _c
}

// Query provides a mock function with given fields: ctx, req
func (_m *MockShardDelegator) Query(ctx context.Context, req *querypb.QueryRequest) ([]*internalpb.RetrieveResults, error) {
	ret := _m.Called(ctx, req)

	var r0 []*internalpb.RetrieveResults
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.QueryRequest) []*internalpb.RetrieveResults); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internalpb.RetrieveResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.QueryRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShardDelegator_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockShardDelegator_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//  - ctx context.Context
//  - req *querypb.QueryRequest
func (_e *MockShardDelegator_Expecter) Query(ctx interface{}, req interface{}) *MockShardDelegator_Query_Call {
	return &MockShardDelegator_Query_Call{Call: _e.mock.On("Query", ctx, req)}
}

func (_c *MockShardDelegator_Query_Call) Run(run func(ctx context.Context, req *querypb.QueryRequest)) *MockShardDelegator_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.QueryRequest))
	})
	return _c
}

func (_c *MockShardDelegator_Query_Call) Return(_a0 []*internalpb.RetrieveResults, _a1 error) *MockShardDelegator_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReleaseSegments provides a mock function with given fields: ctx, req, force
func (_m *MockShardDelegator) ReleaseSegments(ctx context.Context, req *querypb.ReleaseSegmentsRequest, force bool) error {
	ret := _m.Called(ctx, req, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.ReleaseSegmentsRequest, bool) error); ok {
		r0 = rf(ctx, req, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockShardDelegator_ReleaseSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseSegments'
type MockShardDelegator_ReleaseSegments_Call struct {
	*mock.Call
}

// ReleaseSegments is a helper method to define mock.On call
//  - ctx context.Context
//  - req *querypb.ReleaseSegmentsRequest
//  - force bool
func (_e *MockShardDelegator_Expecter) ReleaseSegments(ctx interface{}, req interface{}, force interface{}) *MockShardDelegator_ReleaseSegments_Call {
	return &MockShardDelegator_ReleaseSegments_Call{Call: _e.mock.On("ReleaseSegments", ctx, req, force)}
}

func (_c *MockShardDelegator_ReleaseSegments_Call) Run(run func(ctx context.Context, req *querypb.ReleaseSegmentsRequest, force bool)) *MockShardDelegator_ReleaseSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.ReleaseSegmentsRequest), args[2].(bool))
	})
	return _c
}

func (_c *MockShardDelegator_ReleaseSegments_Call) Return(_a0 error) *MockShardDelegator_ReleaseSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

// Search provides a mock function with given fields: ctx, req
func (_m *MockShardDelegator) Search(ctx context.Context, req *querypb.SearchRequest) ([]*internalpb.SearchResults, error) {
	ret := _m.Called(ctx, req)

	var r0 []*internalpb.SearchResults
	if rf, ok := ret.Get(0).(func(context.Context, *querypb.SearchRequest) []*internalpb.SearchResults); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*internalpb.SearchResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *querypb.SearchRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShardDelegator_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockShardDelegator_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - ctx context.Context
//  - req *querypb.SearchRequest
func (_e *MockShardDelegator_Expecter) Search(ctx interface{}, req interface{}) *MockShardDelegator_Search_Call {
	return &MockShardDelegator_Search_Call{Call: _e.mock.On("Search", ctx, req)}
}

func (_c *MockShardDelegator_Search_Call) Run(run func(ctx context.Context, req *querypb.SearchRequest)) *MockShardDelegator_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*querypb.SearchRequest))
	})
	return _c
}

func (_c *MockShardDelegator_Search_Call) Return(_a0 []*internalpb.SearchResults, _a1 error) *MockShardDelegator_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Serviceable provides a mock function with given fields:
func (_m *MockShardDelegator) Serviceable() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockShardDelegator_Serviceable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serviceable'
type MockShardDelegator_Serviceable_Call struct {
	*mock.Call
}

// Serviceable is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) Serviceable() *MockShardDelegator_Serviceable_Call {
	return &MockShardDelegator_Serviceable_Call{Call: _e.mock.On("Serviceable")}
}

func (_c *MockShardDelegator_Serviceable_Call) Run(run func()) *MockShardDelegator_Serviceable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_Serviceable_Call) Return(_a0 bool) *MockShardDelegator_Serviceable_Call {
	_c.Call.Return(_a0)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockShardDelegator) Start() {
	_m.Called()
}

// MockShardDelegator_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockShardDelegator_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) Start() *MockShardDelegator_Start_Call {
	return &MockShardDelegator_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockShardDelegator_Start_Call) Run(run func()) *MockShardDelegator_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_Start_Call) Return() *MockShardDelegator_Start_Call {
	_c.Call.Return()
	return _c
}

// SyncDistribution provides a mock function with given fields: ctx, entries
func (_m *MockShardDelegator) SyncDistribution(ctx context.Context, entries ...SegmentEntry) {
	_va := make([]interface{}, len(entries))
	for _i := range entries {
		_va[_i] = entries[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockShardDelegator_SyncDistribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncDistribution'
type MockShardDelegator_SyncDistribution_Call struct {
	*mock.Call
}

// SyncDistribution is a helper method to define mock.On call
//  - ctx context.Context
//  - entries ...SegmentEntry
func (_e *MockShardDelegator_Expecter) SyncDistribution(ctx interface{}, entries ...interface{}) *MockShardDelegator_SyncDistribution_Call {
	return &MockShardDelegator_SyncDistribution_Call{Call: _e.mock.On("SyncDistribution",
		append([]interface{}{ctx}, entries...)...)}
}

func (_c *MockShardDelegator_SyncDistribution_Call) Run(run func(ctx context.Context, entries ...SegmentEntry)) *MockShardDelegator_SyncDistribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentEntry, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentEntry)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockShardDelegator_SyncDistribution_Call) Return() *MockShardDelegator_SyncDistribution_Call {
	_c.Call.Return()
	return _c
}

// Version provides a mock function with given fields:
func (_m *MockShardDelegator) Version() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockShardDelegator_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type MockShardDelegator_Version_Call struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
func (_e *MockShardDelegator_Expecter) Version() *MockShardDelegator_Version_Call {
	return &MockShardDelegator_Version_Call{Call: _e.mock.On("Version")}
}

func (_c *MockShardDelegator_Version_Call) Run(run func()) *MockShardDelegator_Version_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockShardDelegator_Version_Call) Return(_a0 int64) *MockShardDelegator_Version_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockShardDelegator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockShardDelegator creates a new instance of MockShardDelegator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockShardDelegator(t mockConstructorTestingTNewMockShardDelegator) *MockShardDelegator {
	mock := &MockShardDelegator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
