package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ItemValueReport struct {
	Sku               string  `json:"sku"`
	Name              string  `json:"name"`
	TotalQtyPurchased int     `json:"total_qty_purchased"`
	TotalItemReceived int     `json:"total_item_received"`
	TotalStock        int     `json:"total_stock"`
	PurchasePrice     float64 `json:"purhase_price"`
}

type SalesReport struct {
	InvoiceNumber string    `json:"invoice_number"`
	CreatedAt     time.Time `json:"created_at"`
	Sku           string    `json:"sku"`
	Name          string    `json:"name"`
	TotalQty      int       `json:"total_qty"`
	SellingPrice  float64   `json:"selling_price"`
	Total         float64   `json:"total"`
	PurchasePrice float64   `json:"purchase_price"`
	Profit        float64   `json:"profit"`
}

func GenerateItemValueReport(db *gorm.DB) []ItemValueReport {
	var itemValues []ItemValueReport
	sql := `
	SELECT 
		i.sku AS sku,
		i.name AS name, 
		(SELECT SUM(qty) FROM purchase_details WHERE sku = i.sku) AS total_qty_purchased,
		(SELECT SUM(item_received) FROM purchase_details WHERE sku = i.sku) AS total_item_received,
		i.stock AS total_stock,
		i.purchase_price AS purchase_price
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
		s.invoice_number AS invoice_number,
		s.created_at AS created_at,
		sd.sku AS sku, 
		sd.name AS name,
		SUM(sd.qty) AS total_qty,
		i.selling_price AS selling_price,	
		SUM(sd.qty) * i.selling_price AS total,
		i.purchase_price AS purchase_price,	
		SUM(sd.profit) AS profit
	FROM sales_details sd
	LEFT JOIN items AS i ON i.sku = sd.sku
	LEFT JOIN sales AS s ON s.id = sd.sales_id
	GROUP BY sd.sku
	`
	db.Raw(sql).Scan(&salesReports)
	return salesReports
}
