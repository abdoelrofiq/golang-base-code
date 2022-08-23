package routes

import (
	controller "golang-base-code/src/app/controllers"

	"github.com/labstack/echo/v4"
)

func userRoutes(routeHandler *routeHandler) {
	user := controller.UserHandler(routeHandler.connection)
	auth := controller.AuthHandler(routeHandler.connection)
	routeHandler.restricted.GET("/users", func(c echo.Context) error {
		cc := c.(*CustomContext)
		auth.TokenValueExtraction(cc)
		return user.GetAll(cc)
	})
}
