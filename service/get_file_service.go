package service

import (
	"context"
	"sftp/interfaces"
	"sftp/util"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger     *zap.SugaredLogger
	repository interfaces.FeeRateRepository
}

// NewGetFileService の初期化処理
func NewGetFileService(feeRateRepository interfaces.FeeRateRepository) interfaces.Service {

	logger := util.NewLogger()
	// サービスの生成
	service := &Service{
		logger:     logger,
		repository: feeRateRepository,
	}
	return service
}

func (s *Service) Execute(ctx context.Context, db *gorm.DB) {

	s.logger.Info("start get fee rate")
	feeRates := s.getFeeRate(ctx, db)
	s.logger.Info("end get fee rate")

	s.logger.Info("start connect sftp")
	sshClient, sftpClient := s.connectSftp()
	s.logger.Info("end connect sftp")

	s.logger.Info("start get file from sftp")
	s.getSalesData(sshClient, sftpClient)
	s.logger.Info("end get file from sftp")

	s.logger.Info("start calc fee amount")
	salesData := s.calcFeeAmount(feeRates)

	s.logger.Info("end calc fee amount")
}
