// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"sync"
)

// SiblingMock is a mock implementation of Sibling.
//
//	func TestSomethingThatUsesSibling(t *testing.T) {
//
//		// make and configure a mocked Sibling
//		mockedSibling := &SiblingMock{
//			DoSomethingFunc: func()  {
//				panic("mock out the DoSomething method")
//			},
//		}
//
//		// use mockedSibling in code that requires Sibling
//		// and then make assertions.
//
//	}
type SiblingMock struct {
	// DoSomethingFunc mocks the DoSomething method.
	DoSomethingFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// DoSomething holds details about calls to the DoSomething method.
		DoSomething []struct {
		}
	}
	lockDoSomething sync.RWMutex
}

// DoSomething calls DoSomethingFunc.
func (mock *SiblingMock) DoSomething() {
	if mock.DoSomethingFunc == nil {
		panic("SiblingMock.DoSomethingFunc: method is nil but Sibling.DoSomething was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDoSomething.Lock()
	mock.calls.DoSomething = append(mock.calls.DoSomething, callInfo)
	mock.lockDoSomething.Unlock()
	mock.DoSomethingFunc()
}

// DoSomethingCalls gets all the calls that were made to DoSomething.
// Check the length with:
//
//	len(mockedSibling.DoSomethingCalls())
func (mock *SiblingMock) DoSomethingCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDoSomething.RLock()
	calls = mock.calls.DoSomething
	mock.lockDoSomething.RUnlock()
	return calls
}

// ResetDoSomethingCalls reset all the calls that were made to DoSomething.
func (mock *SiblingMock) ResetDoSomethingCalls() {
	mock.lockDoSomething.Lock()
	mock.calls.DoSomething = nil
	mock.lockDoSomething.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *SiblingMock) ResetCalls() {
	mock.lockDoSomething.Lock()
	mock.calls.DoSomething = nil
	mock.lockDoSomething.Unlock()
}