package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-app/internal/models"
	"go-app/internal/services"
	"io"
	"net/http"
)

func LoginUser(c echo.Context) error {
	var user models.UserDao

	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	loginUser, err := services.LoginUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, loginUser)
}
