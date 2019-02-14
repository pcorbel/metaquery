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
	"google.golang.org/api/iterator"
)

// DatasetScraper represents a dataset scraper instance
type DatasetScraper struct {
	Context   context.Context
	DB        *gorm.DB
	Client    *bigquery.Client
	Config    *configs.Config
	Dataset   models.Dataset
	Logger    *log.Logger
	WaitGroup *sync.WaitGroup
}

func (s *DatasetScraper) scrap() {

	defer s.WaitGroup.Done()

	// Filter out undeclared datasets
	if contains(s.Config.Datasets, s.Dataset.Name) {

		defer utils.TimeTrack(time.Now(), fmt.Sprintf("Scraping dataset %s", s.Dataset.FullID), s.Logger)

		s.Logger.Infof("Scraping dataset %s", s.Dataset.FullID)

		// Call API
		apiMetadata, err := s.Client.Dataset(s.Dataset.Name).Metadata(s.Context)
		if err != nil {
			s.Logger.Errorf("cannot get metadata for dataset %s ", s.Dataset.FullID)
			return
		}

		// Merge data
		if err = mergo.Merge(&s.Dataset, models.Dataset{
			Name:             getDatasetName(*apiMetadata),
			Description:      getDatasetDescription(*apiMetadata),
			Location:         getDatasetLocation(*apiMetadata),
			Expiration:       getDatasetExpiration(*apiMetadata),
			CreationTime:     getDatasetCreationTime(*apiMetadata),
			LastModifiedTime: getDatasetLastModifiedTime(*apiMetadata),
			Etag:             getDatasetEtag(*apiMetadata),
			FullID:           getDatasetFullID(*apiMetadata),
			Owners:           getDatasetOwners(*apiMetadata),
			Readers:          getDatasetReaders(*apiMetadata),
			Writers:          getDatasetWriters(*apiMetadata),
			Labels:           getDatasetLabels(*apiMetadata),
			CreatedAt:        s.Dataset.CreatedAt,
		}, mergo.WithOverride); err != nil {
			s.Logger.Errorf("cannot merge metadata for dataset %s ", s.Dataset.FullID)
			return
		}

		// Save into DB
		s.DB.Save(&s.Dataset)

		// Call to deeper level
		s.child()
	}
}

func (s *DatasetScraper) child() {

	var wg sync.WaitGroup
	it := s.Client.Dataset(s.Dataset.Name).Tables(s.Context)

	// Iterate over tables
	for {
		tb, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			s.Logger.Errorf("Error iterating over tables: %s", err)
			break
		}

		// Load from DB
		table := models.Table{}
		s.DB.FirstOrInit(&table, models.Table{
			FullID: s.Dataset.FullID + "." + tb.TableID,
			Name:   tb.TableID,
		})

		// Scrap
		tableScraper := &TableScraper{
			Context:   s.Context,
			DB:        s.DB,
			Client:    s.Client,
			Config:    s.Config,
			Table:     table,
			Logger:    s.Logger,
			WaitGroup: &wg,
		}
		wg.Add(1)
		go tableScraper.scrap()
	}
	wg.Wait()
}

func getDatasetName(metadata bigquery.DatasetMetadata) string {
	return strings.Split(metadata.FullID, ":")[1]
}

func getDatasetDescription(metadata bigquery.DatasetMetadata) string {
	return metadata.Description
}

func getDatasetLocation(metadata bigquery.DatasetMetadata) string {
	return metadata.Location
}

func getDatasetExpiration(metadata bigquery.DatasetMetadata) string {
	return metadata.DefaultTableExpiration.String()
}

func getDatasetCreationTime(metadata bigquery.DatasetMetadata) time.Time {
	return metadata.CreationTime.UTC()
}

func getDatasetLastModifiedTime(metadata bigquery.DatasetMetadata) time.Time {
	return metadata.LastModifiedTime.UTC()
}

func getDatasetEtag(metadata bigquery.DatasetMetadata) string {
	return metadata.ETag
}

func getDatasetFullID(metadata bigquery.DatasetMetadata) string {
	return metadata.FullID
}

func getDatasetOwners(metadata bigquery.DatasetMetadata) []string {
	var result []string
	for _, entry := range metadata.Access {
		if entry.Role == bigquery.OwnerRole {
			result = append(result, entry.Entity)
		}
	}
	return result
}

func getDatasetReaders(metadata bigquery.DatasetMetadata) []string {
	var result []string
	for _, entry := range metadata.Access {
		if entry.Role == bigquery.ReaderRole {
			result = append(result, entry.Entity)
		}
	}
	return result
}

func getDatasetWriters(metadata bigquery.DatasetMetadata) []string {
	var result []string
	for _, entry := range metadata.Access {
		if entry.Role == bigquery.WriterRole {
			result = append(result, entry.Entity)
		}
	}
	return result
}

func getDatasetLabels(metadata bigquery.DatasetMetadata) []string {
	var result []string
	for key, value := range metadata.Labels {
		result = append(result, fmt.Sprintf("%s:%s", key, value))
	}
	return result
}

func contains(datasetFilter []string, dataset string) bool {
	for _, datasetToKeep := range datasetFilter {
		if datasetToKeep == dataset {
			return true
		}
	}
	return false
}
