package model

import (
	"sorabel/helpers"
	ItemModel "sorabel/src/item/model"
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
	SalesID       uint       `json:"sales_id"`
}

func GetSales(db *gorm.DB, context echo.Context) ([]Sales, error) {
	var sales []Sales
	_, limit, offset, order := helpers.QueryString(context)
	result := db.Limit(limit).Offset(offset).Order(order, true).Find(&sales)
	if result.Error != nil {
		return []Sales{}, result.Error
	}
	return sales, nil
}

func GetSalesDetail(db *gorm.DB, sales Sales) (Sales, error) {
	result := db.First(&sales)
	salesDetails, _ := GetSalesDetailItems(db, sales.ID)
	sales.SalesDetails = salesDetails
	if result.Error != nil {
		return Sales{}, result.Error
	}
	return sales, nil
}

func CreateSales(db *gorm.DB, sales Sales) (Sales, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return Sales{}, err
	}
	row := new(Sales)
	result := db.Create(&sales).Scan(&row)
	if result.Error != nil {
		tx.Rollback()
		return Sales{}, result.Error
	}

	for _, salesDetail := range sales.SalesDetails {
		var item ItemModel.Item
		search := db.Where("sku = ?", salesDetail.Sku).First(&item)
		if search.Error != nil {
			tx.Rollback()
			return Sales{}, search.Error
		}

		total := float64(salesDetail.Qty) * item.SellingPrice
		profit := total - (float64(salesDetail.Qty) * item.PurchasePrice)

		salesDetailItem := SalesDetail{
			Sku:           salesDetail.Sku,
			Name:          item.Name,
			Qty:           salesDetail.Qty,
			SellingPrice:  item.SellingPrice,
			PurchasePrice: item.PurchasePrice,
			Total:         total,
			Profit:        profit,
			Note:          salesDetail.Note,
			SalesID:       row.ID,
		}
		insertDetail := db.Create(&salesDetailItem)
		item.Stock = item.Stock - salesDetail.Qty
		updateItem := db.Save(&item)
		if insertDetail.Error != nil || updateItem.Error != nil {
			tx.Rollback()
			return Sales{}, insertDetail.Error
		}
	}

	tx.Commit()

	dataSales, _ := GetSalesDetail(db, sales)
	salesItems, _ := GetSalesDetailItems(db, sales.ID)
	dataSales.SalesDetails = salesItems

	return dataSales, nil
}

func EditSales(db *gorm.DB, sales Sales) (Sales, error) {
	_, errorExist := GetSalesDetail(db, sales)
	if errorExist != nil {
		return Sales{}, errorExist
	}
	result := db.Save(&sales)
	if result.Error != nil {
		return Sales{}, result.Error
	}
	return sales, nil
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
	result := db.Where("sales_id = ?", id).Find(&salesDetails)
	if result.Error != nil {
		return []SalesDetail{}, result.Error
	}
	return salesDetails, nil
}
