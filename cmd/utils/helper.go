package utils

import (
	"regexp"
)

// ParseFullID parse a table spec in the form project:dataset.table
func ParseFullID(fullID string) (string, string, string) {
	r := regexp.MustCompile(`^(.*?):(.*?)(?:\.(.*))?$`)
	results := r.FindStringSubmatch(fullID)
	return results[1], results[2], results[3]
}
