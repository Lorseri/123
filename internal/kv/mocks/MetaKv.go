// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	clientv3 "go.etcd.io/etcd/client/v3"

	mock "github.com/stretchr/testify/mock"
)

// MetaKv is an autogenerated mock type for the MetaKv type
type MetaKv struct {
	mock.Mock
}

type MetaKv_Expecter struct {
	mock *mock.Mock
}

func (_m *MetaKv) EXPECT() *MetaKv_Expecter {
	return &MetaKv_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MetaKv) Close() {
	_m.Called()
}

// MetaKv_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MetaKv_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MetaKv_Expecter) Close() *MetaKv_Close_Call {
	return &MetaKv_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MetaKv_Close_Call) Run(run func()) *MetaKv_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MetaKv_Close_Call) Return() *MetaKv_Close_Call {
	_c.Call.Return()
	return _c
}

// CompareValueAndSwap provides a mock function with given fields: key, value, target, opts
func (_m *MetaKv) CompareValueAndSwap(key string, value string, target string, opts ...clientv3.OpOption) (bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, value, target)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, ...clientv3.OpOption) bool); ok {
		r0 = rf(key, value, target, opts...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, ...clientv3.OpOption) error); ok {
		r1 = rf(key, value, target, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_CompareValueAndSwap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompareValueAndSwap'
type MetaKv_CompareValueAndSwap_Call struct {
	*mock.Call
}

// CompareValueAndSwap is a helper method to define mock.On call
//   - key string
//   - value string
//   - target string
//   - opts ...clientv3.OpOption
func (_e *MetaKv_Expecter) CompareValueAndSwap(key interface{}, value interface{}, target interface{}, opts ...interface{}) *MetaKv_CompareValueAndSwap_Call {
	return &MetaKv_CompareValueAndSwap_Call{Call: _e.mock.On("CompareValueAndSwap",
		append([]interface{}{key, value, target}, opts...)...)}
}

func (_c *MetaKv_CompareValueAndSwap_Call) Run(run func(key string, value string, target string, opts ...clientv3.OpOption)) *MetaKv_CompareValueAndSwap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(string), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MetaKv_CompareValueAndSwap_Call) Return(_a0 bool, _a1 error) *MetaKv_CompareValueAndSwap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CompareVersionAndSwap provides a mock function with given fields: key, version, target, opts
func (_m *MetaKv) CompareVersionAndSwap(key string, version int64, target string, opts ...clientv3.OpOption) (bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, version, target)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64, string, ...clientv3.OpOption) bool); ok {
		r0 = rf(key, version, target, opts...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64, string, ...clientv3.OpOption) error); ok {
		r1 = rf(key, version, target, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_CompareVersionAndSwap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompareVersionAndSwap'
type MetaKv_CompareVersionAndSwap_Call struct {
	*mock.Call
}

// CompareVersionAndSwap is a helper method to define mock.On call
//   - key string
//   - version int64
//   - target string
//   - opts ...clientv3.OpOption
func (_e *MetaKv_Expecter) CompareVersionAndSwap(key interface{}, version interface{}, target interface{}, opts ...interface{}) *MetaKv_CompareVersionAndSwap_Call {
	return &MetaKv_CompareVersionAndSwap_Call{Call: _e.mock.On("CompareVersionAndSwap",
		append([]interface{}{key, version, target}, opts...)...)}
}

func (_c *MetaKv_CompareVersionAndSwap_Call) Run(run func(key string, version int64, target string, opts ...clientv3.OpOption)) *MetaKv_CompareVersionAndSwap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]clientv3.OpOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(clientv3.OpOption)
			}
		}
		run(args[0].(string), args[1].(int64), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MetaKv_CompareVersionAndSwap_Call) Return(_a0 bool, _a1 error) *MetaKv_CompareVersionAndSwap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPath provides a mock function with given fields: key
func (_m *MetaKv) GetPath(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MetaKv_GetPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPath'
type MetaKv_GetPath_Call struct {
	*mock.Call
}

// GetPath is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) GetPath(key interface{}) *MetaKv_GetPath_Call {
	return &MetaKv_GetPath_Call{Call: _e.mock.On("GetPath", key)}
}

func (_c *MetaKv_GetPath_Call) Run(run func(key string)) *MetaKv_GetPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_GetPath_Call) Return(_a0 string) *MetaKv_GetPath_Call {
	_c.Call.Return(_a0)
	return _c
}

// Grant provides a mock function with given fields: ttl
func (_m *MetaKv) Grant(ttl int64) (clientv3.LeaseID, error) {
	ret := _m.Called(ttl)

	var r0 clientv3.LeaseID
	if rf, ok := ret.Get(0).(func(int64) clientv3.LeaseID); ok {
		r0 = rf(ttl)
	} else {
		r0 = ret.Get(0).(clientv3.LeaseID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_Grant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Grant'
type MetaKv_Grant_Call struct {
	*mock.Call
}

// Grant is a helper method to define mock.On call
//   - ttl int64
func (_e *MetaKv_Expecter) Grant(ttl interface{}) *MetaKv_Grant_Call {
	return &MetaKv_Grant_Call{Call: _e.mock.On("Grant", ttl)}
}

func (_c *MetaKv_Grant_Call) Run(run func(ttl int64)) *MetaKv_Grant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MetaKv_Grant_Call) Return(id clientv3.LeaseID, err error) *MetaKv_Grant_Call {
	_c.Call.Return(id, err)
	return _c
}

// KeepAlive provides a mock function with given fields: id
func (_m *MetaKv) KeepAlive(id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	ret := _m.Called(id)

	var r0 <-chan *clientv3.LeaseKeepAliveResponse
	if rf, ok := ret.Get(0).(func(clientv3.LeaseID) <-chan *clientv3.LeaseKeepAliveResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *clientv3.LeaseKeepAliveResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(clientv3.LeaseID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_KeepAlive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KeepAlive'
type MetaKv_KeepAlive_Call struct {
	*mock.Call
}

// KeepAlive is a helper method to define mock.On call
//   - id clientv3.LeaseID
func (_e *MetaKv_Expecter) KeepAlive(id interface{}) *MetaKv_KeepAlive_Call {
	return &MetaKv_KeepAlive_Call{Call: _e.mock.On("KeepAlive", id)}
}

func (_c *MetaKv_KeepAlive_Call) Run(run func(id clientv3.LeaseID)) *MetaKv_KeepAlive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(clientv3.LeaseID))
	})
	return _c
}

func (_c *MetaKv_KeepAlive_Call) Return(_a0 <-chan *clientv3.LeaseKeepAliveResponse, _a1 error) *MetaKv_KeepAlive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Load provides a mock function with given fields: key
func (_m *MetaKv) Load(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MetaKv_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Load(key interface{}) *MetaKv_Load_Call {
	return &MetaKv_Load_Call{Call: _e.mock.On("Load", key)}
}

func (_c *MetaKv_Load_Call) Run(run func(key string)) *MetaKv_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Load_Call) Return(_a0 string, _a1 error) *MetaKv_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadWithPrefix provides a mock function with given fields: key
func (_m *MetaKv) LoadWithPrefix(key string) ([]string, []string, error) {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MetaKv_LoadWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithPrefix'
type MetaKv_LoadWithPrefix_Call struct {
	*mock.Call
}

// LoadWithPrefix is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) LoadWithPrefix(key interface{}) *MetaKv_LoadWithPrefix_Call {
	return &MetaKv_LoadWithPrefix_Call{Call: _e.mock.On("LoadWithPrefix", key)}
}

func (_c *MetaKv_LoadWithPrefix_Call) Run(run func(key string)) *MetaKv_LoadWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_LoadWithPrefix_Call) Return(_a0 []string, _a1 []string, _a2 error) *MetaKv_LoadWithPrefix_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

// LoadWithPrefix2 provides a mock function with given fields: key
func (_m *MetaKv) LoadWithPrefix2(key string) ([]string, []string, []int64, error) {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 []int64
	if rf, ok := ret.Get(2).(func(string) []int64); ok {
		r2 = rf(key)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]int64)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(key)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MetaKv_LoadWithPrefix2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithPrefix2'
type MetaKv_LoadWithPrefix2_Call struct {
	*mock.Call
}

// LoadWithPrefix2 is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) LoadWithPrefix2(key interface{}) *MetaKv_LoadWithPrefix2_Call {
	return &MetaKv_LoadWithPrefix2_Call{Call: _e.mock.On("LoadWithPrefix2", key)}
}

func (_c *MetaKv_LoadWithPrefix2_Call) Run(run func(key string)) *MetaKv_LoadWithPrefix2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_LoadWithPrefix2_Call) Return(_a0 []string, _a1 []string, _a2 []int64, _a3 error) *MetaKv_LoadWithPrefix2_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

// LoadWithRevision provides a mock function with given fields: key
func (_m *MetaKv) LoadWithRevision(key string) ([]string, []string, int64, error) {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 int64
	if rf, ok := ret.Get(2).(func(string) int64); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Get(2).(int64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(key)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MetaKv_LoadWithRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithRevision'
type MetaKv_LoadWithRevision_Call struct {
	*mock.Call
}

// LoadWithRevision is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) LoadWithRevision(key interface{}) *MetaKv_LoadWithRevision_Call {
	return &MetaKv_LoadWithRevision_Call{Call: _e.mock.On("LoadWithRevision", key)}
}

func (_c *MetaKv_LoadWithRevision_Call) Run(run func(key string)) *MetaKv_LoadWithRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_LoadWithRevision_Call) Return(_a0 []string, _a1 []string, _a2 int64, _a3 error) *MetaKv_LoadWithRevision_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

// LoadWithRevisionAndVersions provides a mock function with given fields: key
func (_m *MetaKv) LoadWithRevisionAndVersions(key string) ([]string, []string, []int64, int64, error) {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 []int64
	if rf, ok := ret.Get(2).(func(string) []int64); ok {
		r2 = rf(key)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]int64)
		}
	}

	var r3 int64
	if rf, ok := ret.Get(3).(func(string) int64); ok {
		r3 = rf(key)
	} else {
		r3 = ret.Get(3).(int64)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(string) error); ok {
		r4 = rf(key)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// MetaKv_LoadWithRevisionAndVersions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithRevisionAndVersions'
type MetaKv_LoadWithRevisionAndVersions_Call struct {
	*mock.Call
}

// LoadWithRevisionAndVersions is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) LoadWithRevisionAndVersions(key interface{}) *MetaKv_LoadWithRevisionAndVersions_Call {
	return &MetaKv_LoadWithRevisionAndVersions_Call{Call: _e.mock.On("LoadWithRevisionAndVersions", key)}
}

func (_c *MetaKv_LoadWithRevisionAndVersions_Call) Run(run func(key string)) *MetaKv_LoadWithRevisionAndVersions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_LoadWithRevisionAndVersions_Call) Return(_a0 []string, _a1 []string, _a2 []int64, _a3 int64, _a4 error) *MetaKv_LoadWithRevisionAndVersions_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3, _a4)
	return _c
}

// MultiLoad provides a mock function with given fields: keys
func (_m *MetaKv) MultiLoad(keys []string) ([]string, error) {
	ret := _m.Called(keys)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaKv_MultiLoad_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiLoad'
type MetaKv_MultiLoad_Call struct {
	*mock.Call
}

// MultiLoad is a helper method to define mock.On call
//   - keys []string
func (_e *MetaKv_Expecter) MultiLoad(keys interface{}) *MetaKv_MultiLoad_Call {
	return &MetaKv_MultiLoad_Call{Call: _e.mock.On("MultiLoad", keys)}
}

func (_c *MetaKv_MultiLoad_Call) Run(run func(keys []string)) *MetaKv_MultiLoad_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiLoad_Call) Return(_a0 []string, _a1 error) *MetaKv_MultiLoad_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// MultiRemove provides a mock function with given fields: keys
func (_m *MetaKv) MultiRemove(keys []string) error {
	ret := _m.Called(keys)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRemove'
type MetaKv_MultiRemove_Call struct {
	*mock.Call
}

// MultiRemove is a helper method to define mock.On call
//   - keys []string
func (_e *MetaKv_Expecter) MultiRemove(keys interface{}) *MetaKv_MultiRemove_Call {
	return &MetaKv_MultiRemove_Call{Call: _e.mock.On("MultiRemove", keys)}
}

func (_c *MetaKv_MultiRemove_Call) Run(run func(keys []string)) *MetaKv_MultiRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiRemove_Call) Return(_a0 error) *MetaKv_MultiRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

// MultiRemoveWithPrefix provides a mock function with given fields: keys
func (_m *MetaKv) MultiRemoveWithPrefix(keys []string) error {
	ret := _m.Called(keys)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiRemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRemoveWithPrefix'
type MetaKv_MultiRemoveWithPrefix_Call struct {
	*mock.Call
}

// MultiRemoveWithPrefix is a helper method to define mock.On call
//   - keys []string
func (_e *MetaKv_Expecter) MultiRemoveWithPrefix(keys interface{}) *MetaKv_MultiRemoveWithPrefix_Call {
	return &MetaKv_MultiRemoveWithPrefix_Call{Call: _e.mock.On("MultiRemoveWithPrefix", keys)}
}

func (_c *MetaKv_MultiRemoveWithPrefix_Call) Run(run func(keys []string)) *MetaKv_MultiRemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiRemoveWithPrefix_Call) Return(_a0 error) *MetaKv_MultiRemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// MultiSave provides a mock function with given fields: kvs
func (_m *MetaKv) MultiSave(kvs map[string]string) error {
	ret := _m.Called(kvs)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(kvs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSave'
type MetaKv_MultiSave_Call struct {
	*mock.Call
}

// MultiSave is a helper method to define mock.On call
//   - kvs map[string]string
func (_e *MetaKv_Expecter) MultiSave(kvs interface{}) *MetaKv_MultiSave_Call {
	return &MetaKv_MultiSave_Call{Call: _e.mock.On("MultiSave", kvs)}
}

func (_c *MetaKv_MultiSave_Call) Run(run func(kvs map[string]string)) *MetaKv_MultiSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *MetaKv_MultiSave_Call) Return(_a0 error) *MetaKv_MultiSave_Call {
	_c.Call.Return(_a0)
	return _c
}

// MultiSaveAndRemove provides a mock function with given fields: saves, removals
func (_m *MetaKv) MultiSaveAndRemove(saves map[string]string, removals []string) error {
	ret := _m.Called(saves, removals)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string) error); ok {
		r0 = rf(saves, removals)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSaveAndRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemove'
type MetaKv_MultiSaveAndRemove_Call struct {
	*mock.Call
}

// MultiSaveAndRemove is a helper method to define mock.On call
//   - saves map[string]string
//   - removals []string
func (_e *MetaKv_Expecter) MultiSaveAndRemove(saves interface{}, removals interface{}) *MetaKv_MultiSaveAndRemove_Call {
	return &MetaKv_MultiSaveAndRemove_Call{Call: _e.mock.On("MultiSaveAndRemove", saves, removals)}
}

func (_c *MetaKv_MultiSaveAndRemove_Call) Run(run func(saves map[string]string, removals []string)) *MetaKv_MultiSaveAndRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiSaveAndRemove_Call) Return(_a0 error) *MetaKv_MultiSaveAndRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

// MultiSaveAndRemoveWithPrefix provides a mock function with given fields: saves, removals
func (_m *MetaKv) MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string) error {
	ret := _m.Called(saves, removals)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string) error); ok {
		r0 = rf(saves, removals)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_MultiSaveAndRemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemoveWithPrefix'
type MetaKv_MultiSaveAndRemoveWithPrefix_Call struct {
	*mock.Call
}

// MultiSaveAndRemoveWithPrefix is a helper method to define mock.On call
//   - saves map[string]string
//   - removals []string
func (_e *MetaKv_Expecter) MultiSaveAndRemoveWithPrefix(saves interface{}, removals interface{}) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	return &MetaKv_MultiSaveAndRemoveWithPrefix_Call{Call: _e.mock.On("MultiSaveAndRemoveWithPrefix", saves, removals)}
}

func (_c *MetaKv_MultiSaveAndRemoveWithPrefix_Call) Run(run func(saves map[string]string, removals []string)) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string), args[1].([]string))
	})
	return _c
}

func (_c *MetaKv_MultiSaveAndRemoveWithPrefix_Call) Return(_a0 error) *MetaKv_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// Remove provides a mock function with given fields: key
func (_m *MetaKv) Remove(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MetaKv_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Remove(key interface{}) *MetaKv_Remove_Call {
	return &MetaKv_Remove_Call{Call: _e.mock.On("Remove", key)}
}

func (_c *MetaKv_Remove_Call) Run(run func(key string)) *MetaKv_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Remove_Call) Return(_a0 error) *MetaKv_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveWithPrefix provides a mock function with given fields: key
func (_m *MetaKv) RemoveWithPrefix(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_RemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveWithPrefix'
type MetaKv_RemoveWithPrefix_Call struct {
	*mock.Call
}

// RemoveWithPrefix is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) RemoveWithPrefix(key interface{}) *MetaKv_RemoveWithPrefix_Call {
	return &MetaKv_RemoveWithPrefix_Call{Call: _e.mock.On("RemoveWithPrefix", key)}
}

func (_c *MetaKv_RemoveWithPrefix_Call) Run(run func(key string)) *MetaKv_RemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_RemoveWithPrefix_Call) Return(_a0 error) *MetaKv_RemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// Save provides a mock function with given fields: key, value
func (_m *MetaKv) Save(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MetaKv_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MetaKv_Expecter) Save(key interface{}, value interface{}) *MetaKv_Save_Call {
	return &MetaKv_Save_Call{Call: _e.mock.On("Save", key, value)}
}

func (_c *MetaKv_Save_Call) Run(run func(key string, value string)) *MetaKv_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MetaKv_Save_Call) Return(_a0 error) *MetaKv_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveWithIgnoreLease provides a mock function with given fields: key, value
func (_m *MetaKv) SaveWithIgnoreLease(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_SaveWithIgnoreLease_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveWithIgnoreLease'
type MetaKv_SaveWithIgnoreLease_Call struct {
	*mock.Call
}

// SaveWithIgnoreLease is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MetaKv_Expecter) SaveWithIgnoreLease(key interface{}, value interface{}) *MetaKv_SaveWithIgnoreLease_Call {
	return &MetaKv_SaveWithIgnoreLease_Call{Call: _e.mock.On("SaveWithIgnoreLease", key, value)}
}

func (_c *MetaKv_SaveWithIgnoreLease_Call) Run(run func(key string, value string)) *MetaKv_SaveWithIgnoreLease_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MetaKv_SaveWithIgnoreLease_Call) Return(_a0 error) *MetaKv_SaveWithIgnoreLease_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveWithLease provides a mock function with given fields: key, value, id
func (_m *MetaKv) SaveWithLease(key string, value string, id clientv3.LeaseID) error {
	ret := _m.Called(key, value, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, clientv3.LeaseID) error); ok {
		r0 = rf(key, value, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_SaveWithLease_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveWithLease'
type MetaKv_SaveWithLease_Call struct {
	*mock.Call
}

// SaveWithLease is a helper method to define mock.On call
//   - key string
//   - value string
//   - id clientv3.LeaseID
func (_e *MetaKv_Expecter) SaveWithLease(key interface{}, value interface{}, id interface{}) *MetaKv_SaveWithLease_Call {
	return &MetaKv_SaveWithLease_Call{Call: _e.mock.On("SaveWithLease", key, value, id)}
}

func (_c *MetaKv_SaveWithLease_Call) Run(run func(key string, value string, id clientv3.LeaseID)) *MetaKv_SaveWithLease_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(clientv3.LeaseID))
	})
	return _c
}

func (_c *MetaKv_SaveWithLease_Call) Return(_a0 error) *MetaKv_SaveWithLease_Call {
	_c.Call.Return(_a0)
	return _c
}

// WalkWithPrefix provides a mock function with given fields: prefix, paginationSize, fn
func (_m *MetaKv) WalkWithPrefix(prefix string, paginationSize int, fn func([]byte, []byte) error) error {
	ret := _m.Called(prefix, paginationSize, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, func([]byte, []byte) error) error); ok {
		r0 = rf(prefix, paginationSize, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaKv_WalkWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WalkWithPrefix'
type MetaKv_WalkWithPrefix_Call struct {
	*mock.Call
}

// WalkWithPrefix is a helper method to define mock.On call
//   - prefix string
//   - paginationSize int
//   - fn func([]byte , []byte) error
func (_e *MetaKv_Expecter) WalkWithPrefix(prefix interface{}, paginationSize interface{}, fn interface{}) *MetaKv_WalkWithPrefix_Call {
	return &MetaKv_WalkWithPrefix_Call{Call: _e.mock.On("WalkWithPrefix", prefix, paginationSize, fn)}
}

func (_c *MetaKv_WalkWithPrefix_Call) Run(run func(prefix string, paginationSize int, fn func([]byte, []byte) error)) *MetaKv_WalkWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int), args[2].(func([]byte, []byte) error))
	})
	return _c
}

func (_c *MetaKv_WalkWithPrefix_Call) Return(_a0 error) *MetaKv_WalkWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// Watch provides a mock function with given fields: key
func (_m *MetaKv) Watch(key string) clientv3.WatchChan {
	ret := _m.Called(key)

	var r0 clientv3.WatchChan
	if rf, ok := ret.Get(0).(func(string) clientv3.WatchChan); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.WatchChan)
		}
	}

	return r0
}

// MetaKv_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MetaKv_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) Watch(key interface{}) *MetaKv_Watch_Call {
	return &MetaKv_Watch_Call{Call: _e.mock.On("Watch", key)}
}

func (_c *MetaKv_Watch_Call) Run(run func(key string)) *MetaKv_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_Watch_Call) Return(_a0 clientv3.WatchChan) *MetaKv_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

// WatchWithPrefix provides a mock function with given fields: key
func (_m *MetaKv) WatchWithPrefix(key string) clientv3.WatchChan {
	ret := _m.Called(key)

	var r0 clientv3.WatchChan
	if rf, ok := ret.Get(0).(func(string) clientv3.WatchChan); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.WatchChan)
		}
	}

	return r0
}

// MetaKv_WatchWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchWithPrefix'
type MetaKv_WatchWithPrefix_Call struct {
	*mock.Call
}

// WatchWithPrefix is a helper method to define mock.On call
//   - key string
func (_e *MetaKv_Expecter) WatchWithPrefix(key interface{}) *MetaKv_WatchWithPrefix_Call {
	return &MetaKv_WatchWithPrefix_Call{Call: _e.mock.On("WatchWithPrefix", key)}
}

func (_c *MetaKv_WatchWithPrefix_Call) Run(run func(key string)) *MetaKv_WatchWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MetaKv_WatchWithPrefix_Call) Return(_a0 clientv3.WatchChan) *MetaKv_WatchWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

// WatchWithRevision provides a mock function with given fields: key, revision
func (_m *MetaKv) WatchWithRevision(key string, revision int64) clientv3.WatchChan {
	ret := _m.Called(key, revision)

	var r0 clientv3.WatchChan
	if rf, ok := ret.Get(0).(func(string, int64) clientv3.WatchChan); ok {
		r0 = rf(key, revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.WatchChan)
		}
	}

	return r0
}

// MetaKv_WatchWithRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchWithRevision'
type MetaKv_WatchWithRevision_Call struct {
	*mock.Call
}

// WatchWithRevision is a helper method to define mock.On call
//   - key string
//   - revision int64
func (_e *MetaKv_Expecter) WatchWithRevision(key interface{}, revision interface{}) *MetaKv_WatchWithRevision_Call {
	return &MetaKv_WatchWithRevision_Call{Call: _e.mock.On("WatchWithRevision", key, revision)}
}

func (_c *MetaKv_WatchWithRevision_Call) Run(run func(key string, revision int64)) *MetaKv_WatchWithRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64))
	})
	return _c
}

func (_c *MetaKv_WatchWithRevision_Call) Return(_a0 clientv3.WatchChan) *MetaKv_WatchWithRevision_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMetaKv interface {
	mock.TestingT
	Cleanup(func())
}

// NewMetaKv creates a new instance of MetaKv. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetaKv(t mockConstructorTestingTNewMetaKv) *MetaKv {
	mock := &MetaKv{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
