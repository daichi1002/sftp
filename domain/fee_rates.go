package domain

type FeeRate struct {
	FeeRateId     string
	FeeRate       int    // 手数料率
	PaymentMethod string // 決済手段
}
