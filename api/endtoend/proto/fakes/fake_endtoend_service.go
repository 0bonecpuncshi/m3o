// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/micro/micro/v3/service/client"
	endtoend "m3o.dev/api/endtoend/proto"
)

type FakeEndtoendService struct {
	CheckStub        func(context.Context, *endtoend.Request, ...client.CallOption) (*endtoend.Response, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 context.Context
		arg2 *endtoend.Request
		arg3 []client.CallOption
	}
	checkReturns struct {
		result1 *endtoend.Response
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 *endtoend.Response
		result2 error
	}
	RunCheckStub        func(context.Context, *endtoend.Request, ...client.CallOption) (*endtoend.Response, error)
	runCheckMutex       sync.RWMutex
	runCheckArgsForCall []struct {
		arg1 context.Context
		arg2 *endtoend.Request
		arg3 []client.CallOption
	}
	runCheckReturns struct {
		result1 *endtoend.Response
		result2 error
	}
	runCheckReturnsOnCall map[int]struct {
		result1 *endtoend.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEndtoendService) Check(arg1 context.Context, arg2 *endtoend.Request, arg3 ...client.CallOption) (*endtoend.Response, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 context.Context
		arg2 *endtoend.Request
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.CheckStub
	fakeReturns := fake.checkReturns
	fake.recordInvocation("Check", []interface{}{arg1, arg2, arg3})
	fake.checkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEndtoendService) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeEndtoendService) CheckCalls(stub func(context.Context, *endtoend.Request, ...client.CallOption) (*endtoend.Response, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeEndtoendService) CheckArgsForCall(i int) (context.Context, *endtoend.Request, []client.CallOption) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEndtoendService) CheckReturns(result1 *endtoend.Response, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 *endtoend.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeEndtoendService) CheckReturnsOnCall(i int, result1 *endtoend.Response, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 *endtoend.Response
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 *endtoend.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeEndtoendService) RunCheck(arg1 context.Context, arg2 *endtoend.Request, arg3 ...client.CallOption) (*endtoend.Response, error) {
	fake.runCheckMutex.Lock()
	ret, specificReturn := fake.runCheckReturnsOnCall[len(fake.runCheckArgsForCall)]
	fake.runCheckArgsForCall = append(fake.runCheckArgsForCall, struct {
		arg1 context.Context
		arg2 *endtoend.Request
		arg3 []client.CallOption
	}{arg1, arg2, arg3})
	stub := fake.RunCheckStub
	fakeReturns := fake.runCheckReturns
	fake.recordInvocation("RunCheck", []interface{}{arg1, arg2, arg3})
	fake.runCheckMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEndtoendService) RunCheckCallCount() int {
	fake.runCheckMutex.RLock()
	defer fake.runCheckMutex.RUnlock()
	return len(fake.runCheckArgsForCall)
}

func (fake *FakeEndtoendService) RunCheckCalls(stub func(context.Context, *endtoend.Request, ...client.CallOption) (*endtoend.Response, error)) {
	fake.runCheckMutex.Lock()
	defer fake.runCheckMutex.Unlock()
	fake.RunCheckStub = stub
}

func (fake *FakeEndtoendService) RunCheckArgsForCall(i int) (context.Context, *endtoend.Request, []client.CallOption) {
	fake.runCheckMutex.RLock()
	defer fake.runCheckMutex.RUnlock()
	argsForCall := fake.runCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEndtoendService) RunCheckReturns(result1 *endtoend.Response, result2 error) {
	fake.runCheckMutex.Lock()
	defer fake.runCheckMutex.Unlock()
	fake.RunCheckStub = nil
	fake.runCheckReturns = struct {
		result1 *endtoend.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeEndtoendService) RunCheckReturnsOnCall(i int, result1 *endtoend.Response, result2 error) {
	fake.runCheckMutex.Lock()
	defer fake.runCheckMutex.Unlock()
	fake.RunCheckStub = nil
	if fake.runCheckReturnsOnCall == nil {
		fake.runCheckReturnsOnCall = make(map[int]struct {
			result1 *endtoend.Response
			result2 error
		})
	}
	fake.runCheckReturnsOnCall[i] = struct {
		result1 *endtoend.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeEndtoendService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.runCheckMutex.RLock()
	defer fake.runCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEndtoendService) recordInvocation(key string, args []interface{}) {
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

var _ endtoend.EndtoendService = new(FakeEndtoendService)
