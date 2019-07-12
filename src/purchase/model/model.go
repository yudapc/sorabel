package model

import (
	"sorabel/helpers"
	ItemModel "sorabel/src/item/model"
	"time"

	"github.com/labstack/echo"

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
	PurchasePrice float64    `json:"purchase_price"`
	Total         float64    `json:"total"`
	Note          string     `json:"note"`
	PurchaseID    uint       `json:"purchase_id"`
}

func GetPurchases(db *gorm.DB, context echo.Context) ([]Purchase, error) {
	_, limit, offset, order := helpers.QueryString(context)
	var purchases []Purchase
	result := db.Limit(limit).Offset(offset).Order(order, true).Find(&purchases)
	if result.Error != nil {
		return []Purchase{}, result.Error
	}

	for i, _ := range purchases {
		purchaseDetails, _ := GetPurchaseDetailItems(db, purchases[i].ID)
		purchases[i].PurchaseDetails = purchaseDetails
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
		var item ItemModel.Item
		search := db.Where("sku = ?", purchaseDetail.Sku).First(&item)
		if search.Error != nil {
			tx.Rollback()
			return Purchase{}, search.Error
		}
		data := PurchaseDetail{
			Sku:           purchaseDetail.Sku,
			Name:          item.Name,
			Qty:           purchaseDetail.Qty,
			ItemReceived:  purchaseDetail.ItemReceived,
			PurchasePrice: item.PurchasePrice,
			Total:         float64(purchaseDetail.Qty) * item.PurchasePrice,
			Note:          purchaseDetail.Note,
			PurchaseID:    row.ID,
		}
		insertDetail := db.Create(&data)
		item.Stock = item.Stock + purchaseDetail.ItemReceived
		db.Save(&item)
		if insertDetail.Error != nil {
			tx.Rollback()
			return Purchase{}, insertDetail.Error
		}
	}

	tx.Commit()

	dataPurchase, _ := GetPurchaseDetail(db, purchase)
	purchaseItems, _ := GetPurchaseDetailItems(db, purchase.ID)
	dataPurchase.PurchaseDetails = purchaseItems

	return dataPurchase, nil
}

func EditPurchase(db *gorm.DB, purchase Purchase) (Purchase, error) {
	_, errorExist := GetPurchaseDetail(db, purchase)
	if errorExist != nil {
		return Purchase{}, errorExist
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return Purchase{}, err
	}
	updateHeader := db.Save(&purchase)
	if updateHeader.Error != nil {
		tx.Rollback()
		return Purchase{}, updateHeader.Error
	}

	for _, purchaseDetail := range purchase.PurchaseDetails {
		var item ItemModel.Item
		search := db.Where("sku = ?", purchaseDetail.Sku).First(&item)
		if search.Error != nil {
			tx.Rollback()
			return Purchase{}, search.Error
		}
		data := PurchaseDetail{
			ID:            purchaseDetail.ID,
			Sku:           purchaseDetail.Sku,
			Name:          purchaseDetail.Name,
			Qty:           purchaseDetail.Qty,
			ItemReceived:  purchaseDetail.ItemReceived,
			PurchasePrice: purchaseDetail.PurchasePrice,
			Total:         purchaseDetail.Total,
			Note:          purchaseDetail.Note,
			PurchaseID:    purchase.ID,
		}
		updateDetail := db.Save(&data)
		if updateDetail.Error != nil {
			tx.Rollback()
			return Purchase{}, updateDetail.Error
		}
	}

	tx.Commit()

	data, _ := GetPurchaseDetail(db, purchase)

	return data, nil
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
