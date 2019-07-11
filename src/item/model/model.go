package model

import (
	"sorabel/helpers"
	"strconv"
	"time"

	"github.com/labstack/echo"

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
	PurchasePrice float64    `json:"purchase_price" validate:"required"`
	SellingPrice  float64    `json:"selling_price" validate:"required"`
}

func GetItems(db *gorm.DB, context echo.Context) ([]Item, error) {
	var items []Item
	_, limit, offset, order := helpers.QueryString(context)
	result := db.Limit(limit).Offset(offset).Order(order, true).Find(&items)
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

func InsertBulkItems(db *gorm.DB, lines [][]string) ([]Item, error) {
	var dataItems []Item
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return []Item{}, err
	}

	for index, line := range lines {
		if index > 0 {
			lenLine := len(line)
			stock, _ := strconv.Atoi(line[2])
			var purchasePrice float64
			var sellingPrice float64
			if (lenLine - 2) == 3 {
				convertPurchasePrice, _ := strconv.ParseFloat(line[3], 64)
				purchasePrice = convertPurchasePrice
			}
			if (lenLine - 1) == 4 {
				convertSellingPrice, _ := strconv.ParseFloat(line[4], 64)
				sellingPrice = convertSellingPrice
			}
			data := Item{
				Sku:           line[0],
				Name:          line[1],
				Stock:         stock,
				PurchasePrice: purchasePrice,
				SellingPrice:  sellingPrice,
			}
			insertDetail := db.Create(&data)
			if insertDetail.Error != nil {
				tx.Rollback()
				return []Item{}, insertDetail.Error
			}
			dataItems = append(dataItems, data)
		}
	}
	tx.Commit()
	return dataItems, nil
}
