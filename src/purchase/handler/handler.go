package handler

import (
	"net/http"
	"sorabel/libraries"
	"sorabel/src/purchase/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetPurchases(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := model.GetPurchases(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetPurchaseDetail(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var purchase model.Purchase
		paramID := uint(id)
		purchase.ID = paramID
		data, err := model.GetPurchaseDetail(db, purchase)
		purchaseDetails, _ := model.GetPurchaseDetailItems(db, paramID)
		data.PurchaseDetails = purchaseDetails
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data model.Purchase
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}
		dataItem, err := model.CreatePurchase(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var data model.Purchase
		data.ID = uint(id)

		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}

		_, err := model.EditPurchase(db, data)
		if err == nil {
			return libraries.ToJson(c, http.StatusOK, "data has been updated!", data)
		} else {
			return err
		}
	}
}

func DeletePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var data model.Purchase
		data.ID = uint(id)
		dataItem, err := model.DeletePurchase(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "data has been deleted!", dataItem)
	}
}

func GetPurchaseDetailItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		purchaseID := uint(id)
		data, err := model.GetPurchaseDetailItems(db, purchaseID)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}
