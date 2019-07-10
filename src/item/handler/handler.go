package itemhandler

import (
	"database/sql"
	"net/http"
	"sorabel/libraries"
	itemmodel "sorabel/src/item/model"
	"strconv"

	"github.com/labstack/echo"
)

func GetItems(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := itemmodel.GetItems(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetItemDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		data, err := itemmodel.GetItemDetail(id, db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data itemmodel.Item
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}

		id, err := itemmodel.CreateItem(db, data.Sku, data.Name, data.Stock)

		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}

		dataID := int(id)
		data.ID = dataID
		return libraries.ToJson(c, http.StatusCreated, "item has been created!", data)
	}
}

func UpdateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var data itemmodel.Item
		data.ID = id

		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}

		_, err := itemmodel.EditItem(db, id, data.Sku, data.Name, data.Stock)
		if err == nil {
			return libraries.ToJson(c, http.StatusOK, "data has been updated!", data)
		} else {
			return err
		}
	}
}

func DeleteItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := itemmodel.DeleteItem(db, id)
		var data = libraries.H{"id": id}
		if err == nil {
			return libraries.ToJson(c, http.StatusOK, "data has been deleted!", data)
		} else {
			return err
		}
	}
}
