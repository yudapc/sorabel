package purchasemodel

import (
	"database/sql"
)

type Purchase struct {
	ID             int              `json:"id"`
	DateTime       string           `json:"date_time" validate:"required"`
	ReceiptNumber  string           `json:"receipt_number" validate:"required"`
	PurchaseDetail []PurchaseDetail `json:"purchase_details" validate:"required"`
}

type PurchaseDetail struct {
	ID            int
	Sku           string `json:"sku"`
	Name          string `json:"name"`
	Qty           int    `json:"qty"`
	ItemReceived  int    `json:"item_received"`
	PurchasePrice int    `json:"purchase_price"`
	Total         int    `json:"total"`
	Note          string `json:"note"`
}

func GetPurchases(db *sql.DB) []Purchase {
	sql := "SELECT * FROM purchases"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := []Purchase{}
	for rows.Next() {
		item := Purchase{}
		err2 := rows.Scan(&item.ID, &item.DateTime, &item.ReceiptNumber)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func CreatePurchase(db *sql.DB, date_time string, receipt_number string) (int64, error) {
	sql := "INSERT INTO purchases(date_time, receipt_number) VALUES(?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(date_time, receipt_number)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
