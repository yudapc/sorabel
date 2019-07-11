package handler

import (
	"net/http"
	"sorabel/helpers"
	ItemModel "sorabel/src/item/model"
	"sorabel/src/report/model"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func ItemValueReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data := model.GenerateItemValueReport(db)
		const layout = "2 Jan 2006"
		t := time.Now()
		DateTime := t.Format(layout)
		var totalItems int
		var totalValues float64
		items, _ := ItemModel.GetItems(db, context)
		totalSku := len(items)
		for _, value := range data {
			totalItems = totalItems + value.TotalItemReceived
			totalValues = totalValues + value.PurchasePrice
		}
		header := helpers.HashObject{
			"DateTime":   DateTime,
			"totalSku":   totalSku,
			"totalItems": totalItems,
		}
		response := helpers.HashObject{
			"header": header,
			"items":  data,
		}
		return helpers.ToJson(context, http.StatusOK, "succesfully", response)
	}
}

func SalesReport(db *gorm.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		data := model.GenerateSalesReport(db)
		const layout = "2 Jan 2006"
		t := time.Now()
		DateTime := t.Format(layout)
		var totalProfit float64
		var totalItems int
		var totalOmzet float64
		totalTransactions := len(data)
		for _, value := range data {
			totalOmzet = totalOmzet + value.Total
			totalProfit = totalProfit + value.Profit
			totalItems = totalItems + value.TotalQty
		}
		header := helpers.HashObject{
			"Date":              DateTime,
			"totalOmzet":        totalOmzet,
			"totalProfit":       totalProfit,
			"totalItems":        totalItems,
			"totalTransactions": totalTransactions,
		}

		response := helpers.HashObject{
			"header": header,
			"items":  data,
		}
		return helpers.ToJson(context, http.StatusOK, "succesfully", response)
	}
}
