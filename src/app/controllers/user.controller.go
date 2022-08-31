package controllers

import (
	response "golang-base-code/src/app/core"
	service "golang-base-code/src/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	Service service.UserService
}

func UserHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		Service: service.UserConnectionMw(db),
	}
}

func (u *userHandler) GetAll(c echo.Context) error {
	user, err := u.Service.Fetch(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success get all data", user))
}
