package entities

import (
	"context"

	"github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type PhotoCutRepository struct {
	*gorm.DB
}

func NewPhotoCutRepository(db *gorm.DB) *PhotoCutRepository {
	return &PhotoCutRepository{db}
}

func (ir *PhotoCutRepository) FindPhotoCut(ctx context.Context, photoCutId uint) (photoCut PhotoCut, err error) {
	tx := ir.DB.First(&photoCut, photoCutId)
	if tx.RowsAffected == 0 {
		err = errors.NewNotFoundError("PhotoCut")
		return
	}
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (ir *PhotoCutRepository) SavePhotoCut(ctx context.Context, photoCut PhotoCut) (PhotoCut, error) {
	tx := ir.DB.Create(&photoCut)
	if tx.Error != nil {
		return photoCut, errors.NewCreateFailedError("PhotoCut")
	}

	return photoCut, nil
}
