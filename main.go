package main

import (
	"context"
	"flag"
	"sftp/config"
	"sftp/constant"
	"sftp/repository"
	"sftp/service"
	"sftp/util"

	"gorm.io/gorm"
)

var logger = util.NewLogger()

func main() {
	// 起動引数の処理
	reqProcessingDate := parseArgs()
	// contextに値を設定
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.ProcessingDateContextKey, reqProcessingDate)
	// DB接続処理
	databaseConfig := config.GetDatabaseConfig()
	db, err := databaseConfig.ConnectDatabaseWithGorm(10)
	if err != nil {
		logger.Fatalf("Failed to connect database")
	}
	defer closeGormDB(db)

	// tmpディレクトリを作成
	mkdirErr := util.MakeDirectory(constant.TMP_DIR)
	if mkdirErr != nil {
		logger.Fatal(mkdirErr.Error())
	}

	// リポジトリ初期化
	feeRateRepository := repository.NewFeeRateRepository(db)
	salesDataRepository := repository.NewSalesDataRepository(db)

	// service初期化処理
	service := service.NewGetFileService(feeRateRepository, salesDataRepository)

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

func parseArgs() string {
	// 処理日時を取得
	reqProcessingDate := flag.String("date", "", "処理日時")

	flag.Parse()

	// 処理日時をフォーマットする
	processingDate, err := util.CreateProcessingDate(*reqProcessingDate)

	if err != nil {
		logger.Fatalf("Failed to parse date", err)
	}

	return processingDate
}
