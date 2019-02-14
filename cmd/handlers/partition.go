package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllPartitions godoc
// @Description Get all the partitions
// @Tags projects
// @Param project path string true "Project ID"
// @Param dataset path string true "Dataset ID"
// @Param table path string true "Table ID"
// @Success 200 {array} models.Partition
// @Router /projects/{project}/datasets/{dataset}/tables/{table}/partitions [get]
func GetAllPartitions(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var partitions []models.Partition
		db.Table("partitions").Where("table_full_id LIKE ?", c.Param("project")+":"+c.Param("dataset")+"."+c.Param("table")).Order("partitiontime desc").Find(&partitions)
		return c.JSON(http.StatusOK, partitions)
	}
}
