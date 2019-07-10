package main

import (
	HomeHandler "sorabel/src/home/handler"
	ItemHandler "sorabel/src/item/handler"
	ItemModel "sorabel/src/item/model"
	PurchaseHandler "sorabel/src/purchase/handler"
	PurchaseModel "sorabel/src/purchase/model"
	SalesHandler "sorabel/src/sales/handler"
	SalesModel "sorabel/src/sales/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	db, err := gorm.Open("sqlite3", "storage.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	migrate(db)

	e.GET("/", HomeHandler.Home())
	e.GET("/items", ItemHandler.GetItems(db))
	e.GET("/items/:id", ItemHandler.GetItemDetail(db))
	e.POST("/items", ItemHandler.CreateItem(db))
	e.PUT("/items/:id", ItemHandler.UpdateItem(db))
	e.DELETE("/items/:id", ItemHandler.DeleteItem(db))
	e.GET("/purchases", PurchaseHandler.GetPurchases(db))
	e.GET("/purchases/:id", PurchaseHandler.GetPurchaseDetail(db))
	e.POST("/purchases", PurchaseHandler.CreatePurchase(db))
	e.PUT("/purchases/:id", PurchaseHandler.UpdatePurchase(db))
	e.DELETE("/purchases/:id", PurchaseHandler.DeletePurchase(db))
	e.GET("/purchases/:id/items", PurchaseHandler.GetPurchaseDetailItems(db))
	e.GET("/sales", SalesHandler.GetSales(db))
	e.GET("/sales/:id", SalesHandler.GetSalesDetail(db))
	e.POST("/sales", SalesHandler.CreateSales(db))
	e.PUT("/sales/:id", SalesHandler.UpdateSales(db))
	e.DELETE("/sales/:id", SalesHandler.DeleteSales(db))
	e.GET("/sales/:id/items", SalesHandler.GetSalesDetailItems(db))
	e.Logger.Fatal(e.Start(":8000"))
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&ItemModel.Item{},
		&PurchaseModel.Purchase{},
		&PurchaseModel.PurchaseDetail{},
		&SalesModel.Sales{},
		&SalesModel.SalesDetail{},
	)
}
