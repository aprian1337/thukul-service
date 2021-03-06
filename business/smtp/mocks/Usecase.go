// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	smtp "aprian1337/thukul-service/business/smtp"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// SendMailSMTP provides a mock function with given fields: ctx, domain
func (_m *Usecase) SendMailSMTP(ctx context.Context, domain smtp.Domain) error {
	ret := _m.Called(ctx, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, smtp.Domain) error); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
