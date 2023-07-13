package grpc

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"golang.org/x/xerrors"
)

func ToModelPoint(p *api.Point) (model.Point, error) {
	if p == nil {
		return model.Point{}, xerrors.Errorf("p: %w", errors.ErrNilArgument)
	}
	return model.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}, nil
}

func ToGRPCPoint(p model.Point) *api.Point {
	return &api.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}
