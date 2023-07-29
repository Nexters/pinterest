package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type GroupService struct {
	repo *entities.GroupRepository
}

func NewGroupService(repo *entities.GroupRepository) *GroupService {
	return &GroupService{repo}
}
func (g *GroupService) FindByGroupId(ctx context.Context, groupId uint) (groupDetailResponse dto.GroupDetailResponse, err error) {
	group, err := g.repo.FindGroup(ctx, groupId)
	if err != nil {
		return
	}

	items, err := dto.ToItemDtoList(group.Items)
	if err != nil {
		return
	}

	groupDetailResponse = dto.GroupDetailResponse{
		GroupID:   group.ID,
		Title:     group.Title,
		Order:     group.Order,
		ItemCount: group.ItemCount,
		Likes:     group.Likes,
		UserID:    group.UserID,
		Items:     items,
	}
	return
}

func (g *GroupService) CreateGroup(ctx context.Context, groupCreationRequest dto.GroupCreationRequest) (groupDetailResponse dto.GroupDetailResponse, err error) {
	order, err := g.repo.CountOrderByUserId(ctx, groupCreationRequest.UserID)
	if err != nil {
		return
	}

	group := entities.Group{
		Title:  groupCreationRequest.Title,
		UserID: groupCreationRequest.UserID,
		Order:  uint(order),
	}

	savedGroup, err := g.repo.SaveGroup(ctx, group)
	if err != nil {
		return
	}

	groupDetailResponse = dto.GroupDetailResponse{
		GroupID:   savedGroup.ID,
		Title:     savedGroup.Title,
		Order:     savedGroup.Order,
		ItemCount: savedGroup.ItemCount,
		Likes:     savedGroup.Likes,
		UserID:    savedGroup.UserID,
	}
	return
}
