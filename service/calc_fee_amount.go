package service

import (
	"os"
	"sftp/domain"
	"sftp/util"

	"github.com/jszwec/csvutil"
)

func (s Service) calcFeeAmount(feeRates []*domain.FeeRate, localFilePath string) (salesData []*domain.SalesData) {
	// tmpディレクトリに生成されたファイルの値をsalesData構造体に格納
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
	for _, data := range salesData {
		s.calcData(feeRates, data)
		// ULIDの生成
		data.SalesDataId = util.GetUlid()
	}

	return
}

func (s Service) calcData(feeRates []*domain.FeeRate, data *domain.SalesData) {
	for _, rate := range feeRates {
		if data.PaymentMethod == rate.PaymentMethod {
			data.FeeAmount = data.SalesAmount.Mul(rate.FeeRate)
			return
		}
	}

	s.logger.Warnf("sales not associated fee. 商品名: %v 取引金額: %v 支払方法: %v 決済日時: %v",
		data.ProductName, data.SalesAmount, data.PaymentMethod, data.TransactionDate)
}
