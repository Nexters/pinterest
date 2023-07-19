package entities

import (
	"context"

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

func (ur *UserRepository) SaveUser(ctx context.Context, user *User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
