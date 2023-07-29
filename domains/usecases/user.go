package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/Nexters/pinterest/domains/errors"
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

func (u *UserService) CreateUser(ctx context.Context, userCreationRequest dto.UserCreationRequest) (userResponse dto.UserCreationResponse, err error) {
	user := entities.User{
		Name:     userCreationRequest.Name,
		UserID:   userCreationRequest.UserID,
		Password: userCreationRequest.Password,
		PageUrl:  userCreationRequest.PageUrl,
	}

	savedUser, err := u.repo.SaveUser(ctx, user)
	if err != nil {
		return
	}

	userResponse = dto.UserCreationResponse{
		Name:       savedUser.Name,
		PageUrl:    savedUser.PageUrl,
		Email:      savedUser.Email,
		Visitors:   savedUser.Visitors,
		ThemeColor: savedUser.ThemeColor,
		Text:       savedUser.Text,
		CreatedAt:  savedUser.CreatedAt,
	}
	return
}

func (u *UserService) LoginUser(ctx context.Context, dto dto.UserLoginRequest) error {
	// find user with id and password
	user, err := u.repo.FindUser(ctx, dto.UserID)
	if err != nil {
		return err
	}

	if user.Password != dto.Password {
		return errors.NewUnauthorizedError()
	}
	return nil
}
