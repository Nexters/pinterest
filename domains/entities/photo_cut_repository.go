package entities

import (
	"context"
	"errors"

	customerrors "github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type PhotoCutRepository struct {
	*gorm.DB
}

func NewPhotoCutRepository(db *gorm.DB) *PhotoCutRepository {
	return &PhotoCutRepository{db}
}

func (pcr *PhotoCutRepository) FindPhotoCut(ctx context.Context, photoCutId uint) (photoCut PhotoCut, err error) {
	tx := pcr.First(&photoCut, photoCutId)
	if tx.RowsAffected == 0 {
		err = customerrors.NewNotFoundError("PhotoCut")
		return
	}

	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func (pcr *PhotoCutRepository) SavePhotoCut(ctx context.Context, photoCut PhotoCut) (PhotoCut, error) {
	tx := pcr.DB.Begin()
	if err := tx.Error; err != nil {
		return PhotoCut{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// film 조회
	var film Film
	err := tx.First(&film, "id = ?", photoCut.FilmID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return photoCut, customerrors.NewNotFoundError("Film")
	}

	// photo cut 저장
	err = tx.Create(&photoCut).Error
	if err != nil {
		tx.Rollback()
		return photoCut, customerrors.NewCreateFailedError("PhotoCut")
	}

	// film의 photo_cut_count 증가
	film.PhotoCutCount++
	err = tx.Save(film).Error
	if err != nil {
		tx.Rollback() // 에러 시 트랜잭션 롤백
		return photoCut, err
	}

	// 트랜잭션 커밋
	err = tx.Commit().Error
	if err != nil {
		return photoCut, err
	}

	return photoCut, nil
}
