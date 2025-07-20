package domain

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	Category  string    `gorm:"type:varchar(50)" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]Product, error)
}

type ProductUsecase interface {
	Create(name string, price float64, stock int, category string) error
	Update(id uint, name string, price float64, stock int, category string) error
	Delete(id uint) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]Product, error)
}
