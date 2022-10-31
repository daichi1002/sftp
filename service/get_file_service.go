package service

import (
	"context"
	"sftp/interfaces"
	"sftp/util"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger *zap.SugaredLogger
}

// NewGetFileService の初期化処理
func NewGetFileService() interfaces.Service {

	logger := util.NewLogger()
	// サービスの生成
	service := &Service{
		logger: logger,
	}
	return service
}

func (s *Service) Execute(ctx context.Context, db *gorm.DB) {}
