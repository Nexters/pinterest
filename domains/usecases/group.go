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
		Type:      group.Type,
		Title:     group.Title,
		Text:      group.Text,
		Image:     group.Image,
		Order:     group.Order,
		ItemCount: group.ItemCount,
		Likes:     group.Likes,
		Link:      group.Link,
		UserID:    group.UserID,
		Items:     items,
	}
	return
}

func (g *GroupService) CreateGroup(ctx context.Context, groupCreationRequest dto.GroupCreationRequest) (groupDetailResponse dto.GroupDetailResponse, err error) {
	group := entities.Group{
		Title: groupCreationRequest.Title,
	}

	savedGroup, err := g.repo.SaveGroup(ctx, group)
	if err != nil {
		return
	}

	groupDetailResponse = dto.GroupDetailResponse{
		GroupID:   group.ID,
		Type:      savedGroup.Type,
		Title:     savedGroup.Title,
		Text:      savedGroup.Text,
		Image:     savedGroup.Image,
		Order:     savedGroup.Order,
		ItemCount: savedGroup.ItemCount,
		Likes:     savedGroup.Likes,
		Link:      savedGroup.Link,
		UserID:    savedGroup.UserID,
	}
	return
}
