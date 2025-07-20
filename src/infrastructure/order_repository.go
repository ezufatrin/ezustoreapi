package infrastructure

import (
	"ezustore/src/domain"
	"ezustore/src/infrastructure/database"
)

type orderRepository struct{}

func NewOrderRepository(db interface{}) domain.OrderRepository { return &orderRepository{} }

func (r *orderRepository) Create(o *domain.Order) error {
	return database.DB.Create(o).Error
}

func (r *orderRepository) GetByUserID(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := database.DB.Preload("Items").Preload("Items.Product").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var o domain.Order
	err := database.DB.Preload("Items").Preload("Items.Product").First(&o, id).Error
	return &o, err
}
