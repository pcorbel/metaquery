package computers

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pcorbel/metaquery/cmd/configs"
	"github.com/pcorbel/metaquery/cmd/models"
	log "github.com/sirupsen/logrus"
)

// PartitionComputer represents a partition computer instance
type PartitionComputer struct {
	DB     *gorm.DB
	Config *configs.Config
	Table  models.Table
	Logger *log.Logger
}

func (c *PartitionComputer) compute() {

	c.Logger.Infof("Computing partitions for table %s", c.Table.FullID)
	var maxDate time.Time
	var minDate time.Time

	c.DB.Table("partitions").Where("table_full_id = ?", c.Table.FullID).Select("max(partitiontime) AS max_date").Row().Scan(&maxDate)
	c.DB.Table("partitions").Where("table_full_id = ?", c.Table.FullID).Select("min(partitiontime) AS min_date").Row().Scan(&minDate)
	query := fmt.Sprintf(`SELECT dates AS missing_partitions
		            FROM generate_series('%s', '%s', interval '1 day') AS dates
	               WHERE dates NOT IN (SELECT partitiontime FROM partitions WHERE table_full_id = '%s');`, minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"), c.Table.FullID)

	rows, _ := c.DB.Raw(query).Rows()
	defer rows.Close()
	for rows.Next() {
		var missingPartition time.Time
		rows.Scan(&missingPartition)
		partition := models.Partition{
			TableFullID:   c.Table.FullID,
			Partitiontime: missingPartition.UTC(),
			Count:         uint64(0),
		}
		c.DB.Save(&partition)
	}
}
