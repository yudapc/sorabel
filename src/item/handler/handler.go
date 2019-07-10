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
		items := itemmodel.GetItems(db)
		return libraries.ToJson(c, http.StatusOK, "successfully", items)
	}
}

func GetItemDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		data := itemmodel.GetItemDetail(id, db)
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data itemmodel.Item
		c.Bind(&data)

		id, err := itemmodel.CreateItem(db, data.Sku, data.Name, data.Stock)
		dataId := int(id)
		data.ID = dataId

		if err == nil {
			return libraries.ToJson(c, http.StatusCreated, "successfully", data)
		} else {
			return err
		}

	}
}

func UpdateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var data itemmodel.Item
		data.ID = id
		c.Bind(&data)

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
