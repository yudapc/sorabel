package purchasehandler

import (
	"database/sql"
	"net/http"
	"sorabel/libraries"
	purchasemodel "sorabel/src/purchase/model"
	"strconv"

	"github.com/labstack/echo"
)

func GetPurchases(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := purchasemodel.GetPurchases(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetPurchaseDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := purchasemodel.GetPurchaseDetail(db, c.Param("id"))
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreatePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data purchasemodel.Purchase
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed!", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed!", errValidate.Error())
		}
		id, err := purchasemodel.CreatePurchase(db, data.DateTime, data.ReceiptNumber)
		dataID := int(id)
		data.ID = dataID
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "purchase has been created!", data)
	}
}

func UpdatePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data purchasemodel.Purchase
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}
		id, err := purchasemodel.EditPurchase(db, c.Param("id"), data.DateTime, data.ReceiptNumber)
		dataID := int(id)
		data.ID = dataID
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "purchase has been created!", data)
	}
}

func DeletePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := purchasemodel.DeletePurchase(db, id)
		var data = libraries.H{"id": id}
		if err == nil {
			return libraries.ToJson(c, http.StatusOK, "data has been deleted!", data)
		} else {
			return err
		}
	}
}
