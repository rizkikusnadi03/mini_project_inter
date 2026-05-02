package port

import (
	"backend_golang/internal/core/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByPhone(phone string) (*domain.User, error)
}

type StoreRepository interface {
	Create(store *domain.Store) error
	Update(store *domain.Store) error
	FindByID(id uint) (*domain.Store, error)
	FindByUserID(userID uint) (*domain.Store, error)
}

type AuthUsecase interface {
	Register(name, email, phone, password string) (*domain.User, error)
	Login(email, password string) (string, error)
}
