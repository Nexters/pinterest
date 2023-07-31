package entities

import (
	"context"

	"github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) FindAllUsers(ctx context.Context) (users []User, err error) {
	tx := ur.DB.Find(&users)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}

func (ur *UserRepository) FindUser(ctx context.Context, userID string) (user User, err error) {
	tx := ur.DB.Where("id = ?", userID).First(&user)
	if tx.RowsAffected == 0 {
		err = errors.NewNotFoundError("User")
		return
	}
	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}

func (ur *UserRepository) SaveUser(ctx context.Context, user User) (createdUser User, err error) {
	// if userId is not null, return error
	// tx := ur.DB.Where("id = ?", user.ID).First(&user)
	// if tx.RowsAffected != 0 {
	// 	err = errors.NewCreateFailedError("User")
	// 	return
	// }

	tx := ur.DB.Create(&user)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	createdUser = user

	return
}
