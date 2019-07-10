package purchasemodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Purchase struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DateTime      string     `json:"date_time" validate:"required"`
	ReceiptNumber string     `json:"receipt_number" validate:"required"`
	// PurchaseDetail []PurchaseDetail `json:"purchase_details"`
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
	if result.Error != nil {
		return Purchase{}, result.Error
	}
	return purchase, nil
}

func CreatePurchase(db *gorm.DB, purchase Purchase) (Purchase, error) {
	result := db.Create(&purchase)
	if result.Error != nil {
		return Purchase{}, result.Error
	}
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
