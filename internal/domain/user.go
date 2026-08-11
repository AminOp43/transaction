package domain

import "time"

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
