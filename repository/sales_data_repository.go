package repository

import (
	"sftp/domain"

	"gorm.io/gorm"
)

type SalesDataRepository struct {
	db *gorm.DB
}

func NewSalesDataRepository(db *gorm.DB) *SalesDataRepository {
	return &SalesDataRepository{
		db: db,
	}
}

func (r *SalesDataRepository) CreaateSalesData(data []*domain.SalesData) (err error) {
	err = r.db.Create(&data).Error

	return
}
