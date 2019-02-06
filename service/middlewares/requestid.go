package middlewares

import (
	"contacts/utils"

	"github.com/labstack/echo"
)

//RequestID middleware for setting request id
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		c.Set("RequestID", utils.RandStringRunes(32))
		return next(c)
	})
}
