package dto

type UserLoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
