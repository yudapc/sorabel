package purchasehandler

import (
	"net/http"
	"sorabel/libraries"
	purchasemodel "sorabel/src/purchase/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetPurchases(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := purchasemodel.GetPurchases(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetPurchaseDetail(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var purchase purchasemodel.Purchase
		purchase.ID = uint(id)
		data, err := purchasemodel.GetPurchaseDetail(db, purchase)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data purchasemodel.Purchase
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}
		dataItem, err := purchasemodel.CreatePurchase(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var data purchasemodel.Purchase
		data.ID = uint(id)

		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}

		_, err := purchasemodel.EditPurchase(db, data)
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
		var data purchasemodel.Purchase
		data.ID = uint(id)
		dataItem, err := purchasemodel.DeletePurchase(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "data has been deleted!", dataItem)
	}
}
