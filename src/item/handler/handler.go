package handler

import (
	"io"
	"net/http"
	"os"
	"sorabel/helpers"
	"sorabel/src/item/model"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

func GetItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data, err := model.GetItems(db)
		if err != nil {
			return helpers.ToJson(context, http.StatusBadRequest, "failed", err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func GetItemDetail(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var item model.Item
		item.ID = uint(id)
		data, err := model.GetItemDetail(db, item)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func CreateItem(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		var data model.Item

		if errJSONValidate := helpers.SchemaValidation(context, "/schemas/item.json"); errJSONValidate != nil {
			return helpers.ToJsonBadRequest(context, errJSONValidate.Error())
		}
		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}
		dataItem, err := model.CreateItem(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdateItem(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Item
		data.ID = uint(id)

		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}

		_, err := model.EditItem(db, data)
		if err == nil {
			return helpers.ToJson(context, http.StatusOK, "data has been updated!", data)
		} else {
			return err
		}
	}
}

func DeleteItem(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Item
		data.ID = uint(id)
		dataItem, err := model.DeleteItem(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "data has been deleted!", dataItem)
	}
}

func ImportItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		file, err := context.FormFile("file")
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		src, errOpenFile := file.Open()
		if errOpenFile != nil {
			return helpers.ToJsonBadRequest(context, errOpenFile.Error())
		}
		defer src.Close()

		// Destination
		projectDirectory := helpers.ProjectDirectory()
		dst, errDestination := os.Create(projectDirectory + "/uploaded/" + file.Filename)
		if errDestination != nil {
			return helpers.ToJsonBadRequest(context, errDestination.Error())
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}

		return helpers.ToJson(context, http.StatusOK, "horee", nil)
	}
}
