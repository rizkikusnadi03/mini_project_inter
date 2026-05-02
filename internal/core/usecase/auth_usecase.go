package usecase

import (
	"errors"
	"backend_golang/internal/core/domain"
	"backend_golang/internal/core/port"
	"backend_golang/pkg/jwtutils"
	"backend_golang/pkg/password"
	"gorm.io/gorm"
)

type authUsecase struct {
	userRepo port.UserRepository
}

func NewAuthUsecase(userRepo port.UserRepository) port.AuthUsecase {
	return &authUsecase{userRepo}
}

func (u *authUsecase) Register(name, email, phone, plainPassword string) (*domain.User, error) {
	// Check if email already exists
	_, err := u.userRepo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Check if phone already exists
	_, err = u.userRepo.FindByPhone(phone)
	if err == nil {
		return nil, errors.New("phone already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := password.HashPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: hashedPassword,
		Role:     "user", // default role
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *authUsecase) Login(email, plainPassword string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !password.CheckPasswordHash(plainPassword, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := jwtutils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
