package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-app/internal/handlers"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
	}))

	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	s := &http.Server{
		Addr: ":8080",
	}
	e.Logger.Fatal(e.StartServer(s))
}
