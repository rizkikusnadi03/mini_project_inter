package usecase

import (
	"errors"

	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
)

type productUsecase struct {
	productRepo port.ProductRepository
	storeRepo   port.StoreRepository
}

func NewProductUsecase(productRepo port.ProductRepository, storeRepo port.StoreRepository) port.ProductUsecase {
	return &productUsecase{productRepo, storeRepo}
}

func (u *productUsecase) CreateProduct(userID uint, categoryID uint, name, description string, price float64, stock int, image string) (*domain.Product, error) {
	// Need to find the store related to this user
	store, err := u.storeRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("store not found for this user")
	}

	product := &domain.Product{
		StoreID:     store.ID,
		CategoryID:  categoryID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		Image:       image,
	}

	err = u.productRepo.Create(product)
	return product, err
}

func (u *productUsecase) GetProducts(page, limit int, search string) ([]domain.Product, int64, error) {
	return u.productRepo.FindAll(page, limit, search)
}
