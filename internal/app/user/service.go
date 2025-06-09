package user

import (
	"github.com/theisaachome/eWallet-platform/internal/domain/user"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
)

type Service interface {
	CreateNewUser(request dto.CreateUserRequest) (*dto.UserResponse, *errors.AppError)
}

type DefaultService struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return DefaultService{repo: repo}
}

func (s DefaultService) CreateNewUser(request dto.CreateUserRequest) (*dto.UserResponse, *errors.AppError) {
	// todo: validate the request body here
	user := user.NewUser(request.FullName, request.PhoneNumber, request.Email)
	if newUser, err := s.repo.RegisterUser(user); err != nil {
		return nil, err
	} else {
		return newUser.ToDto(), nil
	}
}
