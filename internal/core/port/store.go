package port

import (
	"backend_golang/internal/core/domain"
)

type StoreUsecase interface {
	GetStoreByUserID(userID uint) (*domain.Store, error)
	UpdateStore(userID uint, storeID uint, name, description, profilePicture string) (*domain.Store, error)
}
