package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) error
	FindAllUser() ([]model.UserCredential, error)
	FindByUsername(username string) (model.UserCredential, error)
	FindByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.UserCredential, error) {
	return u.repo.List()
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.repo.GetUsername(username)
}

// FindByUsernamePassword implements UserUseCase.
func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.UserCredential, error) {
	return u.repo.GetUsernamePassword(username, password)
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload model.UserCredential) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)

	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
