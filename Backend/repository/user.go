package repository 

import (
	"context"
	"net/http"

)

type UserRepository interface {
 createUser(ctx context.Context,user *types.UserModel) (*types.UserModel,error)
}