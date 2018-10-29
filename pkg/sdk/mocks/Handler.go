// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import sdk "github.com/operator-framework/operator-sdk/pkg/sdk"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, event
func (_m *Handler) Handle(ctx context.Context, event sdk.Event) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sdk.Event) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
