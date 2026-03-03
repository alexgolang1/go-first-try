package mw

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Token")
		if token != "my-auth-token" {
			ctx.JSON(http.StatusUnauthorized, "Not allowed")
		}
		return next(ctx)
	}
}
