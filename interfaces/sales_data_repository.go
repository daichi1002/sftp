package interfaces

import "sftp/domain"

type SalesDataRepository interface {
	CreaateSalesData([]*domain.SalesData) error
}
