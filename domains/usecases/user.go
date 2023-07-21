package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
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

func (u *UserService) FindByUserId(ctx context.Context, userId int) (userResponse dto.UserDetailResponse, err error) {
	user, err := u.repo.FindUser(ctx, userId)
	if err != nil {
		return
	}

	var groups []dto.Group = dto.ToGroupDtoList(user.Group)
	var visitLogs []dto.VisitLog = dto.ToVisitLogDtoList(user.VisitLog)

	userResponse = dto.UserDetailResponse{
		Name:       user.Name,
		PageUrl:    user.PageUrl,
		Group:      groups,
		VisitLog:   visitLogs,
		ThemeColor: user.ThemeColor,
		Text:       user.Text,
	}
	return userResponse, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	return u.repo.SaveUser(ctx, user)
}
