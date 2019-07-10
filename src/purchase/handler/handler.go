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
		purchases := []int{1, 2, 3}
		return libraries.ToJson(c, http.StatusOK, "successfully", purchases)
	}
}

func GetPurchaseDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return libraries.ToJson(c, http.StatusOK, "successfully", id)
	}
}

func CreatePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data purchasemodel.Purchase
		err := c.Bind(&data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed!", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "purchase has been created!", err)
	}
}

func UpdatePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data purchasemodel.Purchase
		err := c.Bind(&data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed!", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "purchase has been updated!", err)
	}
}

func DeletePurchase(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var data = libraries.H{"id": id}
		return libraries.ToJson(c, http.StatusOK, "data has been deleted!", data)
	}
}
