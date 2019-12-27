package model

import (
	"sorabel/helpers"
	"time"

	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
)

type Sales struct {
	ID            uint          `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	DeletedAt     *time.Time    `json:"deleted_at"`
	DateTime      string        `json:"date_time" validate:"required"`
	InvoiceNumber string        `json:"invoice_number" validate:"required"`
	SalesDetails  []SalesDetail `gorm:"foreignkey:SalesRefer" json:"sales_details"`
}

type SalesDetail struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Sku           string     `json:"sku"`
	Name          string     `json:"name"`
	Qty           int        `json:"qty"`
	PurchasePrice float64    `json:"purchase_price"`
	SellingPrice  float64    `json:"selling_price"`
	Total         float64    `json:"total"`
	Profit        float64    `json:"profit"`
	Note          string     `json:"note"`
	SalesRefer    uint       `json:"sales_refer"`
	Sales         *Sales     `gorm:"foreignkey:SalesRefer" json:"sales"`
}

func GetSales(db *gorm.DB, context echo.Context) ([]Sales, error) {
	_, limit, offset, order := helpers.QueryString(context)
	var sales []Sales
	result := db.Limit(limit).Offset(offset).Order(order, true).Preload("SalesDetails").Find(&sales)
	if result.Error != nil {
		return []Sales{}, result.Error
	}

	return sales, nil
}

func GetSalesDetail(db *gorm.DB, sales Sales) (Sales, error) {
	result := db.Preload("SalesDetails").Find(&sales)
	if result.Error != nil {
		return Sales{}, result.Error
	}
	return sales, nil
}

func CreateSales(db *gorm.DB, sales Sales) (Sales, error) {
	row := new(Sales)
	result := db.Create(&sales).Scan(&row)
	if result.Error != nil {
		return Sales{}, result.Error
	}
	dataSales, _ := GetSalesDetail(db, sales)
	return dataSales, nil
}

func EditSales(db *gorm.DB, sales Sales) (Sales, error) {
	_, errorExist := GetSalesDetail(db, sales)
	if errorExist != nil {
		return Sales{}, errorExist
	}

	updateHeader := db.Save(&sales)
	if updateHeader.Error != nil {
		return Sales{}, updateHeader.Error
	}

	dataSales, _ := GetSalesDetail(db, sales)

	return dataSales, nil
}

func DeleteSales(db *gorm.DB, sales Sales) (Sales, error) {
	dataSales, errorExist := GetSalesDetail(db, sales)
	if errorExist != nil {
		return Sales{}, errorExist
	}
	result := db.Delete(&sales)
	if result.Error != nil {
		return Sales{}, result.Error
	}
	return dataSales, nil
}

func GetSalesDetailItems(db *gorm.DB, id uint) ([]SalesDetail, error) {
	salesDetails := []SalesDetail{}
	result := db.Where("sales_refer = ?", id).Find(&salesDetails)
	if result.Error != nil {
		return []SalesDetail{}, result.Error
	}
	return salesDetails, nil
}
