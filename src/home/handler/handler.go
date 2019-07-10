package handler

import (
	"net/http"
	"sorabel/libraries"

	"github.com/labstack/echo"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, http.StatusOK, "successfully", "Hello Sorabel")
	}
}
