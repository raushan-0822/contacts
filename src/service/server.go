package contacts

import (
	"github.com/raushan-0822/contacts/service/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

// CustomValidator a custom validator for request
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// StartContactAPI starts contact apis
func StartContactAPI() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1024KB"))
	e.Use(middleware.Secure())
	// e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
	// Generator: middlewares.RequestID,
	// }))
	e.Use(middlewares.RequestID)
	e.Use(middleware.RemoveTrailingSlash())
	g := e.Group("/accounts/:accountSid/")
	g.Use(middleware.BasicAuth(middlewares.BasicAuth))
	AddRoutes(g)
	e.Logger.Fatal(e.Start(":8080"))
}
