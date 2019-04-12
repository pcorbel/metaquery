package scrapers

import (
	"fmt"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/bigquery"
	"github.com/imdario/mergo"
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	"github.com/pcorbel/metaquery/cmd/utils"
	"golang.org/x/net/context"
)

// TableScraper represents a table scraper instance
type TableScraper struct {
	Context   context.Context
	DB        *gorm.DB
	Client    *bigquery.Client
	Config    *configs.Config
	Table     models.Table
	Logger    *log.Logger
	WaitGroup *sync.WaitGroup
}

func (s *TableScraper) scrap() {

	defer s.WaitGroup.Done()
	defer utils.TimeTrack(time.Now(), fmt.Sprintf("Scraping table %s", s.Table.FullID), s.Logger)

	s.Logger.Infof("Scraping table %s", s.Table.FullID)

	// Call API
	_, datasetName, tableName := utils.ParseFullID(s.Table.FullID)
	apiMetadata, err := s.Client.Dataset(datasetName).Table(tableName).Metadata(s.Context)
	if err != nil {
		s.Logger.Errorf("cannot get metadata for table %s ", s.Table.FullID)
		return
	}

	// Merge data
	if err = mergo.Merge(&s.Table, models.Table{
		Name:             getTableName(*apiMetadata),
		Description:      getTableDescription(*apiMetadata),
		LegacySQL:        getTableLegacySQL(*apiMetadata),
		CreationTime:     getTableCreationTime(*apiMetadata),
		LastModifiedTime: getTableLastModifiedTime(*apiMetadata),
		Etag:             getTableEtag(*apiMetadata),
		FullID:           getTableFullID(*apiMetadata),
		TimePartitioning: getTableTimePartitioning(*apiMetadata),
		RowCount:         getTableRowCount(*apiMetadata),
		ByteCount:        getTableByteCount(*apiMetadata),
		Labels:           getTableLabels(*apiMetadata),
		CreatedAt:        s.Table.CreatedAt,
	}, mergo.WithOverride); err != nil {
		s.Logger.Errorf("cannot get metadata for table %s ", s.Table.FullID)
		return
	}

	// Save into DB
	s.DB.Save(&s.Table)

	// Call to deeper level
	s.childField(apiMetadata)
	s.childPartition(apiMetadata)
}

func (s *TableScraper) childPartition(apiMetadata *bigquery.TableMetadata) {
	// Scrap
	partitionScraper := &PartitionScraper{
		Context:     s.Context,
		DB:          s.DB,
		Client:      s.Client,
		Config:      s.Config,
		APIMetadata: apiMetadata,
		Logger:      s.Logger,
	}
	partitionScraper.scrap()
}

func (s *TableScraper) childField(apiMetadata *bigquery.TableMetadata) {
	// Scrap
	fieldScraper := &FieldScraper{
		DB:           s.DB,
		APIMetadata:  apiMetadata,
		Logger:       s.Logger,
		ColumnNumber: 0,
	}
	fieldScraper.scrap()
}

func getTableName(metadata bigquery.TableMetadata) string {
	return strings.Split(metadata.FullID, ".")[1]
}

func getTableDescription(metadata bigquery.TableMetadata) string {
	return metadata.Description
}

func getTableLegacySQL(metadata bigquery.TableMetadata) bool {
	return metadata.UseLegacySQL
}

func getTableCreationTime(metadata bigquery.TableMetadata) time.Time {
	return metadata.CreationTime.UTC()
}

func getTableLastModifiedTime(metadata bigquery.TableMetadata) time.Time {
	return metadata.LastModifiedTime.UTC()
}

func getTableEtag(metadata bigquery.TableMetadata) string {
	return metadata.ETag
}

func getTableFullID(metadata bigquery.TableMetadata) string {
	return metadata.FullID
}

func getTableTimePartitioning(metadata bigquery.TableMetadata) string {
	var result string
	if metadata.TimePartitioning != nil {
		if metadata.TimePartitioning.Field == "" {
			result = "DAY"
		}
	} else {
		result = "TOTAL"
	}
	return result
}

func getTableByteCount(metadata bigquery.TableMetadata) uint64 {
	return uint64(metadata.NumBytes)
}

func getTableRowCount(metadata bigquery.TableMetadata) uint64 {
	return uint64(metadata.NumRows)
}

func getTableLabels(metadata bigquery.TableMetadata) []string {
	var result []string
	for key, value := range metadata.Labels {
		result = append(result, fmt.Sprintf("%s:%s", key, value))
	}
	return result
}
