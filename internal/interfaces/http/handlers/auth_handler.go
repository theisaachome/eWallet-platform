package handlers

import (
	"encoding/json"
	"github.com/theisaachome/eWallet-platform/internal/app/auth"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/response"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"net/http"
)

// log-in

type AuthHandler struct {
	authService auth.Service
}

// localhost:8080/api/v1/auth/register post method
func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errors.ErrValidation, err.Error())
	} else {
		newUsr, err := h.authService.Register(request)
		if err != nil {
			response.WriteError(w, err.Code, errors.ErrInternalServer, err.Message)
		} else {
			response.WriteJSON(w, http.StatusCreated, newUsr)
		}
	}
}

// localhost:8080/wallet/api/v1/auth/login post method
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errors.ErrValidation, err.Error())
	} else {
		loginResponse, err := h.authService.Login(request)
		if err != nil {
			response.WriteError(w, err.Code, errors.ErrUnauthorized, err.Message)
		} else {
			response.WriteJSON(w, http.StatusOK, loginResponse)
		}
	}
}

func NewAuthHandler(authService auth.Service) *AuthHandler {
	return &AuthHandler{authService}
}
