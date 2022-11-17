package service

import (
	"context"
	"sftp/domain"
)

func (s Service) getFeeRate(ctx context.Context) []*domain.FeeRate {

	// 手数料率の取得
	feeRates, err := s.feeRateRepository.ListFeeRates(ctx)

	if err != nil {
		s.logger.Fatal("Failed to get fee rate")
	}

	return feeRates
}
