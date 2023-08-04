package usecase

import (
	"enigma-laundry-apps/utils/security"
	"fmt"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	usecase UserUseCase
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(username string, password string) (string, error) {
	user, err := a.usecase.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed generate token")
	}
	return token, nil
}

func NewAuthUseCase(usecase UserUseCase) AuthUseCase {
	return &authUseCase{
		usecase: usecase,
	}
}
