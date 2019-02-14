package scrapers

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/bigquery"
	"github.com/imdario/mergo"
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/models"
	"github.com/pcorbel/metaquery/cmd/utils"
)

// FieldScraper represents a field scraper instance
type FieldScraper struct {
	DB           *gorm.DB
	APIMetadata  *bigquery.TableMetadata
	Logger       *log.Logger
	ColumnNumber int
}

func (s *FieldScraper) scrap() {
	defer utils.TimeTrack(time.Now(), fmt.Sprintf("Scraping fields for table %s", s.APIMetadata.FullID), s.Logger)
	s.Logger.Infof("Scraping fields for table %s", s.APIMetadata.FullID)
	s.work(s.APIMetadata.Schema, 0, "")
}

func (s *FieldScraper) work(schema bigquery.Schema, level int, parent string) {

	// Iterate over fields
	for _, field := range schema {

		var mode string
		var parentName string

		switch {
		case field.Repeated:
			mode = "REPEATED"
		case field.Required:
			mode = "REQUIRED"
		default:
			mode = "OPTIONAL"
		}

		if parent != "" {
			parentName = parent + "."
		}

		// Load from DB
		fld := models.Field{}
		s.DB.FirstOrInit(&fld, models.Field{
			TableFullID: s.APIMetadata.FullID,
			Name:        parentName + field.Name,
		})

		// Merge data
		if err := mergo.Merge(&fld, models.Field{
			TableFullID:  s.APIMetadata.FullID,
			Name:         parentName + field.Name,
			Type:         fmt.Sprintf("%s", field.Type),
			Mode:         mode,
			Description:  field.Description,
			ColumnNumber: s.ColumnNumber,
			CreatedAt:    fld.CreatedAt,
		}, mergo.WithOverride); err != nil {
			s.Logger.Errorf("cannot merge metadata for field %s.%s ", s.APIMetadata.FullID, parentName+field.Name)
			return
		}

		// Save into DB
		s.DB.Save(&fld)

		s.ColumnNumber++

		// Recursive call for record type
		if field.Type == bigquery.RecordFieldType {
			s.work(field.Schema, level+1, parentName+field.Name)
		}
	}
}
