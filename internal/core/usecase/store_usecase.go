package usecase

import (
	"errors"

	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
)

type storeUsecase struct {
	storeRepo port.StoreRepository
}

func NewStoreUsecase(storeRepo port.StoreRepository) port.StoreUsecase {
	return &storeUsecase{storeRepo}
}

func (u *storeUsecase) GetStoreByUserID(userID uint) (*domain.Store, error) {
	return u.storeRepo.FindByUserID(userID)
}

func (u *storeUsecase) UpdateStore(userID uint, storeID uint, name, description, profilePicture string) (*domain.Store, error) {
	store, err := u.storeRepo.FindByID(storeID)
	if err != nil {
		return nil, errors.New("store not found")
	}

	if store.UserID != userID {
		return nil, errors.New("unauthorized to update this store")
	}

	if name != "" {
		store.Name = name
	}
	if description != "" {
		store.Description = description
	}
	if profilePicture != "" {
		store.ProfilePicture = profilePicture
	}

	err = u.storeRepo.Update(store)
	return store, err
}
