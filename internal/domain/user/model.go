package user

import (
	"github.com/google/uuid"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"time"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

type User struct {
	ID          int64     `db:"id" json:"id"`
	ReferenceID uuid.UUID `db:"reference_id"`
	FullName    string    `db:"full_name"`
	PhoneNumber string    `db:"phone_number"`
	Email       string    `db:"email" `
	Status      string    `db:"status" `
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (u User) statusAsText() string {
	switch u.Status {
	case StatusActive:
		return "active"
	default:
		return "suspended"
	}
}

func (u User) ToDto() *dto.UserResponse {
	return &dto.UserResponse{
		Message:  "Created user",
		Status:   u.statusAsText(),
		FullName: u.FullName,
		PublicID: u.ReferenceID.String(),
	}
}

func NewUser(name string, phone string, email string) User {
	now := time.Now()

	return User{
		ReferenceID: uuid.Must(uuid.NewRandom()),
		Status:      StatusActive,
		FullName:    name,
		Email:       email,
		PhoneNumber: phone,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
