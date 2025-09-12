package services

import(
    "Backend/types"
    "Backend/repository"
    "context"
)

type UserService interface {
    CreateUser(ctx context.Context, user *types.UserModel) (*types.UserModel, error)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{
        userRepo: userRepo,
    }
}

func (s *userService) CreateUser(ctx context.Context, user *types.UserModel) (*types.UserModel, error) {
    return s.userRepo.CreateUser(ctx, user) 
}