package handler

import (
	"net/http"
	"sorabel/helpers"
	"sorabel/src/report/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ItemValueReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data := model.GenerateItemValueReport(db)
		return helpers.ToJson(context, http.StatusOK, "succesfully", data)
	}
}

func SalesReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data := model.GenerateSalesReport(db)
		return helpers.ToJson(context, http.StatusOK, "succesfully", data)
	}
}
