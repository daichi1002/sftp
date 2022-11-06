package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type SalesData struct {
	SalesDataId       string          // 売上ID
	ProductName       string          `csv:"商品名"`  // 商品名
	SalesAmount       decimal.Decimal `csv:"取引金額"` // 売上金額
	FeeAmount         decimal.Decimal // 手数料額
	Quantity          int             `csv:"数量"`      // 数量
	PaymentMethod     string          `csv:"支払方法"`    // 決済手段
	TransactionStatus string          `csv:"取引ステータス"` // 取引ステータス
	TransactionDate   string          `csv:"決済日時"`
	CreatedAt         time.Time       // 作成日時
	UpdatedAt         time.Time       // 更新日時
}
