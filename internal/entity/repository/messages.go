package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Messages interface {
	Repository[model.Message]

	FindByFeature(ctx context.Context, f *model.Feature, orderBy string) (result []*model.Message, err error)
}
