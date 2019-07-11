package handler

import (
	"net/http"
	"sorabel/helpers"

	"github.com/labstack/echo"
)

func Home() echo.HandlerFunc {
	return func(context echo.Context) error {
		return helpers.ToJson(context, http.StatusOK, "successfully", "Hello Sorabel")
	}
}
