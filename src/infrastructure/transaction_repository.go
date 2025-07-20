package infrastructure

import (
	"ezustore/src/domain"
	"ezustore/src/infrastructure/database"
)

type transactionRepo struct{}

func NewTransactionRepo(db interface{}) domain.TransactionRepository { return &transactionRepo{} }

func (r *transactionRepo) Create(t *domain.Transaction) error {
	return database.DB.Create(t).Error
}

func (r *transactionRepo) GetByUser(userID uint) ([]domain.Transaction, error) {
	var txs []domain.Transaction
	err := database.DB.Preload("Details").Where("user_id = ?", userID).Order("created_at desc").Find(&txs).Error
	return txs, err
}

func (r *transactionRepo) GetByID(id uint) (*domain.Transaction, error) {
	var t domain.Transaction
	err := database.DB.Preload("Details").First(&t, id).Error
	return &t, err
}
