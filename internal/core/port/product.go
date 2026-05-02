package port

import (
	"backend_golang/internal/core/domain"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(id uint) error
	FindAll() ([]domain.Category, error)
	FindByID(id uint) (*domain.Category, error)
}

type CategoryUsecase interface {
	CreateCategory(name, description string) (*domain.Category, error)
	UpdateCategory(id uint, name, description string) (*domain.Category, error)
	DeleteCategory(id uint) error
	GetCategories() ([]domain.Category, error)
}

type ProductRepository interface {
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	FindAll(page, limit int, search string) ([]domain.Product, int64, error)
	FindByID(id uint) (*domain.Product, error)
}

type ProductUsecase interface {
	CreateProduct(userID uint, categoryID uint, name, description string, price float64, stock int, image string) (*domain.Product, error)
	GetProducts(page, limit int, search string) ([]domain.Product, int64, error)
}
