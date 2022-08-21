package routes

import (
	controller "golang-base-code/src/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func bookRoutes(e *echo.Echo, restricted *echo.Group, connection *gorm.DB) {
	boostore := controller.BookHandler(connection)
	e.GET("/books", boostore.GetAll)
	e.GET("/books/:id", boostore.GetById)
	e.POST("/books", boostore.Create)
	e.PUT("/books", boostore.Update)
	e.DELETE("/books/:id", boostore.Delete)
}
