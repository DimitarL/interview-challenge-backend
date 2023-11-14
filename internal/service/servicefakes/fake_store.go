// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/DimitarL/rental/internal/model"
	"github.com/DimitarL/rental/internal/service"
)

type FakeStore struct {
	GetRentalStub        func(context.Context, int) (model.Rental, error)
	getRentalMutex       sync.RWMutex
	getRentalArgsForCall []struct {
		arg1 context.Context
		arg2 int
	}
	getRentalReturns struct {
		result1 model.Rental
		result2 error
	}
	getRentalReturnsOnCall map[int]struct {
		result1 model.Rental
		result2 error
	}
	GetRentalsStub        func(context.Context, service.SearchCriteria) ([]model.Rental, error)
	getRentalsMutex       sync.RWMutex
	getRentalsArgsForCall []struct {
		arg1 context.Context
		arg2 service.SearchCriteria
	}
	getRentalsReturns struct {
		result1 []model.Rental
		result2 error
	}
	getRentalsReturnsOnCall map[int]struct {
		result1 []model.Rental
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) GetRental(arg1 context.Context, arg2 int) (model.Rental, error) {
	fake.getRentalMutex.Lock()
	ret, specificReturn := fake.getRentalReturnsOnCall[len(fake.getRentalArgsForCall)]
	fake.getRentalArgsForCall = append(fake.getRentalArgsForCall, struct {
		arg1 context.Context
		arg2 int
	}{arg1, arg2})
	stub := fake.GetRentalStub
	fakeReturns := fake.getRentalReturns
	fake.recordInvocation("GetRental", []interface{}{arg1, arg2})
	fake.getRentalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetRentalCallCount() int {
	fake.getRentalMutex.RLock()
	defer fake.getRentalMutex.RUnlock()
	return len(fake.getRentalArgsForCall)
}

func (fake *FakeStore) GetRentalCalls(stub func(context.Context, int) (model.Rental, error)) {
	fake.getRentalMutex.Lock()
	defer fake.getRentalMutex.Unlock()
	fake.GetRentalStub = stub
}

func (fake *FakeStore) GetRentalArgsForCall(i int) (context.Context, int) {
	fake.getRentalMutex.RLock()
	defer fake.getRentalMutex.RUnlock()
	argsForCall := fake.getRentalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) GetRentalReturns(result1 model.Rental, result2 error) {
	fake.getRentalMutex.Lock()
	defer fake.getRentalMutex.Unlock()
	fake.GetRentalStub = nil
	fake.getRentalReturns = struct {
		result1 model.Rental
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetRentalReturnsOnCall(i int, result1 model.Rental, result2 error) {
	fake.getRentalMutex.Lock()
	defer fake.getRentalMutex.Unlock()
	fake.GetRentalStub = nil
	if fake.getRentalReturnsOnCall == nil {
		fake.getRentalReturnsOnCall = make(map[int]struct {
			result1 model.Rental
			result2 error
		})
	}
	fake.getRentalReturnsOnCall[i] = struct {
		result1 model.Rental
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetRentals(arg1 context.Context, arg2 service.SearchCriteria) ([]model.Rental, error) {
	fake.getRentalsMutex.Lock()
	ret, specificReturn := fake.getRentalsReturnsOnCall[len(fake.getRentalsArgsForCall)]
	fake.getRentalsArgsForCall = append(fake.getRentalsArgsForCall, struct {
		arg1 context.Context
		arg2 service.SearchCriteria
	}{arg1, arg2})
	stub := fake.GetRentalsStub
	fakeReturns := fake.getRentalsReturns
	fake.recordInvocation("GetRentals", []interface{}{arg1, arg2})
	fake.getRentalsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) GetRentalsCallCount() int {
	fake.getRentalsMutex.RLock()
	defer fake.getRentalsMutex.RUnlock()
	return len(fake.getRentalsArgsForCall)
}

func (fake *FakeStore) GetRentalsCalls(stub func(context.Context, service.SearchCriteria) ([]model.Rental, error)) {
	fake.getRentalsMutex.Lock()
	defer fake.getRentalsMutex.Unlock()
	fake.GetRentalsStub = stub
}

func (fake *FakeStore) GetRentalsArgsForCall(i int) (context.Context, service.SearchCriteria) {
	fake.getRentalsMutex.RLock()
	defer fake.getRentalsMutex.RUnlock()
	argsForCall := fake.getRentalsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) GetRentalsReturns(result1 []model.Rental, result2 error) {
	fake.getRentalsMutex.Lock()
	defer fake.getRentalsMutex.Unlock()
	fake.GetRentalsStub = nil
	fake.getRentalsReturns = struct {
		result1 []model.Rental
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetRentalsReturnsOnCall(i int, result1 []model.Rental, result2 error) {
	fake.getRentalsMutex.Lock()
	defer fake.getRentalsMutex.Unlock()
	fake.GetRentalsStub = nil
	if fake.getRentalsReturnsOnCall == nil {
		fake.getRentalsReturnsOnCall = make(map[int]struct {
			result1 []model.Rental
			result2 error
		})
	}
	fake.getRentalsReturnsOnCall[i] = struct {
		result1 []model.Rental
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRentalMutex.RLock()
	defer fake.getRentalMutex.RUnlock()
	fake.getRentalsMutex.RLock()
	defer fake.getRentalsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
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

var _ service.Store = new(FakeStore)
