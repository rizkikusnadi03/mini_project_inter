package repository

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) port.AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address *domain.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) FindByUserID(userID uint) ([]domain.Address, error) {
	var addresses []domain.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
