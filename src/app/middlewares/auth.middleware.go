package middlewares

import (
	"errors"
	model "golang-base-code/src/app/models"
	"golang-base-code/src/app/utilities"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var user model.User

type AuthMiddleware interface {
	Login(c echo.Context) (model.JWTClaims, error)
	TokenValueExtraction(c echo.Context)
}

type authMiddlewareBuilder struct {
	DB *gorm.DB
}

func AuthConnectionMw(connection *gorm.DB) AuthMiddleware {
	return &authMiddlewareBuilder{
		DB: connection,
	}
}

func (conn *authMiddlewareBuilder) Login(c echo.Context) (model.JWTClaims, error) {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "abdul@gmail.com" && password == "password" {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(utilities.GetEnvValue("JWT_TOKEN_SECRET")))
		if err != nil {
			return model.JWTClaims{}, errors.New("failed to generate token")
		}

		return model.JWTClaims{Username: username, Token: t}, nil
	}

	return model.JWTClaims{}, errors.New("user not found")
}

func (conn *authMiddlewareBuilder) TokenValueExtraction(c echo.Context) {
	tokenExtraction := c.Get("user").(*jwt.Token)
	claims := tokenExtraction.Claims.(jwt.MapClaims)
	conn.DB.Where("email = ?", claims["username"]).Find(&user)

	c.Set("currentUser", user)
}
