package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-app/internal/models"
	"go-app/internal/services"
	"io"
	"net/http"
	"strconv"
)

func AddProduct(c echo.Context) error {
	var product models.Product
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &product)
	if err != nil {
		return err
	}

	addedProduct, err := services.AddProduct(&product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, addedProduct)
}

func GetProducts(c echo.Context) error {
	products, err := services.GetProducts()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func GetProductByID(c echo.Context) error {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return err
	}

	product, err := services.GetProductByID(productID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	var product models.Product
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &product)
	if err != nil {
		return err
	}

	updatedProduct, err := services.UpdateProduct(&product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c echo.Context) error {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return err
	}

	err = services.DeleteProduct(productID)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
