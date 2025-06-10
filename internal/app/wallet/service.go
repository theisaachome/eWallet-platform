package wallet

import (
	"github.com/theisaachome/eWallet-platform/internal/domain/wallet"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
)

type Service interface {
	CreateUserWallet(userId int64) (*wallet.Wallet, *errors.AppError)
	GetWalletBalance(userId int64) (dto.WalletBalance, *errors.AppError)
	AddMoney(userId int64, amount float64)
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

func (s *DefaultService) GetWalletBalance(userId int64) (dto.WalletBalance, *errors.AppError) {
	wallet, err := s.repo.GetWalletBalance(userId)
	if err != nil {
		errors.NewUnexpectedError("error getting wallet balance")
	}
	return wallet.ToBalanceResponse(), nil
}

func (s *DefaultService) AddMoney(userId int64, amount float64) {}
