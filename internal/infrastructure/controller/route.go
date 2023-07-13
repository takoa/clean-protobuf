package controller

import (
	"io"
	"time"

	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/controller/grpc"
	"golang.org/x/xerrors"
)

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points, number of known features visited, total distance traveled, and
// total time spent.
func (s *RouteGuideServer) RecordRoute(stream api.RouteGuide_RecordRouteServer) error {
	var pointCount, featureCount, distance int32
	var previousPoint model.Point
	startTime := time.Now()

	for loops := true; loops; {
		point, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
			return stream.SendAndClose(&api.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(time.Since(startTime).Seconds()),
			})
		}

		p, err := grpc.ToModelPoint(point)
		if err != nil {
			return xerrors.Errorf("point: %w", errors.ErrNilArgument)
		}
		onPointAdded := func(p model.Point, matchedFeature *model.Feature, addedDistance int32) error {
			pointCount++
			if matchedFeature != nil {
				featureCount++
			}
			distance += addedDistance
			previousPoint = p

			return nil
		}
		if err := s.routeRecorder.GetInformation(
			stream.Context(),
			previousPoint,
			p,
			onPointAdded,
		); err != nil {
			return err
		}
	}

	return nil
}

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *RouteGuideServer) RouteChat(stream api.RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		p, err := grpc.ToModelPoint(in.Location)
		if err != nil {
			return xerrors.Errorf("point: %w", errors.ErrNilArgument)
		}
		onPostedRouteMessage := func(message string) error {
			if err := stream.Send(&api.RouteNote{Location: in.Location, Message: message}); err != nil {
				return err
			}
			return nil
		}
		if err := s.routeMessagePoster.PostMessage(
			stream.Context(),
			p,
			in.Message,
			onPostedRouteMessage,
		); err != nil {
			return err
		}
	}
}
