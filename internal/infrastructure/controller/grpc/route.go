package controller

import (
	"math"
	"time"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/protobuf/converter"
)

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points, number of known features visited, total distance traveled, and
// total time spent.
func (s *RouteGuideServer) RecordRoute(stream api.RouteGuide_RecordRouteServer) error {
	receive := func() (model.Point, error) {
		p, err := stream.Recv()
		if err != nil {
			return model.Point{}, err
		}

		point, err := converter.ToModelPoint(p)
		if err != nil {
			return model.Point{}, err
		}

		return point, nil
	}
	onFinished := func(pointCount int32, matchedFeatures []*model.Feature, totalDistance int32, elapsedTime time.Duration) error {
		return stream.SendAndClose(&api.RouteSummary{
			PointCount:   pointCount,
			FeatureCount: int32(len(matchedFeatures)),
			Distance:     totalDistance,
			ElapsedTime:  int32(math.Round(elapsedTime.Seconds())),
		})
	}
	if err := s.recordRouteHandler.Invoke(
		stream.Context(),
		receive,
		onFinished,
	); err != nil {
		return err
	}

	return nil
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *RouteGuideServer) RouteChat(stream api.RouteGuide_RouteChatServer) error {
	receive := func() (model.Point, string, error) {
		in, err := stream.Recv()
		if err != nil {
			return model.Point{}, "", err
		}

		point, err := converter.ToModelPoint(in.Location)
		if err != nil {
			return model.Point{}, "", err
		}

		return point, in.Message, nil
	}
	send := func(location model.Point, message string) error {
		return stream.Send(&api.RouteNote{Location: converter.ToGRPCPoint(location), Message: message})
	}
	if err := s.routeChatHandler.Invoke(
		stream.Context(),
		receive,
		send,
	); err != nil {
		return err
	}

	return nil
}
