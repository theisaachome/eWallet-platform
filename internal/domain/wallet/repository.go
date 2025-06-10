package wallet

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"github.com/theisaachome/eWallet-platform/pkg/utils/logger"
)

type Repository interface {
	CreateUserWallet(userId int64) (*Wallet, *errors.AppError)
	GetWalletBalance(userId int64) (*Wallet, *errors.AppError)
}
type RepositoryDb struct {
	db *sqlx.DB
}

func (r *RepositoryDb) GetWalletBalance(userId int64) (*Wallet, *errors.AppError) {
	const query = `SELECT * FROM wallets WHERE user_id = $1`
	var wallet Wallet
	err := r.db.Get(&wallet, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundException("User Wallet not found")
		}
		return nil, errors.NewUnexpectedError("Database error: " + err.Error())
	}
	return &wallet, nil
}

func (r *RepositoryDb) CreateUserWallet(userId int64) (*Wallet, *errors.AppError) {
	var w Wallet
	query := `INSERT INTO wallets (user_id) VALUES ($1)	
			RETURNING id, public_id, user_id, balance, currency, status, created_at, updated_at`
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
