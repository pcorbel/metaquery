package computers

import (
	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	log "github.com/sirupsen/logrus"
)

// ProjectComputer represents a project computer instance
type ProjectComputer struct {
	DB      *gorm.DB
	Config  *configs.Config
	Project models.Project
	Logger  *log.Logger
}

func (c *ProjectComputer) compute() {

	c.DB.Where("full_id = ?", c.Config.Project).Find(&c.Project)

	c.Logger.Infof("Computing project %s", c.Project.FullID)

	c.DB.Model(&models.Dataset{}).Count(&c.Project.DatasetCount)
	c.DB.Model(&models.Table{}).Count(&c.Project.TableCount)
	c.DB.Model(&models.Partition{}).Count(&c.Project.PartitionCount)
	c.DB.Model(&models.Field{}).Count(&c.Project.FieldCount)
	c.DB.Table("tables").Select("sum(row_count) as total").Row().Scan(&c.Project.RowCount)
	c.DB.Table("tables").Select("sum(byte_count) as total").Row().Scan(&c.Project.ByteCount)
	c.DB.Table("partitions").Select("max(partitiontime) as latest_partition").Row().Scan(&c.Project.LatestPartition)
	c.DB.Table("partitions").Where("partitiontime = ?", c.Project.LatestPartition).Select("sum(count) as latest_partition_row_count").Row().Scan(&c.Project.LatestPartitionRowCount)

	c.DB.Save(&c.Project)
}
