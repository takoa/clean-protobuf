package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"gorm.io/gorm"
)

type Features struct {
	Repository[model.Feature]
}

func NewFeatures(
	db *gorm.DB,
) *Features {
	return &Features{
		Repository: Repository[model.Feature]{
			DB: db,
		},
	}
}

func (r *Features) FindByPoint(ctx context.Context, p model.Point) (result *model.Feature, err error) {
	tx := r.DB.WithContext(ctx).
		Where("latitude = ? and longitude = ?", p.Latitude, p.Longitude).
		First(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}
