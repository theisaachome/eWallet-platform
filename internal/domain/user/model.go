package user

import (
	"github.com/google/uuid"
	"github.com/theisaachome/eWallet-platform/internal/domain/wallet"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"time"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

type User struct {
	ID           int64     `db:"id" json:"id"`
	ReferenceID  uuid.UUID `db:"reference_id"`
	Username     string    `db:"username"`
	PhoneNumber  string    `db:"phone_number"`
	HashPassWord string    `db:"hash_password"`
	Status       string    `db:"status" `
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (u User) statusAsText() string {
	switch u.Status {
	case StatusActive:
		return "active"
	default:
		return "suspended"
	}
}

func (u User) ToDto(wallet *wallet.Wallet) *dto.UserResponse {
	return &dto.UserResponse{
		Message:  "Created user",
		Status:   u.statusAsText(),
		PublicID: u.ReferenceID.String(),
		Wallet: &dto.WalletResponse{
			PublicID: wallet.PublicID,
			Balance:  wallet.Balance,
			Currency: wallet.Currency,
			Status:   wallet.Status,
		},
	}
}

func NewUser(phone string, password string) User {
	return User{
		ReferenceID:  uuid.Must(uuid.NewRandom()),
		Status:       StatusActive,
		HashPassWord: password,
		PhoneNumber:  phone,
	}
}
