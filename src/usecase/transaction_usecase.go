package usecase

import (
	"errors"
	"ezustore/src/domain"
)

type transactionUsecase struct {
	repo        domain.TransactionRepository
	productRepo domain.ProductRepository
}

func NewTransactionUsecase(r domain.TransactionRepository, p domain.ProductRepository) domain.TransactionUsecase {
	return &transactionUsecase{repo: r, productRepo: p}
}

func (u *transactionUsecase) Create(userID uint, items []domain.CreateTransactionItemInput) error {
	var details []domain.TransactionItem
	var total float64

	for _, it := range items {
		prod, err := u.productRepo.GetByID(it.ProductID)
		if err != nil {
			return err
		}
		if prod.Stock < it.Quantity {
			return errors.New("stok kurang: " + prod.Name)
		}
		subtotal := float64(it.Quantity) * prod.Price
		total += subtotal
		details = append(details, domain.TransactionItem{
			ProductID:   prod.ID,
			ProductName: prod.Name,
			Price:       prod.Price,
			Quantity:    it.Quantity,
			Subtotal:    subtotal,
		})
		prod.Stock -= it.Quantity
		if err := u.productRepo.Update(prod); err != nil {
			return err
		}
	}

	trx := &domain.Transaction{UserID: userID, TotalPrice: total, Details: details}
	return u.repo.Create(trx)
}

func (u *transactionUsecase) GetByUser(userID uint) ([]domain.Transaction, error) {
	return u.repo.GetByUser(userID)
}

func (u *transactionUsecase) GetByID(userID, trxID uint) (*domain.Transaction, error) {
	t, err := u.repo.GetByID(trxID)
	if err != nil {
		return nil, err
	}
	if t.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	return t, nil
}
