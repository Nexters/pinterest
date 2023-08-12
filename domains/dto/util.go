package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
)

func ToFilmDtoList(FilmList []entities.Film) (FilmDtoList []FilmDetailResponse, err error) {
	for _, film := range FilmList {
		FilmInfo := FilmDetailResponse{
			FilmID:        film.ID,
			Title:         film.Title,
			Order:         film.Order,
			PhotoCutCount: film.PhotoCutCount,
			Likes:         film.Likes,
			UserID:        film.UserID,
		}
		FilmInfo.PhotoCuts, err = ToPhotoCutDtoList(film.PhotoCuts)
		if err != nil {
			return
		}

		if len(FilmInfo.PhotoCuts) == 0 {
			FilmInfo.PhotoCuts = []PhotoCutDetailResponse{}
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
