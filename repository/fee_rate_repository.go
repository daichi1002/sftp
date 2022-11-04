package repository

import (
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

func (r *FeeRateRepository) ListFeeRates() ([]*domain.FeeRate, error) {
	FeeRates := make([]*domain.FeeRate, 0)
	err := r.db.Find(&FeeRates).Error

	return FeeRates, err
}
