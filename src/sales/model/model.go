package model

import (
	ItemModel "sorabel/src/item/model"
	"time"

	"github.com/jinzhu/gorm"
)

type Sales struct {
	ID            uint          `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	DeletedAt     *time.Time    `json:"deleted_at"`
	DateTime      string        `json:"date_time" validate:"required"`
	ReceiptNumber string        `json:"receipt_number" validate:"required"`
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
	PurchasePrice int        `json:"purchase_price"`
	SellingPrice  int        `json:"selling_price"`
	Total         int        `json:"total"`
	Profit        int        `json:"profit"`
	Note          string     `json:"note"`
	SalesID       uint       `json:"sales_id"`
}

func GetSales(db *gorm.DB) ([]Sales, error) {
	var sales []Sales
	result := db.Find(&sales)
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
		db.Where("sku = ?", salesDetail.Sku).First(&item)
		salesDetailItem := SalesDetail{
			Sku:           salesDetail.Sku,
			Name:          item.Name,
			Qty:           salesDetail.Qty,
			SellingPrice:  salesDetail.SellingPrice,
			PurchasePrice: item.PurchasePrice,
			Total:         salesDetail.Qty * salesDetail.SellingPrice,
			Profit:        salesDetail.Qty*salesDetail.SellingPrice - salesDetail.Qty*item.PurchasePrice,
			Note:          salesDetail.Note,
			SalesID:       row.ID,
		}
		insertDetail := db.Create(&salesDetailItem)
		if insertDetail.Error != nil {
			tx.Rollback()
			return Sales{}, insertDetail.Error
		}
	}

	tx.Commit()

	return sales, nil
}

func GetSalesDetailItems(db *gorm.DB, id uint) ([]SalesDetail, error) {
	salesDetails := []SalesDetail{}
	result := db.Where("sales_id = ?", id).Find(&salesDetails)
	if result.Error != nil {
		return []SalesDetail{}, result.Error
	}
	return salesDetails, nil
}
