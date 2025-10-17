package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorHandler handles errors consistently
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	// Log the error for debugging
	c.Logger().Error(err)

	// Return a consistent JSON response
	if !c.Response().Committed {
		if c.Request().Header.Get("Content-Type") == "application/json" {
			c.JSON(code, map[string]string{"error": message})
		} else {
			c.String(code, message)
		}
	}
}

// AuthMiddleware - Placeholder for future authentication
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement authentication
		// For now, allow all requests to pass through
		return next(c)
	}
}

// RequireAuth - Middleware for routes that require authentication
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Verify JWT token or other auth method
		// For now, allow all requests to pass through
		return next(c)
	}
}
