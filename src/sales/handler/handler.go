package handler

import (
	"encoding/csv"
	"net/http"
	"os"
	"sorabel/helpers"
	"sorabel/src/sales/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data, err := model.GetSales(db, context)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func GetSalesDetail(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var sales model.Sales
		sales.ID = uint(id)
		data, err := model.GetSalesDetail(db, sales)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func CreateSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		var data model.Sales
		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}
		dataItem, err := model.CreateSales(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdateSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Sales
		data.ID = uint(id)

		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}

		_, err := model.EditSales(db, data)
		if err == nil {
			return helpers.ToJson(context, http.StatusOK, "data has been updated!", data)
		} else {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
	}
}

func DeleteSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Sales
		data.ID = uint(id)
		dataItem, err := model.DeleteSales(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "data has been deleted!", helpers.HashObject{
			"id":             dataItem.ID,
			"date_time":      dataItem.DateTime,
			"invoice_number": dataItem.InvoiceNumber,
		})
	}
}

func GetSalesDetailItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		dataID := uint(id)
		data, err := model.GetSalesDetailItems(db, dataID)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func ImportSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		return helpers.ToJson(context, http.StatusOK, "successfully", nil)
	}
}

func ExportSales(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		fileName := "sales.csv"
		uploadPath := helpers.ProjectDirectory() + "/public/"
		fullPath := uploadPath + fileName

		os.Remove(fullPath)
		file, _ := os.Create(fullPath)
		defer file.Close()
		writer := csv.NewWriter(file)

		data := [][]string{}
		data = append(data, []string{"Waktu", "SKU", "Nama Barang", "Jumlah Keluar", "Harga Jual", "Total", "Catatan"})
		sales, _ := model.GetSales(db, context)

		for _, sale := range sales {
			salesDetails, _ := model.GetSalesDetailItems(db, sale.ID)
			for _, salesDetail := range salesDetails {
				qty := strconv.Itoa(salesDetail.Qty)
				sellingPrice := helpers.FormatRupiah(float64(salesDetail.SellingPrice))
				total := helpers.FormatRupiah(float64(salesDetail.Total))
				data = append(data, []string{
					sale.DateTime,
					salesDetail.Sku,
					salesDetail.Name,
					qty,
					sellingPrice,
					total,
					salesDetail.Note,
				})
			}
		}

		writer.WriteAll(data)
		defer writer.Flush()

		return context.Attachment(fullPath, fileName)
	}
}
