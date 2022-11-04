package service

import (
	"context"
	"sftp/domain"

	"gorm.io/gorm"
)

func (s Service) getFeeRate(ctx context.Context, db *gorm.DB) []*domain.FeeRate {

	// 手数料率の取得
	feeRates, err := s.repository.ListFeeRates(ctx)

	if err != nil {
		s.logger.Fatal("Failed to get fee rate")
	}

	return feeRates
}
