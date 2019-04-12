package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllFields godoc
// @Description Get all the fields
// @Tags projects
// @Param project path string true "Project ID"
// @Param dataset path string true "Dataset ID"
// @Param table path string true "Table ID"
// @Success 200 {array} models.Field
// @Router /projects/{project}/datasets/{dataset}/tables/{table}/fields [get]
func GetAllFields(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var fields []models.Field
		db.Table("fields").Where("table_full_id LIKE ?", c.Param("project")+":"+c.Param("dataset")+"."+c.Param("table")).Order("column_number").Find(&fields)
		return c.JSON(http.StatusOK, fields)
	}
}
