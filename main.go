package main

import (
	"sftp/config"
	"sftp/util"

	"gorm.io/gorm"
)

var logger = util.NewLogger()

func main() {
	// DB接続処理
	databaseConfig := config.GetDatabaseConfig()
	db, err := databaseConfig.ConnectDatabaseWithGorm(10)
	if err != nil {
		logger.Fatalf("Failed to connect database")
	}
	defer closeGormDB(db)
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
