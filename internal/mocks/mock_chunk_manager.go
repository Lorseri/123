// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mmap "golang.org/x/exp/mmap"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/milvus-io/milvus/internal/storage"
)

// ChunkManager is an autogenerated mock type for the ChunkManager type
type ChunkManager struct {
	mock.Mock
}

type ChunkManager_Expecter struct {
	mock *mock.Mock
}

func (_m *ChunkManager) EXPECT() *ChunkManager_Expecter {
	return &ChunkManager_Expecter{mock: &_m.Mock}
}

// Exist provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Exist(ctx context.Context, filePath string) (bool, error) {
	ret := _m.Called(ctx, filePath)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, filePath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Exist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exist'
type ChunkManager_Exist_Call struct {
	*mock.Call
}

// Exist is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Exist(ctx interface{}, filePath interface{}) *ChunkManager_Exist_Call {
	return &ChunkManager_Exist_Call{Call: _e.mock.On("Exist", ctx, filePath)}
}

func (_c *ChunkManager_Exist_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Exist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Exist_Call) Return(_a0 bool, _a1 error) *ChunkManager_Exist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Exist_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *ChunkManager_Exist_Call {
	_c.Call.Return(run)
	return _c
}

// Mmap provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Mmap(ctx context.Context, filePath string) (*mmap.ReaderAt, error) {
	ret := _m.Called(ctx, filePath)

	var r0 *mmap.ReaderAt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*mmap.ReaderAt, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *mmap.ReaderAt); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mmap.ReaderAt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Mmap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mmap'
type ChunkManager_Mmap_Call struct {
	*mock.Call
}

// Mmap is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Mmap(ctx interface{}, filePath interface{}) *ChunkManager_Mmap_Call {
	return &ChunkManager_Mmap_Call{Call: _e.mock.On("Mmap", ctx, filePath)}
}

func (_c *ChunkManager_Mmap_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Mmap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Mmap_Call) Return(_a0 *mmap.ReaderAt, _a1 error) *ChunkManager_Mmap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Mmap_Call) RunAndReturn(run func(context.Context, string) (*mmap.ReaderAt, error)) *ChunkManager_Mmap_Call {
	_c.Call.Return(run)
	return _c
}

// MultiRead provides a mock function with given fields: ctx, filePaths
func (_m *ChunkManager) MultiRead(ctx context.Context, filePaths []string) ([][]byte, error) {
	ret := _m.Called(ctx, filePaths)

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([][]byte, error)); ok {
		return rf(ctx, filePaths)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) [][]byte); ok {
		r0 = rf(ctx, filePaths)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, filePaths)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_MultiRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRead'
type ChunkManager_MultiRead_Call struct {
	*mock.Call
}

// MultiRead is a helper method to define mock.On call
//   - ctx context.Context
//   - filePaths []string
func (_e *ChunkManager_Expecter) MultiRead(ctx interface{}, filePaths interface{}) *ChunkManager_MultiRead_Call {
	return &ChunkManager_MultiRead_Call{Call: _e.mock.On("MultiRead", ctx, filePaths)}
}

func (_c *ChunkManager_MultiRead_Call) Run(run func(ctx context.Context, filePaths []string)) *ChunkManager_MultiRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *ChunkManager_MultiRead_Call) Return(_a0 [][]byte, _a1 error) *ChunkManager_MultiRead_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_MultiRead_Call) RunAndReturn(run func(context.Context, []string) ([][]byte, error)) *ChunkManager_MultiRead_Call {
	_c.Call.Return(run)
	return _c
}

// MultiRemove provides a mock function with given fields: ctx, filePaths
func (_m *ChunkManager) MultiRemove(ctx context.Context, filePaths []string) error {
	ret := _m.Called(ctx, filePaths)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, filePaths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_MultiRemove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiRemove'
type ChunkManager_MultiRemove_Call struct {
	*mock.Call
}

// MultiRemove is a helper method to define mock.On call
//   - ctx context.Context
//   - filePaths []string
func (_e *ChunkManager_Expecter) MultiRemove(ctx interface{}, filePaths interface{}) *ChunkManager_MultiRemove_Call {
	return &ChunkManager_MultiRemove_Call{Call: _e.mock.On("MultiRemove", ctx, filePaths)}
}

func (_c *ChunkManager_MultiRemove_Call) Run(run func(ctx context.Context, filePaths []string)) *ChunkManager_MultiRemove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *ChunkManager_MultiRemove_Call) Return(_a0 error) *ChunkManager_MultiRemove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_MultiRemove_Call) RunAndReturn(run func(context.Context, []string) error) *ChunkManager_MultiRemove_Call {
	_c.Call.Return(run)
	return _c
}

// MultiWrite provides a mock function with given fields: ctx, contents
func (_m *ChunkManager) MultiWrite(ctx context.Context, contents map[string][]byte) error {
	ret := _m.Called(ctx, contents)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string][]byte) error); ok {
		r0 = rf(ctx, contents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_MultiWrite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiWrite'
type ChunkManager_MultiWrite_Call struct {
	*mock.Call
}

// MultiWrite is a helper method to define mock.On call
//   - ctx context.Context
//   - contents map[string][]byte
func (_e *ChunkManager_Expecter) MultiWrite(ctx interface{}, contents interface{}) *ChunkManager_MultiWrite_Call {
	return &ChunkManager_MultiWrite_Call{Call: _e.mock.On("MultiWrite", ctx, contents)}
}

func (_c *ChunkManager_MultiWrite_Call) Run(run func(ctx context.Context, contents map[string][]byte)) *ChunkManager_MultiWrite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string][]byte))
	})
	return _c
}

func (_c *ChunkManager_MultiWrite_Call) Return(_a0 error) *ChunkManager_MultiWrite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_MultiWrite_Call) RunAndReturn(run func(context.Context, map[string][]byte) error) *ChunkManager_MultiWrite_Call {
	_c.Call.Return(run)
	return _c
}

// Path provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Path(ctx context.Context, filePath string) (string, error) {
	ret := _m.Called(ctx, filePath)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'
type ChunkManager_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Path(ctx interface{}, filePath interface{}) *ChunkManager_Path_Call {
	return &ChunkManager_Path_Call{Call: _e.mock.On("Path", ctx, filePath)}
}

func (_c *ChunkManager_Path_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Path_Call) Return(_a0 string, _a1 error) *ChunkManager_Path_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Path_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ChunkManager_Path_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Read(ctx context.Context, filePath string) ([]byte, error) {
	ret := _m.Called(ctx, filePath)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type ChunkManager_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Read(ctx interface{}, filePath interface{}) *ChunkManager_Read_Call {
	return &ChunkManager_Read_Call{Call: _e.mock.On("Read", ctx, filePath)}
}

func (_c *ChunkManager_Read_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Read_Call) Return(_a0 []byte, _a1 error) *ChunkManager_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Read_Call) RunAndReturn(run func(context.Context, string) ([]byte, error)) *ChunkManager_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadAt provides a mock function with given fields: ctx, filePath, off, length
func (_m *ChunkManager) ReadAt(ctx context.Context, filePath string, off int64, length int64) ([]byte, error) {
	ret := _m.Called(ctx, filePath, off, length)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) ([]byte, error)); ok {
		return rf(ctx, filePath, off, length)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []byte); ok {
		r0 = rf(ctx, filePath, off, length)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, filePath, off, length)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_ReadAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadAt'
type ChunkManager_ReadAt_Call struct {
	*mock.Call
}

// ReadAt is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
//   - off int64
//   - length int64
func (_e *ChunkManager_Expecter) ReadAt(ctx interface{}, filePath interface{}, off interface{}, length interface{}) *ChunkManager_ReadAt_Call {
	return &ChunkManager_ReadAt_Call{Call: _e.mock.On("ReadAt", ctx, filePath, off, length)}
}

func (_c *ChunkManager_ReadAt_Call) Run(run func(ctx context.Context, filePath string, off int64, length int64)) *ChunkManager_ReadAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64), args[3].(int64))
	})
	return _c
}

func (_c *ChunkManager_ReadAt_Call) Return(p []byte, err error) *ChunkManager_ReadAt_Call {
	_c.Call.Return(p, err)
	return _c
}

func (_c *ChunkManager_ReadAt_Call) RunAndReturn(run func(context.Context, string, int64, int64) ([]byte, error)) *ChunkManager_ReadAt_Call {
	_c.Call.Return(run)
	return _c
}

// Reader provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Reader(ctx context.Context, filePath string) (storage.FileReader, error) {
	ret := _m.Called(ctx, filePath)

	var r0 storage.FileReader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (storage.FileReader, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) storage.FileReader); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.FileReader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Reader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reader'
type ChunkManager_Reader_Call struct {
	*mock.Call
}

// Reader is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Reader(ctx interface{}, filePath interface{}) *ChunkManager_Reader_Call {
	return &ChunkManager_Reader_Call{Call: _e.mock.On("Reader", ctx, filePath)}
}

func (_c *ChunkManager_Reader_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Reader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Reader_Call) Return(_a0 storage.FileReader, _a1 error) *ChunkManager_Reader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Reader_Call) RunAndReturn(run func(context.Context, string) (storage.FileReader, error)) *ChunkManager_Reader_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Remove(ctx context.Context, filePath string) error {
	ret := _m.Called(ctx, filePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type ChunkManager_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Remove(ctx interface{}, filePath interface{}) *ChunkManager_Remove_Call {
	return &ChunkManager_Remove_Call{Call: _e.mock.On("Remove", ctx, filePath)}
}

func (_c *ChunkManager_Remove_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Remove_Call) Return(_a0 error) *ChunkManager_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_Remove_Call) RunAndReturn(run func(context.Context, string) error) *ChunkManager_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveWithPrefix provides a mock function with given fields: ctx, prefix
func (_m *ChunkManager) RemoveWithPrefix(ctx context.Context, prefix string) error {
	ret := _m.Called(ctx, prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_RemoveWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveWithPrefix'
type ChunkManager_RemoveWithPrefix_Call struct {
	*mock.Call
}

// RemoveWithPrefix is a helper method to define mock.On call
//   - ctx context.Context
//   - prefix string
func (_e *ChunkManager_Expecter) RemoveWithPrefix(ctx interface{}, prefix interface{}) *ChunkManager_RemoveWithPrefix_Call {
	return &ChunkManager_RemoveWithPrefix_Call{Call: _e.mock.On("RemoveWithPrefix", ctx, prefix)}
}

func (_c *ChunkManager_RemoveWithPrefix_Call) Run(run func(ctx context.Context, prefix string)) *ChunkManager_RemoveWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_RemoveWithPrefix_Call) Return(_a0 error) *ChunkManager_RemoveWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_RemoveWithPrefix_Call) RunAndReturn(run func(context.Context, string) error) *ChunkManager_RemoveWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// RootPath provides a mock function with given fields:
func (_m *ChunkManager) RootPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ChunkManager_RootPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RootPath'
type ChunkManager_RootPath_Call struct {
	*mock.Call
}

// RootPath is a helper method to define mock.On call
func (_e *ChunkManager_Expecter) RootPath() *ChunkManager_RootPath_Call {
	return &ChunkManager_RootPath_Call{Call: _e.mock.On("RootPath")}
}

func (_c *ChunkManager_RootPath_Call) Run(run func()) *ChunkManager_RootPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ChunkManager_RootPath_Call) Return(_a0 string) *ChunkManager_RootPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_RootPath_Call) RunAndReturn(run func() string) *ChunkManager_RootPath_Call {
	_c.Call.Return(run)
	return _c
}

// Size provides a mock function with given fields: ctx, filePath
func (_m *ChunkManager) Size(ctx context.Context, filePath string) (int64, error) {
	ret := _m.Called(ctx, filePath)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, filePath)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkManager_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type ChunkManager_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *ChunkManager_Expecter) Size(ctx interface{}, filePath interface{}) *ChunkManager_Size_Call {
	return &ChunkManager_Size_Call{Call: _e.mock.On("Size", ctx, filePath)}
}

func (_c *ChunkManager_Size_Call) Run(run func(ctx context.Context, filePath string)) *ChunkManager_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ChunkManager_Size_Call) Return(_a0 int64, _a1 error) *ChunkManager_Size_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChunkManager_Size_Call) RunAndReturn(run func(context.Context, string) (int64, error)) *ChunkManager_Size_Call {
	_c.Call.Return(run)
	return _c
}

// WalkWithPrefix provides a mock function with given fields: ctx, prefix, recursive, cb
func (_m *ChunkManager) WalkWithPrefix(ctx context.Context, prefix string, recursive bool, cb func(*storage.ChunkObjectInfo) error) error {
	ret := _m.Called(ctx, prefix, recursive, cb)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, func(*storage.ChunkObjectInfo) error) error); ok {
		r0 = rf(ctx, prefix, recursive, cb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_WalkWithPrefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WalkWithPrefix'
type ChunkManager_WalkWithPrefix_Call struct {
	*mock.Call
}

// WalkWithPrefix is a helper method to define mock.On call
//   - ctx context.Context
//   - prefix string
//   - recursive bool
//   - cb func(*storage.ChunkObjectInfo) error
func (_e *ChunkManager_Expecter) WalkWithPrefix(ctx interface{}, prefix interface{}, recursive interface{}, cb interface{}) *ChunkManager_WalkWithPrefix_Call {
	return &ChunkManager_WalkWithPrefix_Call{Call: _e.mock.On("WalkWithPrefix", ctx, prefix, recursive, cb)}
}

func (_c *ChunkManager_WalkWithPrefix_Call) Run(run func(ctx context.Context, prefix string, recursive bool, cb func(*storage.ChunkObjectInfo) error)) *ChunkManager_WalkWithPrefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool), args[3].(func(*storage.ChunkObjectInfo) error))
	})
	return _c
}

func (_c *ChunkManager_WalkWithPrefix_Call) Return(_a0 error) *ChunkManager_WalkWithPrefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_WalkWithPrefix_Call) RunAndReturn(run func(context.Context, string, bool, func(*storage.ChunkObjectInfo) error) error) *ChunkManager_WalkWithPrefix_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: ctx, filePath, content
func (_m *ChunkManager) Write(ctx context.Context, filePath string, content []byte) error {
	ret := _m.Called(ctx, filePath, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, filePath, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChunkManager_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type ChunkManager_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
//   - content []byte
func (_e *ChunkManager_Expecter) Write(ctx interface{}, filePath interface{}, content interface{}) *ChunkManager_Write_Call {
	return &ChunkManager_Write_Call{Call: _e.mock.On("Write", ctx, filePath, content)}
}

func (_c *ChunkManager_Write_Call) Run(run func(ctx context.Context, filePath string, content []byte)) *ChunkManager_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *ChunkManager_Write_Call) Return(_a0 error) *ChunkManager_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChunkManager_Write_Call) RunAndReturn(run func(context.Context, string, []byte) error) *ChunkManager_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewChunkManager creates a new instance of ChunkManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChunkManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChunkManager {
	mock := &ChunkManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
