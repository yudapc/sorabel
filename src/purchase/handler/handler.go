package handler

import (
	"encoding/csv"
	"net/http"
	"os"
	"sorabel/helpers"
	"sorabel/src/purchase/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func GetPurchases(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data, err := model.GetPurchases(db, context)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func GetPurchaseDetail(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var purchase model.Purchase
		purchase.ID = uint(id)
		data, err := model.GetPurchaseDetail(db, purchase)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func CreatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		var data model.Purchase
		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}
		dataItem, err := model.CreatePurchase(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusCreated, "data has been created!", dataItem)
	}
}

func UpdatePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Purchase
		data.ID = uint(id)

		if errBind := context.Bind(&data); errBind != nil {
			return helpers.ToJsonBadRequest(context, errBind.Error())
		}
		if errValidate := context.Validate(&data); errValidate != nil {
			return helpers.ToJsonBadRequest(context, errValidate.Error())
		}

		_, err := model.EditPurchase(db, data)
		if err == nil {
			return helpers.ToJson(context, http.StatusOK, "data has been updated!", data)
		} else {
			return err
		}
	}
}

func DeletePurchase(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		var data model.Purchase
		data.ID = uint(id)
		dataItem, err := model.DeletePurchase(db, data)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "data has been deleted!", dataItem)
	}
}

func GetPurchaseDetailItems(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
		dataID := uint(id)
		data, err := model.GetPurchaseDetailItems(db, dataID)
		if err != nil {
			return helpers.ToJsonBadRequest(context, err.Error())
		}
		return helpers.ToJson(context, http.StatusOK, "successfully", data)
	}
}

func ImportPurchases(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		return helpers.ToJson(context, http.StatusOK, "successfully", nil)
	}
}

func ExportPurchases(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		fileName := "purchase.csv"
		uploadPath := helpers.ProjectDirectory() + "/public/"
		fullPath := uploadPath + fileName

		os.Remove(fullPath)
		file, _ := os.Create(fullPath)
		defer file.Close()
		writer := csv.NewWriter(file)

		data := [][]string{}
		data = append(data, []string{"Waktu", "SKU", "Nama Barang", "Jumlah Pemesanan", "Jumlah Diterima", "Harga Beli", "Total", "Nomor Kwitansi", "Catatan"})
		purchases, _ := model.GetPurchases(db, context)

		for _, purchase := range purchases {
			purchaseDetails, _ := model.GetPurchaseDetailItems(db, purchase.ID)
			for _, purchaseDetail := range purchaseDetails {
				qty := strconv.Itoa(purchaseDetail.Qty)
				itemReceived := strconv.Itoa(purchaseDetail.ItemReceived)
				purchasePrice := helpers.FormatRupiah(float64(purchaseDetail.PurchasePrice))
				total := helpers.FormatRupiah(float64(purchaseDetail.Total))
				data = append(data, []string{
					purchase.DateTime,
					purchaseDetail.Sku,
					purchaseDetail.Name,
					qty,
					itemReceived,
					purchasePrice,
					total,
					purchase.ReceiptNumber,
					purchaseDetail.Note,
				})
			}
		}

		writer.WriteAll(data)
		defer writer.Flush()

		return context.Attachment(fullPath, fileName)
	}
}
