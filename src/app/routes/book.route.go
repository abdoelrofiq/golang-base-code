package routes

import (
	controller "golang-base-code/src/app/controllers"
)

func bookRoutes(routeHandler *routeHandler) {
	boostore := controller.BookHandler(routeHandler.connection)
	routeHandler.open.GET("/books", boostore.GetAll)
	routeHandler.open.GET("/books/:id", boostore.GetById)
	routeHandler.restricted.POST("/books", boostore.Create)
	routeHandler.open.PUT("/books", boostore.Update)
	routeHandler.open.DELETE("/books/:id", boostore.Delete)
}
