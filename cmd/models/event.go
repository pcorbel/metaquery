package models

import "time"

// Event is a struct containing event data
type Event struct {
	FullID    string    `json:"full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name"`
	Type      string    `json:"type" example:"dataset"`
	CreatedAt time.Time `json:"_created_at" example:"2000-01-01T00:00:00.000Z"`
	UpdatedAt time.Time `json:"_updated_at" example:"2000-01-01T00:00:00.000Z"`
	Message   string    `json:"message" example:"The partition my-awesome-gcp-project-name:my-awesome-dataset-name.my-awesome-table-name$YYYYMMDD has been updated"`
	Before    string    `json:"before" example:""`
	After     string    `json:"after" example:""`
}
