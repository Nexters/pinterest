package entities

import (
	"context"
	"errors"

	customerrors "github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type FilmRepository struct {
	*gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepository {
	return &FilmRepository{db}
}

func (fr *FilmRepository) FindFilm(ctx context.Context, filmId uint) (film Film, err error) {
	err = fr.DB.Preload("PhotoCuts").First(&film, filmId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = customerrors.NewNotFoundError("Film")
		return
	}
	return
}

func (fr *FilmRepository) CountOrderByUserId(ctx context.Context, userId string) (count int64, err error) {
	tx := fr.DB.Model(&Film{}).Where("user_id = ?", userId).Unscoped().Count(&count)
	if tx.Error != nil {
		return
	}
	return
}

func (fr *FilmRepository) SaveFilm(ctx context.Context, film Film) (Film, error) {
	tx := fr.DB.Create(&film)
	if tx.Error != nil {
		return film, customerrors.NewCreateFailedError("Film")
	}
	return film, nil
}

func (fr *FilmRepository) FindAllFilmsInOrder(ctx context.Context, userId string) (films []Film, err error) {
	tx := fr.DB.Preload("PhotoCuts").Where("user_id = ?", userId).Order("`order` DESC").Find(&films)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
