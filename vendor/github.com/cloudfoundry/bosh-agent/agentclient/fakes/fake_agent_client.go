// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agentclient"
	"github.com/cloudfoundry/bosh-agent/agentclient/applyspec"
)

type FakeAgentClient struct {
	PingStub        func() (string, error)
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 string
		result2 error
	}
	pingReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	DrainStub        func(string) (int64, error)
	drainMutex       sync.RWMutex
	drainArgsForCall []struct {
		arg1 string
	}
	drainReturns struct {
		result1 int64
		result2 error
	}
	drainReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	ApplyStub        func(applyspec.ApplySpec) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 applyspec.ApplySpec
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	GetStateStub        func() (agentclient.AgentState, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct{}
	getStateReturns     struct {
		result1 agentclient.AgentState
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 agentclient.AgentState
		result2 error
	}
	AddPersistentDiskStub        func(string, interface{}) error
	addPersistentDiskMutex       sync.RWMutex
	addPersistentDiskArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	addPersistentDiskReturns struct {
		result1 error
	}
	addPersistentDiskReturnsOnCall map[int]struct {
		result1 error
	}
	MountDiskStub        func(string, interface{}) error
	mountDiskMutex       sync.RWMutex
	mountDiskArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	mountDiskReturns struct {
		result1 error
	}
	mountDiskReturnsOnCall map[int]struct {
		result1 error
	}
	UnmountDiskStub        func(string) error
	unmountDiskMutex       sync.RWMutex
	unmountDiskArgsForCall []struct {
		arg1 string
	}
	unmountDiskReturns struct {
		result1 error
	}
	unmountDiskReturnsOnCall map[int]struct {
		result1 error
	}
	ListDiskStub        func() ([]string, error)
	listDiskMutex       sync.RWMutex
	listDiskArgsForCall []struct{}
	listDiskReturns     struct {
		result1 []string
		result2 error
	}
	listDiskReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MigrateDiskStub        func() error
	migrateDiskMutex       sync.RWMutex
	migrateDiskArgsForCall []struct{}
	migrateDiskReturns     struct {
		result1 error
	}
	migrateDiskReturnsOnCall map[int]struct {
		result1 error
	}
	CompilePackageStub        func(packageSource agentclient.BlobRef, compiledPackageDependencies []agentclient.BlobRef) (compiledPackageRef agentclient.BlobRef, err error)
	compilePackageMutex       sync.RWMutex
	compilePackageArgsForCall []struct {
		packageSource               agentclient.BlobRef
		compiledPackageDependencies []agentclient.BlobRef
	}
	compilePackageReturns struct {
		result1 agentclient.BlobRef
		result2 error
	}
	compilePackageReturnsOnCall map[int]struct {
		result1 agentclient.BlobRef
		result2 error
	}
	DeleteARPEntriesStub        func(ips []string) error
	deleteARPEntriesMutex       sync.RWMutex
	deleteARPEntriesArgsForCall []struct {
		ips []string
	}
	deleteARPEntriesReturns struct {
		result1 error
	}
	deleteARPEntriesReturnsOnCall map[int]struct {
		result1 error
	}
	SyncDNSStub        func(blobID, sha1 string, version uint64) (string, error)
	syncDNSMutex       sync.RWMutex
	syncDNSArgsForCall []struct {
		blobID  string
		sha1    string
		version uint64
	}
	syncDNSReturns struct {
		result1 string
		result2 error
	}
	syncDNSReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunScriptStub        func(scriptName string, options map[string]interface{}) error
	runScriptMutex       sync.RWMutex
	runScriptArgsForCall []struct {
		scriptName string
		options    map[string]interface{}
	}
	runScriptReturns struct {
		result1 error
	}
	runScriptReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAgentClient) Ping() (string, error) {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.pingReturns.result1, fake.pingReturns.result2
}

func (fake *FakeAgentClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeAgentClient) PingReturns(result1 string, result2 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) PingReturnsOnCall(i int, result1 string, result2 error) {
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopReturns.result1
}

func (fake *FakeAgentClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeAgentClient) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) StopReturnsOnCall(i int, result1 error) {
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Drain(arg1 string) (int64, error) {
	fake.drainMutex.Lock()
	ret, specificReturn := fake.drainReturnsOnCall[len(fake.drainArgsForCall)]
	fake.drainArgsForCall = append(fake.drainArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Drain", []interface{}{arg1})
	fake.drainMutex.Unlock()
	if fake.DrainStub != nil {
		return fake.DrainStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.drainReturns.result1, fake.drainReturns.result2
}

func (fake *FakeAgentClient) DrainCallCount() int {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	return len(fake.drainArgsForCall)
}

func (fake *FakeAgentClient) DrainArgsForCall(i int) string {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	return fake.drainArgsForCall[i].arg1
}

func (fake *FakeAgentClient) DrainReturns(result1 int64, result2 error) {
	fake.DrainStub = nil
	fake.drainReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) DrainReturnsOnCall(i int, result1 int64, result2 error) {
	fake.DrainStub = nil
	if fake.drainReturnsOnCall == nil {
		fake.drainReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.drainReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) Apply(arg1 applyspec.ApplySpec) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 applyspec.ApplySpec
	}{arg1})
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.applyReturns.result1
}

func (fake *FakeAgentClient) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeAgentClient) ApplyArgsForCall(i int) applyspec.ApplySpec {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].arg1
}

func (fake *FakeAgentClient) ApplyReturns(result1 error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) ApplyReturnsOnCall(i int, result1 error) {
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *FakeAgentClient) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeAgentClient) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) StartReturnsOnCall(i int, result1 error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) GetState() (agentclient.AgentState, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct{}{})
	fake.recordInvocation("GetState", []interface{}{})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateReturns.result1, fake.getStateReturns.result2
}

func (fake *FakeAgentClient) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *FakeAgentClient) GetStateReturns(result1 agentclient.AgentState, result2 error) {
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 agentclient.AgentState
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) GetStateReturnsOnCall(i int, result1 agentclient.AgentState, result2 error) {
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 agentclient.AgentState
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 agentclient.AgentState
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) AddPersistentDisk(arg1 string, arg2 interface{}) error {
	fake.addPersistentDiskMutex.Lock()
	ret, specificReturn := fake.addPersistentDiskReturnsOnCall[len(fake.addPersistentDiskArgsForCall)]
	fake.addPersistentDiskArgsForCall = append(fake.addPersistentDiskArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("AddPersistentDisk", []interface{}{arg1, arg2})
	fake.addPersistentDiskMutex.Unlock()
	if fake.AddPersistentDiskStub != nil {
		return fake.AddPersistentDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addPersistentDiskReturns.result1
}

func (fake *FakeAgentClient) AddPersistentDiskCallCount() int {
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	return len(fake.addPersistentDiskArgsForCall)
}

func (fake *FakeAgentClient) AddPersistentDiskArgsForCall(i int) (string, interface{}) {
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	return fake.addPersistentDiskArgsForCall[i].arg1, fake.addPersistentDiskArgsForCall[i].arg2
}

func (fake *FakeAgentClient) AddPersistentDiskReturns(result1 error) {
	fake.AddPersistentDiskStub = nil
	fake.addPersistentDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) AddPersistentDiskReturnsOnCall(i int, result1 error) {
	fake.AddPersistentDiskStub = nil
	if fake.addPersistentDiskReturnsOnCall == nil {
		fake.addPersistentDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPersistentDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MountDisk(arg1 string, arg2 interface{}) error {
	fake.mountDiskMutex.Lock()
	ret, specificReturn := fake.mountDiskReturnsOnCall[len(fake.mountDiskArgsForCall)]
	fake.mountDiskArgsForCall = append(fake.mountDiskArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("MountDisk", []interface{}{arg1, arg2})
	fake.mountDiskMutex.Unlock()
	if fake.MountDiskStub != nil {
		return fake.MountDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mountDiskReturns.result1
}

func (fake *FakeAgentClient) MountDiskCallCount() int {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	return len(fake.mountDiskArgsForCall)
}

func (fake *FakeAgentClient) MountDiskArgsForCall(i int) (string, interface{}) {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	return fake.mountDiskArgsForCall[i].arg1, fake.mountDiskArgsForCall[i].arg2
}

func (fake *FakeAgentClient) MountDiskReturns(result1 error) {
	fake.MountDiskStub = nil
	fake.mountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MountDiskReturnsOnCall(i int, result1 error) {
	fake.MountDiskStub = nil
	if fake.mountDiskReturnsOnCall == nil {
		fake.mountDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) UnmountDisk(arg1 string) error {
	fake.unmountDiskMutex.Lock()
	ret, specificReturn := fake.unmountDiskReturnsOnCall[len(fake.unmountDiskArgsForCall)]
	fake.unmountDiskArgsForCall = append(fake.unmountDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnmountDisk", []interface{}{arg1})
	fake.unmountDiskMutex.Unlock()
	if fake.UnmountDiskStub != nil {
		return fake.UnmountDiskStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unmountDiskReturns.result1
}

func (fake *FakeAgentClient) UnmountDiskCallCount() int {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	return len(fake.unmountDiskArgsForCall)
}

func (fake *FakeAgentClient) UnmountDiskArgsForCall(i int) string {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	return fake.unmountDiskArgsForCall[i].arg1
}

func (fake *FakeAgentClient) UnmountDiskReturns(result1 error) {
	fake.UnmountDiskStub = nil
	fake.unmountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) UnmountDiskReturnsOnCall(i int, result1 error) {
	fake.UnmountDiskStub = nil
	if fake.unmountDiskReturnsOnCall == nil {
		fake.unmountDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmountDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) ListDisk() ([]string, error) {
	fake.listDiskMutex.Lock()
	ret, specificReturn := fake.listDiskReturnsOnCall[len(fake.listDiskArgsForCall)]
	fake.listDiskArgsForCall = append(fake.listDiskArgsForCall, struct{}{})
	fake.recordInvocation("ListDisk", []interface{}{})
	fake.listDiskMutex.Unlock()
	if fake.ListDiskStub != nil {
		return fake.ListDiskStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDiskReturns.result1, fake.listDiskReturns.result2
}

func (fake *FakeAgentClient) ListDiskCallCount() int {
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	return len(fake.listDiskArgsForCall)
}

func (fake *FakeAgentClient) ListDiskReturns(result1 []string, result2 error) {
	fake.ListDiskStub = nil
	fake.listDiskReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) ListDiskReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ListDiskStub = nil
	if fake.listDiskReturnsOnCall == nil {
		fake.listDiskReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listDiskReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) MigrateDisk() error {
	fake.migrateDiskMutex.Lock()
	ret, specificReturn := fake.migrateDiskReturnsOnCall[len(fake.migrateDiskArgsForCall)]
	fake.migrateDiskArgsForCall = append(fake.migrateDiskArgsForCall, struct{}{})
	fake.recordInvocation("MigrateDisk", []interface{}{})
	fake.migrateDiskMutex.Unlock()
	if fake.MigrateDiskStub != nil {
		return fake.MigrateDiskStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.migrateDiskReturns.result1
}

func (fake *FakeAgentClient) MigrateDiskCallCount() int {
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	return len(fake.migrateDiskArgsForCall)
}

func (fake *FakeAgentClient) MigrateDiskReturns(result1 error) {
	fake.MigrateDiskStub = nil
	fake.migrateDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MigrateDiskReturnsOnCall(i int, result1 error) {
	fake.MigrateDiskStub = nil
	if fake.migrateDiskReturnsOnCall == nil {
		fake.migrateDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) CompilePackage(packageSource agentclient.BlobRef, compiledPackageDependencies []agentclient.BlobRef) (compiledPackageRef agentclient.BlobRef, err error) {
	var compiledPackageDependenciesCopy []agentclient.BlobRef
	if compiledPackageDependencies != nil {
		compiledPackageDependenciesCopy = make([]agentclient.BlobRef, len(compiledPackageDependencies))
		copy(compiledPackageDependenciesCopy, compiledPackageDependencies)
	}
	fake.compilePackageMutex.Lock()
	ret, specificReturn := fake.compilePackageReturnsOnCall[len(fake.compilePackageArgsForCall)]
	fake.compilePackageArgsForCall = append(fake.compilePackageArgsForCall, struct {
		packageSource               agentclient.BlobRef
		compiledPackageDependencies []agentclient.BlobRef
	}{packageSource, compiledPackageDependenciesCopy})
	fake.recordInvocation("CompilePackage", []interface{}{packageSource, compiledPackageDependenciesCopy})
	fake.compilePackageMutex.Unlock()
	if fake.CompilePackageStub != nil {
		return fake.CompilePackageStub(packageSource, compiledPackageDependencies)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.compilePackageReturns.result1, fake.compilePackageReturns.result2
}

func (fake *FakeAgentClient) CompilePackageCallCount() int {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	return len(fake.compilePackageArgsForCall)
}

func (fake *FakeAgentClient) CompilePackageArgsForCall(i int) (agentclient.BlobRef, []agentclient.BlobRef) {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	return fake.compilePackageArgsForCall[i].packageSource, fake.compilePackageArgsForCall[i].compiledPackageDependencies
}

func (fake *FakeAgentClient) CompilePackageReturns(result1 agentclient.BlobRef, result2 error) {
	fake.CompilePackageStub = nil
	fake.compilePackageReturns = struct {
		result1 agentclient.BlobRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) CompilePackageReturnsOnCall(i int, result1 agentclient.BlobRef, result2 error) {
	fake.CompilePackageStub = nil
	if fake.compilePackageReturnsOnCall == nil {
		fake.compilePackageReturnsOnCall = make(map[int]struct {
			result1 agentclient.BlobRef
			result2 error
		})
	}
	fake.compilePackageReturnsOnCall[i] = struct {
		result1 agentclient.BlobRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) DeleteARPEntries(ips []string) error {
	var ipsCopy []string
	if ips != nil {
		ipsCopy = make([]string, len(ips))
		copy(ipsCopy, ips)
	}
	fake.deleteARPEntriesMutex.Lock()
	ret, specificReturn := fake.deleteARPEntriesReturnsOnCall[len(fake.deleteARPEntriesArgsForCall)]
	fake.deleteARPEntriesArgsForCall = append(fake.deleteARPEntriesArgsForCall, struct {
		ips []string
	}{ipsCopy})
	fake.recordInvocation("DeleteARPEntries", []interface{}{ipsCopy})
	fake.deleteARPEntriesMutex.Unlock()
	if fake.DeleteARPEntriesStub != nil {
		return fake.DeleteARPEntriesStub(ips)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteARPEntriesReturns.result1
}

func (fake *FakeAgentClient) DeleteARPEntriesCallCount() int {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	return len(fake.deleteARPEntriesArgsForCall)
}

func (fake *FakeAgentClient) DeleteARPEntriesArgsForCall(i int) []string {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	return fake.deleteARPEntriesArgsForCall[i].ips
}

func (fake *FakeAgentClient) DeleteARPEntriesReturns(result1 error) {
	fake.DeleteARPEntriesStub = nil
	fake.deleteARPEntriesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) DeleteARPEntriesReturnsOnCall(i int, result1 error) {
	fake.DeleteARPEntriesStub = nil
	if fake.deleteARPEntriesReturnsOnCall == nil {
		fake.deleteARPEntriesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteARPEntriesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) SyncDNS(blobID string, sha1 string, version uint64) (string, error) {
	fake.syncDNSMutex.Lock()
	ret, specificReturn := fake.syncDNSReturnsOnCall[len(fake.syncDNSArgsForCall)]
	fake.syncDNSArgsForCall = append(fake.syncDNSArgsForCall, struct {
		blobID  string
		sha1    string
		version uint64
	}{blobID, sha1, version})
	fake.recordInvocation("SyncDNS", []interface{}{blobID, sha1, version})
	fake.syncDNSMutex.Unlock()
	if fake.SyncDNSStub != nil {
		return fake.SyncDNSStub(blobID, sha1, version)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.syncDNSReturns.result1, fake.syncDNSReturns.result2
}

func (fake *FakeAgentClient) SyncDNSCallCount() int {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	return len(fake.syncDNSArgsForCall)
}

func (fake *FakeAgentClient) SyncDNSArgsForCall(i int) (string, string, uint64) {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	return fake.syncDNSArgsForCall[i].blobID, fake.syncDNSArgsForCall[i].sha1, fake.syncDNSArgsForCall[i].version
}

func (fake *FakeAgentClient) SyncDNSReturns(result1 string, result2 error) {
	fake.SyncDNSStub = nil
	fake.syncDNSReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) SyncDNSReturnsOnCall(i int, result1 string, result2 error) {
	fake.SyncDNSStub = nil
	if fake.syncDNSReturnsOnCall == nil {
		fake.syncDNSReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.syncDNSReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) RunScript(scriptName string, options map[string]interface{}) error {
	fake.runScriptMutex.Lock()
	ret, specificReturn := fake.runScriptReturnsOnCall[len(fake.runScriptArgsForCall)]
	fake.runScriptArgsForCall = append(fake.runScriptArgsForCall, struct {
		scriptName string
		options    map[string]interface{}
	}{scriptName, options})
	fake.recordInvocation("RunScript", []interface{}{scriptName, options})
	fake.runScriptMutex.Unlock()
	if fake.RunScriptStub != nil {
		return fake.RunScriptStub(scriptName, options)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runScriptReturns.result1
}

func (fake *FakeAgentClient) RunScriptCallCount() int {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	return len(fake.runScriptArgsForCall)
}

func (fake *FakeAgentClient) RunScriptArgsForCall(i int) (string, map[string]interface{}) {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	return fake.runScriptArgsForCall[i].scriptName, fake.runScriptArgsForCall[i].options
}

func (fake *FakeAgentClient) RunScriptReturns(result1 error) {
	fake.RunScriptStub = nil
	fake.runScriptReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) RunScriptReturnsOnCall(i int, result1 error) {
	fake.RunScriptStub = nil
	if fake.runScriptReturnsOnCall == nil {
		fake.runScriptReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runScriptReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAgentClient) recordInvocation(key string, args []interface{}) {
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

var _ agentclient.AgentClient = new(FakeAgentClient)
