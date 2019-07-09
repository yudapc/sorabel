package main

import (
	"database/sql"
	"net/http"
	"sorabel/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)
	e.GET("/", HelloWorld())
	e.GET("/items", handlers.GetItems(db))
	e.Logger.Fatal(e.Start(":8000"))
}

type H map[string]interface{}

func HelloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"code": http.StatusOK,
			"data": "Hello World",
		})
	}
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER
	);
	CREATE TABLE IF NOT EXISTS items(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		sku VARCHAR NOT NULL,
		name VARCHAR NOT NULL,
		stock INTEGER
		);
		CREATE TABLE IF NOT EXISTS purchases(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		date_time VARCHAR NOT NULL,
		receipt_number VARCHAR NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_items_sku ON items (sku);
	CREATE TABLE IF NOT EXISTS purchase_details(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		sku VARCHAR NOT NULL,
		name VARCHAR NOT NULL,
		qty INTEGER,
		item_received INTEGER,
		purchase_price INTEGER,
		total INTEGER,
		note TEXT
	);
	CREATE TABLE IF NOT EXISTS sales(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		date_time NUMERIC NOT NULL,
		invoice_number VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS sales_details(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		sku VARCHAR NOT NULL,
		name VARCHAR NOT NULL,
		qty INTEGER,
		item_received INTEGER,
		selling_price INTEGER,
		total INTEGER,
		note TEXT
	);
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
