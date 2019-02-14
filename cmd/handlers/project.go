package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllProjects godoc
// @Description Get all the projects
// @Tags projects
// @Success 200 {array} models.Project
// @Router /projects [get]
func GetAllProjects(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var projects []models.Project
		db.Order("full_id").Find(&projects)
		return c.JSON(http.StatusOK, projects)
	}
}

// GetProject godoc
// @Description Get the specified project
// @Tags projects
// @Param project path string true "Project ID"
// @Success 200 {object} models.Project
// @Router /projects/{project} [get]
func GetProject(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var project models.Project
		db.Where("full_id = ?", c.Param("project")).Find(&project)
		return c.JSON(http.StatusOK, project)
	}
}
