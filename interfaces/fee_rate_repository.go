package interfaces

import (
	"sftp/domain"
)

type FeeRateRepository interface {
	ListFeeRates() ([]*domain.FeeRate, error)
}
