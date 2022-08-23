package routes

import (
	controller "golang-base-code/src/app/controllers"
)

func userRoutes(routeHandler *routeHandler) {
	user := controller.UserHandler(routeHandler.connection)
	routeHandler.restricted.GET("/users", user.GetAll)
}
