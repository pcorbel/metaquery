package models

// Entry is a struct which define a project, a dataset or a table from the DB
type Entry struct {
	FullID string `json:"full_id" example:"my-awesome-gcp-project-name:my-awesome-dataset-name"`
	Type   string `json:"type" example:"dataset"`
}
