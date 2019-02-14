package scrapers

import (
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	"github.com/pcorbel/metaquery/cmd/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// Scraper represents a scraper instance
type Scraper struct {
	Context context.Context
	DB      *gorm.DB
	Client  *bigquery.Client
	Config  *configs.Config
	Logger  *log.Logger
}

// Scrap get data from BigQuery API
func (s *Scraper) Scrap() {

	defer utils.TimeTrack(time.Now(), "Scraper", s.Logger)

	s.Logger.Info("Instantiating Scraper")

	// Load from DB
	project := models.Project{}
	s.DB.FirstOrCreate(&project, models.Project{FullID: s.Config.Project})

	// Call to deeper level
	s.child(project)

	s.Logger.Info("Scraper done")
}

func (s *Scraper) child(project models.Project) {
	// Scrap
	projectScraper := &ProjectScraper{
		Context: s.Context,
		DB:      s.DB,
		Client:  s.Client,
		Config:  s.Config,
		Project: project,
		Logger:  s.Logger,
	}
	projectScraper.scrap()
}
