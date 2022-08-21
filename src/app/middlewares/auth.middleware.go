package middlewares

import (
	"errors"
	"golang-base-code/src/app/models"
	"golang-base-code/src/app/utilities"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthMiddleware interface {
	Login(c echo.Context) (models.JWTClaims, error)
}

type authMiddlewareBuilder struct {
	DB *gorm.DB
}

func AuthConnectionMw(connection *gorm.DB) AuthMiddleware {
	return &authMiddlewareBuilder{
		DB: connection,
	}
}

func (m *authMiddlewareBuilder) Login(c echo.Context) (models.JWTClaims, error) {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon_doe" && password == "password" {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Doe"
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(utilities.GetEnvValue("JWT_TOKEN_SECRET")))
		if err != nil {
			return models.JWTClaims{}, errors.New("failed to generate token")
		}

		return models.JWTClaims{Name: "Jon Doe", Token: t}, nil
	}

	return models.JWTClaims{}, errors.New("user not found")
}
