package dtos

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignupParams struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserUpdateParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
}

type CheckTokenResponse struct {
	ID       string `json:"id"`
	UserType string `json:"user_type"`
}

type ForgotPasswordParams struct {
	Email string `json:"email" example:"user@gmail.com"`
}

type ResetPasswordParams struct {
	Token    string `json:"token" example:"abcdtoken"`
	Password string `json:"password" example:"Password@123"`
}
