// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"sync"

	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// Ensure, that VariadicReturnFuncMoq does implement test.VariadicReturnFunc.
// If this is not the case, regenerate this file with moq.
var _ test.VariadicReturnFunc = &VariadicReturnFuncMoq{}

// VariadicReturnFuncMoq is a mock implementation of test.VariadicReturnFunc.
//
//	func TestSomethingThatUsesVariadicReturnFunc(t *testing.T) {
//
//		// make and configure a mocked test.VariadicReturnFunc
//		mockedVariadicReturnFunc := &VariadicReturnFuncMoq{
//			SampleMethodFunc: func(str string) func(str string, arr []int, a ...interface{}) {
//				panic("mock out the SampleMethod method")
//			},
//		}
//
//		// use mockedVariadicReturnFunc in code that requires test.VariadicReturnFunc
//		// and then make assertions.
//
//	}
type VariadicReturnFuncMoq struct {
	// SampleMethodFunc mocks the SampleMethod method.
	SampleMethodFunc func(str string) func(str string, arr []int, a ...interface{})

	// calls tracks calls to the methods.
	calls struct {
		// SampleMethod holds details about calls to the SampleMethod method.
		SampleMethod []struct {
			// Str is the str argument value.
			Str string
		}
	}
	lockSampleMethod sync.RWMutex
}

// SampleMethod calls SampleMethodFunc.
func (mock *VariadicReturnFuncMoq) SampleMethod(str string) func(str string, arr []int, a ...interface{}) {
	if mock.SampleMethodFunc == nil {
		panic("VariadicReturnFuncMoq.SampleMethodFunc: method is nil but VariadicReturnFunc.SampleMethod was just called")
	}
	callInfo := struct {
		Str string
	}{
		Str: str,
	}
	mock.lockSampleMethod.Lock()
	mock.calls.SampleMethod = append(mock.calls.SampleMethod, callInfo)
	mock.lockSampleMethod.Unlock()
	return mock.SampleMethodFunc(str)
}

// SampleMethodCalls gets all the calls that were made to SampleMethod.
// Check the length with:
//
//	len(mockedVariadicReturnFunc.SampleMethodCalls())
func (mock *VariadicReturnFuncMoq) SampleMethodCalls() []struct {
	Str string
} {
	var calls []struct {
		Str string
	}
	mock.lockSampleMethod.RLock()
	calls = mock.calls.SampleMethod
	mock.lockSampleMethod.RUnlock()
	return calls
}

// ResetSampleMethodCalls reset all the calls that were made to SampleMethod.
func (mock *VariadicReturnFuncMoq) ResetSampleMethodCalls() {
	mock.lockSampleMethod.Lock()
	mock.calls.SampleMethod = nil
	mock.lockSampleMethod.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *VariadicReturnFuncMoq) ResetCalls() {
	mock.lockSampleMethod.Lock()
	mock.calls.SampleMethod = nil
	mock.lockSampleMethod.Unlock()
}