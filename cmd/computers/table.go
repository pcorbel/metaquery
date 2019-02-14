package computers

import (
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	log "github.com/sirupsen/logrus"
)

// TableComputer represents a table computer instance
type TableComputer struct {
	DB     *gorm.DB
	Config *configs.Config
	Table  models.Table
	Logger *log.Logger
}

func (c *TableComputer) compute() {

	c.Logger.Infof("Computing table %s", c.Table.FullID)

	var tableID string
	c.DB.Table("tables").Select("full_id").Where("full_id = ?", c.Table.FullID).Row().Scan(&tableID)

	c.DB.Table("fields").Where("table_full_id = ?", tableID).Count(&c.Table.FieldCount)
	c.DB.Table("partitions").Where("table_full_id = ?", tableID).Count(&c.Table.PartitionCount)
	c.DB.Table("partitions").Where("table_full_id = ?", tableID).Select("max(partitiontime) as latest_partition").Row().Scan(&c.Table.LatestPartition)
	c.DB.Table("partitions").Where("table_full_id = ?", tableID).Where("partitiontime = ?", c.Table.LatestPartition).Select("sum(count) as latest_partition_row_count").Row().Scan(&c.Table.LatestPartitionRowCount)

	c.DB.Save(&c.Table)
}
