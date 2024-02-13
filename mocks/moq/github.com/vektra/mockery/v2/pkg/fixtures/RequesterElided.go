// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"sync"
)

// RequesterElidedMock is a mock implementation of RequesterElided.
//
//	func TestSomethingThatUsesRequesterElided(t *testing.T) {
//
//		// make and configure a mocked RequesterElided
//		mockedRequesterElided := &RequesterElidedMock{
//			GetFunc: func(path string, url string) error {
//				panic("mock out the Get method")
//			},
//		}
//
//		// use mockedRequesterElided in code that requires RequesterElided
//		// and then make assertions.
//
//	}
type RequesterElidedMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(path string, url string) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Path is the path argument value.
			Path string
			// URL is the url argument value.
			URL string
		}
	}
	lockGet sync.RWMutex
}

// Get calls GetFunc.
func (mock *RequesterElidedMock) Get(path string, url string) error {
	if mock.GetFunc == nil {
		panic("RequesterElidedMock.GetFunc: method is nil but RequesterElided.Get was just called")
	}
	callInfo := struct {
		Path string
		URL  string
	}{
		Path: path,
		URL:  url,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(path, url)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRequesterElided.GetCalls())
func (mock *RequesterElidedMock) GetCalls() []struct {
	Path string
	URL  string
} {
	var calls []struct {
		Path string
		URL  string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// ResetGetCalls reset all the calls that were made to Get.
func (mock *RequesterElidedMock) ResetGetCalls() {
	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *RequesterElidedMock) ResetCalls() {
	mock.lockGet.Lock()
	mock.calls.Get = nil
	mock.lockGet.Unlock()
}