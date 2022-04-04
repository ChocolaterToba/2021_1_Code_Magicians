package domain

// UserLoginInput is used when parsing JSON in auth/login handler
type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
