package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Features interface {
	Find(ctx context.Context) ([]*model.Feature, error)
}
