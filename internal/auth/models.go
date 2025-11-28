package auth

type SignupRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	UserID   string `bson:"user_id"`
}
