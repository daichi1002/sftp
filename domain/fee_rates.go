package domain

type FeeRate struct {
	FeeRateId     string
	FeeRate       int    // 手数料率
	PaymentMethod string // 決済手段
	StartDate     string // 適用開始日
	EndDate       string // 提供終了日
}
