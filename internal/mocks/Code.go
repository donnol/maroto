// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	internal "github.com/donnol/maroto/internal"
	mock "github.com/stretchr/testify/mock"

	props "github.com/donnol/maroto/pkg/props"
)

// Code is an autogenerated mock type for the Code type
type Code struct {
	mock.Mock
}

// AddBar provides a mock function with given fields: code, cell, prop
func (_m *Code) AddBar(code string, cell internal.Cell, prop props.Barcode) error {
	ret := _m.Called(code, cell, prop)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, internal.Cell, props.Barcode) error); ok {
		r0 = rf(code, cell, prop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDataMatrix provides a mock function with given fields: code, cell, prop
func (_m *Code) AddDataMatrix(code string, cell internal.Cell, prop props.Rect) {
	_m.Called(code, cell, prop)
}

// AddQr provides a mock function with given fields: code, cell, prop
func (_m *Code) AddQr(code string, cell internal.Cell, prop props.Rect) {
	_m.Called(code, cell, prop)
}

type mockConstructorTestingTNewCode interface {
	mock.TestingT
	Cleanup(func())
}

// NewCode creates a new instance of Code. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCode(t mockConstructorTestingTNewCode) *Code {
	mock := &Code{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
