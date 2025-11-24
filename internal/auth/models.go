package auth

type SignupRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
