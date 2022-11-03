package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Sale struct {
	SaleId        string          // 売上ID
	ProductName   string          // 商品名
	SalesAmount   decimal.Decimal // 売上金額
	FeeAmount     decimal.Decimal // 手数料額
	PaymentMethod string          // 決済手段
	CreatedAt     time.Time       // 作成日時
	UpdatedAt     time.Time       // 更新日時
}
