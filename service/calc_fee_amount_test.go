package service

import (
	"sftp/domain"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCalcData(t *testing.T) {
	type Args struct {
		feeRates []*domain.FeeRate
		data     *domain.SalesData
	}

	testCases := []struct {
		desc string
		args Args
		exps *domain.SalesData
	}{
		{
			desc: "売上データが手数料率に紐づくこと",
			args: Args{
				feeRates: []*domain.FeeRate{
					{
						FeeRateId:     "TEST1",
						FeeRate:       decimal.NewFromFloat(0.3),
						PaymentMethod: "paypay",
						StartDate:     "2022-01-01",
						EndDate:       "2030-12-31",
					},
					{
						FeeRateId:     "TEST2",
						FeeRate:       decimal.NewFromFloat(0.3),
						PaymentMethod: "linepay",
						StartDate:     "2022-01-01",
						EndDate:       "2030-12-31",
					},
				},
				data: &domain.SalesData{
					SalesDataId:       "TEST1",
					ProductName:       "トマト",
					SalesAmount:       decimal.NewFromInt(1000),
					FeeAmount:         decimal.NewFromInt(0),
					Quantity:          1,
					PaymentMethod:     "paypay",
					TransactionStatus: "支払",
					TransactionDate:   "2022-11-08 10:00:00",
				},
			},
			exps: &domain.SalesData{
				SalesDataId:       "TEST1",
				ProductName:       "トマト",
				SalesAmount:       decimal.NewFromInt(1000),
				FeeAmount:         decimal.NewFromInt(300),
				Quantity:          1,
				PaymentMethod:     "paypay",
				TransactionStatus: "支払",
				TransactionDate:   "2022-11-08 10:00:00",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := Service{}
			s.calcData(tC.args.feeRates, tC.args.data)

			tc, _ := tC.args.data.FeeAmount.Value()
			exp, _ := tC.exps.FeeAmount.Value()
			assert.Equal(t, tc, exp)
		})
	}
}
