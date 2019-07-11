package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Sku           string     `json:"sku" validate:"required" gorm:"type:varchar(100);unique_index"`
	Name          string     `json:"name" validate:"required" gorm:"size:255"`
	Stock         int        `json:"stock" validate:"required"`
	PurchasePrice int        `json:"purchase_price" validate:"required"`
	SellingPrice  int        `json:"selling_price" validate:"required"`
}

func GetItems(db *gorm.DB) ([]Item, error) {
	var items []Item
	result := db.Find(&items)
	if result.Error != nil {
		return []Item{}, result.Error
	}
	return items, nil
}

func GetItemDetail(db *gorm.DB, item Item) (Item, error) {
	result := db.First(&item)
	if result.Error != nil {
		return Item{}, result.Error
	}
	return item, nil
}

func CreateItem(db *gorm.DB, item Item) (Item, error) {
	result := db.Create(&item)
	if result.Error != nil {
		return Item{}, result.Error
	}
	return item, nil
}

func EditItem(db *gorm.DB, item Item) (Item, error) {
	_, errorExist := GetItemDetail(db, item)
	if errorExist != nil {
		return Item{}, errorExist
	}
	result := db.Save(&item)
	if result.Error != nil {
		return Item{}, result.Error
	}
	return item, nil
}

func DeleteItem(db *gorm.DB, item Item) (Item, error) {
	_, errorExist := GetItemDetail(db, item)
	if errorExist != nil {
		return Item{}, errorExist
	}
	result := db.Delete(&item)
	if result.Error != nil {
		return Item{}, result.Error
	}
	return item, nil
}
