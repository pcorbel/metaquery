package handlers

import (
	"net/http"

	"github.com/pcorbel/metaquery/cmd/utils"

	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// GetAllEntries godoc
// @Description GetAllEntries return the list of projects, datasets and tables into the DB
// @Tags db
// @Success 201
// @Router /entries [get]
func GetAllEntries(op *models.DBOperator) echo.HandlerFunc {
	return func(c echo.Context) error {
		responses := utils.Entries(op)
		return c.JSON(http.StatusOK, responses)
	}
}
