package auth

import (
	"github.com/theisaachome/eWallet-platform/internal/app/wallet"
	"github.com/theisaachome/eWallet-platform/internal/domain/user"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"github.com/theisaachome/eWallet-platform/pkg/security/jwt"
	"github.com/theisaachome/eWallet-platform/pkg/utils"
	"strconv"
)

type Service interface {
	Register(req dto.RegisterRequest) (dto.AuthResponse, *errors.AppError)
	Login(req dto.LoginRequest) (dto.AuthResponse, *errors.AppError)
}

type DefaultService struct {
	userRepo      user.Repository
	walletService wallet.Service
	jwt           jwt.Service
}

func (s DefaultService) Register(req dto.RegisterRequest) (dto.AuthResponse, *errors.AppError) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		errors.NewUnexpectedError("Unable to hash password")
	}
	newUser := user.NewUser(req.PhoneNumber, hashPassword)
	savedUser, _ := s.userRepo.SaveNewUser(newUser)
	//if err != nil {
	//	errors.NewUnexpectedError("Unable to register user")
	//}
	s.walletService.CreateUserWallet(savedUser.ID)
	// generate jwt token
	token, err := s.jwt.GenerateToken(strconv.FormatInt(newUser.ID, 10), "wallet-user")
	if err != nil {
		errors.NewUnexpectedError("Unable to generate token")
	}
	return dto.AuthResponse{
		Token:  token,
		Status: "Register success",
	}, nil

}

func (s DefaultService) Login(req dto.LoginRequest) (dto.AuthResponse, *errors.AppError) {
	user, err := s.userRepo.FindUserByPhone(req.Phone)

	if err != nil {
		errors.NewNotFoundException("Invalid Credentials")
	}
	isMatch := utils.CheckPasswordHash(req.Password, user.HashPassWord)
	if !isMatch {
		errors.NewNotFoundException("Invalid Credentials")
	}
	jwtToken, _ := s.jwt.GenerateToken(strconv.FormatInt(user.ID, 10), "wallet-user")

	return dto.AuthResponse{
		Token:  jwtToken,
		Status: "Login success",
	}, nil
}

func NewAuthService(r user.Repository, walletService wallet.Service, jwtService jwt.Service) Service {
	return &DefaultService{userRepo: r, walletService: walletService, jwt: jwtService}
}
