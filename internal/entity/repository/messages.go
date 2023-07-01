package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Messages interface {
	Find(ctx context.Context, p *model.Point) ([]string, error)
	Create(ctx context.Context, p *model.Point, message string) error
}
