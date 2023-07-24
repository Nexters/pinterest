package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/go-playground/validator"
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

func (u *UserService) FindByUserId(ctx context.Context, userId int) (userResponse dto.UserDetailResponse, err error) {
	user, err := u.repo.FindUser(ctx, userId)
	if err != nil {
		return
	}

	groups, err := dto.ToGroupDtoList(user.Group)
	if err != nil {
		return
	}
	visitLogs, err := dto.ToVisitLogDtoList(user.VisitLog)
	if err != nil {
		return
	}

	userResponse = dto.UserDetailResponse{
		Name:       user.Name,
		PageUrl:    user.PageUrl,
		Group:      groups,
		VisitLog:   visitLogs,
		ThemeColor: user.ThemeColor,
		Text:       user.Text,
	}
	validate := validator.New()
	err = validate.Struct(userResponse)
	if err != nil {
		return
	}
	return
}

func (u *UserService) CreateUser(ctx context.Context, userCreationRequest dto.UserCreationRequest) (userResponse dto.UserCreationResponse, err error) {
	user := entities.User{
		Name:     userCreationRequest.Name,
		Password: userCreationRequest.Password,
		PageUrl:  userCreationRequest.PageUrl,
	}

	err = u.repo.SaveUser(ctx, user)
	if err != nil {
		return
	}

	userResponse = dto.UserCreationResponse{
		Name:    user.Name,
		PageUrl: user.PageUrl,
	}
	validate := validator.New()
	err = validate.Struct(userResponse)
	if err != nil {
		return
	}
	return
}
