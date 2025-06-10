package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"github.com/theisaachome/eWallet-platform/pkg/utils/logger"
)

// Repository --- INTERFACE ---
type Repository interface {
	SaveNewUser(User) (*User, *errors.AppError)
	FindUserByPhone(phone string) (*User, *errors.AppError)
}

// RepositoryDb --- STRUCT IMPLEMENTATION ---
type RepositoryDb struct {
	client *sqlx.DB
}

func (db *RepositoryDb) FindUserByPhone(phone string) (*User, *errors.AppError) {
	query := `SELECT * FROM users WHERE phone_number = $1`
	var user User
	err := db.client.Get(&user, query, phone)
	if err != nil {
		errors.NewNotFoundException("user not found" + err.Error())
	}
	return &user, nil
}

// NewRepositoryDb --- CONSTRUCTOR ---
func NewRepositoryDb(db *sqlx.DB) Repository {
	return &RepositoryDb{client: db}
}

func (db *RepositoryDb) SaveNewUser(u User) (*User, *errors.AppError) {
	insertQuery := `
		INSERT INTO users(phone_number,hash_password)
		VALUES ($1, $2)
		RETURNING id;
	`

	var id int64
	err := db.client.QueryRow(insertQuery,
		u.PhoneNumber,
		u.HashPassWord,
	).Scan(&id)

	if err != nil {
		logger.Error("Error while  creating new wallet-user: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")
	}
	u.ID = id
	return &u, nil
}
