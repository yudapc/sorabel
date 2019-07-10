package purchasemodel

import (
	"database/sql"
	"fmt"
)

type Purchase struct {
	ID             int              `json:"id"`
	DateTime       string           `json:"date_time" validate:"required"`
	ReceiptNumber  string           `json:"receipt_number" validate:"required"`
	PurchaseDetail []PurchaseDetail `json:"purchase_details"`
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

func GetPurchases(db *sql.DB) ([]Purchase, error) {
	sql := "SELECT * FROM purchases"
	rows, err := db.Query(sql)
	if err != nil {
		return []Purchase{}, err
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
	return result, nil
}

func GetPurchaseDetail(db *sql.DB, id string) (Purchase, error) {
	sql := fmt.Sprintf("SELECT * FROM purchases WHERE id = %s", id)
	rows, err := db.Query(sql)
	if err != nil {
		return Purchase{}, err
	}
	defer rows.Close()

	result := Purchase{}
	for rows.Next() {
		item := Purchase{}
		err2 := rows.Scan(&item.ID, &item.DateTime, &item.ReceiptNumber)
		if err2 != nil {
			panic(err2)
		}
		result = item
	}
	return result, err
}

func CreatePurchase(db *sql.DB, date_time string, receipt_number string) (int64, error) {
	sql := "INSERT INTO purchases(date_time, receipt_number) VALUES(?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(date_time, receipt_number)
	if err2 != nil {
		return 0, err2
	}

	return result.LastInsertId()
}

func EditPurchase(db *sql.DB, id string, date_time string, receipt_number string) (int64, error) {
	sql := "UPDATE purchases SET date_time = ?, receipt_number = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	result, err2 := stmt.Exec(date_time, receipt_number, id)

	if err2 != nil {
		return 0, err2
	}

	return result.RowsAffected()
}

func DeletePurchase(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM purchases WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		return 0, err2
	}

	return result.RowsAffected()
}
