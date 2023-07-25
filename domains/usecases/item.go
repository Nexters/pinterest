package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type ItemService struct {
	repo *entities.ItemRepository
}

func NewItemService(repo *entities.ItemRepository) *ItemService {
	return &ItemService{repo}
}

func (i *ItemService) FindByItemId(ctx context.Context, itemId uint) (itemResponse dto.ItemDetailResponse, err error) {
	item, err := i.repo.FindItem(ctx, itemId)
	if err != nil {
		return
	}

	itemResponse = dto.ItemDetailResponse{
		Title:     item.Title,
		Text:      item.Text,
		Link:      item.Link,
		Image:     item.Image,
		Likes:     item.Likes,
		GroupID:   item.GroupID,
		CreatedAt: item.CreatedAt,
	}
	return
}

func (i *ItemService) CreateItem(
	ctx context.Context,
	itemCreationRequest dto.ItemCreationRequest,
) (itemResponse dto.ItemDetailResponse, err error) {
	item := entities.Item{
		Title:   itemCreationRequest.Title,
		Text:    itemCreationRequest.Text,
		Link:    itemCreationRequest.Link,
		Image:   itemCreationRequest.Image,
		GroupID: itemCreationRequest.GroupID,
	}

	savedItem, err := i.repo.SaveItem(ctx, item)
	if err != nil {
		return
	}

	itemResponse = dto.ItemDetailResponse{
		Title:     savedItem.Title,
		Text:      savedItem.Text,
		Link:      savedItem.Link,
		Image:     savedItem.Image,
		Likes:     savedItem.Likes,
		GroupID:   savedItem.GroupID,
		CreatedAt: savedItem.CreatedAt,
	}
	return
}
