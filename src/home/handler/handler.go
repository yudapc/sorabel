package handler

import (
	"net/http"
	"sorabel/libraries"

	"github.com/labstack/echo"
)

func Home() echo.HandlerFunc {
	return func(context echo.Context) error {
		return libraries.ToJson(context, http.StatusOK, "successfully", "Hello Sorabel")
	}
}
