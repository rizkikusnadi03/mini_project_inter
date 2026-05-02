package repository

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	// Atomic transaction for user and store creation
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		store := &domain.Store{
			UserID: user.ID,
			Name:   user.Name + "'s Store", // Default store name
		}

		if err := tx.Create(store).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
