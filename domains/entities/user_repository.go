package entities

import (
	"context"
	"errors"

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

// tx: transaction 정보
func (ur *UserRepository) FindUser(ctx context.Context, userId int) (user User, err error) {
	tx := ur.DB.First(&user, userId)
	if tx.RowsAffected == 0 {
		err = errors.New("값이 없습니다")
		return
	}
	if tx.Error != nil {
		err = tx.Error
		return
	}

	return user, nil
}

func (ur *UserRepository) SaveUser(ctx context.Context, user User) error {
	tx := ur.DB.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
