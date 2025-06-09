package dto

type CreateUserRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Email       string `json:"email" binding:"omitempty,email"`
}

type UserResponse struct {
	PublicID string `json:"public_id"`
	FullName string `json:"full_name"`
	Message  string `json:"message"`
	Status   string `json:"status"`
}
