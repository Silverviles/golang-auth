package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-app/internal/db"
	"go-app/internal/middleware"
	"go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func LoginUser(c echo.Context) error {
	var user models.UserDTO
	var existingUser models.UserDTO
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	DB := db.GetDatabaseConnection()
	userSQLStatement := "SELECT * FROM users WHERE username = $1"
	err = DB.QueryRow(userSQLStatement, user.Username).Scan(&existingUser.Username, &existingUser.Password)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateToken(existingUser.ID)
	if err != nil {
		return err
	}

	response := models.UserDTO{
		ID:       existingUser.ID,
		Username: existingUser.Username,
		Email:    existingUser.Email,
		Token:    token,
	}

	return c.JSON(http.StatusOK, response)
}
