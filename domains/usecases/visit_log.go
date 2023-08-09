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

func (v *VisitLogService) FindAll(ctx context.Context, UserID string) (visitLogs []dto.VisitLogResponse, err error) {
	logs, err := v.repo.FindAllVisitLogsByUserID(ctx, UserID)
	if err != nil {
		return
	}

	for _, log := range logs {
		v := dto.VisitLogResponse{}
		visitLogs = append(visitLogs, v.FromEntity(log))
	}

	return
}

func (v *VisitLogService) Create(ctx context.Context, dto dto.VisitLogCreationRequest) (res dto.VisitLogCreationResponse, err error) {
	logParam := dto.ToEntity()
	log, err := v.repo.CreateVisitLog(ctx, logParam)
	res = res.FromEntity(log)
	return
}

func (v *VisitLogService) Delete(ctx context.Context, logID uint) error {
	log := entities.VisitLog{}
	log.ID = logID

	return v.repo.DeleteVisitLog(ctx, log)
}
