// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	auth "github.com/absmach/supermq/auth"

	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// CheckScope provides a mock function with given fields: ctx, userID, patID, optionalDomainID, entityType, operation, entityID
func (_m *Cache) CheckScope(ctx context.Context, userID string, patID string, optionalDomainID string, entityType auth.EntityType, operation auth.Operation, entityID string) bool {
	ret := _m.Called(ctx, userID, patID, optionalDomainID, entityType, operation, entityID)

	if len(ret) == 0 {
		panic("no return value specified for CheckScope")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, auth.EntityType, auth.Operation, string) bool); ok {
		r0 = rf(ctx, userID, patID, optionalDomainID, entityType, operation, entityID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Remove provides a mock function with given fields: ctx, userID, scopesID
func (_m *Cache) Remove(ctx context.Context, userID string, scopesID []string) error {
	ret := _m.Called(ctx, userID, scopesID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, userID, scopesID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllScope provides a mock function with given fields: ctx, userID, patID
func (_m *Cache) RemoveAllScope(ctx context.Context, userID string, patID string) error {
	ret := _m.Called(ctx, userID, patID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveAllScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, patID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserAllScope provides a mock function with given fields: ctx, userID
func (_m *Cache) RemoveUserAllScope(ctx context.Context, userID string) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveUserAllScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, userID, scopes
func (_m *Cache) Save(ctx context.Context, userID string, scopes []auth.Scope) error {
	ret := _m.Called(ctx, userID, scopes)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []auth.Scope) error); ok {
		r0 = rf(ctx, userID, scopes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
