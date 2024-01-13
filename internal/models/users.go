package models

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES(@name, @email, @hashedPassword, NOW())`
	args := pgx.NamedArgs{
		"name":           name,
		"email":          email,
		"hashedPassword": string(hashedPassword),
	}

	_, err = m.DB.Exec(context.Background(), stmt, args)
	if err != nil {
		var postgresError *pgconn.PgError
		if errors.As(err, &postgresError) {
			if postgresError.Code == "23505" && strings.Contains(postgresError.Message, "users_email_key") {
				return ErrDuplicateEmail
			}
		}

		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
