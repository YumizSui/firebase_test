package router

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

// Routing ルーティングを張ります
func Routing(e *echo.Echo) {

	// Static Files
	// e.Static("/", "client/dist")

	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong\n")
		})
		api.GET("/private", private, AuthMiddleware)
		api.GET("/public", public)
	}

}

func public(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	return c.String(http.StatusOK, "public is success.")
}

func private(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	return c.String(http.StatusOK, "private is success.")
}
