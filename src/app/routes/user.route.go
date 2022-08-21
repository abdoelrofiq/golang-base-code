package routes

import (
	controller "golang-base-code/src/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func userRoutes(e *echo.Echo, restricted *echo.Group, connection *gorm.DB) {
	user := controller.UserHandler(connection)
	restricted.GET("/users", user.GetAll)
}
