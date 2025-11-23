package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null;size:100"`
	Surname  string `gorm:"not null;size:100"`
	Email    string `gorm:"not null;unique;size:100"`
	Username string `gorm:"not null;unique;size:100"`
	Password string `gorm:"not null;size:100"`
	Todos    []Todo `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	RegisterUser(c context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	DeleteUser(c context.Context, id int) error
	UpdateUser(c context.Context, user *User) error
}
