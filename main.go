package main

import (
	"context"
	"sftp/config"
	"sftp/service"
	"sftp/util"

	"gorm.io/gorm"
)

var logger = util.NewLogger()

func main() {
	// contextに値を設定
	ctx := context.Background()
	// DB接続処理
	databaseConfig := config.GetDatabaseConfig()
	db, err := databaseConfig.ConnectDatabaseWithGorm(10)
	if err != nil {
		logger.Fatalf("Failed to connect database")
	}
	defer closeGormDB(db)

	// service初期化処理
	service := service.NewGetFileService()

	// バッチ処理実行
	service.Execute(ctx, db)
}

func closeGormDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("failed to extract sql db from gorm.DB instance, error: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		logger.Errorf("failed to close sql db, error: %v", err)
	}
}
