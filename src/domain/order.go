package domain

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	User      User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Total     float64     `json:"total"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByUserID(userID uint) ([]Order, error)
	GetByID(id uint) (*Order, error)
}

type OrderUsecase interface {
	Checkout(userID uint, items []OrderItemInput) error
	GetHistory(userID uint) ([]Order, error)
	GetDetail(userID, orderID uint) (*Order, error)
}

type OrderItemInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
