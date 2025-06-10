package wallet

import (
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"time"
)

type Wallet struct {
	ID        int64     `db:"id" json:"id"`
	PublicID  string    `db:"public_id" json:"public_id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Balance   float64   `db:"balance" json:"balance"`
	Currency  string    `db:"currency" json:"currency"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (w Wallet) ToBalanceResponse() dto.WalletBalance {
	return dto.WalletBalance{
		PublicID: w.PublicID,
		Balance:  w.Balance,
		Currency: w.Currency,
		Status:   w.Status,
	}
}
