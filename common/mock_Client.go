/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.23.1. DO NOT EDIT.

package common

import (
	mock "github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// NewGrpcConn provides a mock function with given fields:
func (_m *MockClient) NewGrpcConn() (grpc.ClientConnInterface, Closer, error) {
	ret := _m.Called()

	var r0 grpc.ClientConnInterface
	var r1 Closer
	var r2 error
	if rf, ok := ret.Get(0).(func() (grpc.ClientConnInterface, Closer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() grpc.ClientConnInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.ClientConnInterface)
		}
	}

	if rf, ok := ret.Get(1).(func() Closer); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(Closer)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewMockClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClient(t mockConstructorTestingTNewMockClient) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}