package usecase

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
)

type addressUsecase struct {
	addressRepo port.AddressRepository
}

func NewAddressUsecase(addressRepo port.AddressRepository) port.AddressUsecase {
	return &addressUsecase{addressRepo}
}

func (u *addressUsecase) GetAddressesByUserID(userID uint) ([]domain.Address, error) {
	return u.addressRepo.FindByUserID(userID)
}

func (u *addressUsecase) CreateAddress(userID uint, title, addressDetails, provID, cityID string, isPrimary bool) (*domain.Address, error) {
	address := &domain.Address{
		UserID:         userID,
		Title:          title,
		AddressDetails: addressDetails,
		ProvID:         provID,
		CityID:         cityID,
		IsPrimary:      isPrimary,
	}

	err := u.addressRepo.Create(address)
	return address, err
}
