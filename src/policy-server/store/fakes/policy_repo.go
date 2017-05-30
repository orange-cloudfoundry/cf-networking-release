// This file was generated by counterfeiter
package fakes

import (
	"policy-server/store"
	"sync"
)

type PolicyRepo struct {
	CreateStub        func(store.Transaction, int, int) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(store.Transaction, int, int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	Delete2Stub        func(store.Transaction, string, string, int, string) error
	delete2Mutex       sync.RWMutex
	delete2ArgsForCall []struct {
		arg1 store.Transaction
		arg2 string
		arg3 string
		arg4 int
		arg5 string
	}
	delete2Returns struct {
		result1 error
	}
	delete2ReturnsOnCall map[int]struct {
		result1 error
	}
	CountWhereGroupIDStub        func(store.Transaction, int) (int, error)
	countWhereGroupIDMutex       sync.RWMutex
	countWhereGroupIDArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
	}
	countWhereGroupIDReturns struct {
		result1 int
		result2 error
	}
	countWhereGroupIDReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CountWhereDestinationIDStub        func(store.Transaction, int) (int, error)
	countWhereDestinationIDMutex       sync.RWMutex
	countWhereDestinationIDArgsForCall []struct {
		arg1 store.Transaction
		arg2 int
	}
	countWhereDestinationIDReturns struct {
		result1 int
		result2 error
	}
	countWhereDestinationIDReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyRepo) Create(arg1 store.Transaction, arg2 int, arg3 int) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *PolicyRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *PolicyRepo) CreateArgsForCall(i int) (store.Transaction, int, int) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3
}

func (fake *PolicyRepo) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) Delete(arg1 store.Transaction, arg2 int, arg3 int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *PolicyRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *PolicyRepo) DeleteArgsForCall(i int) (store.Transaction, int, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2, fake.deleteArgsForCall[i].arg3
}

func (fake *PolicyRepo) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) Delete2(arg1 store.Transaction, arg2 string, arg3 string, arg4 int, arg5 string) error {
	fake.delete2Mutex.Lock()
	ret, specificReturn := fake.delete2ReturnsOnCall[len(fake.delete2ArgsForCall)]
	fake.delete2ArgsForCall = append(fake.delete2ArgsForCall, struct {
		arg1 store.Transaction
		arg2 string
		arg3 string
		arg4 int
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Delete2", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.delete2Mutex.Unlock()
	if fake.Delete2Stub != nil {
		return fake.Delete2Stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.delete2Returns.result1
}

func (fake *PolicyRepo) Delete2CallCount() int {
	fake.delete2Mutex.RLock()
	defer fake.delete2Mutex.RUnlock()
	return len(fake.delete2ArgsForCall)
}

func (fake *PolicyRepo) Delete2ArgsForCall(i int) (store.Transaction, string, string, int, string) {
	fake.delete2Mutex.RLock()
	defer fake.delete2Mutex.RUnlock()
	return fake.delete2ArgsForCall[i].arg1, fake.delete2ArgsForCall[i].arg2, fake.delete2ArgsForCall[i].arg3, fake.delete2ArgsForCall[i].arg4, fake.delete2ArgsForCall[i].arg5
}

func (fake *PolicyRepo) Delete2Returns(result1 error) {
	fake.Delete2Stub = nil
	fake.delete2Returns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) Delete2ReturnsOnCall(i int, result1 error) {
	fake.Delete2Stub = nil
	if fake.delete2ReturnsOnCall == nil {
		fake.delete2ReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.delete2ReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyRepo) CountWhereGroupID(arg1 store.Transaction, arg2 int) (int, error) {
	fake.countWhereGroupIDMutex.Lock()
	ret, specificReturn := fake.countWhereGroupIDReturnsOnCall[len(fake.countWhereGroupIDArgsForCall)]
	fake.countWhereGroupIDArgsForCall = append(fake.countWhereGroupIDArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CountWhereGroupID", []interface{}{arg1, arg2})
	fake.countWhereGroupIDMutex.Unlock()
	if fake.CountWhereGroupIDStub != nil {
		return fake.CountWhereGroupIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countWhereGroupIDReturns.result1, fake.countWhereGroupIDReturns.result2
}

func (fake *PolicyRepo) CountWhereGroupIDCallCount() int {
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	return len(fake.countWhereGroupIDArgsForCall)
}

func (fake *PolicyRepo) CountWhereGroupIDArgsForCall(i int) (store.Transaction, int) {
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	return fake.countWhereGroupIDArgsForCall[i].arg1, fake.countWhereGroupIDArgsForCall[i].arg2
}

func (fake *PolicyRepo) CountWhereGroupIDReturns(result1 int, result2 error) {
	fake.CountWhereGroupIDStub = nil
	fake.countWhereGroupIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyRepo) CountWhereGroupIDReturnsOnCall(i int, result1 int, result2 error) {
	fake.CountWhereGroupIDStub = nil
	if fake.countWhereGroupIDReturnsOnCall == nil {
		fake.countWhereGroupIDReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countWhereGroupIDReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyRepo) CountWhereDestinationID(arg1 store.Transaction, arg2 int) (int, error) {
	fake.countWhereDestinationIDMutex.Lock()
	ret, specificReturn := fake.countWhereDestinationIDReturnsOnCall[len(fake.countWhereDestinationIDArgsForCall)]
	fake.countWhereDestinationIDArgsForCall = append(fake.countWhereDestinationIDArgsForCall, struct {
		arg1 store.Transaction
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CountWhereDestinationID", []interface{}{arg1, arg2})
	fake.countWhereDestinationIDMutex.Unlock()
	if fake.CountWhereDestinationIDStub != nil {
		return fake.CountWhereDestinationIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countWhereDestinationIDReturns.result1, fake.countWhereDestinationIDReturns.result2
}

func (fake *PolicyRepo) CountWhereDestinationIDCallCount() int {
	fake.countWhereDestinationIDMutex.RLock()
	defer fake.countWhereDestinationIDMutex.RUnlock()
	return len(fake.countWhereDestinationIDArgsForCall)
}

func (fake *PolicyRepo) CountWhereDestinationIDArgsForCall(i int) (store.Transaction, int) {
	fake.countWhereDestinationIDMutex.RLock()
	defer fake.countWhereDestinationIDMutex.RUnlock()
	return fake.countWhereDestinationIDArgsForCall[i].arg1, fake.countWhereDestinationIDArgsForCall[i].arg2
}

func (fake *PolicyRepo) CountWhereDestinationIDReturns(result1 int, result2 error) {
	fake.CountWhereDestinationIDStub = nil
	fake.countWhereDestinationIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyRepo) CountWhereDestinationIDReturnsOnCall(i int, result1 int, result2 error) {
	fake.CountWhereDestinationIDStub = nil
	if fake.countWhereDestinationIDReturnsOnCall == nil {
		fake.countWhereDestinationIDReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countWhereDestinationIDReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.delete2Mutex.RLock()
	defer fake.delete2Mutex.RUnlock()
	fake.countWhereGroupIDMutex.RLock()
	defer fake.countWhereGroupIDMutex.RUnlock()
	fake.countWhereDestinationIDMutex.RLock()
	defer fake.countWhereDestinationIDMutex.RUnlock()
	return fake.invocations
}

func (fake *PolicyRepo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ store.PolicyRepo = new(PolicyRepo)
