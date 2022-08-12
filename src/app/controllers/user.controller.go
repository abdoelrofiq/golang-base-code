package controllers

import (
	response "golang-base-code/src/app/core"
	middleware "golang-base-code/src/app/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	Repo middleware.UserMiddleware
}

func UsersHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		Repo: middleware.UserConnectionMw(db),
	}
}

func (u *userHandler) GetAll(c echo.Context) error {
	// Get all users and validate if no error
	users, err := u.Repo.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success get all data", users))
}
