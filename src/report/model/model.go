package model

import (
	"github.com/jinzhu/gorm"
)

type ItemValueReport struct {
	Sku               string `json:"sku"`
	Name              string `json:"name"`
	TotalQtyPurchased int    `json:"total_qty_purchased"`
	TotalItemReceived int    `json:"total_item_received"`
	TotalStock        int    `json:"total_stock"`
}

type SalesReport struct {
	Sku              string  `json:"sku"`
	Name             string  `json:"name"`
	TotalQty         int     `json:"total_qty"`
	TotalProfit      float64 `json:"total_profit"`
	TotalTransaction float64 `json:"total_transaction"`
}

func GenerateItemValueReport(db *gorm.DB) []ItemValueReport {
	var itemValues []ItemValueReport
	sql := `
	SELECT 
		i.sku AS sku,
		i.name AS name, 
		(SELECT SUM(qty) FROM purchase_details WHERE sku = i.sku) AS total_qty_purchased,
		(SELECT SUM(item_received) FROM purchase_details WHERE sku = i.sku) AS total_item_received,
		i.stock AS total_stock
	FROM
		items i
	`

	db.Raw(sql).Scan(&itemValues)
	return itemValues
}

func GenerateSalesReport(db *gorm.DB) []SalesReport {
	var salesReports []SalesReport
	sql := `
	SELECT 
		i.sku AS sku,
		i.name AS name,
		(SELECT SUM(qty) FROM sales_details WHERE sku = i.sku) AS total_qty,
		(SELECT SUM(profit) FROM sales_details WHERE sku = i.sku) AS total_profit,
		(SELECT COUNT(*) FROM sales_details WHERE sku = i.sku) AS total_transaction
	FROM
		items i
	`
	db.Raw(sql).Scan(&salesReports)
	return salesReports
}
