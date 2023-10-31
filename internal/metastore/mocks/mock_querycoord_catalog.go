// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
	mock "github.com/stretchr/testify/mock"
)

// QueryCoordCatalog is an autogenerated mock type for the QueryCoordCatalog type
type QueryCoordCatalog struct {
	mock.Mock
}

type QueryCoordCatalog_Expecter struct {
	mock *mock.Mock
}

func (_m *QueryCoordCatalog) EXPECT() *QueryCoordCatalog_Expecter {
	return &QueryCoordCatalog_Expecter{mock: &_m.Mock}
}

// GetCollections provides a mock function with given fields:
func (_m *QueryCoordCatalog) GetCollections() ([]*querypb.CollectionLoadInfo, error) {
	ret := _m.Called()

	var r0 []*querypb.CollectionLoadInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*querypb.CollectionLoadInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*querypb.CollectionLoadInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.CollectionLoadInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollections'
type QueryCoordCatalog_GetCollections_Call struct {
	*mock.Call
}

// GetCollections is a helper method to define mock.On call
func (_e *QueryCoordCatalog_Expecter) GetCollections() *QueryCoordCatalog_GetCollections_Call {
	return &QueryCoordCatalog_GetCollections_Call{Call: _e.mock.On("GetCollections")}
}

func (_c *QueryCoordCatalog_GetCollections_Call) Run(run func()) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryCoordCatalog_GetCollections_Call) Return(_a0 []*querypb.CollectionLoadInfo, _a1 error) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetCollections_Call) RunAndReturn(run func() ([]*querypb.CollectionLoadInfo, error)) *QueryCoordCatalog_GetCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetPartitions provides a mock function with given fields:
func (_m *QueryCoordCatalog) GetPartitions() (map[int64][]*querypb.PartitionLoadInfo, error) {
	ret := _m.Called()

	var r0 map[int64][]*querypb.PartitionLoadInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[int64][]*querypb.PartitionLoadInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[int64][]*querypb.PartitionLoadInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]*querypb.PartitionLoadInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type QueryCoordCatalog_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
func (_e *QueryCoordCatalog_Expecter) GetPartitions() *QueryCoordCatalog_GetPartitions_Call {
	return &QueryCoordCatalog_GetPartitions_Call{Call: _e.mock.On("GetPartitions")}
}

func (_c *QueryCoordCatalog_GetPartitions_Call) Run(run func()) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryCoordCatalog_GetPartitions_Call) Return(_a0 map[int64][]*querypb.PartitionLoadInfo, _a1 error) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetPartitions_Call) RunAndReturn(run func() (map[int64][]*querypb.PartitionLoadInfo, error)) *QueryCoordCatalog_GetPartitions_Call {
	_c.Call.Return(run)
	return _c
}

// GetReplicas provides a mock function with given fields:
func (_m *QueryCoordCatalog) GetReplicas() ([]*querypb.Replica, error) {
	ret := _m.Called()

	var r0 []*querypb.Replica
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*querypb.Replica, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*querypb.Replica); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.Replica)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReplicas'
type QueryCoordCatalog_GetReplicas_Call struct {
	*mock.Call
}

// GetReplicas is a helper method to define mock.On call
func (_e *QueryCoordCatalog_Expecter) GetReplicas() *QueryCoordCatalog_GetReplicas_Call {
	return &QueryCoordCatalog_GetReplicas_Call{Call: _e.mock.On("GetReplicas")}
}

func (_c *QueryCoordCatalog_GetReplicas_Call) Run(run func()) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryCoordCatalog_GetReplicas_Call) Return(_a0 []*querypb.Replica, _a1 error) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetReplicas_Call) RunAndReturn(run func() ([]*querypb.Replica, error)) *QueryCoordCatalog_GetReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// GetResourceGroups provides a mock function with given fields:
func (_m *QueryCoordCatalog) GetResourceGroups() ([]*querypb.ResourceGroup, error) {
	ret := _m.Called()

	var r0 []*querypb.ResourceGroup
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*querypb.ResourceGroup, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*querypb.ResourceGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.ResourceGroup)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryCoordCatalog_GetResourceGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResourceGroups'
type QueryCoordCatalog_GetResourceGroups_Call struct {
	*mock.Call
}

// GetResourceGroups is a helper method to define mock.On call
func (_e *QueryCoordCatalog_Expecter) GetResourceGroups() *QueryCoordCatalog_GetResourceGroups_Call {
	return &QueryCoordCatalog_GetResourceGroups_Call{Call: _e.mock.On("GetResourceGroups")}
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) Run(run func()) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) Return(_a0 []*querypb.ResourceGroup, _a1 error) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QueryCoordCatalog_GetResourceGroups_Call) RunAndReturn(run func() ([]*querypb.ResourceGroup, error)) *QueryCoordCatalog_GetResourceGroups_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseCollection provides a mock function with given fields: collection
func (_m *QueryCoordCatalog) ReleaseCollection(collection int64) error {
	ret := _m.Called(collection)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(collection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseCollection'
type QueryCoordCatalog_ReleaseCollection_Call struct {
	*mock.Call
}

// ReleaseCollection is a helper method to define mock.On call
//   - collection int64
func (_e *QueryCoordCatalog_Expecter) ReleaseCollection(collection interface{}) *QueryCoordCatalog_ReleaseCollection_Call {
	return &QueryCoordCatalog_ReleaseCollection_Call{Call: _e.mock.On("ReleaseCollection", collection)}
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) Run(run func(collection int64)) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseCollection_Call) RunAndReturn(run func(int64) error) *QueryCoordCatalog_ReleaseCollection_Call {
	_c.Call.Return(run)
	return _c
}

// ReleasePartition provides a mock function with given fields: collection, partitions
func (_m *QueryCoordCatalog) ReleasePartition(collection int64, partitions ...int64) error {
	_va := make([]interface{}, len(partitions))
	for _i := range partitions {
		_va[_i] = partitions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, ...int64) error); ok {
		r0 = rf(collection, partitions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleasePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleasePartition'
type QueryCoordCatalog_ReleasePartition_Call struct {
	*mock.Call
}

// ReleasePartition is a helper method to define mock.On call
//   - collection int64
//   - partitions ...int64
func (_e *QueryCoordCatalog_Expecter) ReleasePartition(collection interface{}, partitions ...interface{}) *QueryCoordCatalog_ReleasePartition_Call {
	return &QueryCoordCatalog_ReleasePartition_Call{Call: _e.mock.On("ReleasePartition",
		append([]interface{}{collection}, partitions...)...)}
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) Run(run func(collection int64, partitions ...int64)) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(int64), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) Return(_a0 error) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleasePartition_Call) RunAndReturn(run func(int64, ...int64) error) *QueryCoordCatalog_ReleasePartition_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseReplica provides a mock function with given fields: collection, replica
func (_m *QueryCoordCatalog) ReleaseReplica(collection int64, replica int64) error {
	ret := _m.Called(collection, replica)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(collection, replica)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplica'
type QueryCoordCatalog_ReleaseReplica_Call struct {
	*mock.Call
}

// ReleaseReplica is a helper method to define mock.On call
//   - collection int64
//   - replica int64
func (_e *QueryCoordCatalog_Expecter) ReleaseReplica(collection interface{}, replica interface{}) *QueryCoordCatalog_ReleaseReplica_Call {
	return &QueryCoordCatalog_ReleaseReplica_Call{Call: _e.mock.On("ReleaseReplica", collection, replica)}
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) Run(run func(collection int64, replica int64)) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplica_Call) RunAndReturn(run func(int64, int64) error) *QueryCoordCatalog_ReleaseReplica_Call {
	_c.Call.Return(run)
	return _c
}

// ReleaseReplicas provides a mock function with given fields: collectionID
func (_m *QueryCoordCatalog) ReleaseReplicas(collectionID int64) error {
	ret := _m.Called(collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_ReleaseReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplicas'
type QueryCoordCatalog_ReleaseReplicas_Call struct {
	*mock.Call
}

// ReleaseReplicas is a helper method to define mock.On call
//   - collectionID int64
func (_e *QueryCoordCatalog_Expecter) ReleaseReplicas(collectionID interface{}) *QueryCoordCatalog_ReleaseReplicas_Call {
	return &QueryCoordCatalog_ReleaseReplicas_Call{Call: _e.mock.On("ReleaseReplicas", collectionID)}
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) Run(run func(collectionID int64)) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) Return(_a0 error) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_ReleaseReplicas_Call) RunAndReturn(run func(int64) error) *QueryCoordCatalog_ReleaseReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveResourceGroup provides a mock function with given fields: rgName
func (_m *QueryCoordCatalog) RemoveResourceGroup(rgName string) error {
	ret := _m.Called(rgName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(rgName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_RemoveResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveResourceGroup'
type QueryCoordCatalog_RemoveResourceGroup_Call struct {
	*mock.Call
}

// RemoveResourceGroup is a helper method to define mock.On call
//   - rgName string
func (_e *QueryCoordCatalog_Expecter) RemoveResourceGroup(rgName interface{}) *QueryCoordCatalog_RemoveResourceGroup_Call {
	return &QueryCoordCatalog_RemoveResourceGroup_Call{Call: _e.mock.On("RemoveResourceGroup", rgName)}
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) Run(run func(rgName string)) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) Return(_a0 error) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_RemoveResourceGroup_Call) RunAndReturn(run func(string) error) *QueryCoordCatalog_RemoveResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// SaveCollection provides a mock function with given fields: collection, partitions
func (_m *QueryCoordCatalog) SaveCollection(collection *querypb.CollectionLoadInfo, partitions ...*querypb.PartitionLoadInfo) error {
	_va := make([]interface{}, len(partitions))
	for _i := range partitions {
		_va[_i] = partitions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*querypb.CollectionLoadInfo, ...*querypb.PartitionLoadInfo) error); ok {
		r0 = rf(collection, partitions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveCollection'
type QueryCoordCatalog_SaveCollection_Call struct {
	*mock.Call
}

// SaveCollection is a helper method to define mock.On call
//   - collection *querypb.CollectionLoadInfo
//   - partitions ...*querypb.PartitionLoadInfo
func (_e *QueryCoordCatalog_Expecter) SaveCollection(collection interface{}, partitions ...interface{}) *QueryCoordCatalog_SaveCollection_Call {
	return &QueryCoordCatalog_SaveCollection_Call{Call: _e.mock.On("SaveCollection",
		append([]interface{}{collection}, partitions...)...)}
}

func (_c *QueryCoordCatalog_SaveCollection_Call) Run(run func(collection *querypb.CollectionLoadInfo, partitions ...*querypb.PartitionLoadInfo)) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.PartitionLoadInfo, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.PartitionLoadInfo)
			}
		}
		run(args[0].(*querypb.CollectionLoadInfo), variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveCollection_Call) Return(_a0 error) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveCollection_Call) RunAndReturn(run func(*querypb.CollectionLoadInfo, ...*querypb.PartitionLoadInfo) error) *QueryCoordCatalog_SaveCollection_Call {
	_c.Call.Return(run)
	return _c
}

// SavePartition provides a mock function with given fields: info
func (_m *QueryCoordCatalog) SavePartition(info ...*querypb.PartitionLoadInfo) error {
	_va := make([]interface{}, len(info))
	for _i := range info {
		_va[_i] = info[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*querypb.PartitionLoadInfo) error); ok {
		r0 = rf(info...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SavePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SavePartition'
type QueryCoordCatalog_SavePartition_Call struct {
	*mock.Call
}

// SavePartition is a helper method to define mock.On call
//   - info ...*querypb.PartitionLoadInfo
func (_e *QueryCoordCatalog_Expecter) SavePartition(info ...interface{}) *QueryCoordCatalog_SavePartition_Call {
	return &QueryCoordCatalog_SavePartition_Call{Call: _e.mock.On("SavePartition",
		append([]interface{}{}, info...)...)}
}

func (_c *QueryCoordCatalog_SavePartition_Call) Run(run func(info ...*querypb.PartitionLoadInfo)) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.PartitionLoadInfo, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.PartitionLoadInfo)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SavePartition_Call) Return(_a0 error) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SavePartition_Call) RunAndReturn(run func(...*querypb.PartitionLoadInfo) error) *QueryCoordCatalog_SavePartition_Call {
	_c.Call.Return(run)
	return _c
}

// SaveReplica provides a mock function with given fields: replica
func (_m *QueryCoordCatalog) SaveReplica(replica *querypb.Replica) error {
	ret := _m.Called(replica)

	var r0 error
	if rf, ok := ret.Get(0).(func(*querypb.Replica) error); ok {
		r0 = rf(replica)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveReplica'
type QueryCoordCatalog_SaveReplica_Call struct {
	*mock.Call
}

// SaveReplica is a helper method to define mock.On call
//   - replica *querypb.Replica
func (_e *QueryCoordCatalog_Expecter) SaveReplica(replica interface{}) *QueryCoordCatalog_SaveReplica_Call {
	return &QueryCoordCatalog_SaveReplica_Call{Call: _e.mock.On("SaveReplica", replica)}
}

func (_c *QueryCoordCatalog_SaveReplica_Call) Run(run func(replica *querypb.Replica)) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*querypb.Replica))
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveReplica_Call) Return(_a0 error) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveReplica_Call) RunAndReturn(run func(*querypb.Replica) error) *QueryCoordCatalog_SaveReplica_Call {
	_c.Call.Return(run)
	return _c
}

// SaveResourceGroup provides a mock function with given fields: rgs
func (_m *QueryCoordCatalog) SaveResourceGroup(rgs ...*querypb.ResourceGroup) error {
	_va := make([]interface{}, len(rgs))
	for _i := range rgs {
		_va[_i] = rgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*querypb.ResourceGroup) error); ok {
		r0 = rf(rgs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryCoordCatalog_SaveResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveResourceGroup'
type QueryCoordCatalog_SaveResourceGroup_Call struct {
	*mock.Call
}

// SaveResourceGroup is a helper method to define mock.On call
//   - rgs ...*querypb.ResourceGroup
func (_e *QueryCoordCatalog_Expecter) SaveResourceGroup(rgs ...interface{}) *QueryCoordCatalog_SaveResourceGroup_Call {
	return &QueryCoordCatalog_SaveResourceGroup_Call{Call: _e.mock.On("SaveResourceGroup",
		append([]interface{}{}, rgs...)...)}
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) Run(run func(rgs ...*querypb.ResourceGroup)) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.ResourceGroup, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.ResourceGroup)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) Return(_a0 error) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QueryCoordCatalog_SaveResourceGroup_Call) RunAndReturn(run func(...*querypb.ResourceGroup) error) *QueryCoordCatalog_SaveResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewQueryCoordCatalog creates a new instance of QueryCoordCatalog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryCoordCatalog(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryCoordCatalog {
	mock := &QueryCoordCatalog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
