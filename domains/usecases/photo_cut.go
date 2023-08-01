package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type PhotoCutService struct {
	repo *entities.PhotoCutRepository
}

func NewPhotoCutService(repo *entities.PhotoCutRepository) *PhotoCutService {
	return &PhotoCutService{repo}
}

func (pc *PhotoCutService) FindByPhotoCutId(ctx context.Context, photoCutId uint) (photoCutResponse dto.PhotoCutDetailResponse, err error) {
	photoCut, err := pc.repo.FindPhotoCut(ctx, photoCutId)
	if err != nil {
		return
	}

	photoCutResponse = dto.PhotoCutDetailResponse{
		Title:     photoCut.Title,
		Text:      photoCut.Text,
		Link:      photoCut.Link,
		Image:     photoCut.Image,
		Likes:     photoCut.Likes,
		FilmID:    photoCut.FilmID,
		CreatedAt: photoCut.CreatedAt,
	}
	return
}

func (pc *PhotoCutService) CreatePhotoCut(
	ctx context.Context,
	photoCutCreationRequest dto.PhotoCutCreationRequest,
) (photoCutResponse dto.PhotoCutDetailResponse, err error) {
	photoCut := entities.PhotoCut{
		Title:  photoCutCreationRequest.Title,
		Text:   photoCutCreationRequest.Text,
		Link:   photoCutCreationRequest.Link,
		Image:  photoCutCreationRequest.Image,
		FilmID: photoCutCreationRequest.FilmID,
	}

	savedPhotoCut, err := pc.repo.SavePhotoCut(ctx, photoCut)
	if err != nil {
		return
	}

	photoCutResponse = dto.PhotoCutDetailResponse{
		Title:     savedPhotoCut.Title,
		Text:      savedPhotoCut.Text,
		Link:      savedPhotoCut.Link,
		Image:     savedPhotoCut.Image,
		Likes:     savedPhotoCut.Likes,
		FilmID:    savedPhotoCut.FilmID,
		CreatedAt: savedPhotoCut.CreatedAt,
	}
	return
}
