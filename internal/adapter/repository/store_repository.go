package repository

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) port.StoreRepository {
	return &storeRepository{db}
}

func (r *storeRepository) Create(store *domain.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) Update(store *domain.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) FindByID(id uint) (*domain.Store, error) {
	var store domain.Store
	err := r.db.First(&store, id).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) FindByUserID(userID uint) (*domain.Store, error) {
	var store domain.Store
	err := r.db.Where("user_id = ?", userID).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}
