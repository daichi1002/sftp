package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sftp/constant"
	"sftp/domain"
)

func (s Service) calcFeeAmount(feeRates []*domain.FeeRate) []*domain.SalesData {
	// tmpディレクトリに生成されたファイルの値をsalesData構造体に格納
	localFilePath := filepath.Join(constant.TMP_DIR, constant.FILE_NAME)
	salesData, err := os.Open(localFilePath)

	if err != nil {
		s.logger.Fatalf("Failed to read file : %v", err)
	}

	// CSVファイルの読み込み
	r := csv.NewReader(salesData)
	rows, err := r.ReadAll()

	if err != nil {
		s.logger.Fatalf("Failed to read rows : %v")
	}

	fmt.Println(rows)

	return make([]*domain.SalesData, 0)
}
