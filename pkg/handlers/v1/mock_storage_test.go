// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/asecurityteam/asset-inventory-api/pkg/domain (interfaces: PartitionGenerator,PartitionsGetter,PartitionsDeleter,CloudAssetStorer,CloudAssetByIPFetcher,CloudAssetByHostnameFetcher,CloudAllAssetsByTimeFetcher)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	domain "github.com/asecurityteam/asset-inventory-api/pkg/domain"
)

// MockPartitionGenerator is a mock of PartitionGenerator interface
type MockPartitionGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockPartitionGeneratorMockRecorder
}

// MockPartitionGeneratorMockRecorder is the mock recorder for MockPartitionGenerator
type MockPartitionGeneratorMockRecorder struct {
	mock *MockPartitionGenerator
}

// NewMockPartitionGenerator creates a new mock instance
func NewMockPartitionGenerator(ctrl *gomock.Controller) *MockPartitionGenerator {
	mock := &MockPartitionGenerator{ctrl: ctrl}
	mock.recorder = &MockPartitionGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPartitionGenerator) EXPECT() *MockPartitionGeneratorMockRecorder {
	return m.recorder
}

// GeneratePartition mocks base method
func (m *MockPartitionGenerator) GeneratePartition(arg0 context.Context, arg1 time.Time, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePartition", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GeneratePartition indicates an expected call of GeneratePartition
func (mr *MockPartitionGeneratorMockRecorder) GeneratePartition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePartition", reflect.TypeOf((*MockPartitionGenerator)(nil).GeneratePartition), arg0, arg1, arg2)
}

// MockPartitionsGetter is a mock of PartitionsGetter interface
type MockPartitionsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPartitionsGetterMockRecorder
}

// MockPartitionsGetterMockRecorder is the mock recorder for MockPartitionsGetter
type MockPartitionsGetterMockRecorder struct {
	mock *MockPartitionsGetter
}

// NewMockPartitionsGetter creates a new mock instance
func NewMockPartitionsGetter(ctrl *gomock.Controller) *MockPartitionsGetter {
	mock := &MockPartitionsGetter{ctrl: ctrl}
	mock.recorder = &MockPartitionsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPartitionsGetter) EXPECT() *MockPartitionsGetterMockRecorder {
	return m.recorder
}

// GetPartitions mocks base method
func (m *MockPartitionsGetter) GetPartitions(arg0 context.Context) ([]domain.Partition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitions", arg0)
	ret0, _ := ret[0].([]domain.Partition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartitions indicates an expected call of GetPartitions
func (mr *MockPartitionsGetterMockRecorder) GetPartitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitions", reflect.TypeOf((*MockPartitionsGetter)(nil).GetPartitions), arg0)
}

// MockPartitionsDeleter is a mock of PartitionsDeleter interface
type MockPartitionsDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPartitionsDeleterMockRecorder
}

// MockPartitionsDeleterMockRecorder is the mock recorder for MockPartitionsDeleter
type MockPartitionsDeleterMockRecorder struct {
	mock *MockPartitionsDeleter
}

// NewMockPartitionsDeleter creates a new mock instance
func NewMockPartitionsDeleter(ctrl *gomock.Controller) *MockPartitionsDeleter {
	mock := &MockPartitionsDeleter{ctrl: ctrl}
	mock.recorder = &MockPartitionsDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPartitionsDeleter) EXPECT() *MockPartitionsDeleterMockRecorder {
	return m.recorder
}

// DeletePartitions mocks base method
func (m *MockPartitionsDeleter) DeletePartitions(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePartitions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePartitions indicates an expected call of DeletePartitions
func (mr *MockPartitionsDeleterMockRecorder) DeletePartitions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePartitions", reflect.TypeOf((*MockPartitionsDeleter)(nil).DeletePartitions), arg0, arg1)
}

// MockCloudAssetStorer is a mock of CloudAssetStorer interface
type MockCloudAssetStorer struct {
	ctrl     *gomock.Controller
	recorder *MockCloudAssetStorerMockRecorder
}

// MockCloudAssetStorerMockRecorder is the mock recorder for MockCloudAssetStorer
type MockCloudAssetStorerMockRecorder struct {
	mock *MockCloudAssetStorer
}

// NewMockCloudAssetStorer creates a new mock instance
func NewMockCloudAssetStorer(ctrl *gomock.Controller) *MockCloudAssetStorer {
	mock := &MockCloudAssetStorer{ctrl: ctrl}
	mock.recorder = &MockCloudAssetStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudAssetStorer) EXPECT() *MockCloudAssetStorerMockRecorder {
	return m.recorder
}

// Store mocks base method
func (m *MockCloudAssetStorer) Store(arg0 context.Context, arg1 domain.CloudAssetChanges) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockCloudAssetStorerMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockCloudAssetStorer)(nil).Store), arg0, arg1)
}

// MockCloudAssetByIPFetcher is a mock of CloudAssetByIPFetcher interface
type MockCloudAssetByIPFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockCloudAssetByIPFetcherMockRecorder
}

// MockCloudAssetByIPFetcherMockRecorder is the mock recorder for MockCloudAssetByIPFetcher
type MockCloudAssetByIPFetcherMockRecorder struct {
	mock *MockCloudAssetByIPFetcher
}

// NewMockCloudAssetByIPFetcher creates a new mock instance
func NewMockCloudAssetByIPFetcher(ctrl *gomock.Controller) *MockCloudAssetByIPFetcher {
	mock := &MockCloudAssetByIPFetcher{ctrl: ctrl}
	mock.recorder = &MockCloudAssetByIPFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudAssetByIPFetcher) EXPECT() *MockCloudAssetByIPFetcherMockRecorder {
	return m.recorder
}

// FetchByIP mocks base method
func (m *MockCloudAssetByIPFetcher) FetchByIP(arg0 context.Context, arg1 time.Time, arg2 string) ([]domain.CloudAssetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByIP", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.CloudAssetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByIP indicates an expected call of FetchByIP
func (mr *MockCloudAssetByIPFetcherMockRecorder) FetchByIP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByIP", reflect.TypeOf((*MockCloudAssetByIPFetcher)(nil).FetchByIP), arg0, arg1, arg2)
}

// MockCloudAssetByHostnameFetcher is a mock of CloudAssetByHostnameFetcher interface
type MockCloudAssetByHostnameFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockCloudAssetByHostnameFetcherMockRecorder
}

// MockCloudAssetByHostnameFetcherMockRecorder is the mock recorder for MockCloudAssetByHostnameFetcher
type MockCloudAssetByHostnameFetcherMockRecorder struct {
	mock *MockCloudAssetByHostnameFetcher
}

// NewMockCloudAssetByHostnameFetcher creates a new mock instance
func NewMockCloudAssetByHostnameFetcher(ctrl *gomock.Controller) *MockCloudAssetByHostnameFetcher {
	mock := &MockCloudAssetByHostnameFetcher{ctrl: ctrl}
	mock.recorder = &MockCloudAssetByHostnameFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudAssetByHostnameFetcher) EXPECT() *MockCloudAssetByHostnameFetcherMockRecorder {
	return m.recorder
}

// FetchByHostname mocks base method
func (m *MockCloudAssetByHostnameFetcher) FetchByHostname(arg0 context.Context, arg1 time.Time, arg2 string) ([]domain.CloudAssetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByHostname", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.CloudAssetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByHostname indicates an expected call of FetchByHostname
func (mr *MockCloudAssetByHostnameFetcherMockRecorder) FetchByHostname(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByHostname", reflect.TypeOf((*MockCloudAssetByHostnameFetcher)(nil).FetchByHostname), arg0, arg1, arg2)
}

// MockCloudAllAssetsByTimeFetcher is a mock of CloudAllAssetsByTimeFetcher interface
type MockCloudAllAssetsByTimeFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockCloudAllAssetsByTimeFetcherMockRecorder
}

// MockCloudAllAssetsByTimeFetcherMockRecorder is the mock recorder for MockCloudAllAssetsByTimeFetcher
type MockCloudAllAssetsByTimeFetcherMockRecorder struct {
	mock *MockCloudAllAssetsByTimeFetcher
}

// NewMockCloudAllAssetsByTimeFetcher creates a new mock instance
func NewMockCloudAllAssetsByTimeFetcher(ctrl *gomock.Controller) *MockCloudAllAssetsByTimeFetcher {
	mock := &MockCloudAllAssetsByTimeFetcher{ctrl: ctrl}
	mock.recorder = &MockCloudAllAssetsByTimeFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudAllAssetsByTimeFetcher) EXPECT() *MockCloudAllAssetsByTimeFetcherMockRecorder {
	return m.recorder
}

// FetchAll mocks base method
func (m *MockCloudAllAssetsByTimeFetcher) FetchAll(arg0 context.Context, arg1 time.Time, arg2, arg3 uint, arg4 string) ([]domain.CloudAssetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]domain.CloudAssetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAll indicates an expected call of FetchAll
func (mr *MockCloudAllAssetsByTimeFetcherMockRecorder) FetchAll(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAll", reflect.TypeOf((*MockCloudAllAssetsByTimeFetcher)(nil).FetchAll), arg0, arg1, arg2, arg3, arg4)
}
