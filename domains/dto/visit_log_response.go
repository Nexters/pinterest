package dto

import (
	"time"

	"github.com/Nexters/pinterest/domains/entities"
)

type VisitLogResponse struct {
	LogID     uint      `json:"log_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
}

func (v VisitLogResponse) FromEntity(log entities.VisitLog) VisitLogResponse {
	return VisitLogResponse{
		LogID:     log.ID,
		UserID:    log.UserID,
		CreatedAt: log.CreatedAt,
		Name:      log.Name,
		Text:      log.Text,
	}
}
