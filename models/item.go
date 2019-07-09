package models

import (
	"database/sql"
	"fmt"
)

type Item struct {
	ID    int    `json:"id"`
	Sku   string `json:"sku"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func GetItems(db *sql.DB) []Item {
	sql := "SELECT * FROM items"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := []Item{}
	for rows.Next() {
		item := Item{}
		err2 := rows.Scan(&item.ID, &item.Sku, &item.Name, &item.Stock)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func GetItemDetail(id string, db *sql.DB) Item {
	sql := fmt.Sprintf("SELECT * FROM items WHERE id = %s", id)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := Item{}
	for rows.Next() {
		item := Item{}
		err2 := rows.Scan(&item.ID, &item.Sku, &item.Name, &item.Stock)
		if err2 != nil {
			panic(err2)
		}
		result = item
	}
	return result
}

func CreateItem(db *sql.DB, sku string, name string, stock int) (int64, error) {
	sql := "INSERT INTO items(sku, name, stock) VALUES(?,?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(sku, name, stock)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditItem(db *sql.DB, id int, sku string, name string, stock int) (int64, error) {
	sql := "UPDATE items set sku = ?, name = ?, stock = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(sku, name, stock, id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
