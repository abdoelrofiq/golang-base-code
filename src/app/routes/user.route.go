package routes

import (
	controller "golang-base-code/src/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func userRoutes(e *echo.Echo, connection *gorm.DB) {
	users := controller.UsersHandler(connection)
	e.GET("/users", users.GetAll)
}
