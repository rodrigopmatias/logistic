package authController

import (
	"github.com/rodrigopmatias/ligistic/framework/db"
	"github.com/rodrigopmatias/ligistic/models"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type RegisterResult struct {
	models.User
}

func (payload RegisterPayload) PasswordMatch() bool {
	return payload.Password == payload.ConfirmPassword
}

func Register(payload *RegisterPayload) (*RegisterResult, error) {
	var countActive int64
	var countAdmin int64
	db, err := db.Open(db.OpenConfig{})

	if err != nil {
		return nil, err
	}

	db.Model(&models.User{}).Where("is_active = ?", true).Count(&countActive)
	db.Model(&models.User{}).Where("is_active = ?", true).Where("is_admin = ?", true).Count(&countAdmin)

	user := models.User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		IsActive:  (countActive == 0),
		IsAdmin:   (countAdmin == 0),
	}

	user.SetPassword(payload.Password)
	result := db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &RegisterResult{
		User: user,
	}, nil
}
