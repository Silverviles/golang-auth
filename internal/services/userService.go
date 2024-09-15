package services

import (
	"go-app/internal/constants"
	"go-app/internal/db"
	"go-app/internal/middleware"
	"go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func isUserExists(username string) (bool, error) {
	DB := db.GetDatabaseConnection()

	var count int
	countSQLStatement := "SELECT COUNT(*) FROM users WHERE username = ?"
	err := DB.QueryRow(countSQLStatement, username).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func RegisterUser(user *models.UserDao) (*models.UserDTO, error) {
	isExists, err := isUserExists(user.Username)
	if err != nil || isExists {
		return nil, err
	}
	DB := db.GetDatabaseConnection()

	transaction, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	user.DateCreated = time.Now().Format("2006-01-02 15:04:05")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	user.Role = constants.Role(string(constants.User))

	insertUserSQL := "INSERT INTO users (username, password, first_name, last_name, email, phone, date_created, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(insertUserSQL, user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.Phone, user.DateCreated, user.Role)
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	token, err := middleware.GenerateToken(int(lastInsertedID))
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	registeredUser := models.UserDTO{
		ID:       int(lastInsertedID),
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return &registeredUser, nil
}

func LoginUser(user *models.UserDao) (*models.UserDTO, error) {
	DB := db.GetDatabaseConnection()

	exists, err := isUserExists(user.Username)
	if err != nil || !exists {
		return nil, err
	}

	var hashedPassword string
	userSQLStatement := "SELECT id, password, email FROM users WHERE username = ?"
	err = DB.QueryRow(userSQLStatement, user.Username).Scan(&user.ID, &hashedPassword, &user.Email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)) != nil {
		return nil, nil
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	loggedUser := models.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return &loggedUser, nil
}
