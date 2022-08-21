package routes

import (
	controller "golang-base-code/src/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func authRoutes(e *echo.Echo, restricted *echo.Group, connection *gorm.DB) {
	auth := controller.AuthHandler(connection)
	e.POST("/login", auth.Login)
}
