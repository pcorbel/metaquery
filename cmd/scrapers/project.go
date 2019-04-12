package scrapers

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/bigquery"
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	"github.com/pcorbel/metaquery/cmd/utils"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// ProjectScraper represents a project scraper instance
type ProjectScraper struct {
	Context context.Context
	DB      *gorm.DB
	Client  *bigquery.Client
	Config  *configs.Config
	Project models.Project
	Logger  *log.Logger
}

func (s *ProjectScraper) scrap() {

	defer utils.TimeTrack(time.Now(), fmt.Sprintf("Scraping project %s", s.Project.FullID), s.Logger)

	s.Logger.Infof("Scraping project %s", s.Project.FullID)

	// Save into DB
	s.DB.Save(&s.Project)

	// Call to deeper level
	s.child()
}

func (s *ProjectScraper) child() {

	it := s.Client.Datasets(s.Context)

	var wg sync.WaitGroup

	// Iterate over childs
	for {
		ds, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			s.Logger.Errorf("Error iterating over datasets: %s ", err)
			break
		}

		// Load from DB
		dataset := models.Dataset{}
		s.DB.FirstOrInit(&dataset, models.Dataset{
			FullID: s.Project.FullID + ":" + ds.DatasetID,
			Name:   ds.DatasetID,
		})

		// Scrap
		datasetScraper := &DatasetScraper{
			Context:   s.Context,
			DB:        s.DB,
			Client:    s.Client,
			Config:    s.Config,
			Dataset:   dataset,
			Logger:    s.Logger,
			WaitGroup: &wg,
		}
		wg.Add(1)
		go datasetScraper.scrap()
	}
	wg.Wait()
}
