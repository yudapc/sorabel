package handler

import (
	"sorabel/libraries"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 200, "/sales", nil)
	}
}

func GetSalesDetail(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 200, "/sales/:id", nil)
	}
}

func CreateSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 201, "create /sales", nil)
	}
}

func UpdateSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 200, "update /sales/:id", nil)
	}
}

func DeleteSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 200, "deelete /sales/:id", nil)
	}
}

func GetSalesDetailItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return libraries.ToJson(c, 200, "/sales/:id/items", nil)
	}
}
