package authController

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/rodrigopmatias/ligistic/framework/config"
	"github.com/rodrigopmatias/ligistic/framework/db"
	"github.com/rodrigopmatias/ligistic/framework/router"
	"github.com/rodrigopmatias/ligistic/models"
)

type UserClaims struct {
	UserID string
	jwt.StandardClaims
}

type AuthenticatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateResult struct {
	Token string `json:"token"`
	router.ContentResult
}

func Authenticate(payload *AuthenticatePayload) (*AuthenticateResult, error) {
	var err error
	var user models.User
	var result *AuthenticateResult

	db, err := db.Open(db.OpenConfig{})

	if err != nil {
		return nil, err
	}

	db.First(&user, "email = ?", payload.Email)

	if user.ID != "" && user.IsActive {
		claims := UserClaims{
			UserID: user.ID,
			StandardClaims: jwt.StandardClaims{
				Issuer: "MyOrg",
			},
		}

		cnf := config.New()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signed, _ := token.SignedString([]byte(cnf.JwtSecret))

		result = &AuthenticateResult{
			Token: signed,
			ContentResult: router.ContentResult{
				Ok: true,
			},
		}
	} else {
		err = errors.New("user or password not match")
	}

	return result, err
}
