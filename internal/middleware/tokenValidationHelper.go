package middleware

//import (
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"strings"
//)
//
//func JWTMiddleware(next echo.HandlerFunc) echo.MiddlewareFunc {
//	return func(c echo.Context) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			if c.Path() == "/login" || c.Path() == "/register" {
//				return next(c)
//			}
//
//			authHeader := c.Request().Header.Get("Authorization")
//			if authHeader == "" {
//				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
//			}
//
//			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//			if tokenString == authHeader {
//				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
//			}
//
//			token, err := ValidateToken(tokenString)
//			if err != nil {
//				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
//			}
//		}
//	}
//}
