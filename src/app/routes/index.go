package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AppRoutes(e *echo.Echo, connection *gorm.DB) {
	bookRoutes(e, connection)
	userRoutes(e, connection)
}
