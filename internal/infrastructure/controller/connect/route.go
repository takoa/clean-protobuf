package connect

import (
	"context"
	"io"
	"math"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
)

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points, number of known features visited, total distance traveled, and
// total time spent.
func (s *RouteGuideServer) RecordRoute(
	ctx context.Context,
	stream *connect.ClientStream[routeguidev1.RecordRouteRequest],
) (*connect.Response[routeguidev1.RecordRouteResponse], error) {
	receive := func() (model.Point, error) {
		if !stream.Receive() {
			if err := stream.Err(); err != nil {
				return model.Point{}, err
			}
			return model.Point{}, io.EOF
		}
		point, err := converter.ToModelPoint(stream.Msg().NewPoint)
		if err != nil {
			return model.Point{}, err
		}
		return point, nil
	}
	var routeSummary routeguidev1.RouteSummary
	onFinished := func(pointCount int32, matchedFeatures []*model.Feature, totalDistance int32, elapsedTime time.Duration) error {
		routeSummary = routeguidev1.RouteSummary{
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

	return connect.NewResponse(&routeguidev1.RecordRouteResponse{
		RouteSummary: &routeSummary,
	}), nil
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *RouteGuideServer) RouteChat(
	ctx context.Context,
	stream *connect.BidiStream[routeguidev1.RouteChatRequest, routeguidev1.RouteChatResponse],
) error {
	receive := func() (model.Point, string, error) {
		newRequest, err := stream.Receive()
		if err != nil {
			return model.Point{}, "", err
		}
		point, err := converter.ToModelPoint(newRequest.NewMessage.Location)
		if err != nil {
			return model.Point{}, "", err
		}
		return point, newRequest.NewMessage.Message, nil
	}
	send := func(location model.Point, message string) error {
		return stream.Send(&routeguidev1.RouteChatResponse{
			Message: &routeguidev1.RouteNote{
				Location: converter.ToGRPCPoint(location),
				Message:  message,
			},
		})
	}
	if err := s.routeChatHandler.Invoke(ctx, receive, send); err != nil {
		return err
	}

	return nil
}
