package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type PhotoCutService struct {
	repo  *entities.PhotoCutRepository
	frepo *entities.FilmRepository
}

func NewPhotoCutService(repo *entities.PhotoCutRepository, frepo *entities.FilmRepository) *PhotoCutService {
	return &PhotoCutService{repo, frepo}
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
	film, err := pc.frepo.FindFilm(ctx, photoCutCreationRequest.FilmID)
	if err != nil {
		return
	}

	photoCut := entities.PhotoCut{
		Title:  photoCutCreationRequest.Title,
		Text:   photoCutCreationRequest.Text,
		Link:   photoCutCreationRequest.Link,
		Image:  photoCutCreationRequest.Image,
		FilmID: film.ID,
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

func (pc *PhotoCutService) UpdatePhotoCut(ctx context.Context, photoCutUpdateRequest dto.PhotoCutUpdateRequest) (err error) {
	photoCut, err := pc.repo.FindPhotoCut(ctx, photoCutUpdateRequest.ID)
	if err != nil {
		return
	}

	if photoCutUpdateRequest.Title != "" {
		photoCut.Title = photoCutUpdateRequest.Title
	}
	if photoCutUpdateRequest.Text != "" {
		photoCut.Text = photoCutUpdateRequest.Text
	}
	if photoCutUpdateRequest.Image != "" {
		photoCut.Image = photoCutUpdateRequest.Image
	}

	err = pc.repo.Save(&photoCut).Error
	return
}

func (pc *PhotoCutService) DeletePhotoCut(ctx context.Context, photoCutId uint) (err error) {
	photoCut, err := pc.repo.FindPhotoCut(ctx, uint(photoCutId))
	if err != nil {
		return
	}

	err = pc.repo.Delete(&photoCut).Error
	return
}
