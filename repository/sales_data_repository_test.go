package repository

import (
	"regexp"
	"sftp/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
)

func TestStore(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	arg := []*domain.SalesData{
		{
			SalesDataId:       "TEST1",
			ProductName:       "トマト",
			SalesAmount:       decimal.NewFromInt(1000),
			FeeAmount:         decimal.NewFromInt(300),
			Quantity:          1,
			PaymentMethod:     "paypay",
			TransactionStatus: "支払",
			TransactionDate:   "2022-11-08 10:00:00",
		},
		{
			SalesDataId:       "TEST2",
			ProductName:       "タマネギ",
			SalesAmount:       decimal.NewFromInt(1000),
			FeeAmount:         decimal.NewFromInt(200),
			Quantity:          1,
			PaymentMethod:     "aupay",
			TransactionStatus: "支払",
			TransactionDate:   "2022-11-08 10:00:00",
		},
	}

	// mock設定
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email") VALUES ($1,$2,$3,$4,$5)`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	salesDataRepository := NewSalesDataRepository(mockDB)
	error := salesDataRepository.CreaateSalesData(arg)

	if error != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create Sales Data: %v", err)
	}
}
