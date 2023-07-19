package converter

import (
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
	"golang.org/x/xerrors"
)

func ToModelPoint(p *routeguidev1.Point) (model.Point, error) {
	if p == nil {
		return model.Point{}, xerrors.Errorf("p: %w", errors.ErrNilArgument)
	}
	return model.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}, nil
}

func ToGRPCPoint(p model.Point) *routeguidev1.Point {
	return &routeguidev1.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}
