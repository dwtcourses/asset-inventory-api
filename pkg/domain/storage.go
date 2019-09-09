package domain

import (
	"context"
	"fmt"
	"time"
)

// PartitionGenerator is used to generate the next time-based partition
type PartitionGenerator interface {
	GeneratePartition(context.Context, time.Time, int) error
}

// PartitionsGetter is used to fetch a list of partitions
type PartitionsGetter interface {
	GetPartitions(context.Context) ([]Partition, error)
}

// PartitionsDeleter is used to drop partitions by name.
type PartitionsDeleter interface {
	DeletePartitions(context.Context, string) error
}

// CloudAssetStorer interface provides functions for inserting cloud assets
type CloudAssetStorer interface {
	Store(context.Context, CloudAssetChanges) error
}

// CloudAssetByIPFetcher fetches details for a cloud asset with a given IP address at a point in time
type CloudAssetByIPFetcher interface {
	FetchByIP(ctx context.Context, when time.Time, ipAddress string) ([]CloudAssetDetails, error)
}

// CloudAssetByHostnameFetcher fetches details for a cloud asset with a given hostname at a point in time
type CloudAssetByHostnameFetcher interface {
	FetchByHostname(ctx context.Context, when time.Time, hostname string) ([]CloudAssetDetails, error)
}

// CloudAssetAllByTimestampFetcher fetches details for all cloud assets based on limit and optional offset with a given point in time
type CloudAssetAllByTimestampFetcher interface {
	FetchAll(ctx context.Context, when time.Time, count uint, offset uint) ([]CloudAssetDetails, error)
}

// Partition represents a database partition with the specified time range
type Partition struct {
	Name      string
	CreatedAt time.Time
	Begin     time.Time
	End       time.Time
	Count     int
}

// PartitionConflict is used to indicate a partition exists which overlaps with a partition requested to be created
type PartitionConflict struct {
	Name string
}

func (e PartitionConflict) Error() string {
	return fmt.Sprintf("A partition already exists which overlaps with the requested partition, %s", e.Name)
}
