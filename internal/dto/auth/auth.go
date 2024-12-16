package authdto

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
