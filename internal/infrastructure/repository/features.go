package repository

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Features struct {
	features []*model.Feature
}

func NewFeatures(features []*model.Feature) *Features {
	return &Features{
		features: features,
	}
}

func (r *Features) Find(ctx context.Context) ([]*model.Feature, error) {
	s := make([]*model.Feature, len(r.features))
	copy(s, r.features)

	return s, nil
}
