package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-app/internal/customMiddleware"
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

	userGroup := e.Group("/user")
	userGroup.Use(customMiddleware.CheckRole("user"))
	userGroup.GET("/products", handlers.GetProducts)
	userGroup.GET("/products/:id", handlers.GetProductByID)

	adminGroup := e.Group("/admin")
	adminGroup.Use(customMiddleware.CheckRole("admin"))
	adminGroup.GET("/products", handlers.GetProducts)
	adminGroup.GET("/products/:id", handlers.GetProductByID)
	adminGroup.POST("/products", handlers.AddProduct)
	adminGroup.PUT("/products/:id", handlers.UpdateProduct)
	adminGroup.DELETE("/products/:id", handlers.DeleteProduct)

	s := &http.Server{
		Addr: ":8080",
	}
	e.Logger.Fatal(e.StartServer(s))
}
