package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type CreateUserRequest struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	NewUserID uuid.UUID `json:"user_id"`
}
