package interfaces

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	Execute(ctx context.Context, db *gorm.DB)
}
