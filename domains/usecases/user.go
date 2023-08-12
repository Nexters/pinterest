package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/Nexters/pinterest/domains/errors"
)

const NickName = "grafi"

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
		Name:     NickName,
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
		Visitors:   savedUser.Visitors,
		ThemeColor: savedUser.ThemeColor,
		Text:       savedUser.Text,
		CreatedAt:  savedUser.CreatedAt,
	}
	return
}

func (u *UserService) FindUserByID(ctx context.Context, userID string) (userDetail dto.UserDetailResponse, err error) {
	user, err := u.repo.FindUser(ctx, userID)

	userDetail.Name = user.Name
	userDetail.Profile = user.Profile
	userDetail.Text = user.Text
	userDetail.UserID = user.ID
	userDetail.Visitors = user.Visitors
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

func (u *UserService) UpdateUser(ctx context.Context, updateDto dto.UserUpdateRequest) (userDetail dto.UserDetailResponse, err error) {
	userUpdateParam := updateDto.ToEntity()

	user, err := u.repo.UpdateUser(ctx, userUpdateParam)
	if err != nil {
		return
	}

	userDetail = userDetail.FromEntity(user)

	return
}
