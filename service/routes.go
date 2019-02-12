package contacts

import (
	"github.com/raushan-0822/contacts/service/handlers"

	"github.com/labstack/echo"
)

// AddRoutes attaches the routes
func AddRoutes(e *echo.Group) {
	e.POST("contacts", handler.ContactHandler)
	e.PUT("contacts/:id", handler.ContactHandler)
	e.GET("contacts/:id", handler.ContactHandler)
	e.DELETE("contacts/:id", handler.ContactHandler)
	e.GET("contacts/search/:name", handler.ContactSearchHandler)
}
