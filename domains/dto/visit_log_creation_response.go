package dto

import (
	"time"

	"github.com/Nexters/pinterest/domains/entities"
)

type VisitLogCreationResponse struct {
	LogID     uint      `json:"log_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
}

func (VisitLogCreationResponse) FromEntity(log entities.VisitLog) VisitLogCreationResponse {
	return VisitLogCreationResponse{
		LogID:     log.ID,
		UserID:    log.UserID,
		CreatedAt: log.CreatedAt,
		Name:      log.Name,
		Text:      log.Text,
	}
}
