package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

