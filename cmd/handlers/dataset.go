package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllDatasets godoc
// @Description Get all the datasets
// @Tags projects
// @Param project path string true "Project ID"
// @Success 200 {array} models.Dataset
// @Router /projects/{project}/datasets [get]
func GetAllDatasets(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var datasets []models.Dataset
		db.Where("full_id LIKE ?", c.Param("project")+":%").Order("full_id").Find(&datasets)
		return c.JSON(http.StatusOK, datasets)
	}
}

// GetDataset godoc
// @Description Get the specified dataset
// @Tags projects
// @Param project path string true "Project ID"
// @Param dataset path string true "Dataset ID"
// @Success 200 {object} models.Dataset
// @Router /projects/{project}/datasets/{dataset} [get]
func GetDataset(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataset models.Dataset
		db.Where("full_id LIKE ?", c.Param("project")+":"+c.Param("dataset")).Find(&dataset)
		return c.JSON(http.StatusOK, dataset)
	}
}
