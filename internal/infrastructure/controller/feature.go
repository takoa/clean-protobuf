package controller

import (
	"context"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/grpc"
)

// GetFeature returns the feature at the given point.
func (s *RouteGuideServer) GetFeature(ctx context.Context, point *api.Point) (*api.Feature, error) {
	p := grpc.ToModelPoint(point)

	var feature *model.Feature
	onFeatureFound := func(f *model.Feature) error {
		feature = f
		return nil
	}
	if err := s.featureFinder.Find(ctx, &model.Rectangle{Hi: p, Lo: p}, onFeatureFound); err != nil {
		return nil, err
	} else if feature == nil {
		return &api.Feature{Location: point}, nil
	}

	return grpc.ToGRPCFeature(feature), nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *RouteGuideServer) ListFeatures(rect *api.Rectangle, stream api.RouteGuide_ListFeaturesServer) error {
	onFeatureFound := func(f *model.Feature) error {
		if err := stream.Send(grpc.ToGRPCFeature(f)); err != nil {
			return err
		}
		return nil
	}
	if err := s.featureFinder.Find(context.Background(), grpc.ToModelRectangle(rect), onFeatureFound); err != nil {
		return err
	}

	return nil
}
