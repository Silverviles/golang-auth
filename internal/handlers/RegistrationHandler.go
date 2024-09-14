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
	"time"
)

func RegisterUser(c echo.Context) error {
	var user models.UserDao
	body, _ := io.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	DB := db.GetDatabaseConnection()
	var count int
	countSQLStatement := "SELECT COUNT(*) FROM users WHERE username = $1"
	err = DB.QueryRow(countSQLStatement, user.Username).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return c.JSON(http.StatusConflict, "User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.DateCreated = time.Now().Format("2006-01-02 15:04:05")

	insertSQLStatement, err := DB.Prepare("INSERT INTO users (username, password, first_name, last_name, email, phone, date_created) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = insertSQLStatement.Exec(user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.Phone, user.DateCreated)
	if err != nil {
		return err
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	response := models.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return c.JSON(http.StatusCreated, response)
}
