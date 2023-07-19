package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/entities"
)

type UserService struct {
	repo *entities.UserRepository
}

func NewUserService(repo *entities.UserRepository) *UserService {
	return &UserService{repo}
}

func (u *UserService) FindAll(ctx context.Context) (users []entities.User, err error) {
	users, err = u.repo.FindAllUsers(ctx)
	return
}

func (u *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	return u.repo.SaveUser(ctx, user)
}
