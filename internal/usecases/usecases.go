package usecases

import (
	"errors"
	"log/slog"
	"simple-rest-api/internal/models"
	"simple-rest-api/internal/repositories"

	"github.com/google/uuid"
)

type UseCases struct {
	repo *repositories.Repositories
}

func New(repo *repositories.Repositories) *UseCases {
	return &UseCases{
		repo: repo,
	}
}

func (u UseCases) GetAll() []models.User {
	users := u.repo.User.GetAll()
	return users
}

func (u UseCases) Add(newUser models.User) (uuid.UUID, error) {
	emailExists := u.repo.User.EmailExists(newUser.Email)

	if emailExists {
		slog.Error("this email", newUser.Email, "already exists")
		return uuid.Nil, errors.New("email already exists")
	}

	repoReq := models.User{
		ID:   uuid.New(),
		Name: newUser.Name,
	}

	u.repo.User.Add(repoReq)
	return repoReq.ID, nil
}
