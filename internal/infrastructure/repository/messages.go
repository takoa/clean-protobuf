package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"gorm.io/gorm"
)

type Messages struct {
	Repository[model.Message]
}

func NewMessages(
	db *gorm.DB,
) *Messages {
	return &Messages{
		Repository: Repository[model.Message]{
			DB: db,
		},
	}
}

func (r *Messages) FindByFeature(ctx context.Context, f *model.Feature, orderBy string) (result []*model.Message, err error) {
	if f == nil {
		return nil, nil
	}

	tx := r.DB.WithContext(ctx).
		Where("feature_id = ?", f.ID).
		Order(orderBy).
		Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}
