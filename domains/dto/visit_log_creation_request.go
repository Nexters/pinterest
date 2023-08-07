package dto

import "github.com/Nexters/pinterest/domains/entities"

type VisitLogCreationRequest struct {
	UserID string `json:"user_id" validate:"required,ascii"`
	Name   string `json:"name" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

func (v VisitLogCreationRequest) ToEntity() entities.VisitLog {
	return entities.VisitLog{
		UserID: v.UserID,
		Text:   v.Text,
		Name:   v.Name,
	}
}
