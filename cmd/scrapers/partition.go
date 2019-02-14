package scrapers

import (
	"fmt"
	"strings"
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

// PartitionScraper represents a partition scraper instance
type PartitionScraper struct {
	Context     context.Context
	DB          *gorm.DB
	Client      *bigquery.Client
	Config      *configs.Config
	APIMetadata *bigquery.TableMetadata
	Logger      *log.Logger
}

// Row is a struct to be merged with BigQuery results
type Row struct {
	Partitiontime time.Time
	Count         int64
}

func (s *PartitionScraper) scrap() {

	defer utils.TimeTrack(time.Now(), fmt.Sprintf("Scraping partitions for table %s", s.APIMetadata.FullID), s.Logger)
	s.Logger.Infof("Scraping partitions for table %s", s.APIMetadata.FullID)

	// If table isn't partitionned
	if s.APIMetadata.TimePartitioning == nil {
		// Load partition from DB
		partition := models.Partition{}
		s.DB.FirstOrInit(&partition, models.Partition{
			TableFullID:   s.APIMetadata.FullID,
			Partitiontime: time.Unix(0, 0).UTC(),
		})

		// Merge data
		if err := mergo.Merge(&partition, models.Partition{
			TableFullID:   s.APIMetadata.FullID,
			Partitiontime: time.Unix(0, 0).UTC(),
			Count:         uint64(s.APIMetadata.NumRows),
		}, mergo.WithOverride); err != nil {
			s.Logger.Errorf("cannot merge metadata for partition %s$%s ", s.APIMetadata.FullID, time.Unix(0, 0).UTC())
			return
		}

		// Save into DB
		s.DB.Save(&partition)
		return
	}

	// Wait until slot on BigQuery are available
	waitForSlot := true
	for waitForSlot {
		itJob := s.Client.Jobs(s.Context)
		itJob.State = bigquery.Running
		var runningJobs int
		for {
			_, err := itJob.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}
				s.Logger.Errorf("Error iterating over jobs: %s", err)
				break
			}
			runningJobs++
		}
		if runningJobs < s.Config.MaxClientConcurrency {
			waitForSlot = false
		} else {
			s.Logger.Infof("Slot not yet available for table %s (%d/%d)", s.APIMetadata.FullID, runningJobs, s.Config.MaxClientConcurrency)
			time.Sleep(1 * time.Second)
		}
	}

	// Query to be executed on BigQuery
	query := s.Client.Query(fmt.Sprintf(`
       #standardSQL
SELECT _PARTITIONTIME AS partitiontime,
       COUNT(1) AS count 
  FROM`+"`%s`"+`
 WHERE _PARTITIONTIME IS NOT NULL
 GROUP BY 1
 ORDER BY 1 DESC`, strings.Split(s.APIMetadata.FullID, ":")[1]))

	it, err := query.Read(s.Context)
	if err != nil {
		s.Logger.Errorf("cannot execute query: %s", err)
		s.Logger.Errorf("api metadat was: %+v", s.APIMetadata)
		return
	}

	// Iterate over rows
	for {
		var row Row
		err := it.Next(&row)
		if err != nil {
			if err == iterator.Done {
				break
			}
			s.Logger.Errorf("Error iterating over rows: %s ", err)
			break
		}

		// Load from DB
		partition := models.Partition{}
		s.DB.FirstOrInit(&partition, models.Partition{
			TableFullID:   s.APIMetadata.FullID,
			Partitiontime: row.Partitiontime.UTC(),
		})

		// Merge data
		if err := mergo.Merge(&partition, models.Partition{
			TableFullID:   s.APIMetadata.FullID,
			Partitiontime: row.Partitiontime.UTC(),
			Count:         uint64(row.Count),
			CreatedAt:     partition.CreatedAt,
		}, mergo.WithOverride); err != nil {
			s.Logger.Errorf("cannot merge metadata for partition %s.%s ", s.APIMetadata.FullID, row.Partitiontime.UTC())
			return
		}

		// Save into DB
		s.DB.Save(&partition)
	}
}
