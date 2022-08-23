package routes

import (
	controller "golang-base-code/src/app/controllers"
)

func authRoutes(routeHandler *routeHandler) {
	auth := controller.AuthHandler(routeHandler.connection)
	routeHandler.open.POST("/login", auth.Login)
}
