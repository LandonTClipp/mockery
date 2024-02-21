// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery

package test

import (
	"sync"

	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// Ensure, that KeyManagerMoq does implement test.KeyManager.
// If this is not the case, regenerate this file with moq.
var _ test.KeyManager = &KeyManagerMoq{}

// KeyManagerMoq is a mock implementation of test.KeyManager.
//
//	func TestSomethingThatUsesKeyManager(t *testing.T) {
//
//		// make and configure a mocked test.KeyManager
//		mockedKeyManager := &KeyManagerMoq{
//			GetKeyFunc: func(s string, v uint16) ([]byte, *test.Err) {
//				panic("mock out the GetKey method")
//			},
//		}
//
//		// use mockedKeyManager in code that requires test.KeyManager
//		// and then make assertions.
//
//	}
type KeyManagerMoq struct {
	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(s string, v uint16) ([]byte, *test.Err)

	// calls tracks calls to the methods.
	calls struct {
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// S is the s argument value.
			S string
			// V is the v argument value.
			V uint16
		}
	}
	lockGetKey sync.RWMutex
}

// GetKey calls GetKeyFunc.
func (mock *KeyManagerMoq) GetKey(s string, v uint16) ([]byte, *test.Err) {
	if mock.GetKeyFunc == nil {
		panic("KeyManagerMoq.GetKeyFunc: method is nil but KeyManager.GetKey was just called")
	}
	callInfo := struct {
		S string
		V uint16
	}{
		S: s,
		V: v,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc(s, v)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//
//	len(mockedKeyManager.GetKeyCalls())
func (mock *KeyManagerMoq) GetKeyCalls() []struct {
	S string
	V uint16
} {
	var calls []struct {
		S string
		V uint16
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// ResetGetKeyCalls reset all the calls that were made to GetKey.
func (mock *KeyManagerMoq) ResetGetKeyCalls() {
	mock.lockGetKey.Lock()
	mock.calls.GetKey = nil
	mock.lockGetKey.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *KeyManagerMoq) ResetCalls() {
	mock.lockGetKey.Lock()
	mock.calls.GetKey = nil
	mock.lockGetKey.Unlock()
}