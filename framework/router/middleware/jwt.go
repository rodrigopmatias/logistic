package middleware

import (
	"fmt"
	"log"
	"regexp"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/rodrigopmatias/ligistic/framework/config"
	"github.com/rodrigopmatias/ligistic/framework/db"
	"github.com/rodrigopmatias/ligistic/framework/router/context"
	"github.com/rodrigopmatias/ligistic/models"
)

func JwtAuthorizationMiddleware(ctx *context.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	parser, _ := regexp.Compile(`^(Bearer|JWT) (.*)$`)

	if authorization != "" && parser.MatchString(authorization) {
		result := parser.FindAllStringSubmatch(authorization, -1)
		token, err := jwt.Parse(result[0][2], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok == !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			cnf := config.New()
			return []byte(cnf.JwtSecret), nil
		})

		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var user models.User
				db, _ := db.Open(db.OpenConfig{})
				db.First(&user, "id = ?", claims["UserID"])

				if user.ID != "" && user.IsActive {
					ctx.Values["user"] = user
				}
			} else {
				log.Println("Token are not valid")
			}
		} else {
			log.Println(err)
		}
	}
}
