package domain

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null;size:100"`
	Surname      string `gorm:"not null;size:100"`
	Email        string `gorm:"not null;unique;size:100"`
	Username     string `gorm:"not null;unique;size:100"`
	PasswordHash string `gorm:"not null;size:300"`
	Todos        []Todo `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	RegisterUser(c context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserByID(ctx context.Context, id uint) (*User, error)
}

var ErrUserAlreadyExists = errors.New("username or email already exists")
