// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	predicates "github.com/milvus-io/milvus/internal/kv/predicates"
	mock "github.com/stretchr/testify/mock"
)

// TxnKV is an autogenerated mock type for the TxnKV type
type TxnKV struct {
	mock.Mock
}

type TxnKV_Expecter struct {
	mock *mock.Mock
}

func (_m *TxnKV) EXPECT() *TxnKV_Expecter {
	return &TxnKV_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *TxnKV) Close() {
	_m.Called()
}

// TxnKV_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type TxnKV_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *TxnKV_Expecter) Close() *TxnKV_Close_Call {
	return &TxnKV_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *TxnKV_Close_Call) Run(run func()) *TxnKV_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TxnKV_Close_Call) Return() *TxnKV_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *TxnKV_Close_Call) RunAndReturn(run func()) *TxnKV_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: key
func (_m *TxnKV) Has(key string) (bool, error) {
	ret := _m.Called(key)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxnKV_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type TxnKV_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//  - key string
func (_e *TxnKV_Expecter) Has(key interface{}) *TxnKV_Has_Call {
	return &TxnKV_Has_Call{Call: _e.mock.On("Has", key)}
}

func (_c *TxnKV_Has_Call) Run(run func(key string)) *TxnKV_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_Has_Call) Return(_a0 bool, _a1 error) *TxnKV_Has_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxnKV_Has_Call) RunAndReturn(run func(string) (bool, error)) *TxnKV_Has_Call {
	_c.Call.Return(run)
	return _c
}

// HasPrefix provides a mock function with given fields: prefix
func (_m *TxnKV) HasPrefix(prefix string) (bool, error) {
	ret := _m.Called(prefix)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(prefix)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxnKV_HasPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasPrefix'
type TxnKV_HasPrefix_Call struct {
	*mock.Call
}

// HasPrefix is a helper method to define mock.On call
//  - prefix string
func (_e *TxnKV_Expecter) HasPrefix(prefix interface{}) *TxnKV_HasPrefix_Call {
	return &TxnKV_HasPrefix_Call{Call: _e.mock.On("HasPrefix", prefix)}
}

func (_c *TxnKV_HasPrefix_Call) Run(run func(prefix string)) *TxnKV_HasPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_HasPrefix_Call) Return(_a0 bool, _a1 error) *TxnKV_HasPrefix_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxnKV_HasPrefix_Call) RunAndReturn(run func(string) (bool, error)) *TxnKV_HasPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: key
func (_m *TxnKV) Load(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxnKV_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type TxnKV_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//  - key string
func (_e *TxnKV_Expecter) Load(key interface{}) *TxnKV_Load_Call {
	return &TxnKV_Load_Call{Call: _e.mock.On("Load", key)}
}

func (_c *TxnKV_Load_Call) Run(run func(key string)) *TxnKV_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_Load_Call) Return(_a0 string, _a1 error) *TxnKV_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxnKV_Load_Call) RunAndReturn(run func(string) (string, error)) *TxnKV_Load_Call {
	_c.Call.Return(run)
	return _c
}

// LoadWithPrefix provides a mock function with given fields: key
func (_m *TxnKV) LoadWithPrefix(key string) ([]string, []string, error) {
	ret := _m.Called(key)

	var r0 []string
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) ([]string, []string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) []string); ok {
		r1 = rf(key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TxnKV_LoadWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadWithPrefix'
type TxnKV_LoadWithPrefix_Call struct {
	*mock.Call
}

// LoadWithPrefix is a helper method to define mock.On call
//  - key string
func (_e *TxnKV_Expecter) LoadWithPrefix(key interface{}) *TxnKV_LoadWithPrefix_Call {
	return &TxnKV_LoadWithPrefix_Call{Call: _e.mock.On("LoadWithPrefix", key)}
}

func (_c *TxnKV_LoadWithPrefix_Call) Run(run func(key string)) *TxnKV_LoadWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_LoadWithPrefix_Call) Return(_a0 []string, _a1 []string, _a2 error) *TxnKV_LoadWithPrefix_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *TxnKV_LoadWithPrefix_Call) RunAndReturn(run func(string) ([]string, []string, error)) *TxnKV_LoadWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// MultiLoad provides a mock function with given fields: keys
func (_m *TxnKV) MultiLoad(keys []string) ([]string, error) {
	ret := _m.Called(keys)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]string, error)); ok {
		return rf(keys)
	}
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxnKV_MultiLoad_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiLoad'
type TxnKV_MultiLoad_Call struct {
	*mock.Call
}

// MultiLoad is a helper method to define mock.On call
//  - keys []string
func (_e *TxnKV_Expecter) MultiLoad(keys interface{}) *TxnKV_MultiLoad_Call {
	return &TxnKV_MultiLoad_Call{Call: _e.mock.On("MultiLoad", keys)}
}

func (_c *TxnKV_MultiLoad_Call) Run(run func(keys []string)) *TxnKV_MultiLoad_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *TxnKV_MultiLoad_Call) Return(_a0 []string, _a1 error) *TxnKV_MultiLoad_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxnKV_MultiLoad_Call) RunAndReturn(run func([]string) ([]string, error)) *TxnKV_MultiLoad_Call {
	_c.Call.Return(run)
	return _c
}

// MultiRemove provides a mock function with given fields: keys
func (_m *TxnKV) MultiRemove(keys []string) error {
	ret := _m.Called(keys)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(keys)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_MultiRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRemove'
type TxnKV_MultiRemove_Call struct {
	*mock.Call
}

// MultiRemove is a helper method to define mock.On call
//  - keys []string
func (_e *TxnKV_Expecter) MultiRemove(keys interface{}) *TxnKV_MultiRemove_Call {
	return &TxnKV_MultiRemove_Call{Call: _e.mock.On("MultiRemove", keys)}
}

func (_c *TxnKV_MultiRemove_Call) Run(run func(keys []string)) *TxnKV_MultiRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *TxnKV_MultiRemove_Call) Return(_a0 error) *TxnKV_MultiRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_MultiRemove_Call) RunAndReturn(run func([]string) error) *TxnKV_MultiRemove_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSave provides a mock function with given fields: kvs
func (_m *TxnKV) MultiSave(kvs map[string]string) error {
	ret := _m.Called(kvs)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(kvs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_MultiSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSave'
type TxnKV_MultiSave_Call struct {
	*mock.Call
}

// MultiSave is a helper method to define mock.On call
//  - kvs map[string]string
func (_e *TxnKV_Expecter) MultiSave(kvs interface{}) *TxnKV_MultiSave_Call {
	return &TxnKV_MultiSave_Call{Call: _e.mock.On("MultiSave", kvs)}
}

func (_c *TxnKV_MultiSave_Call) Run(run func(kvs map[string]string)) *TxnKV_MultiSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]string))
	})
	return _c
}

func (_c *TxnKV_MultiSave_Call) Return(_a0 error) *TxnKV_MultiSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_MultiSave_Call) RunAndReturn(run func(map[string]string) error) *TxnKV_MultiSave_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSaveAndRemove provides a mock function with given fields: saves, removals, preds
func (_m *TxnKV) MultiSaveAndRemove(saves map[string]string, removals []string, preds ...predicates.Predicate) error {
	_va := make([]interface{}, len(preds))
	for _i := range preds {
		_va[_i] = preds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, saves, removals)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string, ...predicates.Predicate) error); ok {
		r0 = rf(saves, removals, preds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_MultiSaveAndRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemove'
type TxnKV_MultiSaveAndRemove_Call struct {
	*mock.Call
}

// MultiSaveAndRemove is a helper method to define mock.On call
//  - saves map[string]string
//  - removals []string
//  - preds ...predicates.Predicate
func (_e *TxnKV_Expecter) MultiSaveAndRemove(saves interface{}, removals interface{}, preds ...interface{}) *TxnKV_MultiSaveAndRemove_Call {
	return &TxnKV_MultiSaveAndRemove_Call{Call: _e.mock.On("MultiSaveAndRemove",
		append([]interface{}{saves, removals}, preds...)...)}
}

func (_c *TxnKV_MultiSaveAndRemove_Call) Run(run func(saves map[string]string, removals []string, preds ...predicates.Predicate)) *TxnKV_MultiSaveAndRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]predicates.Predicate, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(predicates.Predicate)
			}
		}
		run(args[0].(map[string]string), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *TxnKV_MultiSaveAndRemove_Call) Return(_a0 error) *TxnKV_MultiSaveAndRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_MultiSaveAndRemove_Call) RunAndReturn(run func(map[string]string, []string, ...predicates.Predicate) error) *TxnKV_MultiSaveAndRemove_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSaveAndRemoveWithPrefix provides a mock function with given fields: saves, removals, preds
func (_m *TxnKV) MultiSaveAndRemoveWithPrefix(saves map[string]string, removals []string, preds ...predicates.Predicate) error {
	_va := make([]interface{}, len(preds))
	for _i := range preds {
		_va[_i] = preds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, saves, removals)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string, []string, ...predicates.Predicate) error); ok {
		r0 = rf(saves, removals, preds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_MultiSaveAndRemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSaveAndRemoveWithPrefix'
type TxnKV_MultiSaveAndRemoveWithPrefix_Call struct {
	*mock.Call
}

// MultiSaveAndRemoveWithPrefix is a helper method to define mock.On call
//  - saves map[string]string
//  - removals []string
//  - preds ...predicates.Predicate
func (_e *TxnKV_Expecter) MultiSaveAndRemoveWithPrefix(saves interface{}, removals interface{}, preds ...interface{}) *TxnKV_MultiSaveAndRemoveWithPrefix_Call {
	return &TxnKV_MultiSaveAndRemoveWithPrefix_Call{Call: _e.mock.On("MultiSaveAndRemoveWithPrefix",
		append([]interface{}{saves, removals}, preds...)...)}
}

func (_c *TxnKV_MultiSaveAndRemoveWithPrefix_Call) Run(run func(saves map[string]string, removals []string, preds ...predicates.Predicate)) *TxnKV_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]predicates.Predicate, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(predicates.Predicate)
			}
		}
		run(args[0].(map[string]string), args[1].([]string), variadicArgs...)
	})
	return _c
}

func (_c *TxnKV_MultiSaveAndRemoveWithPrefix_Call) Return(_a0 error) *TxnKV_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_MultiSaveAndRemoveWithPrefix_Call) RunAndReturn(run func(map[string]string, []string, ...predicates.Predicate) error) *TxnKV_MultiSaveAndRemoveWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: key
func (_m *TxnKV) Remove(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type TxnKV_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//  - key string
func (_e *TxnKV_Expecter) Remove(key interface{}) *TxnKV_Remove_Call {
	return &TxnKV_Remove_Call{Call: _e.mock.On("Remove", key)}
}

func (_c *TxnKV_Remove_Call) Run(run func(key string)) *TxnKV_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_Remove_Call) Return(_a0 error) *TxnKV_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_Remove_Call) RunAndReturn(run func(string) error) *TxnKV_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveWithPrefix provides a mock function with given fields: key
func (_m *TxnKV) RemoveWithPrefix(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_RemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveWithPrefix'
type TxnKV_RemoveWithPrefix_Call struct {
	*mock.Call
}

// RemoveWithPrefix is a helper method to define mock.On call
//  - key string
func (_e *TxnKV_Expecter) RemoveWithPrefix(key interface{}) *TxnKV_RemoveWithPrefix_Call {
	return &TxnKV_RemoveWithPrefix_Call{Call: _e.mock.On("RemoveWithPrefix", key)}
}

func (_c *TxnKV_RemoveWithPrefix_Call) Run(run func(key string)) *TxnKV_RemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TxnKV_RemoveWithPrefix_Call) Return(_a0 error) *TxnKV_RemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_RemoveWithPrefix_Call) RunAndReturn(run func(string) error) *TxnKV_RemoveWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: key, value
func (_m *TxnKV) Save(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxnKV_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type TxnKV_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//  - key string
//  - value string
func (_e *TxnKV_Expecter) Save(key interface{}, value interface{}) *TxnKV_Save_Call {
	return &TxnKV_Save_Call{Call: _e.mock.On("Save", key, value)}
}

func (_c *TxnKV_Save_Call) Run(run func(key string, value string)) *TxnKV_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *TxnKV_Save_Call) Return(_a0 error) *TxnKV_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TxnKV_Save_Call) RunAndReturn(run func(string, string) error) *TxnKV_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxnKV creates a new instance of TxnKV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxnKV(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxnKV {
	mock := &TxnKV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
