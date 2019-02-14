package computers

import (
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	log "github.com/sirupsen/logrus"
)

// DatasetComputer represents a dataset computer instance
type DatasetComputer struct {
	DB      *gorm.DB
	Config  *configs.Config
	Dataset models.Dataset
	Logger  *log.Logger
}

func (c *DatasetComputer) compute() {

	c.Logger.Infof("Computing dataset %s", c.Dataset.FullID)

	var tablesID []string
	rows, _ := c.DB.Table("tables").Select("full_id").Where("full_id LIKE ?", c.Dataset.FullID+"%").Rows()
	defer rows.Close()
	for rows.Next() {
		var fullID string
		rows.Scan(&fullID)
		tablesID = append(tablesID, fullID)
	}

	c.DB.Table("tables").Where("full_id IN (?)", tablesID).Count(&c.Dataset.TableCount)
	c.DB.Table("tables").Where("full_id IN (?)", tablesID).Select("sum(row_count) as total").Row().Scan(&c.Dataset.RowCount)
	c.DB.Table("tables").Where("full_id IN (?)", tablesID).Select("sum(byte_count) as total").Row().Scan(&c.Dataset.ByteCount)
	c.DB.Table("fields").Where("table_full_id IN (?)", tablesID).Count(&c.Dataset.FieldCount)
	c.DB.Table("partitions").Where("table_full_id IN (?)", tablesID).Count(&c.Dataset.PartitionCount)
	c.DB.Table("partitions").Where("table_full_id IN (?)", tablesID).Select("max(partitiontime) as latest_partition").Row().Scan(&c.Dataset.LatestPartition)
	c.DB.Table("partitions").Where("table_full_id IN (?)", tablesID).Where("partitiontime = ?", c.Dataset.LatestPartition).Select("sum(count) as latest_partition_row_count").Row().Scan(&c.Dataset.LatestPartitionRowCount)

	c.DB.Save(&c.Dataset)
}
