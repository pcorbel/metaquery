package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Field is a struct containing Field data
type Field struct {
	TableFullID  string    `json:"table_full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name.my-awesome-table-name" gorm:"primary_key"`
	Name         string    `json:"name" example:"my-awesome-field-name" gorm:"primary_key"`
	CreatedAt    time.Time `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt    time.Time `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Type         string    `json:"type" example:"INTEGER"`
	Mode         string    `json:"mode" example:"NULLABLE"`
	Description  string    `json:"description" example:"This is my awesome field"`
	ColumnNumber int       `json:"-"`
}

// AfterCreate hook call to fire up a creation event
func (field *Field) AfterCreate(db *gorm.DB) {
	FullID := fmt.Sprintf("%s.%s", field.TableFullID, field.Name)
	event := Event{
		FullID:  FullID,
		Type:    "field",
		Message: fmt.Sprintf("The field %s has been created", FullID),
		Before:  fmt.Sprintf("%+v", nil),
		After:   fmt.Sprintf("%+v", field),
	}
	db.Save(&event)
}
