package service

import (
	"sftp/domain"

	"gorm.io/gorm"
)

func (s Service) getFeeRate(db *gorm.DB) []*domain.FeeRate {

	// 手数料率の全件取得
	feeRates, err := s.repository.ListFeeRates()

	if err != nil {
		s.logger.Fatal("Failed to get fee rate")
	}

	return feeRates
}
