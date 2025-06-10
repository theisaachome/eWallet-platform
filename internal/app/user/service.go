package user

import (
	"github.com/theisaachome/eWallet-platform/internal/app/wallet"
	"github.com/theisaachome/eWallet-platform/internal/domain/user"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
)

type Service interface {
	CreateNewUser(request dto.RegisterRequest) (*dto.UserResponse, *errors.AppError)
}

type DefaultService struct {
	repo          user.Repository
	walletService wallet.Service
}

func NewService(repo user.Repository, walletService wallet.Service) Service {
	return DefaultService{
		repo:          repo,
		walletService: walletService,
	}
}

func (s DefaultService) CreateNewUser(request dto.RegisterRequest) (*dto.UserResponse, *errors.AppError) {
	// todo: validate the request body here
	//newUser := user.NewUser(request.FullName, request.PhoneNumber, request.Email)
	//
	//savedUser, err := s.repo.RegisterUser(newUser)
	//if err != nil {
	//	return nil, err
	//}
	//// Create wallet after user creation
	//newWallet, err := s.walletService.CreateUserWallet(savedUser.ID)
	//if err != nil {
	//	return nil, err
	//}
	//return savedUser.ToDto(newWallet), nil

	return nil, errors.NewUnexpectedError("")

}
