package connect

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
	"golang.org/x/xerrors"
)

// GetFeature returns the feature at the given point.
func (s *RouteGuideServer) GetFeature(ctx context.Context, request *connect.Request[api.Point]) (*connect.Response[api.Feature], error) {
	p, err := converter.ToModelPoint(request.Msg)
	if err != nil {
		return nil, xerrors.Errorf("point: %w", errors.ErrNilArgument)
	}

	feature, err := s.getFeatureHandler.Invoke(ctx, p)
	if err != nil {
		return nil, xerrors.Errorf("failed to get feature at %v: %w", p, err)
	}

	return connect.NewResponse(converter.ToGRPCFeature(feature)), nil
}

// ListFeatures lists all features contained within the given bounding Rectangle.
func (s *RouteGuideServer) ListFeatures(ctx context.Context, request *connect.Request[api.Rectangle], stream *connect.ServerStream[api.Feature]) error {
	r, err := converter.ToModelRectangle(request.Msg)
	if err != nil {
		return xerrors.Errorf("rect: %w", errors.ErrNilArgument)
	}
	send := func(f *model.Feature) error {
		return stream.Send(converter.ToGRPCFeature(f))
	}
	if err := s.listFeaturesHandler.Invoke(ctx, r, send); err != nil {
		return err
	}

	return nil
}
