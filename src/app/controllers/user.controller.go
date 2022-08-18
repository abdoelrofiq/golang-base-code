package controllers

import (
	response "golang-base-code/src/app/core"
	middleware "golang-base-code/src/app/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	Middleware middleware.UserMiddleware
}

func UserHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		Middleware: middleware.UserConnectionMw(db),
	}
}

func (u *userHandler) GetAll(c echo.Context) error {
	user, err := u.Middleware.Fetch(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success get all data", user))
}
