package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pcorbel/metaquery/cmd/configs"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBOperator represents a database operator instance
type DBOperator struct {
	DB     *gorm.DB
	Config *configs.Config
	Logger *log.Logger
}

// ConnectDB connect to the database
func (op *DBOperator) ConnectDB() (*gorm.DB, error) {
	op.Logger.Infof("Connecting to Postgre database on %s:%d", op.Config.DBHostname, op.Config.DBPort)
	var err error
	op.DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=postgres dbname=postgres password=postgres sslmode=disable", op.Config.DBHostname, op.Config.DBPort))
	if err != nil {
		return nil, errors.Wrap(err, "cannot open DB")
	}
	op.DB.LogMode(op.Logger.GetLevel() == log.TraceLevel)
	return op.DB, nil
}

// MigrateDB create the underlying tables if they don't exist
func (op *DBOperator) MigrateDB() {
	op.Logger.Infof("Migrating Postgre database on %s:%d", op.Config.DBHostname, op.Config.DBPort)
	op.DB.AutoMigrate(&Project{})
	op.DB.AutoMigrate(&Dataset{})
	op.DB.AutoMigrate(&Table{})
	op.DB.AutoMigrate(&Field{})
	op.DB.AutoMigrate(&Partition{})
	op.DB.AutoMigrate(&Event{})
}

// DropDB drop the underlying tables if they exist
func (op *DBOperator) DropDB() {
	if op.Config.DBReset {
		op.Logger.Infof("Dropping Postgre database on %s:%d", op.Config.DBHostname, op.Config.DBPort)
		op.DB.DropTableIfExists(&Project{})
		op.DB.DropTableIfExists(&Dataset{})
		op.DB.DropTableIfExists(&Table{})
		op.DB.DropTableIfExists(&Field{})
		op.DB.DropTableIfExists(&Partition{})
		op.DB.DropTableIfExists(&Event{})
	}
}
