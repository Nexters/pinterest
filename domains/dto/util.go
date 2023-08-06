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

func ToVisitLogDtoList(visitLogList []entities.VisitLog) (visitLogDtoList []VisitLog, err error) {
	for _, visitLog := range visitLogList {
		visitLogInfo := VisitLog{
			UserID: visitLog.UserID,
			Name:   visitLog.Name,
			Text:   visitLog.Text,
		}
		validate := validator.New()
		err := validate.Struct(visitLogInfo)
		if err != nil {
			return visitLogDtoList, err
		}
		visitLogDtoList = append(visitLogDtoList, visitLogInfo)
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
