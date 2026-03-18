package models

import "time"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	UserType string `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	User        User   `json:"user"`
}