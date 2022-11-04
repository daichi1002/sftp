package service

import (
	"context"
	"fmt"
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
	feeRates := s.getFeeRate(db)
	fmt.Println(feeRates)
	s.logger.Info("end get fee rate")

	s.logger.Info("start connect sftp")
	sshClient, sftpClient := s.connectSftp()
	s.logger.Info("end connect sftp")

	s.logger.Info("start get file from sftp")
	file := s.getFile(sshClient, sftpClient)
	fmt.Println(file)
	s.logger.Info("end get file from sftp")
}
