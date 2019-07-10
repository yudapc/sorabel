package main

import (
	homehandler "sorabel/src/home/handler"
	itemhandler "sorabel/src/item/handler"
	itemmodel "sorabel/src/item/model"
	purchasehandler "sorabel/src/purchase/handler"
	purchasemodel "sorabel/src/purchase/model"

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

	e.GET("/", homehandler.Home())
	e.GET("/items", itemhandler.GetItems(db))
	e.GET("/items/:id", itemhandler.GetItemDetail(db))
	e.POST("/items", itemhandler.CreateItem(db))
	e.PUT("/items/:id", itemhandler.UpdateItem(db))
	e.DELETE("/items/:id", itemhandler.DeleteItem(db))
	e.GET("/purchases", purchasehandler.GetPurchases(db))
	e.GET("/purchases/:id", purchasehandler.GetPurchaseDetail(db))
	e.POST("/purchases", purchasehandler.CreatePurchase(db))
	e.PUT("/purchases/:id", purchasehandler.UpdatePurchase(db))
	e.DELETE("/purchases/:id", purchasehandler.DeletePurchase(db))
	e.Logger.Fatal(e.Start(":8000"))
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&itemmodel.Item{},
		&purchasemodel.Purchase{},
		&purchasemodel.PurchaseDetail{},
	)
}
