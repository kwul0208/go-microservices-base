package request

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
	Role     string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
}
