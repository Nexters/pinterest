package usecases

import (
	"context"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/entities"
)

type VisitLogService struct {
	repo *entities.VisitLogRepository
}

func NewVisitLogService(repo *entities.VisitLogRepository) *VisitLogService {
	return &VisitLogService{repo}
}

func (v *VisitLogService) FindAll(ctx context.Context, UserID string) ([]entities.VisitLog, error) {
	return v.repo.FindAllVisitLogsByUserID(ctx, UserID)
}

func (v *VisitLogService) Create(ctx context.Context, dto dto.VisitLogCreationRequest) (log entities.VisitLog, err error) {
	logParam := dto.ToEntity()
	log, err = v.repo.CreateVisitLog(ctx, logParam)
	return
}

func (v *VisitLogService) Delete(ctx context.Context, logID uint) error {
	log := entities.VisitLog{}
	log.ID = logID

	return v.repo.DeleteVisitLog(ctx, log)
}
