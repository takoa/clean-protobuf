package connect

import (
	"context"
	"io"
	"math"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
)

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points, number of known features visited, total distance traveled, and
// total time spent.
func (s *RouteGuideServer) RecordRoute(ctx context.Context, stream *connect.ClientStream[api.Point]) (*connect.Response[api.RouteSummary], error) {
	receive := func() (model.Point, error) {
		if !stream.Receive() {
			if err := stream.Err(); err != nil {
				return model.Point{}, err
			}
			return model.Point{}, io.EOF
		}
		point, err := converter.ToModelPoint(stream.Msg())
		if err != nil {
			return model.Point{}, err
		}
		return point, nil
	}
	var routeSummary api.RouteSummary
	onFinished := func(pointCount int32, matchedFeatures []*model.Feature, totalDistance int32, elapsedTime time.Duration) error {
		routeSummary = api.RouteSummary{
			PointCount:   pointCount,
			FeatureCount: int32(len(matchedFeatures)),
			Distance:     totalDistance,
			ElapsedTime:  int32(math.Round(elapsedTime.Seconds())),
		}
		return nil
	}
	if err := s.recordRouteHandler.Invoke(
		ctx,
		receive,
		onFinished,
	); err != nil {
		return nil, err
	}

	return connect.NewResponse(&routeSummary), nil
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *RouteGuideServer) RouteChat(ctx context.Context, stream *connect.BidiStream[api.RouteNote, api.RouteNote]) error {
	receive := func() (model.Point, string, error) {
		newRequest, err := stream.Receive()
		if err != nil {
			return model.Point{}, "", err
		}
		point, err := converter.ToModelPoint(newRequest.Location)
		if err != nil {
			return model.Point{}, "", err
		}
		return point, newRequest.Message, nil
	}
	send := func(location model.Point, message string) error {
		return stream.Send(&api.RouteNote{
			Location: converter.ToGRPCPoint(location),
			Message:  message,
		})
	}
	if err := s.routeChatHandler.Invoke(ctx, receive, send); err != nil {
		return err
	}

	return nil
}
