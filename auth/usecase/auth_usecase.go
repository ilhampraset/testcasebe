package usecase

import (
	"github.com/ilhampraset/testcasebe/domain"
)

type AuthUseCase struct {
	repo domain.UserRepository
}

func NewAuthUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &AuthUseCase{repo}
}

func (uc *AuthUseCase) VerifyLogin(username, password string) (bool, error) {
	_, err := uc.repo.GetUserWithCredential(username, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uc *AuthUseCase) ParseToken(accessToken string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (uc *AuthUseCase) Me(username string) (*domain.User, error) {
	byUsername, _ := uc.repo.GetUserByUsername(username)
	return uc.repo.GetUserByID(byUsername.ID)
}
