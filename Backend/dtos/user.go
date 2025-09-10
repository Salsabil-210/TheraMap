package dtos

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignupParams struct {
    Name     string `json:"name" binding:"required,alpha,max=85"`
    Surname  string `json:"surname" binding:"required,alpha,max=100"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
    Role     string `json:"role,omitempty" binding:"omitempty,oneof=admin therapist patient user"`
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
