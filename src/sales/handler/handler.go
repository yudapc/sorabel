package handler

import (
	"net/http"
	"sorabel/helpers"
	"sorabel/src/sales/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data, err := model.GetSales(db)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func GetSalesDetail(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var sales model.Sales
		sales.ID = uint(id)
		data, err := model.GetSalesDetail(db, sales)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func CreateSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		var data model.Sales
		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}
		dataItem, err := model.CreateSales(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdateSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Sales
		data.ID = uint(id)

		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}

		_, err := model.EditSales(db, data)
		if err == nil {
			return helpers.ToJson(context, http.StatusOK, "data has been updated!", data)
		} else {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
	}
}

func DeleteSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Sales
		data.ID = uint(id)
		dataItem, err := model.DeleteSales(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "data has been deleted!", helpers.H{
			"id":             dataItem.ID,
			"date_time":      dataItem.DateTime,
			"invoice_number": dataItem.InvoiceNumber,
		})
	}
}

func GetSalesDetailItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		dataID := uint(id)
		data, err := model.GetSalesDetailItems(db, dataID)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}
