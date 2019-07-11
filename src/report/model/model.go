package model

import (
	"github.com/jinzhu/gorm"
)

type ItemValueReport struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	TotalQty int    `json:"total_qty"`
}

type SalesReport struct {
	Sku           string  `json:"sku"`
	Name          string  `json:"name"`
	TotalQty      int     `json:"total_qty"`
	SellingPrice  float64 `json:"selling_price"`
	PurchasePrice float64 `json:"purchase_price"`
	Profit        float64 `json:"profit"`
}

func GenerateItemValueReport(db *gorm.DB) []ItemValueReport {
	var itemValues []ItemValueReport
	sql := `
		SELECT pd.sku AS sku, pd.name As name, SUM(pd.qty) AS total_qty 
		FROM purchase_details pd 
		LEFT JOIN purchases as p ON p.id = pd.purchase_id GROUP BY pd.sku
	`
	db.Raw(sql).Scan(&itemValues)
	return itemValues
}

func GenerateSalesReport(db *gorm.DB) []SalesReport {
	var salesReports []SalesReport
	sql := `
	SELECT sd.sku AS sku, sd.name AS name, SUM(sd.qty) AS total_qty, sd.selling_price AS selling_price, i.purchase_price AS purchase_price, (SUM(sd.qty) * sd.selling_price) - (SUM(sd.qty) * i.purchase_price) AS total_profit 
	FROM sales_details sd
	LEFT JOIN items AS i ON i.sku = sd.sku
	LEFT JOIN sales AS s ON s.id = sd.sales_id
	GROUP BY sd.sku
	`
	db.Raw(sql).Scan(&salesReports)
	return salesReports
}
