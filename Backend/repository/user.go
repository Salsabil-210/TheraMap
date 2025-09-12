package repository

import (
	"context"
	"Backend/types"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *types.UserModel) (*types.UserModel, error)
}