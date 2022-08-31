package controllers

import (
	response "golang-base-code/src/app/core"
	service "golang-base-code/src/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type authHandler struct {
	Service service.AuthService
}

func AuthHandler(db *gorm.DB) *authHandler {
	return &authHandler{
		Service: service.AuthConnectionMw(db),
	}
}

func (a *authHandler) Login(c echo.Context) error {
	token, err := a.Service.Login(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success to login", token))
}

func (a *authHandler) TokenValueExtraction(c echo.Context) {
	a.Service.TokenValueExtraction(c)
}
