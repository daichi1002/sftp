package assertion

import (
	"fmt"
	"sftp/domain"
)

type SalesData struct {
	ProductName       string
	SalesAmount       string
	FeeAmount         string
	Quantity          int
	PaymentMethod     string
	TransactionStatus string
	TransactionDate   string
}

func newSalesData(sale *domain.SalesData) SalesData {

	data := SalesData{
		ProductName:       sale.ProductName,
		SalesAmount:       sale.SalesAmount.String(),
		FeeAmount:         sale.FeeAmount.String(),
		Quantity:          sale.Quantity,
		PaymentMethod:     sale.PaymentMethod,
		TransactionStatus: sale.TransactionStatus,
		TransactionDate:   sale.TransactionDate,
	}

	return data
}

func AssertSales(exps, acts []*domain.SalesData) error {
	if exps == nil && acts == nil {
		return nil
	}

	for i, exp := range exps {
		e := newSalesData(exp)
		a := newSalesData(acts[i])

		if e != a {
			return fmt.Errorf("-expected: %v\n+actual: %v", e, a)
		}
	}

	return nil
}
