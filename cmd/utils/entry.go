package utils

import (
	"github.com/pcorbel/metaquery/cmd/models"
)

// Entries return the list of projects, datasets and tables into the DB
func Entries(op *models.DBOperator) []models.Entry {
	op.Logger.Debugf("Fetching index from the DB")

	var entries []models.Entry

	// Get projects
	rows, _ := op.DB.Table("projects").Select("full_id").Rows()
	defer rows.Close()
	for rows.Next() {
		var fullID string
		rows.Scan(&fullID)
		entries = append(entries, models.Entry{
			FullID: fullID,
			Type:   "project"})
	}

	// Get datasets
	rows, _ = op.DB.Table("datasets").Select("full_id").Rows()
	for rows.Next() {
		var fullID string
		rows.Scan(&fullID)
		entries = append(entries, models.Entry{
			FullID: fullID,
			Type:   "dataset"})
	}

	// Get tables
	rows, _ = op.DB.Table("tables").Select("full_id").Rows()
	for rows.Next() {
		var fullID string
		rows.Scan(&fullID)
		entries = append(entries, models.Entry{
			FullID: fullID,
			Type:   "table"})
	}

	return entries
}
