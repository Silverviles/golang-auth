package services

import (
	"go-app/internal/constants"
	"go-app/internal/customMiddleware"
	"go-app/internal/db"
	"go-app/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func isUserExists(username string) (bool, error) {
	DB := db.GetDatabaseConnection()

	var count int
	err := DB.QueryRow(db.CountUserSQL, username).Scan(&count)
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
	user.Role = constants.User

	result, err := DB.Exec(db.InsertUserSQL, user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.Phone, user.DateCreated, user.Role)
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	token, err := customMiddleware.GenerateToken(int(lastInsertedID), string(user.Role))
	if err != nil {
		err := transaction.Rollback()
		return nil, err
	}

	registeredUser := models.UserDTO{
		ID:       int(lastInsertedID),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
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
	err = DB.QueryRow(db.SelectUserSQL, user.Username).Scan(&user.ID, &hashedPassword, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)) != nil {
		return nil, nil
	}

	token, err := customMiddleware.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	loggedUser := models.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	return &loggedUser, nil
}
