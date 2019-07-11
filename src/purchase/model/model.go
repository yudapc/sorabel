package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Purchase struct {
	ID              uint             `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       *time.Time       `json:"deleted_at"`
	DateTime        string           `json:"date_time" validate:"required"`
	ReceiptNumber   string           `json:"receipt_number" validate:"required"`
	PurchaseDetails []PurchaseDetail `gorm:"foreignkey:PurchaseRefer" json:"purchase_details"`
}

type PurchaseDetail struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Sku           string     `json:"sku"`
	Name          string     `json:"name"`
	Qty           int        `json:"qty"`
	ItemReceived  int        `json:"item_received"`
	PurchasePrice int        `json:"purchase_price"`
	Total         int        `json:"total"`
	Note          string     `json:"note"`
	PurchaseID    uint       `json:"purchase_id"`
}

func GetPurchases(db *gorm.DB) ([]Purchase, error) {
	var purchases []Purchase
	result := db.Find(&purchases)
	if result.Error != nil {
		return []Purchase{}, result.Error
	}
	return purchases, nil
}

func GetPurchaseDetail(db *gorm.DB, purchase Purchase) (Purchase, error) {
	result := db.First(&purchase)
	purchaseDetails, _ := GetPurchaseDetailItems(db, purchase.ID)
	purchase.PurchaseDetails = purchaseDetails
	if result.Error != nil {
		return Purchase{}, result.Error
	}
	return purchase, nil
}

func CreatePurchase(db *gorm.DB, purchase Purchase) (Purchase, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return Purchase{}, err
	}
	row := new(Purchase)
	result := db.Create(&purchase).Scan(&row)
	if result.Error != nil {
		tx.Rollback()
		return Purchase{}, result.Error
	}

	for _, purchaseDetail := range purchase.PurchaseDetails {
		item := PurchaseDetail{
			Sku:           purchaseDetail.Sku,
			Name:          purchaseDetail.Name,
			Qty:           purchaseDetail.Qty,
			ItemReceived:  purchaseDetail.ItemReceived,
			PurchasePrice: purchaseDetail.PurchasePrice,
			Total:         purchaseDetail.Qty * purchaseDetail.PurchasePrice,
			Note:          purchaseDetail.Note,
			PurchaseID:    row.ID,
		}
		insertDetail := db.Create(&item)
		if insertDetail.Error != nil {
			tx.Rollback()
			return Purchase{}, insertDetail.Error
		}
	}

	tx.Commit()

	return purchase, nil
}

func EditPurchase(db *gorm.DB, purchase Purchase) (Purchase, error) {
	_, errorExist := GetPurchaseDetail(db, purchase)
	if errorExist != nil {
		return Purchase{}, errorExist
	}
	result := db.Save(&purchase)
	if result.Error != nil {
		return Purchase{}, result.Error
	}
	return purchase, nil
}

func DeletePurchase(db *gorm.DB, purchase Purchase) (Purchase, error) {
	_, errorExist := GetPurchaseDetail(db, purchase)
	if errorExist != nil {
		return Purchase{}, errorExist
	}
	result := db.Delete(&purchase)
	if result.Error != nil {
		return Purchase{}, result.Error
	}
	return purchase, nil
}

func GetPurchaseDetailItems(db *gorm.DB, id uint) ([]PurchaseDetail, error) {
	purchaseDetails := []PurchaseDetail{}
	result := db.Where("purchase_id = ?", id).Find(&purchaseDetails)
	if result.Error != nil {
		return []PurchaseDetail{}, result.Error
	}
	return purchaseDetails, nil
}
