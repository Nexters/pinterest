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
	// validate 로직 추가
	return
}
