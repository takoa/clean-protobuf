package protobuf

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/usecase/route"
)

type RecordRouteHandler struct {
	informationGetter route.InformationGetter
}

func NewRecordRouteHandler(
	informationGetter route.InformationGetter,
) *RecordRouteHandler {
	return &RecordRouteHandler{
		informationGetter: informationGetter,
	}
}

func (s *RecordRouteHandler) Invoke(
	ctx context.Context,
	receive func() (model.Point, error),
	onFinished func(pointCount int32, matchedFeatures []*model.Feature, totalDistance int32, elapsedTime time.Duration) error,
) error {
	var pointCount, distance int32
	var matchedFeatures []*model.Feature
	var previousPoint model.Point
	startTime := time.Now()

	for {
		point, err := receive()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}
			return onFinished(pointCount, matchedFeatures, distance, time.Since(startTime))
		}

		onPointAdded := func(p model.Point, matchedFeature *model.Feature, addedDistance int32) error {
			pointCount++
			if matchedFeature != nil {
				matchedFeatures = append(matchedFeatures, matchedFeature)
			}
			distance += addedDistance
			previousPoint = p

			return nil
		}
		if err := s.informationGetter.GetInformation(
			ctx,
			previousPoint,
			point,
			onPointAdded,
		); err != nil {
			return err
		}
	}
}

type RouteChatHandler struct {
	messagePoster route.MessagePoster
}

func NewRouteChatHandler(
	messagePoster route.MessagePoster,
) *RouteChatHandler {
	return &RouteChatHandler{
		messagePoster: messagePoster,
	}
}

func (s *RouteChatHandler) Invoke(
	ctx context.Context,
	receive func() (model.Point, string, error),
	send func(location model.Point, message string) error,
) error {
	for {
		point, newMessage, err := receive()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		onPostedRouteMessage := func(message string) error {
			if err := send(point, message); err != nil {
				return err
			}
			return nil
		}
		if err := s.messagePoster.PostMessage(
			ctx,
			point,
			newMessage,
			onPostedRouteMessage,
		); err != nil {
			return err
		}
	}
}
