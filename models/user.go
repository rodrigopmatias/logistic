package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Email     string `json:"email" gorm:"varchar(200);not null;unique"`
	FirstName string `json:"firstName" gorm:"varchar(100);not null"`
	LastName  string `json:"lastName" gorm:"varchar(100);not null"`
	Password  string `json:"-"`
	IsActive  bool   `json:"isActive"`
	IsAdmin   bool   `json:"isAdmin"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Base.BeforeCreate(tx)
	_, err := bcrypt.Cost([]byte(user.Password))

	if err != nil {
		return err
	}

	return nil
}

func (user *User) SetPassword(plain string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(password)
	return nil
}
