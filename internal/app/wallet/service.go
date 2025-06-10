package wallet

import (
	"github.com/theisaachome/eWallet-platform/internal/domain/wallet"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
)

type Service interface {
	CreateUserWallet(userId int64) (*wallet.Wallet, *errors.AppError)
}
type DefaultService struct {
	repo wallet.Repository
}

func NewService(repo wallet.Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

func (s *DefaultService) CreateUserWallet(userId int64) (*wallet.Wallet, *errors.AppError) {
	newWallet, err := s.repo.CreateUserWallet(userId)
	if err != nil {
		return nil, err
	} else {
		return newWallet, nil
	}
}
