package connect

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
	"golang.org/x/xerrors"
)

// GetFeature returns the feature at the given point.
func (s *RouteGuideServer) GetFeature(
	ctx context.Context,
	request *connect.Request[routeguidev1.GetFeatureRequest],
) (*connect.Response[routeguidev1.GetFeatureResponse], error) {
	p, err := converter.ToModelPoint(request.Msg.Point)
	if err != nil {
		return nil, xerrors.Errorf("point: %w", errors.ErrNilArgument)
	}

	feature, err := s.getFeatureHandler.Invoke(ctx, p)
	if err != nil {
		return nil, xerrors.Errorf("failed to get feature at %v: %w", p, err)
	}

	return connect.NewResponse(&routeguidev1.GetFeatureResponse{
		Feature: converter.ToGRPCFeature(feature),
	}), nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *RouteGuideServer) ListFeatures(
	ctx context.Context,
	request *connect.Request[routeguidev1.ListFeaturesRequest],
	stream *connect.ServerStream[routeguidev1.ListFeaturesResponse],
) error {
	r, err := converter.ToModelRectangle(request.Msg.SearchArea)
	if err != nil {
		return xerrors.Errorf("rect: %w", errors.ErrNilArgument)
	}
	send := func(f *model.Feature) error {
		return stream.Send(&routeguidev1.ListFeaturesResponse{
			Feature: converter.ToGRPCFeature(f),
		})
	}
	if err := s.listFeaturesHandler.Invoke(ctx, r, send); err != nil {
		return err
	}

	return nil
}
