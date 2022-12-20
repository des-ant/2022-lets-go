package mocks

import (
	"time"

	"github.com/des-ant/2022-lets-go/code/snippetbox/internal/models"
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
	if email == "alice@example.com" && password == "pa$$word" {
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
	const layout = "2006-01-02 15:04:05"
	created, _ := time.Parse(layout, "2022-01-01 10:00:00")

	user := models.User{
		ID:             1,
		Name:           "Alice Jones",
		Email:          "alice@example.com",
		HashedPassword: []byte("$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG"),
		Created:        created,
	}

	switch id {
	case 1:
		return &user, nil
	default:
		return nil, models.ErrNoRecord
	}
}
