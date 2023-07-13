package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Features interface {
	Repository[model.Feature]

	FindByPoint(ctx context.Context, p model.Point) (result *model.Feature, err error)
}
