package handlers

import (
	"net/http"

	"github.com/pcorbel/metaquery/cmd/computers"
	"github.com/pcorbel/metaquery/cmd/scrapers"

	"github.com/labstack/echo"
	"github.com/pcorbel/metaquery/cmd/models"
)

// Drop godoc
// @Description Drop the DB
// @Tags db
// @Success 201
// @Router /drop [post]
func Drop(op *models.DBOperator) echo.HandlerFunc {
	return func(c echo.Context) error {
		op.DropDB()
		return c.NoContent(http.StatusOK)
	}
}

// Migrate godoc
// @Description Migrate the DB
// @Tags db
// @Success 201
// @Router /migrate [post]
func Migrate(op *models.DBOperator) echo.HandlerFunc {
	return func(c echo.Context) error {
		op.MigrateDB()
		return c.NoContent(http.StatusOK)
	}
}

// Scrap godoc
// @Description Scrap data to fill the DB
// @Tags db
// @Success 201
// @Router /scrap [post]
func Scrap(scraper *scrapers.Scraper) echo.HandlerFunc {
	return func(c echo.Context) error {
		scraper.Scrap()
		return c.NoContent(http.StatusOK)
	}
}

// Compute godoc
// @Description Compute the data
// @Tags db
// @Success 201
// @Router /scrap [post]
func Compute(computer *computers.Computer) echo.HandlerFunc {
	return func(c echo.Context) error {
		computer.Compute()
		return c.NoContent(http.StatusOK)
	}
}
