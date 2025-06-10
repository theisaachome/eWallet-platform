package handlers

import (
	"github.com/theisaachome/eWallet-platform/internal/app/wallet"
	"github.com/theisaachome/eWallet-platform/internal/interfaces/http/response"
	"net/http"
	"strconv"
)

//
//| Endpoint               | Who Uses It | Purpose                                                             |
//| ---------------------- | ----------- | ------------------------------------------------------------------- |
//| `GET /wallets/{id}`    | Admin       | View wallet by internal DB ID or public ID (e.g., customer support) |
//| `PATCH /wallets/{id}`  | Admin       | Freeze or unfreeze wallet for compliance/fraud                      |
//| `GET /wallets/balance` | Wallet User | **User-friendly balance display** for current authenticated user    |

type WalletHandler struct {
	service wallet.Service
}

// GetUserWallet  /api/wallets/{id} (admin operation)
//func (h WalletHandler) GetUserWallet(w http.ResponseWriter, r *http.Request) error {}

// GetWalletBalance /api/wallets/balance ( Wallet user -> home-page balance show)
func (h WalletHandler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	walletBalance, err := h.service.GetWalletBalance(6)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, strconv.Itoa(err.Code), err.Message)
	} else {
		response.WriteJSON(w, http.StatusOK, walletBalance)
	}

}

func NewWalletHandler(service wallet.Service) *WalletHandler {
	return &WalletHandler{service: service}
}
