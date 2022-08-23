package routes

import (
	"golang-base-code/src/app/utilities"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomContext struct {
	echo.Context
}

type routeHandler struct {
	connection *gorm.DB
	restricted *echo.Group
	open       *echo.Echo
}

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func AppRoutes(e *echo.Echo, connection *gorm.DB) {
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})
	e.Validator = &RequestValidator{validator: validator.New()}

	restricted := e.Group("/restricted")
	echoConfig := middleware.JWTConfig{
		SigningKey: []byte(utilities.GetEnvValue("JWT_TOKEN_SECRET")),
	}
	restricted.Use(middleware.JWTWithConfig(echoConfig))
	restricted.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	routeHandler := &routeHandler{
		connection: connection,
		open:       e,
		restricted: restricted,
	}

	bookRoutes(routeHandler)
	userRoutes(routeHandler)
	authRoutes(routeHandler)
}
