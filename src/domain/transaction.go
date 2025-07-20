package domain

import "time"

type Transaction struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	UserID     uint              `json:"user_id"`
	TotalPrice float64           `json:"total_price"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Details    []TransactionItem `gorm:"foreignKey:TransactionID" json:"details"`
}

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	Subtotal      float64 `json:"subtotal"`
}

type TransactionRepository interface {
	Create(t *Transaction) error
	GetByUser(userID uint) ([]Transaction, error)
	GetByID(id uint) (*Transaction, error)
}

type TransactionUsecase interface {
	Create(userID uint, items []CreateTransactionItemInput) error
	GetByUser(userID uint) ([]Transaction, error)
	GetByID(userID, trxID uint) (*Transaction, error)
}

type CreateTransactionItemInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
