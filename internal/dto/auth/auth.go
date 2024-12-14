package authdto

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
type LoginResponse struct {
	Username string `json:"username" `
	Name     string `json:"name" `
}
