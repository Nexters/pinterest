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
		ID:       userCreationRequest.UserID,
		Password: userCreationRequest.Password,
	}

	savedUser, err := u.repo.SaveUser(ctx, user)
	if err != nil {
		return
	}

	userResponse = dto.UserCreationResponse{
		Name:       savedUser.Name,
		UserID:     savedUser.ID,
		Email:      savedUser.Email,
		Visitors:   savedUser.Visitors,
		ThemeColor: savedUser.ThemeColor,
		Text:       savedUser.Text,
		CreatedAt:  savedUser.CreatedAt,
	}
	return
}

func (u *UserService) FindUserByID(ctx context.Context, userID string) (userDetail dto.UserDetailResponse, err error) {
	return
}

func (u *UserService) LoginUser(ctx context.Context, loginDto dto.UserLoginRequest) (userDetail dto.UserDetailResponse, err error) {
	// find user with id and password
	user, err := u.repo.FindUser(ctx, loginDto.UserID)
	if err != nil {
		return
	}

	if user.Password != loginDto.Password {
		err = errors.NewUnauthorizedError()
		return
	}

	userDetail = dto.UserDetailResponse{
		Name:     user.Name,
		Text:     user.Text,
		Profile:  user.Profile,
		Visitors: user.Visitors,
		UserID:   user.ID,
	}
	return
}
