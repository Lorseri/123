// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	internalpb "github.com/milvus-io/milvus/internal/proto/internalpb"

	metastore "github.com/milvus-io/milvus/internal/metastore"

	mock "github.com/stretchr/testify/mock"
)

// DataCoordCatalog is an autogenerated mock type for the DataCoordCatalog type
type DataCoordCatalog struct {
	mock.Mock
}

type DataCoordCatalog_Expecter struct {
	mock *mock.Mock
}

func (_m *DataCoordCatalog) EXPECT() *DataCoordCatalog_Expecter {
	return &DataCoordCatalog_Expecter{mock: &_m.Mock}
}

// AddSegment provides a mock function with given fields: ctx, segment
func (_m *DataCoordCatalog) AddSegment(ctx context.Context, segment *datapb.SegmentInfo) error {
	ret := _m.Called(ctx, segment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, segment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_AddSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSegment'
type DataCoordCatalog_AddSegment_Call struct {
	*mock.Call
}

// AddSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - segment *datapb.SegmentInfo
func (_e *DataCoordCatalog_Expecter) AddSegment(ctx interface{}, segment interface{}) *DataCoordCatalog_AddSegment_Call {
	return &DataCoordCatalog_AddSegment_Call{Call: _e.mock.On("AddSegment", ctx, segment)}
}

func (_c *DataCoordCatalog_AddSegment_Call) Run(run func(ctx context.Context, segment *datapb.SegmentInfo)) *DataCoordCatalog_AddSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.SegmentInfo))
	})
	return _c
}

func (_c *DataCoordCatalog_AddSegment_Call) Return(_a0 error) *DataCoordCatalog_AddSegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_AddSegment_Call) RunAndReturn(run func(context.Context, *datapb.SegmentInfo) error) *DataCoordCatalog_AddSegment_Call {
	_c.Call.Return(run)
	return _c
}

// AlterSegment provides a mock function with given fields: ctx, newSegment, oldSegment
func (_m *DataCoordCatalog) AlterSegment(ctx context.Context, newSegment *datapb.SegmentInfo, oldSegment *datapb.SegmentInfo) error {
	ret := _m.Called(ctx, newSegment, oldSegment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.SegmentInfo, *datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, newSegment, oldSegment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_AlterSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AlterSegment'
type DataCoordCatalog_AlterSegment_Call struct {
	*mock.Call
}

// AlterSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - newSegment *datapb.SegmentInfo
//   - oldSegment *datapb.SegmentInfo
func (_e *DataCoordCatalog_Expecter) AlterSegment(ctx interface{}, newSegment interface{}, oldSegment interface{}) *DataCoordCatalog_AlterSegment_Call {
	return &DataCoordCatalog_AlterSegment_Call{Call: _e.mock.On("AlterSegment", ctx, newSegment, oldSegment)}
}

func (_c *DataCoordCatalog_AlterSegment_Call) Run(run func(ctx context.Context, newSegment *datapb.SegmentInfo, oldSegment *datapb.SegmentInfo)) *DataCoordCatalog_AlterSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.SegmentInfo), args[2].(*datapb.SegmentInfo))
	})
	return _c
}

func (_c *DataCoordCatalog_AlterSegment_Call) Return(_a0 error) *DataCoordCatalog_AlterSegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_AlterSegment_Call) RunAndReturn(run func(context.Context, *datapb.SegmentInfo, *datapb.SegmentInfo) error) *DataCoordCatalog_AlterSegment_Call {
	_c.Call.Return(run)
	return _c
}

// AlterSegments provides a mock function with given fields: ctx, newSegments, binlogs
func (_m *DataCoordCatalog) AlterSegments(ctx context.Context, newSegments []*datapb.SegmentInfo, binlogs ...metastore.BinlogsIncrement) error {
	_va := make([]interface{}, len(binlogs))
	for _i := range binlogs {
		_va[_i] = binlogs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, newSegments)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*datapb.SegmentInfo, ...metastore.BinlogsIncrement) error); ok {
		r0 = rf(ctx, newSegments, binlogs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_AlterSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AlterSegments'
type DataCoordCatalog_AlterSegments_Call struct {
	*mock.Call
}

// AlterSegments is a helper method to define mock.On call
//   - ctx context.Context
//   - newSegments []*datapb.SegmentInfo
//   - binlogs ...metastore.BinlogsIncrement
func (_e *DataCoordCatalog_Expecter) AlterSegments(ctx interface{}, newSegments interface{}, binlogs ...interface{}) *DataCoordCatalog_AlterSegments_Call {
	return &DataCoordCatalog_AlterSegments_Call{Call: _e.mock.On("AlterSegments",
		append([]interface{}{ctx, newSegments}, binlogs...)...)}
}

func (_c *DataCoordCatalog_AlterSegments_Call) Run(run func(ctx context.Context, newSegments []*datapb.SegmentInfo, binlogs ...metastore.BinlogsIncrement)) *DataCoordCatalog_AlterSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]metastore.BinlogsIncrement, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(metastore.BinlogsIncrement)
			}
		}
		run(args[0].(context.Context), args[1].([]*datapb.SegmentInfo), variadicArgs...)
	})
	return _c
}

func (_c *DataCoordCatalog_AlterSegments_Call) Return(_a0 error) *DataCoordCatalog_AlterSegments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_AlterSegments_Call) RunAndReturn(run func(context.Context, []*datapb.SegmentInfo, ...metastore.BinlogsIncrement) error) *DataCoordCatalog_AlterSegments_Call {
	_c.Call.Return(run)
	return _c
}

// AlterSegmentsAndAddNewSegment provides a mock function with given fields: ctx, segments, newSegment
func (_m *DataCoordCatalog) AlterSegmentsAndAddNewSegment(ctx context.Context, segments []*datapb.SegmentInfo, newSegment *datapb.SegmentInfo) error {
	ret := _m.Called(ctx, segments, newSegment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*datapb.SegmentInfo, *datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, segments, newSegment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AlterSegmentsAndAddNewSegment'
type DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call struct {
	*mock.Call
}

// AlterSegmentsAndAddNewSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - segments []*datapb.SegmentInfo
//   - newSegment *datapb.SegmentInfo
func (_e *DataCoordCatalog_Expecter) AlterSegmentsAndAddNewSegment(ctx interface{}, segments interface{}, newSegment interface{}) *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call {
	return &DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call{Call: _e.mock.On("AlterSegmentsAndAddNewSegment", ctx, segments, newSegment)}
}

func (_c *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call) Run(run func(ctx context.Context, segments []*datapb.SegmentInfo, newSegment *datapb.SegmentInfo)) *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*datapb.SegmentInfo), args[2].(*datapb.SegmentInfo))
	})
	return _c
}

func (_c *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call) Return(_a0 error) *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call) RunAndReturn(run func(context.Context, []*datapb.SegmentInfo, *datapb.SegmentInfo) error) *DataCoordCatalog_AlterSegmentsAndAddNewSegment_Call {
	_c.Call.Return(run)
	return _c
}

// ChannelExists provides a mock function with given fields: ctx, channel
func (_m *DataCoordCatalog) ChannelExists(ctx context.Context, channel string) bool {
	ret := _m.Called(ctx, channel)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DataCoordCatalog_ChannelExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelExists'
type DataCoordCatalog_ChannelExists_Call struct {
	*mock.Call
}

// ChannelExists is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *DataCoordCatalog_Expecter) ChannelExists(ctx interface{}, channel interface{}) *DataCoordCatalog_ChannelExists_Call {
	return &DataCoordCatalog_ChannelExists_Call{Call: _e.mock.On("ChannelExists", ctx, channel)}
}

func (_c *DataCoordCatalog_ChannelExists_Call) Run(run func(ctx context.Context, channel string)) *DataCoordCatalog_ChannelExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_ChannelExists_Call) Return(_a0 bool) *DataCoordCatalog_ChannelExists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_ChannelExists_Call) RunAndReturn(run func(context.Context, string) bool) *DataCoordCatalog_ChannelExists_Call {
	_c.Call.Return(run)
	return _c
}

// DropChannel provides a mock function with given fields: ctx, channel
func (_m *DataCoordCatalog) DropChannel(ctx context.Context, channel string) error {
	ret := _m.Called(ctx, channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_DropChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropChannel'
type DataCoordCatalog_DropChannel_Call struct {
	*mock.Call
}

// DropChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *DataCoordCatalog_Expecter) DropChannel(ctx interface{}, channel interface{}) *DataCoordCatalog_DropChannel_Call {
	return &DataCoordCatalog_DropChannel_Call{Call: _e.mock.On("DropChannel", ctx, channel)}
}

func (_c *DataCoordCatalog_DropChannel_Call) Run(run func(ctx context.Context, channel string)) *DataCoordCatalog_DropChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_DropChannel_Call) Return(_a0 error) *DataCoordCatalog_DropChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_DropChannel_Call) RunAndReturn(run func(context.Context, string) error) *DataCoordCatalog_DropChannel_Call {
	_c.Call.Return(run)
	return _c
}

// DropChannelCheckpoint provides a mock function with given fields: ctx, vChannel
func (_m *DataCoordCatalog) DropChannelCheckpoint(ctx context.Context, vChannel string) error {
	ret := _m.Called(ctx, vChannel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, vChannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_DropChannelCheckpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropChannelCheckpoint'
type DataCoordCatalog_DropChannelCheckpoint_Call struct {
	*mock.Call
}

// DropChannelCheckpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - vChannel string
func (_e *DataCoordCatalog_Expecter) DropChannelCheckpoint(ctx interface{}, vChannel interface{}) *DataCoordCatalog_DropChannelCheckpoint_Call {
	return &DataCoordCatalog_DropChannelCheckpoint_Call{Call: _e.mock.On("DropChannelCheckpoint", ctx, vChannel)}
}

func (_c *DataCoordCatalog_DropChannelCheckpoint_Call) Run(run func(ctx context.Context, vChannel string)) *DataCoordCatalog_DropChannelCheckpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_DropChannelCheckpoint_Call) Return(_a0 error) *DataCoordCatalog_DropChannelCheckpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_DropChannelCheckpoint_Call) RunAndReturn(run func(context.Context, string) error) *DataCoordCatalog_DropChannelCheckpoint_Call {
	_c.Call.Return(run)
	return _c
}

// DropSegment provides a mock function with given fields: ctx, segment
func (_m *DataCoordCatalog) DropSegment(ctx context.Context, segment *datapb.SegmentInfo) error {
	ret := _m.Called(ctx, segment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, segment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_DropSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropSegment'
type DataCoordCatalog_DropSegment_Call struct {
	*mock.Call
}

// DropSegment is a helper method to define mock.On call
//   - ctx context.Context
//   - segment *datapb.SegmentInfo
func (_e *DataCoordCatalog_Expecter) DropSegment(ctx interface{}, segment interface{}) *DataCoordCatalog_DropSegment_Call {
	return &DataCoordCatalog_DropSegment_Call{Call: _e.mock.On("DropSegment", ctx, segment)}
}

func (_c *DataCoordCatalog_DropSegment_Call) Run(run func(ctx context.Context, segment *datapb.SegmentInfo)) *DataCoordCatalog_DropSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*datapb.SegmentInfo))
	})
	return _c
}

func (_c *DataCoordCatalog_DropSegment_Call) Return(_a0 error) *DataCoordCatalog_DropSegment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_DropSegment_Call) RunAndReturn(run func(context.Context, *datapb.SegmentInfo) error) *DataCoordCatalog_DropSegment_Call {
	_c.Call.Return(run)
	return _c
}

// GcConfirm provides a mock function with given fields: ctx, collectionID, partitionID
func (_m *DataCoordCatalog) GcConfirm(ctx context.Context, collectionID int64, partitionID int64) bool {
	ret := _m.Called(ctx, collectionID, partitionID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) bool); ok {
		r0 = rf(ctx, collectionID, partitionID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DataCoordCatalog_GcConfirm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GcConfirm'
type DataCoordCatalog_GcConfirm_Call struct {
	*mock.Call
}

// GcConfirm is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
//   - partitionID int64
func (_e *DataCoordCatalog_Expecter) GcConfirm(ctx interface{}, collectionID interface{}, partitionID interface{}) *DataCoordCatalog_GcConfirm_Call {
	return &DataCoordCatalog_GcConfirm_Call{Call: _e.mock.On("GcConfirm", ctx, collectionID, partitionID)}
}

func (_c *DataCoordCatalog_GcConfirm_Call) Run(run func(ctx context.Context, collectionID int64, partitionID int64)) *DataCoordCatalog_GcConfirm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *DataCoordCatalog_GcConfirm_Call) Return(_a0 bool) *DataCoordCatalog_GcConfirm_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_GcConfirm_Call) RunAndReturn(run func(context.Context, int64, int64) bool) *DataCoordCatalog_GcConfirm_Call {
	_c.Call.Return(run)
	return _c
}

// ListChannelCheckpoint provides a mock function with given fields: ctx
func (_m *DataCoordCatalog) ListChannelCheckpoint(ctx context.Context) (map[string]*internalpb.MsgPosition, error) {
	ret := _m.Called(ctx)

	var r0 map[string]*internalpb.MsgPosition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]*internalpb.MsgPosition, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]*internalpb.MsgPosition); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*internalpb.MsgPosition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataCoordCatalog_ListChannelCheckpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListChannelCheckpoint'
type DataCoordCatalog_ListChannelCheckpoint_Call struct {
	*mock.Call
}

// ListChannelCheckpoint is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DataCoordCatalog_Expecter) ListChannelCheckpoint(ctx interface{}) *DataCoordCatalog_ListChannelCheckpoint_Call {
	return &DataCoordCatalog_ListChannelCheckpoint_Call{Call: _e.mock.On("ListChannelCheckpoint", ctx)}
}

func (_c *DataCoordCatalog_ListChannelCheckpoint_Call) Run(run func(ctx context.Context)) *DataCoordCatalog_ListChannelCheckpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DataCoordCatalog_ListChannelCheckpoint_Call) Return(_a0 map[string]*internalpb.MsgPosition, _a1 error) *DataCoordCatalog_ListChannelCheckpoint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataCoordCatalog_ListChannelCheckpoint_Call) RunAndReturn(run func(context.Context) (map[string]*internalpb.MsgPosition, error)) *DataCoordCatalog_ListChannelCheckpoint_Call {
	_c.Call.Return(run)
	return _c
}

// ListSegments provides a mock function with given fields: ctx
func (_m *DataCoordCatalog) ListSegments(ctx context.Context) ([]*datapb.SegmentInfo, error) {
	ret := _m.Called(ctx)

	var r0 []*datapb.SegmentInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*datapb.SegmentInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*datapb.SegmentInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.SegmentInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataCoordCatalog_ListSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSegments'
type DataCoordCatalog_ListSegments_Call struct {
	*mock.Call
}

// ListSegments is a helper method to define mock.On call
//   - ctx context.Context
func (_e *DataCoordCatalog_Expecter) ListSegments(ctx interface{}) *DataCoordCatalog_ListSegments_Call {
	return &DataCoordCatalog_ListSegments_Call{Call: _e.mock.On("ListSegments", ctx)}
}

func (_c *DataCoordCatalog_ListSegments_Call) Run(run func(ctx context.Context)) *DataCoordCatalog_ListSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DataCoordCatalog_ListSegments_Call) Return(_a0 []*datapb.SegmentInfo, _a1 error) *DataCoordCatalog_ListSegments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataCoordCatalog_ListSegments_Call) RunAndReturn(run func(context.Context) ([]*datapb.SegmentInfo, error)) *DataCoordCatalog_ListSegments_Call {
	_c.Call.Return(run)
	return _c
}

// MarkChannelAdded provides a mock function with given fields: ctx, channel
func (_m *DataCoordCatalog) MarkChannelAdded(ctx context.Context, channel string) error {
	ret := _m.Called(ctx, channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_MarkChannelAdded_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkChannelAdded'
type DataCoordCatalog_MarkChannelAdded_Call struct {
	*mock.Call
}

// MarkChannelAdded is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *DataCoordCatalog_Expecter) MarkChannelAdded(ctx interface{}, channel interface{}) *DataCoordCatalog_MarkChannelAdded_Call {
	return &DataCoordCatalog_MarkChannelAdded_Call{Call: _e.mock.On("MarkChannelAdded", ctx, channel)}
}

func (_c *DataCoordCatalog_MarkChannelAdded_Call) Run(run func(ctx context.Context, channel string)) *DataCoordCatalog_MarkChannelAdded_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_MarkChannelAdded_Call) Return(_a0 error) *DataCoordCatalog_MarkChannelAdded_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_MarkChannelAdded_Call) RunAndReturn(run func(context.Context, string) error) *DataCoordCatalog_MarkChannelAdded_Call {
	_c.Call.Return(run)
	return _c
}

// MarkChannelDeleted provides a mock function with given fields: ctx, channel
func (_m *DataCoordCatalog) MarkChannelDeleted(ctx context.Context, channel string) error {
	ret := _m.Called(ctx, channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_MarkChannelDeleted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkChannelDeleted'
type DataCoordCatalog_MarkChannelDeleted_Call struct {
	*mock.Call
}

// MarkChannelDeleted is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *DataCoordCatalog_Expecter) MarkChannelDeleted(ctx interface{}, channel interface{}) *DataCoordCatalog_MarkChannelDeleted_Call {
	return &DataCoordCatalog_MarkChannelDeleted_Call{Call: _e.mock.On("MarkChannelDeleted", ctx, channel)}
}

func (_c *DataCoordCatalog_MarkChannelDeleted_Call) Run(run func(ctx context.Context, channel string)) *DataCoordCatalog_MarkChannelDeleted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_MarkChannelDeleted_Call) Return(_a0 error) *DataCoordCatalog_MarkChannelDeleted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_MarkChannelDeleted_Call) RunAndReturn(run func(context.Context, string) error) *DataCoordCatalog_MarkChannelDeleted_Call {
	_c.Call.Return(run)
	return _c
}

// SaveChannelCheckpoint provides a mock function with given fields: ctx, vChannel, pos
func (_m *DataCoordCatalog) SaveChannelCheckpoint(ctx context.Context, vChannel string, pos *internalpb.MsgPosition) error {
	ret := _m.Called(ctx, vChannel, pos)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *internalpb.MsgPosition) error); ok {
		r0 = rf(ctx, vChannel, pos)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_SaveChannelCheckpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveChannelCheckpoint'
type DataCoordCatalog_SaveChannelCheckpoint_Call struct {
	*mock.Call
}

// SaveChannelCheckpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - vChannel string
//   - pos *internalpb.MsgPosition
func (_e *DataCoordCatalog_Expecter) SaveChannelCheckpoint(ctx interface{}, vChannel interface{}, pos interface{}) *DataCoordCatalog_SaveChannelCheckpoint_Call {
	return &DataCoordCatalog_SaveChannelCheckpoint_Call{Call: _e.mock.On("SaveChannelCheckpoint", ctx, vChannel, pos)}
}

func (_c *DataCoordCatalog_SaveChannelCheckpoint_Call) Run(run func(ctx context.Context, vChannel string, pos *internalpb.MsgPosition)) *DataCoordCatalog_SaveChannelCheckpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*internalpb.MsgPosition))
	})
	return _c
}

func (_c *DataCoordCatalog_SaveChannelCheckpoint_Call) Return(_a0 error) *DataCoordCatalog_SaveChannelCheckpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_SaveChannelCheckpoint_Call) RunAndReturn(run func(context.Context, string, *internalpb.MsgPosition) error) *DataCoordCatalog_SaveChannelCheckpoint_Call {
	_c.Call.Return(run)
	return _c
}

// SaveDroppedSegmentsInBatch provides a mock function with given fields: ctx, segments
func (_m *DataCoordCatalog) SaveDroppedSegmentsInBatch(ctx context.Context, segments []*datapb.SegmentInfo) error {
	ret := _m.Called(ctx, segments)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*datapb.SegmentInfo) error); ok {
		r0 = rf(ctx, segments)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataCoordCatalog_SaveDroppedSegmentsInBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveDroppedSegmentsInBatch'
type DataCoordCatalog_SaveDroppedSegmentsInBatch_Call struct {
	*mock.Call
}

// SaveDroppedSegmentsInBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - segments []*datapb.SegmentInfo
func (_e *DataCoordCatalog_Expecter) SaveDroppedSegmentsInBatch(ctx interface{}, segments interface{}) *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call {
	return &DataCoordCatalog_SaveDroppedSegmentsInBatch_Call{Call: _e.mock.On("SaveDroppedSegmentsInBatch", ctx, segments)}
}

func (_c *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call) Run(run func(ctx context.Context, segments []*datapb.SegmentInfo)) *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*datapb.SegmentInfo))
	})
	return _c
}

func (_c *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call) Return(_a0 error) *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call) RunAndReturn(run func(context.Context, []*datapb.SegmentInfo) error) *DataCoordCatalog_SaveDroppedSegmentsInBatch_Call {
	_c.Call.Return(run)
	return _c
}

// ShouldDropChannel provides a mock function with given fields: ctx, channel
func (_m *DataCoordCatalog) ShouldDropChannel(ctx context.Context, channel string) bool {
	ret := _m.Called(ctx, channel)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DataCoordCatalog_ShouldDropChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShouldDropChannel'
type DataCoordCatalog_ShouldDropChannel_Call struct {
	*mock.Call
}

// ShouldDropChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - channel string
func (_e *DataCoordCatalog_Expecter) ShouldDropChannel(ctx interface{}, channel interface{}) *DataCoordCatalog_ShouldDropChannel_Call {
	return &DataCoordCatalog_ShouldDropChannel_Call{Call: _e.mock.On("ShouldDropChannel", ctx, channel)}
}

func (_c *DataCoordCatalog_ShouldDropChannel_Call) Run(run func(ctx context.Context, channel string)) *DataCoordCatalog_ShouldDropChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DataCoordCatalog_ShouldDropChannel_Call) Return(_a0 bool) *DataCoordCatalog_ShouldDropChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DataCoordCatalog_ShouldDropChannel_Call) RunAndReturn(run func(context.Context, string) bool) *DataCoordCatalog_ShouldDropChannel_Call {
	_c.Call.Return(run)
	return _c
}

// NewDataCoordCatalog creates a new instance of DataCoordCatalog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataCoordCatalog(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataCoordCatalog {
	mock := &DataCoordCatalog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
