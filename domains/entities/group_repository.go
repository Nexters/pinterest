package entities

import (
	"context"
	"errors"

	customerrors "github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type GroupRepository struct {
	*gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db}
}

func (gr *GroupRepository) FindGroup(ctx context.Context, groupId uint) (group Group, err error) {
	err = gr.DB.Preload("Items").First(&group, groupId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customerrors.NewNotFoundError("Group")
			return
		}
		return
	}
	return
}

func (gr *GroupRepository) CountOrderByUserId(ctx context.Context, userId string) (count int64, err error) {
	tx := gr.DB.Model(&Group{}).Where("user_id = ?", userId).Count(&count)
	if tx.Error != nil {
		return
	}
	return
}

func (gr *GroupRepository) SaveGroup(ctx context.Context, group Group) (Group, error) {
	tx := gr.DB.Create(&group)
	if tx.Error != nil {
		return group, customerrors.NewCreateFailedError("Group")
	}
	return group, nil
}

func (gr *GroupRepository) FindAllFilmsInOrder(ctx context.Context, userId string) (films []Group, err error) {
	tx := gr.DB.Where("user_id = ?", userId).Order("order DESC").Find(&films)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
