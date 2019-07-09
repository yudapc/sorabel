package handlers

import (
	"database/sql"
	"net/http"
	"sorabel/libraries"
	"sorabel/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetItems(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		items := models.GetItems(db)
		return c.JSON(http.StatusOK, libraries.H{
			"code": http.StatusOK,
			"data": items,
		})
	}
}

func GetItemDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		data := models.GetItemDetail(id, db)
		return c.JSON(http.StatusOK, libraries.H{
			"code": http.StatusOK,
			"data": data,
		})
	}
}

func CreateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data models.Item
		c.Bind(&data)

		id, err := models.CreateItem(db, data.Sku, data.Name, data.Stock)
		dataId := int(id)
		data.ID = dataId

		if err == nil {
			return c.JSON(http.StatusCreated, libraries.H{
				"code": http.StatusCreated,
				"data": data,
			})
		} else {
			return err
		}

	}
}

func UpdateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")
		id, errorConvert := strconv.Atoi(paramId)
		if errorConvert != nil {
			return errorConvert
		}

		var data models.Item
		data.ID = id
		c.Bind(&data)

		_, err := models.EditItem(db, id, data.Sku, data.Name, data.Stock)
		if err == nil {
			return c.JSON(http.StatusOK, libraries.H{
				"code": http.StatusOK,
				"data": data,
			})
		} else {
			return err
		}
	}
}
