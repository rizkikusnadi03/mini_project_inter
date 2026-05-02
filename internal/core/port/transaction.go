package port

import (
	"backend_golang/internal/core/domain"
)

type TransactionRepository interface {
	CreateWithLogs(transaction *domain.Transaction, logs []domain.ProductLog) error
	FindByUserID(userID uint) ([]domain.Transaction, error)
	FindByID(id uint) (*domain.Transaction, error)
}

type TransactionUsecase interface {
	Checkout(userID uint, productID uint, quantity int, paymentMethod string) (*domain.Transaction, error)
	GetUserTransactions(userID uint) ([]domain.Transaction, error)
}
