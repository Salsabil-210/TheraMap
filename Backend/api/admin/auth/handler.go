package auth

import (
	"Backend/dtos"
	"Backend/services"
	"Backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler struct handles authentication endpoints
type AuthHandler struct {
	userService services.UserService
}

// NewAuthHandler creates a new AuthHandler with dependency injection
func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// SignUp handles user registration
func (h *AuthHandler) SignUp(ctx *gin.Context) {
	newUser := dtos.UserSignupParams{} // for json body dtos that we will receive

	// Bind JSON request to DTO
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return

	}

	// Convert DTO to Model
	user := &types.UserModel{
		Name:     newUser.Name,
		Surname:  newUser.Surname,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     newUser.Role,
	}

	// Call service to create user
	createdUser, err := h.userService.CreateUser(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create user",
			"details": err.Error(),
		})
		return
	}

	// Prepare response data (without password)
	responseData := gin.H{
		"id":        createdUser.ID,
		"name":      createdUser.Name,
		"surname":   createdUser.Surname,
		"email":     createdUser.Email,
		"role":      createdUser.Role,
		"createdAt": createdUser.CreatedAt,
		"updatedAt": createdUser.UpdatedAt,
	}

	// Return success response
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
		"data":    responseData,
	})
}
