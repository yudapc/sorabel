package handler

import (
	"net/http"
	"sorabel/libraries"
	"sorabel/src/item/model"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

func GetItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := model.GetItems(db)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func GetItemDetail(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var item model.Item
		item.ID = uint(id)
		data, err := model.GetItemDetail(db, item)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "successfully", data)
	}
}

func CreateItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data model.Item
		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}
		dataItem, err := model.CreateItem(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdateItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var data model.Item
		data.ID = uint(id)

		if errBind := c.Bind(&data); errBind != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errBind.Error())
		}
		if errValidate := c.Validate(&data); errValidate != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", errValidate.Error())
		}

		_, err := model.EditItem(db, data)
		if err == nil {
			return libraries.ToJson(c, http.StatusOK, "data has been updated!", data)
		} else {
			return err
		}
	}
}

func DeleteItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var data model.Item
		data.ID = uint(id)
		dataItem, err := model.DeleteItem(db, data)
		if err != nil {
			return libraries.ToJson(c, http.StatusBadRequest, "failed", err.Error())
		}
		return libraries.ToJson(c, http.StatusOK, "data has been deleted!", dataItem)
	}
}
