package controllers

import (
	response "golang-base-code/src/app/core"
	middleware "golang-base-code/src/app/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authHandler struct {
	Middleware middleware.AuthMiddleware
}

func AuthHandler(db *gorm.DB) *authHandler {
	return &authHandler{
		Middleware: middleware.AuthConnectionMw(db),
	}
}

func (a *authHandler) Login(c echo.Context) error {
	token, err := a.Middleware.Login(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success to login", token))
}

func (a *authHandler) TokenValueExtraction(c echo.Context) {
	a.Middleware.TokenValueExtraction(c)
}
