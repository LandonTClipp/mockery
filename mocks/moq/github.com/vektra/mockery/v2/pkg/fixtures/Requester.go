// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"sync"
)

// RequesterMock is a mock implementation of Requester.
//
//	func TestSomethingThatUsesRequester(t *testing.T) {
//
//		// make and configure a mocked Requester
//		mockedRequester := &RequesterMock{
//			GetFunc: func(path string) (string, error) {
//				panic("mock out the Get method")
//			},
//		}
//
//		// use mockedRequester in code that requires Requester
//		// and then make assertions.
//
//	}
type RequesterMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(path string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Path is the path argument value.
			Path string
		}
	}
	lockGet sync.RWMutex
}

// Get calls GetFunc.
func (mock *RequesterMock) Get(path string) (string, error) {
	if mock.GetFunc == nil {
		panic("RequesterMock.GetFunc: method is nil but Requester.Get was just called")
	}
	callInfo := struct {
		Path string
	}{
		Path: path,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(path)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRequester.GetCalls())
func (mock *RequesterMock) GetCalls() []struct {
	Path string
} {
	var calls []struct {
		Path string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// ResetGetCalls reset all the calls that were made to Get.
func (mock *RequesterMock) ResetGetCalls() {
	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *RequesterMock) ResetCalls() {
	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}