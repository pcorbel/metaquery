package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Partition is a struct containing Partition data
type Partition struct {
	TableFullID   string    `json:"table_full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name.my-awesome-table-name" gorm:"primary_key"`
	Partitiontime time.Time `json:"partitiontime" example:"2000-01-01 00:00:00 +0000 UTC" gorm:"primary_key"`
	CreatedAt     time.Time `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt     time.Time `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Count         uint64    `json:"count" example:"10000"`
}

// AfterCreate hook call to fire up a creation event
func (partition *Partition) AfterCreate(db *gorm.DB) {
	FullID := fmt.Sprintf("%s$%s", partition.TableFullID, partition.Partitiontime.Format("20060102"))
	event := Event{
		FullID:  FullID,
		Type:    "partition",
		Message: fmt.Sprintf("The partition %s has been created", FullID),
		Before:  fmt.Sprintf("%+v", nil),
		After:   fmt.Sprintf("%+v", partition),
	}
	db.Save(&event)
}
