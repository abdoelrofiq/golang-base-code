package routes

import (
	"golang-base-code/src/app/utilities"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func AppRoutes(e *echo.Echo, connection *gorm.DB) {
	e.Use(middleware.Recover())

	restricted := e.Group("/restricted")
	echoConfig := middleware.JWTConfig{
		SigningKey: []byte(utilities.GetEnvValue("JWT_TOKEN_SECRET")),
	}
	restricted.Use(middleware.JWTWithConfig(echoConfig))

	e.Validator = &RequestValidator{validator: validator.New()}

	bookRoutes(e, restricted, connection)
	userRoutes(e, restricted, connection)
	authRoutes(e, restricted, connection)
}
