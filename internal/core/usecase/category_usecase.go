package usecase

import (
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
)

type categoryUsecase struct {
	categoryRepo port.CategoryRepository
}

func NewCategoryUsecase(categoryRepo port.CategoryRepository) port.CategoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (u *categoryUsecase) CreateCategory(name, description string) (*domain.Category, error) {
	category := &domain.Category{
		Name:        name,
		Description: description,
	}
	err := u.categoryRepo.Create(category)
	return category, err
}

func (u *categoryUsecase) UpdateCategory(id uint, name, description string) (*domain.Category, error) {
	category, err := u.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Description = description

	err = u.categoryRepo.Update(category)
	return category, err
}

func (u *categoryUsecase) DeleteCategory(id uint) error {
	return u.categoryRepo.Delete(id)
}

func (u *categoryUsecase) GetCategories() ([]domain.Category, error) {
	return u.categoryRepo.FindAll()
}
