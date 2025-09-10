package auth

import (
	"Backend/dtos"
	// "crypto/ecdh"
    "net/http"
	"github.com/gin-gonic/gin"
)

// Handler defines auth endpoints
type Handler interface {
	Login(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	GoogleAuth(ctx *gin.Context)
	CheckToken(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

func SignUp(ctx *gin.Context) {
	 newUser := dtos.UserSignupParams {}// for json body dtos that we will receive
	//reading from the body
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": "Invalid request",
			"details": err.Error(), 


		})
		return

	}
	user :={
		Name: newUser.Name,
		surname : newUser.Surname,
		email: newUser.Email,
		password: newUser.Password,
		role: newUser.Role,
	}
	createdUser ,err := CreateUser(user)

	h.JSON(ctx,http.statusOk,"Registration successful")

}
