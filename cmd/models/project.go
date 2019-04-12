package models

import (
	"time"
)

// Project is a struct containing Project stats
type Project struct {
	FullID                  string    `json:"full_id" example:"my-awesome-gcp-project-name" gorm:"primary_key"`
	CreatedAt               time.Time `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt               time.Time `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	DatasetCount            int64     `json:"dataset_count" example:"100"`
	TableCount              int64     `json:"table_count" example:"1000"`
	PartitionCount          int64     `json:"partition_count" example:"10000"`
	FieldCount              int64     `json:"field_count" example:"10000"`
	RowCount                int64     `json:"row_count" example:"10000"`
	ByteCount               int64     `json:"byte_count" example:"100000"`
	LatestPartition         time.Time `json:"latest_partition" example:"2000-01-01T00:00:00.000Z"`
	LatestPartitionRowCount int64     `json:"latest_partition_row_count" example:"1000000"`
}
