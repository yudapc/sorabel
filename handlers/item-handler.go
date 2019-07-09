package handlers

import (
	"database/sql"
	"net/http"
	"sorabel/libraries"
	"sorabel/models"

	"github.com/labstack/echo"
)

func GetItems(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		items := models.GetItems(db)
		return c.JSON(http.StatusOK, libraries.H{
			"code": http.StatusOK,
			"data": items,
		})
	}
}
