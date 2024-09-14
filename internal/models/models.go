package models

type UserDao struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateCreated string `json:"date_created"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type Error struct {
	ResponseCode int    `json:"response_code"`
	MessageCode  string `json:"message_code"`
	Message      string `json:"message"`
}
