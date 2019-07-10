package purchasemodel

type Purchase struct {
	ID            int
	DateTime      string `json:"date_time"`
	ReceiptNumber string `json:"receipt_number"`
}
