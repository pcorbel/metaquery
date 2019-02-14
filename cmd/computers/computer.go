package computers

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
)

// Computer represents a computer instance
type Computer struct {
	DB       *gorm.DB
	Config   *configs.Config
	Project  models.Project
	Datasets []models.Dataset
	Tables   []models.Table
	Logger   *log.Logger
}

// Compute calculate data based on scraped data
func (c *Computer) Compute() {

	c.Logger.Info("Instantiating Computer")

	projectComputer := &ProjectComputer{
		DB:     c.DB,
		Config: c.Config,
		Logger: c.Logger,
	}
	projectComputer.compute()

	c.DB.Where("full_id LIKE ?", c.Config.Project+":%").Find(&c.Datasets)
	for _, dataset := range c.Datasets {
		datasetComputer := &DatasetComputer{
			DB:      c.DB,
			Config:  c.Config,
			Dataset: dataset,
			Logger:  c.Logger,
		}
		datasetComputer.compute()

		c.DB.Where("full_id LIKE ?", dataset.FullID+".%").Find(&c.Tables)
		for _, table := range c.Tables {

			tableComputer := &TableComputer{
				DB:     c.DB,
				Config: c.Config,
				Table:  table,
				Logger: c.Logger,
			}
			tableComputer.compute()

			partitionComputer := &PartitionComputer{
				DB:     c.DB,
				Config: c.Config,
				Table:  table,
				Logger: c.Logger,
			}
			partitionComputer.compute()
		}
	}
	c.Logger.Info("Computer done")
}
