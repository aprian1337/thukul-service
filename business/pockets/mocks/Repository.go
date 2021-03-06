// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	pockets "aprian1337/thukul-service/business/pockets"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// PocketsCreate provides a mock function with given fields: ctx, domain
func (_m *Repository) PocketsCreate(ctx context.Context, domain pockets.Domain) (pockets.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 pockets.Domain
	if rf, ok := ret.Get(0).(func(context.Context, pockets.Domain) pockets.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(pockets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pockets.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketsDelete provides a mock function with given fields: ctx, userId, pocketId
func (_m *Repository) PocketsDelete(ctx context.Context, userId int, pocketId int) (int64, error) {
	ret := _m.Called(ctx, userId, pocketId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int, int) int64); ok {
		r0 = rf(ctx, userId, pocketId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, userId, pocketId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketsGetById provides a mock function with given fields: ctx, userId, pocketId
func (_m *Repository) PocketsGetById(ctx context.Context, userId int, pocketId int) (pockets.Domain, error) {
	ret := _m.Called(ctx, userId, pocketId)

	var r0 pockets.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) pockets.Domain); ok {
		r0 = rf(ctx, userId, pocketId)
	} else {
		r0 = ret.Get(0).(pockets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, userId, pocketId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketsGetList provides a mock function with given fields: ctx, id
func (_m *Repository) PocketsGetList(ctx context.Context, id int) ([]pockets.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 []pockets.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []pockets.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pockets.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PocketsUpdate provides a mock function with given fields: ctx, domain, userId, pocketId
func (_m *Repository) PocketsUpdate(ctx context.Context, domain pockets.Domain, userId int, pocketId int) (pockets.Domain, error) {
	ret := _m.Called(ctx, domain, userId, pocketId)

	var r0 pockets.Domain
	if rf, ok := ret.Get(0).(func(context.Context, pockets.Domain, int, int) pockets.Domain); ok {
		r0 = rf(ctx, domain, userId, pocketId)
	} else {
		r0 = ret.Get(0).(pockets.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pockets.Domain, int, int) error); ok {
		r1 = rf(ctx, domain, userId, pocketId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
