package mysql

import (
	"errors"

	"github.com/ilhampraset/testcasebe/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.UserRepository {
	return &AuthRepository{db}
}

func (u AuthRepository) GetUserByID(ID int) (*domain.User, error) {
	var user *domain.User
	err := u.db.Preload("Merchant").Preload("Merchant.Outlet").First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u AuthRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where("user_name=?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u AuthRepository) GetUserWithCredential(username, password string) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where("user_name=? ", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if !checkPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return user, nil

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
