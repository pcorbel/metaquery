package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllTables godoc
// @Description Get all the tables
// @Tags projects
// @Param project path string true "Project ID"
// @Param dataset path string true "Dataset ID"
// @Success 200 {array} models.Table
// @Router /projects/{project}/datasets/{dataset}/tables [get]
func GetAllTables(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tables []models.Table
		db.Where("full_id LIKE ?", c.Param("project")+":"+c.Param("dataset")+".%").Order("full_id").Find(&tables)
		return c.JSON(http.StatusOK, tables)
	}
}

// GetTable godoc
// @Description Get the specified table
// @Tags projects
// @Param project path string true "Project ID"
// @Param dataset path string true "Dataset ID"
// @Param table path string true "Table ID"
// @Success 200 {object} models.Table
// @Router /projects/{project}/datasets/{dataset}/tables/{table} [get]
func GetTable(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var table models.Table
		db.Where("full_id LIKE ?", c.Param("project")+":"+c.Param("dataset")+"."+c.Param("table")).Find(&table)
		return c.JSON(http.StatusOK, table)
	}
}
