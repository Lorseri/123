// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	commonpb "github.com/milvus-io/milvus-proto/go-api/v2/commonpb"

	grpc "google.golang.org/grpc"

	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"

	mock "github.com/stretchr/testify/mock"

	proxypb "github.com/milvus-io/milvus/internal/proto/proxypb"
)

// MockProxyClient is an autogenerated mock type for the ProxyClient type
type MockProxyClient struct {
	mock.Mock
}

type MockProxyClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProxyClient) EXPECT() *MockProxyClient_Expecter {
	return &MockProxyClient_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockProxyClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProxyClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockProxyClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockProxyClient_Expecter) Close() *MockProxyClient_Close_Call {
	return &MockProxyClient_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockProxyClient_Close_Call) Run(run func()) *MockProxyClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProxyClient_Close_Call) Return(_a0 error) *MockProxyClient_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProxyClient_Close_Call) RunAndReturn(run func() error) *MockProxyClient_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetComponentStates provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) GetComponentStates(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption) (*milvuspb.ComponentStates, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *milvuspb.ComponentStates
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetComponentStatesRequest, ...grpc.CallOption) (*milvuspb.ComponentStates, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetComponentStatesRequest, ...grpc.CallOption) *milvuspb.ComponentStates); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.ComponentStates)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetComponentStatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_GetComponentStates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComponentStates'
type MockProxyClient_GetComponentStates_Call struct {
	*mock.Call
}

// GetComponentStates is a helper method to define mock.On call
//   - ctx context.Context
//   - in *milvuspb.GetComponentStatesRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) GetComponentStates(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_GetComponentStates_Call {
	return &MockProxyClient_GetComponentStates_Call{Call: _e.mock.On("GetComponentStates",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_GetComponentStates_Call) Run(run func(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption)) *MockProxyClient_GetComponentStates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*milvuspb.GetComponentStatesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_GetComponentStates_Call) Return(_a0 *milvuspb.ComponentStates, _a1 error) *MockProxyClient_GetComponentStates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_GetComponentStates_Call) RunAndReturn(run func(context.Context, *milvuspb.GetComponentStatesRequest, ...grpc.CallOption) (*milvuspb.ComponentStates, error)) *MockProxyClient_GetComponentStates_Call {
	_c.Call.Return(run)
	return _c
}

// GetDdChannel provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) GetDdChannel(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *milvuspb.StringResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetDdChannelRequest, ...grpc.CallOption) (*milvuspb.StringResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetDdChannelRequest, ...grpc.CallOption) *milvuspb.StringResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.GetDdChannelRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_GetDdChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDdChannel'
type MockProxyClient_GetDdChannel_Call struct {
	*mock.Call
}

// GetDdChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - in *internalpb.GetDdChannelRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) GetDdChannel(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_GetDdChannel_Call {
	return &MockProxyClient_GetDdChannel_Call{Call: _e.mock.On("GetDdChannel",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_GetDdChannel_Call) Run(run func(ctx context.Context, in *internalpb.GetDdChannelRequest, opts ...grpc.CallOption)) *MockProxyClient_GetDdChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*internalpb.GetDdChannelRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_GetDdChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockProxyClient_GetDdChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_GetDdChannel_Call) RunAndReturn(run func(context.Context, *internalpb.GetDdChannelRequest, ...grpc.CallOption) (*milvuspb.StringResponse, error)) *MockProxyClient_GetDdChannel_Call {
	_c.Call.Return(run)
	return _c
}

// GetProxyMetrics provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) GetProxyMetrics(ctx context.Context, in *milvuspb.GetMetricsRequest, opts ...grpc.CallOption) (*milvuspb.GetMetricsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *milvuspb.GetMetricsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest, ...grpc.CallOption) (*milvuspb.GetMetricsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *milvuspb.GetMetricsRequest, ...grpc.CallOption) *milvuspb.GetMetricsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.GetMetricsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *milvuspb.GetMetricsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_GetProxyMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProxyMetrics'
type MockProxyClient_GetProxyMetrics_Call struct {
	*mock.Call
}

// GetProxyMetrics is a helper method to define mock.On call
//   - ctx context.Context
//   - in *milvuspb.GetMetricsRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) GetProxyMetrics(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_GetProxyMetrics_Call {
	return &MockProxyClient_GetProxyMetrics_Call{Call: _e.mock.On("GetProxyMetrics",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_GetProxyMetrics_Call) Run(run func(ctx context.Context, in *milvuspb.GetMetricsRequest, opts ...grpc.CallOption)) *MockProxyClient_GetProxyMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*milvuspb.GetMetricsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_GetProxyMetrics_Call) Return(_a0 *milvuspb.GetMetricsResponse, _a1 error) *MockProxyClient_GetProxyMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_GetProxyMetrics_Call) RunAndReturn(run func(context.Context, *milvuspb.GetMetricsRequest, ...grpc.CallOption) (*milvuspb.GetMetricsResponse, error)) *MockProxyClient_GetProxyMetrics_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatisticsChannel provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) GetStatisticsChannel(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption) (*milvuspb.StringResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *milvuspb.StringResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetStatisticsChannelRequest, ...grpc.CallOption) (*milvuspb.StringResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internalpb.GetStatisticsChannelRequest, ...grpc.CallOption) *milvuspb.StringResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*milvuspb.StringResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internalpb.GetStatisticsChannelRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_GetStatisticsChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatisticsChannel'
type MockProxyClient_GetStatisticsChannel_Call struct {
	*mock.Call
}

// GetStatisticsChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - in *internalpb.GetStatisticsChannelRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) GetStatisticsChannel(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_GetStatisticsChannel_Call {
	return &MockProxyClient_GetStatisticsChannel_Call{Call: _e.mock.On("GetStatisticsChannel",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_GetStatisticsChannel_Call) Run(run func(ctx context.Context, in *internalpb.GetStatisticsChannelRequest, opts ...grpc.CallOption)) *MockProxyClient_GetStatisticsChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*internalpb.GetStatisticsChannelRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_GetStatisticsChannel_Call) Return(_a0 *milvuspb.StringResponse, _a1 error) *MockProxyClient_GetStatisticsChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_GetStatisticsChannel_Call) RunAndReturn(run func(context.Context, *internalpb.GetStatisticsChannelRequest, ...grpc.CallOption) (*milvuspb.StringResponse, error)) *MockProxyClient_GetStatisticsChannel_Call {
	_c.Call.Return(run)
	return _c
}

// InvalidateCollectionMetaCache provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) InvalidateCollectionMetaCache(ctx context.Context, in *proxypb.InvalidateCollMetaCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.InvalidateCollMetaCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.InvalidateCollMetaCacheRequest, ...grpc.CallOption) *commonpb.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.InvalidateCollMetaCacheRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_InvalidateCollectionMetaCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InvalidateCollectionMetaCache'
type MockProxyClient_InvalidateCollectionMetaCache_Call struct {
	*mock.Call
}

// InvalidateCollectionMetaCache is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.InvalidateCollMetaCacheRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) InvalidateCollectionMetaCache(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_InvalidateCollectionMetaCache_Call {
	return &MockProxyClient_InvalidateCollectionMetaCache_Call{Call: _e.mock.On("InvalidateCollectionMetaCache",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_InvalidateCollectionMetaCache_Call) Run(run func(ctx context.Context, in *proxypb.InvalidateCollMetaCacheRequest, opts ...grpc.CallOption)) *MockProxyClient_InvalidateCollectionMetaCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.InvalidateCollMetaCacheRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_InvalidateCollectionMetaCache_Call) Return(_a0 *commonpb.Status, _a1 error) *MockProxyClient_InvalidateCollectionMetaCache_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_InvalidateCollectionMetaCache_Call) RunAndReturn(run func(context.Context, *proxypb.InvalidateCollMetaCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)) *MockProxyClient_InvalidateCollectionMetaCache_Call {
	_c.Call.Return(run)
	return _c
}

// InvalidateCredentialCache provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) InvalidateCredentialCache(ctx context.Context, in *proxypb.InvalidateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.InvalidateCredCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.InvalidateCredCacheRequest, ...grpc.CallOption) *commonpb.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.InvalidateCredCacheRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_InvalidateCredentialCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InvalidateCredentialCache'
type MockProxyClient_InvalidateCredentialCache_Call struct {
	*mock.Call
}

// InvalidateCredentialCache is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.InvalidateCredCacheRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) InvalidateCredentialCache(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_InvalidateCredentialCache_Call {
	return &MockProxyClient_InvalidateCredentialCache_Call{Call: _e.mock.On("InvalidateCredentialCache",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_InvalidateCredentialCache_Call) Run(run func(ctx context.Context, in *proxypb.InvalidateCredCacheRequest, opts ...grpc.CallOption)) *MockProxyClient_InvalidateCredentialCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.InvalidateCredCacheRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_InvalidateCredentialCache_Call) Return(_a0 *commonpb.Status, _a1 error) *MockProxyClient_InvalidateCredentialCache_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_InvalidateCredentialCache_Call) RunAndReturn(run func(context.Context, *proxypb.InvalidateCredCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)) *MockProxyClient_InvalidateCredentialCache_Call {
	_c.Call.Return(run)
	return _c
}

// ListClientInfos provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) ListClientInfos(ctx context.Context, in *proxypb.ListClientInfosRequest, opts ...grpc.CallOption) (*proxypb.ListClientInfosResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proxypb.ListClientInfosResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.ListClientInfosRequest, ...grpc.CallOption) (*proxypb.ListClientInfosResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.ListClientInfosRequest, ...grpc.CallOption) *proxypb.ListClientInfosResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proxypb.ListClientInfosResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.ListClientInfosRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_ListClientInfos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListClientInfos'
type MockProxyClient_ListClientInfos_Call struct {
	*mock.Call
}

// ListClientInfos is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.ListClientInfosRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) ListClientInfos(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_ListClientInfos_Call {
	return &MockProxyClient_ListClientInfos_Call{Call: _e.mock.On("ListClientInfos",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_ListClientInfos_Call) Run(run func(ctx context.Context, in *proxypb.ListClientInfosRequest, opts ...grpc.CallOption)) *MockProxyClient_ListClientInfos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.ListClientInfosRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_ListClientInfos_Call) Return(_a0 *proxypb.ListClientInfosResponse, _a1 error) *MockProxyClient_ListClientInfos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_ListClientInfos_Call) RunAndReturn(run func(context.Context, *proxypb.ListClientInfosRequest, ...grpc.CallOption) (*proxypb.ListClientInfosResponse, error)) *MockProxyClient_ListClientInfos_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshPolicyInfoCache provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) RefreshPolicyInfoCache(ctx context.Context, in *proxypb.RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.RefreshPolicyInfoCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.RefreshPolicyInfoCacheRequest, ...grpc.CallOption) *commonpb.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.RefreshPolicyInfoCacheRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_RefreshPolicyInfoCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshPolicyInfoCache'
type MockProxyClient_RefreshPolicyInfoCache_Call struct {
	*mock.Call
}

// RefreshPolicyInfoCache is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.RefreshPolicyInfoCacheRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) RefreshPolicyInfoCache(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_RefreshPolicyInfoCache_Call {
	return &MockProxyClient_RefreshPolicyInfoCache_Call{Call: _e.mock.On("RefreshPolicyInfoCache",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_RefreshPolicyInfoCache_Call) Run(run func(ctx context.Context, in *proxypb.RefreshPolicyInfoCacheRequest, opts ...grpc.CallOption)) *MockProxyClient_RefreshPolicyInfoCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.RefreshPolicyInfoCacheRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_RefreshPolicyInfoCache_Call) Return(_a0 *commonpb.Status, _a1 error) *MockProxyClient_RefreshPolicyInfoCache_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_RefreshPolicyInfoCache_Call) RunAndReturn(run func(context.Context, *proxypb.RefreshPolicyInfoCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)) *MockProxyClient_RefreshPolicyInfoCache_Call {
	_c.Call.Return(run)
	return _c
}

// SetRates provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) SetRates(ctx context.Context, in *proxypb.SetRatesRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.SetRatesRequest, ...grpc.CallOption) (*commonpb.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.SetRatesRequest, ...grpc.CallOption) *commonpb.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.SetRatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_SetRates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRates'
type MockProxyClient_SetRates_Call struct {
	*mock.Call
}

// SetRates is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.SetRatesRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) SetRates(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_SetRates_Call {
	return &MockProxyClient_SetRates_Call{Call: _e.mock.On("SetRates",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_SetRates_Call) Run(run func(ctx context.Context, in *proxypb.SetRatesRequest, opts ...grpc.CallOption)) *MockProxyClient_SetRates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.SetRatesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_SetRates_Call) Return(_a0 *commonpb.Status, _a1 error) *MockProxyClient_SetRates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_SetRates_Call) RunAndReturn(run func(context.Context, *proxypb.SetRatesRequest, ...grpc.CallOption) (*commonpb.Status, error)) *MockProxyClient_SetRates_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCredentialCache provides a mock function with given fields: ctx, in, opts
func (_m *MockProxyClient) UpdateCredentialCache(ctx context.Context, in *proxypb.UpdateCredCacheRequest, opts ...grpc.CallOption) (*commonpb.Status, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *commonpb.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.UpdateCredCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proxypb.UpdateCredCacheRequest, ...grpc.CallOption) *commonpb.Status); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*commonpb.Status)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proxypb.UpdateCredCacheRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProxyClient_UpdateCredentialCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCredentialCache'
type MockProxyClient_UpdateCredentialCache_Call struct {
	*mock.Call
}

// UpdateCredentialCache is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proxypb.UpdateCredCacheRequest
//   - opts ...grpc.CallOption
func (_e *MockProxyClient_Expecter) UpdateCredentialCache(ctx interface{}, in interface{}, opts ...interface{}) *MockProxyClient_UpdateCredentialCache_Call {
	return &MockProxyClient_UpdateCredentialCache_Call{Call: _e.mock.On("UpdateCredentialCache",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockProxyClient_UpdateCredentialCache_Call) Run(run func(ctx context.Context, in *proxypb.UpdateCredCacheRequest, opts ...grpc.CallOption)) *MockProxyClient_UpdateCredentialCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proxypb.UpdateCredCacheRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockProxyClient_UpdateCredentialCache_Call) Return(_a0 *commonpb.Status, _a1 error) *MockProxyClient_UpdateCredentialCache_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProxyClient_UpdateCredentialCache_Call) RunAndReturn(run func(context.Context, *proxypb.UpdateCredCacheRequest, ...grpc.CallOption) (*commonpb.Status, error)) *MockProxyClient_UpdateCredentialCache_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProxyClient creates a new instance of MockProxyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProxyClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProxyClient {
	mock := &MockProxyClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
