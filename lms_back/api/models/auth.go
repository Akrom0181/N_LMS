package models

type StudentLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type StudentLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type StudentRegisterRequest struct {
	Mail string `json:"mail"`
}