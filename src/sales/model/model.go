package model

import "time"

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
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Sku          string     `json:"sku"`
	Name         string     `json:"name"`
	Qty          int        `json:"qty"`
	SellingPrice int        `json:"selling_price"`
	Total        int        `json:"total"`
	Note         string     `json:"note"`
	SalesID      uint       `json:"sales_id"`
}
