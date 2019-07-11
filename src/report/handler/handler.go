package handler

import (
	"sorabel/helpers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ItemValueReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		return helpers.ToJson(context, 200, "item value report", nil)
	}
}

func SalesReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		return helpers.ToJson(context, 200, "sales report", nil)
	}
}
