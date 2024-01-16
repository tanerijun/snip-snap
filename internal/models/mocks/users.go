package mocks

import (
	"time"

	"github.com/tanerijun/snip-snap/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "bob@example.com" && password == "bobbobbob" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	if id == 1 {
		usr := &models.User{
			ID:      1,
			Name:    "Bob",
			Email:   "bob@example.com",
			Created: time.Now(),
		}

		return usr, nil
	}

	return nil, models.ErrNoRecord
}
