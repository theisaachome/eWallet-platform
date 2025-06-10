package dto

type RegisterRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Password    string `json:"password" binding:"omitempty,min=8,max=16"`
}

type LoginRequest struct {
	Phone    string `json:"phoneOrEmail" validate:"required,emailOrEmail"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

type UserResponse struct {
	Message  string          `json:"message"`
	Status   string          `json:"status"`
	FullName string          `json:"full_name"`
	PublicID string          `json:"public_id"`
	Wallet   *WalletResponse `json:"wallet"`
}

type WalletResponse struct {
	PublicID string  `json:"public_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Status   string  `json:"status"`
}

type WalletBalance struct {
	PublicID string  `json:"public_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Status   string  `json:"status"`
}
