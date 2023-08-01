package usecases

import (
	"context"
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type FilmService struct {
	repo *entities.FilmRepository
}

func NewFilmService(repo *entities.FilmRepository) *FilmService {
	return &FilmService{repo}
}

func (f *FilmService) FindByFilmId(ctx context.Context, filmId uint) (filmDetailResponse dto.FilmDetailResponse, err error) {
	film, err := f.repo.FindFilm(ctx, filmId)
	if err != nil {
		return
	}

	photoCuts, err := dto.ToPhotoCutDtoList(film.PhotoCuts)
	if err != nil {
		return
	}

	filmDetailResponse = dto.FilmDetailResponse{
		FilmID:        film.ID,
		Title:         film.Title,
		Order:         film.Order,
		PhotoCutCount: uint(len(film.PhotoCuts)),
		Likes:         film.Likes,
		UserID:        film.UserID,
		PhotoCuts:     photoCuts,
	}
	return
}

func (f *FilmService) CreateFilm(ctx context.Context, filmCreationRequest dto.FilmCreationRequest) (filmDetailResponse dto.FilmDetailResponse, err error) {
	order, err := f.repo.CountOrderByUserId(ctx, filmCreationRequest.UserID)
	if err != nil {
		return
	}

	film := entities.Film{
		Title:  filmCreationRequest.Title,
		UserID: filmCreationRequest.UserID,
		Order:  uint(order),
	}

	savedFilm, err := f.repo.SaveFilm(ctx, film)
	if err != nil {
		return
	}

	filmDetailResponse = dto.FilmDetailResponse{
		FilmID:        savedFilm.ID,
		Title:         savedFilm.Title,
		Order:         savedFilm.Order,
		PhotoCutCount: savedFilm.PhotoCutCount,
		Likes:         savedFilm.Likes,
		UserID:        savedFilm.UserID,
	}
	return
}

func (f *FilmService) FindAllFilms(ctx context.Context, userId string) (filmList []dto.Film, err error) {
	films, err := f.repo.FindAllFilmsInOrder(ctx, userId)
	if err != nil {
		return
	}

	filmList, err = dto.ToFilmDtoList(films)
	if err != nil {
		return
	}
	return
}

func (f *FilmService) UpdateFilm(ctx context.Context, filmUpdateRequest dto.FilmUpdateRequest) (err error) {
	filmId, err := strconv.Atoi(filmUpdateRequest.FilmID)

	film, err := f.repo.FindFilm(ctx, uint(filmId))
	if err != nil {
		return
	}

	film.Title = filmUpdateRequest.Title

	err = f.repo.Save(&film).Error
	if err != nil {
		return
	}
	return
}

func (f *FilmService) DeleteFilm(ctx context.Context, filmId uint) (err error) {
	film, err := f.repo.FindFilm(ctx, uint(filmId))
	if err != nil {
		return
	}

	err = f.repo.Delete(&film).Error
	if err != nil {
		return
	}
	return
}
