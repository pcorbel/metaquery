package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pcorbel/metaquery/cmd/utils"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/bigquery"
	"github.com/pcorbel/metaquery/cmd/computers"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/handlers"
	"github.com/pcorbel/metaquery/cmd/models"
	"github.com/pcorbel/metaquery/cmd/scrapers"
	"github.com/tylerb/graceful"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/pcorbel/metaquery/cmd/docs"
	"github.com/robfig/cron"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	ctx        context.Context
	err        error
	db         *gorm.DB
	config     *configs.Config
	scraper    *scrapers.Scraper
	computer   *computers.Computer
	dbOperator *models.DBOperator
	logger     *log.Logger
)

func init() {
	logger = log.New()
	logger.SetFormatter(&utils.Formatter{})
	logger.SetLevel(log.InfoLevel)
	logger.Info("Starting metaquery server")
	ctx = context.Background()
	config = &configs.Config{Logger: logger}
	configFile, _ := os.LookupEnv("CONFIG_FILE")
	err = config.LoadFromFile(configFile)
	if err != nil {
		logger.Panic(err)
	}
	level, _ := log.ParseLevel(config.LogLevel)
	logger.SetLevel(level)
	config.PrintConfig()
	dbOperator = &models.DBOperator{
		DB:     db,
		Config: config,
		Logger: logger,
	}
	db, err = dbOperator.ConnectDB()
	if err != nil {
		logger.Panic(err)
	}
	dbOperator.DropDB()
	dbOperator.MigrateDB()
	client, err := bigquery.NewClient(ctx, config.Project)
	if err != nil {
		logger.Panic(err)
	}
	scraper = &scrapers.Scraper{
		Context: ctx,
		DB:      db,
		Config:  config,
		Client:  client,
		Logger:  logger,
	}
	computer = &computers.Computer{
		DB:     db,
		Config: config,
		Logger: logger,
	}
}

// @title GoBigQueryStats API
// @version 1.0
// @host localhost:8081
// @BasePath /api/v1
func main() {

	// Init echo
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Server.Addr = ":" + fmt.Sprint(config.ServerPort)

	// Add handlers
	v1 := e.Group("/api/v1")

	v1.GET("/projects", handlers.GetAllProjects(db))
	v1.GET("/projects/:project", handlers.GetProject(db))

	v1.GET("/projects/:project/datasets", handlers.GetAllDatasets(db))
	v1.GET("/projects/:project/datasets/:dataset", handlers.GetDataset(db))

	v1.GET("/projects/:project/datasets/:dataset/tables", handlers.GetAllTables(db))
	v1.GET("/projects/:project/datasets/:dataset/tables/:table", handlers.GetTable(db))

	v1.GET("/projects/:project/datasets/:dataset/tables/:table/partitions", handlers.GetAllPartitions(db))
	v1.GET("/projects/:project/datasets/:dataset/tables/:table/fields", handlers.GetAllFields(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Entries endpoint
	v1.GET("/entries", handlers.GetAllEntries(dbOperator))

	// Events endpoint
	v1.GET("/events", handlers.GetAllEvents(db))

	// Manual ops endpoint
	v1.POST("/drop", handlers.Drop(dbOperator))
	v1.POST("/migrate", handlers.Migrate(dbOperator))
	v1.POST("/scrap", handlers.Scrap(scraper))
	v1.POST("/compute", handlers.Compute(computer))

	// Add Frontend
	e.Static("/", config.WebDirectory)

	// Schedule cron refresh
	c := cron.NewWithLocation(time.UTC)
	c.AddFunc(config.CronRefresh, func() {
		logger.Info("Refreshing data...")
		scraper.Scrap()
		computer.Compute()
		for _, entry := range c.Entries() {
			logger.Infof("Next refresh scheduled at: %s", entry.Next)
		}
	})
	logger.Info("Starting cron server")
	c.Start()

	// Scraping at startup
	logger.Info("Refreshing data...")
	scraper.Scrap()
	computer.Compute()

	// Start server
	logger.Info("Starting echo server")
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
