// Code generated by mockery v1.0.0. DO NOT EDIT.
package automock

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// nsLister is an autogenerated mock type for the nsLister type
type nsLister struct {
	mock.Mock
}

// List provides a mock function with given fields:
func (_m *nsLister) List() ([]v1.Namespace, error) {
	ret := _m.Called()

	var r0 []v1.Namespace
	if rf, ok := ret.Get(0).(func() []v1.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
