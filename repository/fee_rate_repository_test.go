package repository

import (
	"context"
	"fmt"
	"regexp"
	"sftp/constant"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestGetFeeRate(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// contextに値を設定
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.ProcessingDateContextKey, "2022-11-07 00:00:00")

	// mock設定
	rows := sqlmock.NewRows([]string{"fee_rate_id", "fee_rate", "payment_method", "start_date", "end_date"}).
		AddRow("TEST1", 0.3, "paypay", "2022-01-01", "2030-12-31")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "fee_rates" WHERE start_date <= $1 AND end_date >= $2`)).
		WillReturnRows(rows)

		// リポジトリ初期化
	feeRateRepository := NewFeeRateRepository(mockDB)
	feeRate, err := feeRateRepository.ListFeeRates(ctx)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(feeRate)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Fee Rate: %v", err)
	}
}
