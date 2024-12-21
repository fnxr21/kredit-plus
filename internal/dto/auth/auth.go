package authdto


type LoginRequest struct {
	Username string `form:"username" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
