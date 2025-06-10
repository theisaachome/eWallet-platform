package handlers

import (
	"encoding/json"
	"github.com/theisaachome/eWallet-platform/internal/app/user"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/dto"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/response"
	"github.com/theisaachome/eWallet-platform/pkg/errors"
	"net/http"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return UserHandler{service: service}
}

// localhost:8080/ewallet/api/users post method
func (h UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errors.ErrValidation, err.Error())
	} else {
		newUser, err := h.service.CreateNewUser(request)
		if err != nil {
			response.WriteError(w, err.Code, errors.ErrInternalServer, err.Message)
		} else {
			response.WriteJSON(w, http.StatusCreated, newUser)
		}
	}
}
