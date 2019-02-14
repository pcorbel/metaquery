package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllEvents godoc
// @Description GetAllEvents return the list of events which have occured on the DB
// @Tags db
// @Success 201
// @Router /events [get]
func GetAllEvents(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var events []models.Event
		db.Order("created_at").Find(&events)
		return c.JSON(http.StatusOK, events)
	}
}
