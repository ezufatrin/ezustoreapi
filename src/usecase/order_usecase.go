package usecase

import (
	"errors"
	"ezustore/src/domain"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderUsecase(o domain.OrderRepository, p domain.ProductRepository) domain.OrderUsecase {
	return &orderUsecase{orderRepo: o, productRepo: p}
}

func (u *orderUsecase) Checkout(userID uint, items []domain.OrderItemInput) error {
	var orderItems []domain.OrderItem
	var total float64

	for _, it := range items {
		prod, err := u.productRepo.GetByID(it.ProductID)
		if err != nil {
			return err
		}
		if prod.Stock < it.Quantity {
			return errors.New("stock not enough: " + prod.Name)
		}
		subtotal := float64(it.Quantity) * prod.Price
		total += subtotal
		orderItems = append(orderItems, domain.OrderItem{
			ProductID: it.ProductID,
			Quantity:  it.Quantity,
			Subtotal:  subtotal,
		})
		prod.Stock -= it.Quantity
		if err := u.productRepo.Update(prod); err != nil {
			return err
		}
	}

	order := &domain.Order{UserID: userID, Total: total, Items: orderItems}
	return u.orderRepo.Create(order)
}

func (u *orderUsecase) GetHistory(userID uint) ([]domain.Order, error) {
	return u.orderRepo.GetByUserID(userID)
}

func (u *orderUsecase) GetDetail(userID, orderID uint) (*domain.Order, error) {
	o, err := u.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, err
	}
	if o.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	return o, nil
}
