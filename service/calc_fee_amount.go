package service

import (
	"os"
	"path/filepath"
	"sftp/constant"
	"sftp/domain"

	"github.com/jszwec/csvutil"
)

func (s Service) calcFeeAmount(feeRates []*domain.FeeRate) (salesData []*domain.SalesData) {
	// tmpディレクトリに生成されたファイルの値をsalesData構造体に格納
	localFilePath := filepath.Join(constant.TMP_DIR, constant.FILE_NAME)
	data, err := os.ReadFile(localFilePath)

	if err != nil {
		s.logger.Fatalf("Failed to read file : %v", err)
	}
	// CSV内データを構造体にマッピング
	err = csvutil.Unmarshal(data, &salesData)

	if err != nil {
		s.logger.Fatalf("Failed to unmarshal data : %v", err)
	}

	// 決済手段に応じて、各手数料を計算
	calcData(feeRates, salesData)

	return
}

func calcData(feeRates []*domain.FeeRate, salesData []*domain.SalesData) {
	for _, data := range salesData {
		for _, rate := range feeRates {
			if data.PaymentMethod == rate.PaymentMethod {
				data.FeeAmount = data.SalesAmount.Mul(rate.FeeRate)
			}
		}
	}
}
