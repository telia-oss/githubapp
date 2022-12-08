// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/google/go-github/v45/github"
	"github.com/telia-oss/githubapp"
)

type FakeAppsJWTAPI struct {
	CreateInstallationTokenStub        func(context.Context, int64, *github.InstallationTokenOptions) (*github.InstallationToken, *github.Response, error)
	createInstallationTokenMutex       sync.RWMutex
	createInstallationTokenArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *github.InstallationTokenOptions
	}
	createInstallationTokenReturns struct {
		result1 *github.InstallationToken
		result2 *github.Response
		result3 error
	}
	createInstallationTokenReturnsOnCall map[int]struct {
		result1 *github.InstallationToken
		result2 *github.Response
		result3 error
	}
	ListInstallationsStub        func(context.Context, *github.ListOptions) ([]*github.Installation, *github.Response, error)
	listInstallationsMutex       sync.RWMutex
	listInstallationsArgsForCall []struct {
		arg1 context.Context
		arg2 *github.ListOptions
	}
	listInstallationsReturns struct {
		result1 []*github.Installation
		result2 *github.Response
		result3 error
	}
	listInstallationsReturnsOnCall map[int]struct {
		result1 []*github.Installation
		result2 *github.Response
		result3 error
	}
	RateLimitsStub        func(context.Context) (*github.RateLimits, *github.Response, error)
	rateLimitsMutex       sync.RWMutex
	rateLimitsArgsForCall []struct {
		arg1 context.Context
	}
	rateLimitsReturns struct {
		result1 *github.RateLimits
		result2 *github.Response
		result3 error
	}
	rateLimitsReturnsOnCall map[int]struct {
		result1 *github.RateLimits
		result2 *github.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppsJWTAPI) CreateInstallationToken(arg1 context.Context, arg2 int64, arg3 *github.InstallationTokenOptions) (*github.InstallationToken, *github.Response, error) {
	fake.createInstallationTokenMutex.Lock()
	ret, specificReturn := fake.createInstallationTokenReturnsOnCall[len(fake.createInstallationTokenArgsForCall)]
	fake.createInstallationTokenArgsForCall = append(fake.createInstallationTokenArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *github.InstallationTokenOptions
	}{arg1, arg2, arg3})
	stub := fake.CreateInstallationTokenStub
	fakeReturns := fake.createInstallationTokenReturns
	fake.recordInvocation("CreateInstallationToken", []interface{}{arg1, arg2, arg3})
	fake.createInstallationTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAppsJWTAPI) CreateInstallationTokenCallCount() int {
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	return len(fake.createInstallationTokenArgsForCall)
}

func (fake *FakeAppsJWTAPI) CreateInstallationTokenCalls(stub func(context.Context, int64, *github.InstallationTokenOptions) (*github.InstallationToken, *github.Response, error)) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = stub
}

func (fake *FakeAppsJWTAPI) CreateInstallationTokenArgsForCall(i int) (context.Context, int64, *github.InstallationTokenOptions) {
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	argsForCall := fake.createInstallationTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAppsJWTAPI) CreateInstallationTokenReturns(result1 *github.InstallationToken, result2 *github.Response, result3 error) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = nil
	fake.createInstallationTokenReturns = struct {
		result1 *github.InstallationToken
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) CreateInstallationTokenReturnsOnCall(i int, result1 *github.InstallationToken, result2 *github.Response, result3 error) {
	fake.createInstallationTokenMutex.Lock()
	defer fake.createInstallationTokenMutex.Unlock()
	fake.CreateInstallationTokenStub = nil
	if fake.createInstallationTokenReturnsOnCall == nil {
		fake.createInstallationTokenReturnsOnCall = make(map[int]struct {
			result1 *github.InstallationToken
			result2 *github.Response
			result3 error
		})
	}
	fake.createInstallationTokenReturnsOnCall[i] = struct {
		result1 *github.InstallationToken
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) ListInstallations(arg1 context.Context, arg2 *github.ListOptions) ([]*github.Installation, *github.Response, error) {
	fake.listInstallationsMutex.Lock()
	ret, specificReturn := fake.listInstallationsReturnsOnCall[len(fake.listInstallationsArgsForCall)]
	fake.listInstallationsArgsForCall = append(fake.listInstallationsArgsForCall, struct {
		arg1 context.Context
		arg2 *github.ListOptions
	}{arg1, arg2})
	stub := fake.ListInstallationsStub
	fakeReturns := fake.listInstallationsReturns
	fake.recordInvocation("ListInstallations", []interface{}{arg1, arg2})
	fake.listInstallationsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAppsJWTAPI) ListInstallationsCallCount() int {
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	return len(fake.listInstallationsArgsForCall)
}

func (fake *FakeAppsJWTAPI) ListInstallationsCalls(stub func(context.Context, *github.ListOptions) ([]*github.Installation, *github.Response, error)) {
	fake.listInstallationsMutex.Lock()
	defer fake.listInstallationsMutex.Unlock()
	fake.ListInstallationsStub = stub
}

func (fake *FakeAppsJWTAPI) ListInstallationsArgsForCall(i int) (context.Context, *github.ListOptions) {
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	argsForCall := fake.listInstallationsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAppsJWTAPI) ListInstallationsReturns(result1 []*github.Installation, result2 *github.Response, result3 error) {
	fake.listInstallationsMutex.Lock()
	defer fake.listInstallationsMutex.Unlock()
	fake.ListInstallationsStub = nil
	fake.listInstallationsReturns = struct {
		result1 []*github.Installation
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) ListInstallationsReturnsOnCall(i int, result1 []*github.Installation, result2 *github.Response, result3 error) {
	fake.listInstallationsMutex.Lock()
	defer fake.listInstallationsMutex.Unlock()
	fake.ListInstallationsStub = nil
	if fake.listInstallationsReturnsOnCall == nil {
		fake.listInstallationsReturnsOnCall = make(map[int]struct {
			result1 []*github.Installation
			result2 *github.Response
			result3 error
		})
	}
	fake.listInstallationsReturnsOnCall[i] = struct {
		result1 []*github.Installation
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) RateLimits(arg1 context.Context) (*github.RateLimits, *github.Response, error) {
	fake.rateLimitsMutex.Lock()
	ret, specificReturn := fake.rateLimitsReturnsOnCall[len(fake.rateLimitsArgsForCall)]
	fake.rateLimitsArgsForCall = append(fake.rateLimitsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.RateLimitsStub
	fakeReturns := fake.rateLimitsReturns
	fake.recordInvocation("RateLimits", []interface{}{arg1})
	fake.rateLimitsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAppsJWTAPI) RateLimitsCallCount() int {
	fake.rateLimitsMutex.RLock()
	defer fake.rateLimitsMutex.RUnlock()
	return len(fake.rateLimitsArgsForCall)
}

func (fake *FakeAppsJWTAPI) RateLimitsCalls(stub func(context.Context) (*github.RateLimits, *github.Response, error)) {
	fake.rateLimitsMutex.Lock()
	defer fake.rateLimitsMutex.Unlock()
	fake.RateLimitsStub = stub
}

func (fake *FakeAppsJWTAPI) RateLimitsArgsForCall(i int) context.Context {
	fake.rateLimitsMutex.RLock()
	defer fake.rateLimitsMutex.RUnlock()
	argsForCall := fake.rateLimitsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAppsJWTAPI) RateLimitsReturns(result1 *github.RateLimits, result2 *github.Response, result3 error) {
	fake.rateLimitsMutex.Lock()
	defer fake.rateLimitsMutex.Unlock()
	fake.RateLimitsStub = nil
	fake.rateLimitsReturns = struct {
		result1 *github.RateLimits
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) RateLimitsReturnsOnCall(i int, result1 *github.RateLimits, result2 *github.Response, result3 error) {
	fake.rateLimitsMutex.Lock()
	defer fake.rateLimitsMutex.Unlock()
	fake.RateLimitsStub = nil
	if fake.rateLimitsReturnsOnCall == nil {
		fake.rateLimitsReturnsOnCall = make(map[int]struct {
			result1 *github.RateLimits
			result2 *github.Response
			result3 error
		})
	}
	fake.rateLimitsReturnsOnCall[i] = struct {
		result1 *github.RateLimits
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAppsJWTAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createInstallationTokenMutex.RLock()
	defer fake.createInstallationTokenMutex.RUnlock()
	fake.listInstallationsMutex.RLock()
	defer fake.listInstallationsMutex.RUnlock()
	fake.rateLimitsMutex.RLock()
	defer fake.rateLimitsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAppsJWTAPI) recordInvocation(key string, args []interface{}) {
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

var _ githubapp.AppsJWTAPI = new(FakeAppsJWTAPI)
