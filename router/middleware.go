package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

// AuthMiddleware 認証のmiddleware
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))

		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error initializing app: %v", err))
		}

		auth, err := app.Auth(context.Background())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error getting auth client: %v", err))
		}

		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		_, err = auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("error verifying ID token: %v", err))
		}

		return next(c)
	}
}
