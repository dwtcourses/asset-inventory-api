package v1

//go:generate mockgen -destination mock_storage_test.go -package v1 github.com/asecurityteam/asset-inventory-api/pkg/domain PartitionGenerator,PartitionsGetter,PartitionsDeleter,CloudAssetStorer,CloudAssetByIPFetcher,CloudAssetByHostnameFetcher,CloudAllAssetsByTimeFetcher
