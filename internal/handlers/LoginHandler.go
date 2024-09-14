package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-app/internal/db"
	"go-app/internal/helper"
	"go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func LoginUser(c echo.Context) error {
	var user models.UserDao
	var hashedPassword string
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	DB := db.GetDatabaseConnection()
	userSQLStatement := "SELECT id, password, email FROM users WHERE username = ?"
	err = DB.QueryRow(userSQLStatement, user.Username).Scan(&user.ID, &hashedPassword, &user.Email)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	response := models.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return c.JSON(http.StatusOK, response)
}
