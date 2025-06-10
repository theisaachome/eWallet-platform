package handlers

import (
	"github.com/theisaachome/eWallet-platform/internal/app/wallet"
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
//func (h WalletHandler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
//
//}
