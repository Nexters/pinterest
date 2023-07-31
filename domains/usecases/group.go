package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type FilmService struct {
	repo *entities.FilmRepository
}

func NewFilmService(repo *entities.FilmRepository) *FilmService {
	return &FilmService{repo}
}

func (g *FilmService) FindByFilmId(ctx context.Context, filmId uint) (filmDetailResponse dto.FilmDetailResponse, err error) {
	film, err := g.repo.FindFilm(ctx, filmId)
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

func (g *FilmService) CreateFilm(ctx context.Context, filmCreationRequest dto.FilmCreationRequest) (filmDetailResponse dto.FilmDetailResponse, err error) {
	order, err := g.repo.CountOrderByUserId(ctx, filmCreationRequest.UserID)
	if err != nil {
		return
	}

	film := entities.Film{
		Title:  filmCreationRequest.Title,
		UserID: filmCreationRequest.UserID,
		Order:  uint(order),
	}

	savedFilm, err := g.repo.SaveFilm(ctx, film)
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

func (g *FilmService) FindAllFilms(ctx context.Context, userId string) (filmList []dto.Film, err error) {
	films, err := g.repo.FindAllFilmsInOrder(ctx, userId)
	if err != nil {
		return
	}

	filmList, err = dto.ToFilmDtoList(films)
	if err != nil {
		return
	}
	return
}
