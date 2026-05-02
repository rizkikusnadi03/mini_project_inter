package repository

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) port.TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateWithLogs(transaction *domain.Transaction, logs []domain.ProductLog) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		for i := range logs {
			logs[i].TransactionID = transaction.ID
			if err := tx.Create(&logs[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *transactionRepository) FindByUserID(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("ProductLogs").Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Preload("ProductLogs").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
