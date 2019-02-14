package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Dataset is a struct containing Dataset data
type Dataset struct {
	FullID                  string         `json:"full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name" gorm:"primary_key"`
	CreatedAt               time.Time      `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt               time.Time      `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Name                    string         `json:"name" example:"my-awesome-dataset-name"`
	Description             string         `json:"description" example:"This is my awesome dataset"`
	Location                string         `json:"location" example:"EU"`
	Expiration              string         `json:"expiration" example:"3600s"`
	CreationTime            time.Time      `json:"creation_time" example:"2000-01-01T00:00:00.000Z"`
	LastModifiedTime        time.Time      `json:"last_modified_time" example:"2000-01-01T00:00:00.000Z"`
	Etag                    string         `json:"etag" example:"QiKZCjFNeH22n9b0iC/ZoQ=="`
	Owners                  pq.StringArray `json:"owners" gorm:"type:text[]" example:"[projectOwners]"`
	Writers                 pq.StringArray `json:"writers" gorm:"type:text[]" example:"[projectWriters]"`
	Readers                 pq.StringArray `json:"readers" gorm:"type:text[]" example:"[projectReaders]"`
	Labels                  pq.StringArray `json:"labels" gorm:"type:text[]" example:"[key:value]"`
	TableCount              int64          `json:"table_count" example:"1000"`
	PartitionCount          int64          `json:"partition_count" example:"10000"`
	FieldCount              int64          `json:"field_count" example:"10000"`
	RowCount                int64          `json:"row_count" example:"10000"`
	ByteCount               int64          `json:"byte_count" example:"100000"`
	LatestPartition         time.Time      `json:"latest_partition" example:"2000-01-01T00:00:00.000Z"`
	LatestPartitionRowCount int64          `json:"latest_partition_row_count" example:"1000000"`
}

// AfterCreate hook call to fire up a creation event
func (dataset *Dataset) AfterCreate(db *gorm.DB) {
	event := Event{
		FullID:  dataset.FullID,
		Type:    "dataset",
		Message: fmt.Sprintf("The dataset %s has been created", dataset.FullID),
		Before:  fmt.Sprintf("%+v", nil),
		After:   fmt.Sprintf("%+v", dataset),
	}
	db.Save(&event)
}
