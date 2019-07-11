package handler

import (
	"encoding/csv"
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
		var errorFileType int = 0
		for _, filetype := range file.Header["Content-Type"] {
			if filetype != "text/csv" {
				errorFileType++
			}
		}
		if errorFileType > 0 {
			return helpers.ToJsonBadRequest(context, "Please check your format file, the file must csv")
		}
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		src, errOpenFile := file.Open()
		if errOpenFile != nil {
			return helpers.ToJsonBadRequest(context, errOpenFile.Error())
		}

		lines, err := csv.NewReader(src).ReadAll()
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}

		importedDataItems, err := model.InsertBulkItems(db, lines)
		defer src.Close()
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "data success imported", importedDataItems)
	}
}

func ExportItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		fileName := "items.csv"
		uploadPath := helpers.ProjectDirectory() + "/uploaded/"
		fullPath := uploadPath + fileName

		os.Remove(fullPath)
		file, _ := os.Create(fullPath)
		defer file.Close()
		writer := csv.NewWriter(file)

		data := [][]string{}
		data = append(data, []string{"SKU", "Nama Item", "Jumlah Sekarang", "Harga Beli", "Harga Jual"})
		items, _ := model.GetItems(db)

		for _, item := range items {
			stock := strconv.Itoa(item.Stock)
			purchasePrice := strconv.Itoa(item.PurchasePrice)
			sellingPrice := strconv.Itoa(item.SellingPrice)
			data = append(data, []string{
				item.Sku,
				item.Name,
				stock,
				purchasePrice,
				sellingPrice,
			})
		}

		writer.WriteAll(data)
		defer writer.Flush()

		return context.Attachment(fullPath, fileName)
	}
}
