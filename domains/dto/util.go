package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/go-playground/validator"
)

func ToFilmDtoList(FilmList []entities.Film) (FilmDtoList []Film, err error) {
	for _, film := range FilmList {
		FilmInfo := Film{
			ID:            film.ID,
			Title:         film.Title,
			Order:         film.Order,
			PhotoCutCount: film.PhotoCutCount,
			Likes:         film.Likes,
			UserID:        film.UserID,
		}
		validate := validator.New()
		err := validate.Struct(FilmInfo)
		if err != nil {
			return FilmDtoList, err
		}
		FilmDtoList = append(FilmDtoList, FilmInfo)
	}
	return
}

func ToPhotoCutDtoList(photoCutList []entities.PhotoCut) (photoCutDtoList []PhotoCutDetailResponse, err error) {
	for _, photoCut := range photoCutList {
		photoCutDetail := PhotoCutDetailResponse{
			ID:        photoCut.ID,
			Title:     photoCut.Title,
			Text:      photoCut.Text,
			Link:      photoCut.Link,
			Image:     photoCut.Image,
			Likes:     photoCut.Likes,
			FilmID:    photoCut.FilmID,
			CreatedAt: photoCut.CreatedAt,
		}
		photoCutDtoList = append(photoCutDtoList, photoCutDetail)
	}
	return
}
