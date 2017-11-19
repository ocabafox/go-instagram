// Code generated by mockery v1.0.0
package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import protos "github.com/hieven/go-instagram/src/protos"

// Timeline is an autogenerated mock type for the Timeline type
type Timeline struct {
	mock.Mock
}

// Feed provides a mock function with given fields: _a0, _a1
func (_m *Timeline) Feed(_a0 context.Context, _a1 *protos.TimelineFeedRequest) (*protos.TimelineFeedResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.TimelineFeedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protos.TimelineFeedRequest) *protos.TimelineFeedResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.TimelineFeedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.TimelineFeedRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
