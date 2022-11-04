package interfaces

import (
	"context"
	"sftp/domain"
)

type FeeRateRepository interface {
	ListFeeRates(ctx context.Context) ([]*domain.FeeRate, error)
}
