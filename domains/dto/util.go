package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/go-playground/validator"
)

func ToFilmDtoList(FilmList []entities.Film) (FilmDtoList []Film, err error) {
	for _, film := range FilmList {
		FilmInfo := Film{
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

func ToPhotoCutDtoList(itemList []entities.PhotoCut) (photoCutDtoList []PhotoCutDetailResponse, err error) {
	for _, item := range itemList {
		photoCutDetail := PhotoCutDetailResponse{
			Title:     item.Title,
			Text:      item.Text,
			Link:      item.Link,
			Image:     item.Image,
			Likes:     item.Likes,
			FilmID:    item.FilmID,
			CreatedAt: item.CreatedAt,
		}
		photoCutDtoList = append(photoCutDtoList, photoCutDetail)
	}
	return
}
