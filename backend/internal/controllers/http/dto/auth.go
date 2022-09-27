package dto

type Registration struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshInput struct {
	Token string `json:"token" binding:"required"`
}
