package entities

import (
	"context"

	"github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type ItemRepository struct {
	*gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (ir *ItemRepository) FindItem(ctx context.Context, itemId uint) (item Item, err error) {
	tx := ir.DB.First(&item, itemId)
	if tx.RowsAffected == 0 {
		err = errors.NewNotFoundError("Item")
		return
	}
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (ir *ItemRepository) SaveItem(ctx context.Context, item Item) (Item, error) {
	tx := ir.DB.Create(&item)
	if tx.Error != nil {
		return item, errors.NewCreateFailedError("Item")
	}

	return item, nil
}
