// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	io "io"
	os "os"
	sync "sync"

	pivnet "github.com/pivotal-cf/go-pivnet"
	commands "github.com/pivotal-cf/om/commands"
)

type PivnetDownloader struct {
	DownloadProductFileStub        func(*os.File, string, int, int, io.Writer) error
	downloadProductFileMutex       sync.RWMutex
	downloadProductFileArgsForCall []struct {
		arg1 *os.File
		arg2 string
		arg3 int
		arg4 int
		arg5 io.Writer
	}
	downloadProductFileReturns struct {
		result1 error
	}
	downloadProductFileReturnsOnCall map[int]struct {
		result1 error
	}
	ProductFilesForReleaseStub        func(string, int) ([]pivnet.ProductFile, error)
	productFilesForReleaseMutex       sync.RWMutex
	productFilesForReleaseArgsForCall []struct {
		arg1 string
		arg2 int
	}
	productFilesForReleaseReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFilesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	ReleaseDependenciesStub        func(string, int) ([]pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		arg1 string
		arg2 int
	}
	releaseDependenciesReturns struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	releaseDependenciesReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	ReleaseForVersionStub        func(string, string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		arg1 string
		arg2 string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	releasesForProductSlugReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PivnetDownloader) DownloadProductFile(arg1 *os.File, arg2 string, arg3 int, arg4 int, arg5 io.Writer) error {
	fake.downloadProductFileMutex.Lock()
	ret, specificReturn := fake.downloadProductFileReturnsOnCall[len(fake.downloadProductFileArgsForCall)]
	fake.downloadProductFileArgsForCall = append(fake.downloadProductFileArgsForCall, struct {
		arg1 *os.File
		arg2 string
		arg3 int
		arg4 int
		arg5 io.Writer
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("DownloadProductFile", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.downloadProductFileMutex.Unlock()
	if fake.DownloadProductFileStub != nil {
		return fake.DownloadProductFileStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadProductFileReturns
	return fakeReturns.result1
}

func (fake *PivnetDownloader) DownloadProductFileCallCount() int {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return len(fake.downloadProductFileArgsForCall)
}

func (fake *PivnetDownloader) DownloadProductFileCalls(stub func(*os.File, string, int, int, io.Writer) error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = stub
}

func (fake *PivnetDownloader) DownloadProductFileArgsForCall(i int) (*os.File, string, int, int, io.Writer) {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	argsForCall := fake.downloadProductFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *PivnetDownloader) DownloadProductFileReturns(result1 error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = nil
	fake.downloadProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *PivnetDownloader) DownloadProductFileReturnsOnCall(i int, result1 error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = nil
	if fake.downloadProductFileReturnsOnCall == nil {
		fake.downloadProductFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadProductFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PivnetDownloader) ProductFilesForRelease(arg1 string, arg2 int) ([]pivnet.ProductFile, error) {
	fake.productFilesForReleaseMutex.Lock()
	ret, specificReturn := fake.productFilesForReleaseReturnsOnCall[len(fake.productFilesForReleaseArgsForCall)]
	fake.productFilesForReleaseArgsForCall = append(fake.productFilesForReleaseArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ProductFilesForRelease", []interface{}{arg1, arg2})
	fake.productFilesForReleaseMutex.Unlock()
	if fake.ProductFilesForReleaseStub != nil {
		return fake.ProductFilesForReleaseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productFilesForReleaseReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetDownloader) ProductFilesForReleaseCallCount() int {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return len(fake.productFilesForReleaseArgsForCall)
}

func (fake *PivnetDownloader) ProductFilesForReleaseCalls(stub func(string, int) ([]pivnet.ProductFile, error)) {
	fake.productFilesForReleaseMutex.Lock()
	defer fake.productFilesForReleaseMutex.Unlock()
	fake.ProductFilesForReleaseStub = stub
}

func (fake *PivnetDownloader) ProductFilesForReleaseArgsForCall(i int) (string, int) {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	argsForCall := fake.productFilesForReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PivnetDownloader) ProductFilesForReleaseReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.productFilesForReleaseMutex.Lock()
	defer fake.productFilesForReleaseMutex.Unlock()
	fake.ProductFilesForReleaseStub = nil
	fake.productFilesForReleaseReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ProductFilesForReleaseReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.productFilesForReleaseMutex.Lock()
	defer fake.productFilesForReleaseMutex.Unlock()
	fake.ProductFilesForReleaseStub = nil
	if fake.productFilesForReleaseReturnsOnCall == nil {
		fake.productFilesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFilesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleaseDependencies(arg1 string, arg2 int) ([]pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	ret, specificReturn := fake.releaseDependenciesReturnsOnCall[len(fake.releaseDependenciesArgsForCall)]
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ReleaseDependencies", []interface{}{arg1, arg2})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseDependenciesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetDownloader) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *PivnetDownloader) ReleaseDependenciesCalls(stub func(string, int) ([]pivnet.ReleaseDependency, error)) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = stub
}

func (fake *PivnetDownloader) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	argsForCall := fake.releaseDependenciesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PivnetDownloader) ReleaseDependenciesReturns(result1 []pivnet.ReleaseDependency, result2 error) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleaseDependenciesReturnsOnCall(i int, result1 []pivnet.ReleaseDependency, result2 error) {
	fake.releaseDependenciesMutex.Lock()
	defer fake.releaseDependenciesMutex.Unlock()
	fake.ReleaseDependenciesStub = nil
	if fake.releaseDependenciesReturnsOnCall == nil {
		fake.releaseDependenciesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseDependency
			result2 error
		})
	}
	fake.releaseDependenciesReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleaseForVersion(arg1 string, arg2 string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ReleaseForVersion", []interface{}{arg1, arg2})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseForVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetDownloader) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *PivnetDownloader) ReleaseForVersionCalls(stub func(string, string) (pivnet.Release, error)) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = stub
}

func (fake *PivnetDownloader) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	argsForCall := fake.releaseForVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *PivnetDownloader) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleasesForProductSlug(arg1 string) ([]pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	ret, specificReturn := fake.releasesForProductSlugReturnsOnCall[len(fake.releasesForProductSlugArgsForCall)]
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releasesForProductSlugReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PivnetDownloader) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *PivnetDownloader) ReleasesForProductSlugCalls(stub func(string) ([]pivnet.Release, error)) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = stub
}

func (fake *PivnetDownloader) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	argsForCall := fake.releasesForProductSlugArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PivnetDownloader) ReleasesForProductSlugReturns(result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) ReleasesForProductSlugReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.releasesForProductSlugMutex.Lock()
	defer fake.releasesForProductSlugMutex.Unlock()
	fake.ReleasesForProductSlugStub = nil
	if fake.releasesForProductSlugReturnsOnCall == nil {
		fake.releasesForProductSlugReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.releasesForProductSlugReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *PivnetDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PivnetDownloader) recordInvocation(key string, args []interface{}) {
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

var _ commands.PivnetDownloader = new(PivnetDownloader)