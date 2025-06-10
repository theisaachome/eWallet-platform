package wallet

import (
	"github.com/jmoiron/sqlx"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"github.com/theisaachome/eWallet-platform/pkg/utils/logger"
)

type Repository interface {
	CreateUserWallet(userId int64) (*Wallet, *errors.AppError)
}
type RepositoryDb struct {
	db *sqlx.DB
}

func (r RepositoryDb) CreateUserWallet(userId int64) (*Wallet, *errors.AppError) {
	//TODO implement me
	var w Wallet
	query := `
			INSERT INTO wallets (user_id) VALUES ($1)	
			RETURNING id, public_id, user_id, balance, currency, status, created_at, updated_at
		`
	err := r.db.Get(&w, query, userId)
	if err != nil {
		logger.Error("Error while creating user wallet " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected error from database")

	}
	return &w, nil
}

func NewRepositoryDb(db *sqlx.DB) Repository {
	return &RepositoryDb{db: db}
}
