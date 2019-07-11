package handler

import (
	"net/http"
	"sorabel/libraries"
	"sorabel/src/sales/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := model.GetSales(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetSalesDetail(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var sales model.Sales
		sales.ID = uint(id)
		data, err := model.GetSalesDetail(db, sales)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreateSales(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data model.Sales
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}
		dataItem, err := model.CreateSales(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "data has been created!", dataItem)
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
