package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-app/internal/constants"
	"go-app/internal/customMiddleware"
	"go-app/internal/handlers"
	"net/http"
)

func main() {
	e := echo.New()

	err := godotenv.Load(constants.EnvironmentVariablePath)
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	//userGroup := e.Group("/user")
	//userGroup.Use(customMiddleware.CheckRole("user"))
	//userGroup.GET("/products", handlers.GetProducts)
	//userGroup.GET("/products/:id", handlers.GetProductByID)
	//
	//adminGroup := e.Group("/admin")
	//adminGroup.Use(customMiddleware.CheckRole("admin"))
	//adminGroup.GET("/products", handlers.GetProducts)
	//adminGroup.GET("/products/:id", handlers.GetProductByID)
	//adminGroup.POST("/addProducts", handlers.AddProduct)
	//adminGroup.PUT("/products/:id", handlers.UpdateProduct)
	//adminGroup.DELETE("/products/:id", handlers.DeleteProduct)

	e.GET("/products", handlers.GetProducts, customMiddleware.CheckRole("user"))
	e.GET("/products/:id", handlers.GetProductByID, customMiddleware.CheckRole("user"))

	e.POST("/products", handlers.AddProduct, customMiddleware.CheckRole("admin"))
	e.PUT("/products/:id", handlers.UpdateProduct, customMiddleware.CheckRole("admin"))
	e.DELETE("/products/:id", handlers.DeleteProduct, customMiddleware.CheckRole("admin"))

	s := &http.Server{
		Addr: ":8080",
	}
	e.Logger.Fatal(e.StartServer(s))
}
