package service

import (
	"context"
	"path/filepath"
	"sftp/constant"
	"sftp/interfaces"
	"sftp/util"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger              *zap.SugaredLogger
	feeRateRepository   interfaces.FeeRateRepository
	salesDataRepository interfaces.SalesDataRepository
}

// NewGetFileService の初期化処理
func NewGetFileService(feeRateRepository interfaces.FeeRateRepository, salesDataRepository interfaces.SalesDataRepository) interfaces.Service {

	logger := util.NewLogger()
	// サービスの生成
	service := &Service{
		logger:              logger,
		feeRateRepository:   feeRateRepository,
		salesDataRepository: salesDataRepository,
	}
	return service
}

func (s *Service) Execute(ctx context.Context, db *gorm.DB) {

	s.logger.Info("start get fee rate")
	feeRates := s.getFeeRate(ctx)
	s.logger.Info("end get fee rate")

	s.logger.Info("start connect sftp")
	sshClient, sftpClient := s.connectSftp()
	s.logger.Info("end connect sftp")

	s.logger.Info("start get file from sftp")
	s.getSalesData(sshClient, sftpClient)
	s.logger.Info("end get file from sftp")

	s.logger.Info("start calc fee amount")
	localFilePath := filepath.Join(constant.TMP_DIR, constant.FILE_NAME)
	salesData := s.calcFeeAmount(feeRates, localFilePath)
	s.logger.Info("end calc fee amount")

	s.logger.Info("start save sales data")
	s.saveSalesData(salesData)
	s.logger.Info("end save sales data")
}
