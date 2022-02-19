package domain

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
	Merchant  Merchant  `json:"merchant"`
}

type UserRepository interface {
	GetUserByID(ID int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserWithCredential(username, password string) (*User, error)
}

type UserUseCase interface {
	VerifyLogin(username, password string) (bool, error)
	ParseToken(accessToken string) (*User, error)
	Me(username string) (*User, error)
}
