package homehandler

import (
	"net/http"
	"sorabel/libraries"

	"github.com/labstack/echo"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, libraries.H{
			"code": http.StatusOK,
			"data": "Hello Sorabel",
		})
	}
}
