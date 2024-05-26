package models

import "time"

type OrderDetails struct {
	Id        int     `gorm:"primaryKey" json:"id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`

	Order   Order   `gorm:"foreignKey:OrderId;references:Id"`
	Product Product `gorm:"foreignKey:ProductId;references:Id"`
}

type Order struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	OrderDate time.Time `json:"order_date"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
}

type OrderDetailPayload struct {
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderPayload struct {
	UserId          int                  `json:"user_id"`
	OrderDetailList []OrderDetailPayload `json:"order_detail_list"`
}

const (
	PENDING_STATUS    = "Pending"
	PROCESSING_STATUS = "Processing"
	COMPLETED_STATUS  = "Completed"
)
