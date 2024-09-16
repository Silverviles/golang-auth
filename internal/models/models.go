package models

import "go-app/internal/constants"

type UserDao struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	DateCreated string         `json:"date_created"`
	Role        constants.Role `json:"role"`
}

type UserDTO struct {
	ID       int            `json:"id"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Role     constants.Role `json:"role"`
	Token    string         `json:"token"`
}

type Error struct {
	ResponseCode int    `json:"response_code"`
	MessageCode  string `json:"message_code"`
	Message      string `json:"message"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
