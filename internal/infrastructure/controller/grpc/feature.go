package controller

import (
	"context"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
	"golang.org/x/xerrors"
)

// GetFeature returns the feature at the given point.
func (s *RouteGuideServer) GetFeature(ctx context.Context, point *api.Point) (*api.Feature, error) {
	p, err := converter.ToModelPoint(point)
	if err != nil {
		return nil, xerrors.Errorf("point: %w", errors.ErrNilArgument)
	}

	feature, err := s.getFeatureHandler.Invoke(ctx, p)
	if err != nil {
		return nil, xerrors.Errorf("failed to get feature at %v: %w", p, err)
	}

	return converter.ToGRPCFeature(feature), nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *RouteGuideServer) ListFeatures(rect *api.Rectangle, stream api.RouteGuide_ListFeaturesServer) error {
	r, err := converter.ToModelRectangle(rect)
	if err != nil {
		return xerrors.Errorf("rect: %w", errors.ErrNilArgument)
	}
	send := func(f *model.Feature) error {
		return stream.Send(converter.ToGRPCFeature(f))
	}
	if err := s.listFeaturesHandler.Invoke(stream.Context(), r, send); err != nil {
		return err
	}

	return nil
}
