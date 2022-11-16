package service

import (
	"os"
	"sftp/domain"
	"sftp/test_util/assertion"
	"sftp/util"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var logger = util.NewLogger()

func TestCalcFeeAmount(t *testing.T) {
	dir, _ := os.Getwd()
	testFilePath := dir + "/testdata/test.csv"

	testCases := []struct {
		desc string
		args []*domain.FeeRate
		exps []*domain.SalesData
	}{
		{
			desc: "ファイルからデータを取得し、インスタンスを生成、手数料の計算ができる",
			args: []*domain.FeeRate{
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
			exps: []*domain.SalesData{
				{
					ProductName:       "トマト",
					SalesAmount:       decimal.NewFromInt(1000),
					FeeAmount:         decimal.NewFromInt(300),
					Quantity:          1,
					PaymentMethod:     "paypay",
					TransactionStatus: "支払",
					TransactionDate:   "2022-11-01 10:00:00",
				},
				{
					ProductName:       "キャベツ",
					SalesAmount:       decimal.NewFromInt(2000),
					FeeAmount:         decimal.NewFromInt(600),
					Quantity:          2,
					PaymentMethod:     "linepay",
					TransactionStatus: "返品",
					TransactionDate:   "2022-11-02 10:00:00",
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := Service{
				logger: logger,
			}
			act := s.calcFeeAmount(tC.args, testFilePath)
			err := assertion.AssertSales(tC.exps, act)

			if err != nil {
				t.Error(err)
			}
		})
	}
}

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
		}, {
			desc: "売上データが手数料率に紐づかないこと",
			args: Args{
				feeRates: []*domain.FeeRate{
					{
						FeeRateId:     "TEST1",
						FeeRate:       decimal.NewFromFloat(0.5),
						PaymentMethod: "paypay",
						StartDate:     "2022-01-01",
						EndDate:       "2030-12-31",
					},
					{
						FeeRateId:     "TEST2",
						FeeRate:       decimal.NewFromFloat(0.5),
						PaymentMethod: "linepay",
						StartDate:     "2022-01-01",
						EndDate:       "2030-12-31",
					},
				},
				data: &domain.SalesData{
					SalesDataId:       "TEST1",
					ProductName:       "タマネギ",
					SalesAmount:       decimal.NewFromInt(1000),
					FeeAmount:         decimal.NewFromInt(0),
					Quantity:          1,
					PaymentMethod:     "aupay",
					TransactionStatus: "支払",
					TransactionDate:   "2022-11-08 10:00:00",
				},
			},
			exps: &domain.SalesData{
				SalesDataId:       "TEST1",
				ProductName:       "トマト",
				SalesAmount:       decimal.NewFromInt(1000),
				FeeAmount:         decimal.NewFromInt(0),
				Quantity:          1,
				PaymentMethod:     "aupay",
				TransactionStatus: "支払",
				TransactionDate:   "2022-11-08 10:00:00",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := Service{
				logger: logger,
			}
			s.calcData(tC.args.feeRates, tC.args.data)

			tc, _ := tC.args.data.FeeAmount.Value()
			exp, _ := tC.exps.FeeAmount.Value()

			assert.Equal(t, tc, exp)
		})
	}
}
