package entities

import (
	"context"

	"gorm.io/gorm"
)

type VisitLogRepository struct {
	*gorm.DB
}

func (v *VisitLogRepository) NewVisitLogRepository(db *gorm.DB) *VisitLogRepository {
	return &VisitLogRepository{db}
}

func (v *VisitLogRepository) FindAllVisitLogsByUserID(ctx context.Context, userID string) (logs []VisitLog, err error) {
	tx := v.DB.Find(&logs)
	if tx.Error != nil {
		err = tx.Error
	}

	return
}

func (v *VisitLogRepository) CreateVisitLog(ctx context.Context, logParam VisitLog) (log VisitLog, err error) {
	tx := v.DB.Save(&logParam)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	log = logParam
	return
}

func (v *VisitLogRepository) DeleteVisitLog(ctx context.Context, log VisitLog) error {
	tx := v.DB.Delete(&log)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil
	}
	return nil
}
