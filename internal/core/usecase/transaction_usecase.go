package usecase

import (
	"errors"

	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
)

type transactionUsecase struct {
	transactionRepo port.TransactionRepository
	productRepo     port.ProductRepository
}

func NewTransactionUsecase(transactionRepo port.TransactionRepository, productRepo port.ProductRepository) port.TransactionUsecase {
	return &transactionUsecase{transactionRepo, productRepo}
}

func (u *transactionUsecase) Checkout(userID uint, productID uint, quantity int, paymentMethod string) (*domain.Transaction, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	product, err := u.productRepo.FindByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if product.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	// Calculate total price
	totalPrice := product.Price * float64(quantity)

	transaction := &domain.Transaction{
		UserID:        userID,
		TotalPrice:    totalPrice,
		Quantity:      quantity,
		PaymentMethod: paymentMethod,
		Status:        "completed", // Automatically completed based on requirements
	}

	log := domain.ProductLog{
		ProductID:    product.ID,
		ProductName:  product.Name,
		ProductPrice: product.Price,
	}

	// Update stock
	product.Stock -= quantity
	err = u.productRepo.Update(product)
	if err != nil {
		return nil, errors.New("failed to update stock")
	}

	// Create transaction and logs
	err = u.transactionRepo.CreateWithLogs(transaction, []domain.ProductLog{log})
	if err != nil {
		// Rollback stock if needed in a real app, though using DB transactions is better.
		return nil, errors.New("failed to create transaction")
	}

	return transaction, nil
}

func (u *transactionUsecase) GetUserTransactions(userID uint) ([]domain.Transaction, error) {
	return u.transactionRepo.FindByUserID(userID)
}
