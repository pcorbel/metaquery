package models

import (
	"time"

	"github.com/lib/pq"
)

// Table is a struct containing Table data
type Table struct {
	FullID                  string         `json:"full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name.my-awesome-table-name" gorm:"primary_key"`
	CreatedAt               time.Time      `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt               time.Time      `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Name                    string         `json:"name" example:"my-awesome-table-name"`
	Description             string         `json:"description" example:"This is my awesome table"`
	LegacySQL               bool           `json:"legacySQL" example:"false"`
	CreationTime            time.Time      `json:"creation_time" example:"2000-01-01T00:00:00.000Z"`
	LastModifiedTime        time.Time      `json:"last_modified_time" example:"2000-01-01T00:00:00.000Z"`
	Etag                    string         `json:"etag" example:"QiKZCjFNeH22n9b0iC/ZoQ=="`
	TimePartitioning        string         `json:"time_partitioning" example:"DAY"`
	PartitionCount          int64          `json:"partition_count" example:"10000"`
	FieldCount              int64          `json:"field_count" example:"10000"`
	RowCount                uint64         `json:"row_count" example:"10000"`
	ByteCount               uint64         `json:"byte_count" example:"100000"`
	Labels                  pq.StringArray `json:"labels" gorm:"type:text[]" example:"[key:value]"`
	LatestPartition         time.Time      `json:"latest_partition" example:"2000-01-01T00:00:00.000Z"`
	LatestPartitionRowCount int64          `json:"latest_partition_row_count" example:"1000000"`
}
