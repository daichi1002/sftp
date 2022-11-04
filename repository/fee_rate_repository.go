package repository

import (
	"context"
	"sftp/constant"
	"sftp/domain"
	"sftp/interfaces"

	"gorm.io/gorm"
)

type FeeRateRepository struct {
	db *gorm.DB
}

func NewFeeRateRepository(db *gorm.DB) interfaces.FeeRateRepository {
	return &FeeRateRepository{
		db: db,
	}
}

func (r *FeeRateRepository) ListFeeRates(ctx context.Context) ([]*domain.FeeRate, error) {
	FeeRates := make([]*domain.FeeRate, 0)
	processingDate := ctx.Value(constant.ProcessingDateContextKey).(string)

	err := r.db.Where(`start_date <= ? AND end_date >= ? `, processingDate, processingDate).Find(&FeeRates).Error
	return FeeRates, err
}
