package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"github.com/theisaachome/eWallet-platform/pkg/utils/logger"
)

// Repository --- INTERFACE ---
type Repository interface {
	RegisterUser(User) (*User, *errors.AppError)
}

// RepositoryDb --- STRUCT IMPLEMENTATION ---
type RepositoryDb struct {
	client *sqlx.DB
}

// NewRepositoryDb --- CONSTRUCTOR ---
func NewRepositoryDb(db *sqlx.DB) Repository {
	return &RepositoryDb{client: db}
}

func (db *RepositoryDb) RegisterUser(u User) (*User, *errors.AppError) {
	insertQuery := `
	INSERT INTO users(reference_id, full_name, phone_number, email, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.client.Exec(insertQuery,
		u.ReferenceID,
		u.FullName,
		u.PhoneNumber,
		u.Email,
		u.Status,
		u.CreatedAt,
		u.UpdatedAt,
	)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}

	return &u, nil
}
