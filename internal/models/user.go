package models

import "time"

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Username  string    `json:"username" bson:"username"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(req UserRequest) User {
	return User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
}
