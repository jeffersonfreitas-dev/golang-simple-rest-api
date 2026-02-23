package users

import (
	"errors"
	"log/slog"
	"simple-rest-api/internal/models"
	"simple-rest-api/internal/repositories"

	"github.com/google/uuid"
)

type UsersCase struct {
	repo *repositories.Repositories
}

func New(repo *repositories.Repositories) *UsersCase {
	return &UsersCase{
		repo: repo,
	}
}

func (u UsersCase) GetAll() []models.User {
	users := u.repo.User.GetAll()
	return users
}

func (u UsersCase) Add(newUser models.CreateUserRequest) (uuid.UUID, error) {
	emailExists := u.repo.User.EmailExists(newUser.Email)

	if emailExists {
		slog.Error("this email", newUser.Email, "already exists")
		return uuid.Nil, errors.New("email already exists")
	}

	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repo.User.Add(repoReq)
	return repoReq.ID, nil
}
