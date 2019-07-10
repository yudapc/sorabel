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
		return c.JSON(http.StatusOK, libraries.H{
			"code":    http.StatusOK,
			"message": "successfully",
			"data":    items,
		})
	}
}

func GetItemDetail(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		data := itemmodel.GetItemDetail(id, db)
		return c.JSON(http.StatusOK, libraries.H{
			"code":    http.StatusOK,
			"message": "successfully",
			"data":    data,
		})
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
			return c.JSON(http.StatusCreated, libraries.H{
				"code":    http.StatusCreated,
				"message": "data has been created!",
				"data":    data,
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

		var data itemmodel.Item
		data.ID = id
		c.Bind(&data)

		_, err := itemmodel.EditItem(db, id, data.Sku, data.Name, data.Stock)
		if err == nil {
			return c.JSON(http.StatusOK, libraries.H{
				"code":    http.StatusOK,
				"message": "data has been updated!",
				"data":    data,
			})
		} else {
			return err
		}
	}
}

func DeleteItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := itemmodel.DeleteItem(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, libraries.H{
				"code":    http.StatusOK,
				"message": "data has been deleted!",
			})
		} else {
			return err
		}
	}
}
