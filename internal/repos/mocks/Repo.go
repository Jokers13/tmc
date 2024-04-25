// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/wot-oss/tmc/internal/model"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repo) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, id
func (_m *Repo) Fetch(ctx context.Context, id string) (string, []byte, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 string
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, []byte, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Index provides a mock function with given fields: ctx, updatedIds
func (_m *Repo) Index(ctx context.Context, updatedIds ...string) error {
	_va := make([]interface{}, len(updatedIds))
	for _i := range updatedIds {
		_va[_i] = updatedIds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Index")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, updatedIds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, search
func (_m *Repo) List(ctx context.Context, search *model.SearchParams) (model.SearchResult, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 model.SearchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) (model.SearchResult, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.SearchParams) model.SearchResult); ok {
		r0 = rf(ctx, search)
	} else {
		r0 = ret.Get(0).(model.SearchResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.SearchParams) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCompletions provides a mock function with given fields: ctx, kind, toComplete
func (_m *Repo) ListCompletions(ctx context.Context, kind string, toComplete string) ([]string, error) {
	ret := _m.Called(ctx, kind, toComplete)

	if len(ret) == 0 {
		panic("no return value specified for ListCompletions")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]string, error)); ok {
		return rf(ctx, kind, toComplete)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, kind, toComplete)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, kind, toComplete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Push provides a mock function with given fields: ctx, id, raw
func (_m *Repo) Push(ctx context.Context, id model.TMID, raw []byte) error {
	ret := _m.Called(ctx, id, raw)

	if len(ret) == 0 {
		panic("no return value specified for Push")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.TMID, []byte) error); ok {
		r0 = rf(ctx, id, raw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Spec provides a mock function with given fields:
func (_m *Repo) Spec() model.RepoSpec {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Spec")
	}

	var r0 model.RepoSpec
	if rf, ok := ret.Get(0).(func() model.RepoSpec); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.RepoSpec)
	}

	return r0
}

// Versions provides a mock function with given fields: ctx, name
func (_m *Repo) Versions(ctx context.Context, name string) ([]model.FoundVersion, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Versions")
	}

	var r0 []model.FoundVersion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.FoundVersion, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.FoundVersion); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.FoundVersion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}