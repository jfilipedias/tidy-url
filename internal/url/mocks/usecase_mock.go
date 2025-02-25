// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	url "github.com/jfilipedias/tidy-url/internal/url"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

type UseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *UseCase) EXPECT() *UseCase_Expecter {
	return &UseCase_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, originalURL, userID
func (_m *UseCase) Create(ctx context.Context, originalURL string, userID uuid.UUID) error {
	ret := _m.Called(ctx, originalURL, userID)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) error); ok {
		r0 = rf(ctx, originalURL, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UseCase_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UseCase_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - originalURL string
//   - userID uuid.UUID
func (_e *UseCase_Expecter) Create(ctx interface{}, originalURL interface{}, userID interface{}) *UseCase_Create_Call {
	return &UseCase_Create_Call{Call: _e.mock.On("Create", ctx, originalURL, userID)}
}

func (_c *UseCase_Create_Call) Run(run func(ctx context.Context, originalURL string, userID uuid.UUID)) *UseCase_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *UseCase_Create_Call) Return(_a0 error) *UseCase_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCase_Create_Call) RunAndReturn(run func(context.Context, string, uuid.UUID) error) *UseCase_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, hash
func (_m *UseCase) Get(ctx context.Context, hash string) (*url.URL, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*url.URL, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *url.URL); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseCase_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type UseCase_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *UseCase_Expecter) Get(ctx interface{}, hash interface{}) *UseCase_Get_Call {
	return &UseCase_Get_Call{Call: _e.mock.On("Get", ctx, hash)}
}

func (_c *UseCase_Get_Call) Run(run func(ctx context.Context, hash string)) *UseCase_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UseCase_Get_Call) Return(_a0 *url.URL, _a1 error) *UseCase_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UseCase_Get_Call) RunAndReturn(run func(context.Context, string) (*url.URL, error)) *UseCase_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
