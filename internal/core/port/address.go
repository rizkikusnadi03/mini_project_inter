package port

import (
	"backend_golang/internal/core/domain"
)

type AddressRepository interface {
	Create(address *domain.Address) error
	FindByUserID(userID uint) ([]domain.Address, error)
}

type AddressUsecase interface {
	GetAddressesByUserID(userID uint) ([]domain.Address, error)
	CreateAddress(userID uint, title, addressDetails, provID, cityID string, isPrimary bool) (*domain.Address, error)
}
