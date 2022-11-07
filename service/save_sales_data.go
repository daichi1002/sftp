package service

import "sftp/domain"

func (s Service) saveSalesData(data []*domain.SalesData) {
	// 売上データの保存
	err := s.salesDataRepository.CreaateSalesData(data)

	if err != nil {
		s.logger.Fatalf("Failed to create sales data : %v", err)
	}
}
