package models

import (
	"time"
)

// Partition is a struct containing Partition data
type Partition struct {
	TableFullID   string    `json:"table_full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name.my-awesome-table-name" gorm:"primary_key"`
	Partitiontime time.Time `json:"partitiontime" example:"2000-01-01 00:00:00 +0000 UTC" gorm:"primary_key"`
	CreatedAt     time.Time `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt     time.Time `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Count         uint64    `json:"count" example:"10000"`
}
