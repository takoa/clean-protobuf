package route

import (
	"context"
	"errors"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/entity/model/generated"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/usecase/feature"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

// InformationGetter is an interface for recording a route composited of a sequence of points.
type InformationGetter interface {
	// GetInformation finds and calculates some information on the given point, then calls onPointRecorded passing the information to it.
	GetInformation(
		ctx context.Context,
		previousPoint model.Point,
		point model.Point,
		onPointAdded OnPointAdded,
	) error
}

type OnPointAdded func(point model.Point, matchedFeature *model.Feature, addedDistance int32) error

type RouteInformationGetter struct {
	featureFinder feature.Finder
}

func NewRouteInformationGetter(
	features repository.Features,
) *RouteInformationGetter {
	return &RouteInformationGetter{
		featureFinder: feature.NewFeatureFinder(features),
	}
}

func (r *RouteInformationGetter) GetInformation(
	ctx context.Context,
	previousPoint model.Point,
	point model.Point,
	onPointAdded OnPointAdded,
) error {
	var matchedFeature *model.Feature
	onFeatureFound := func(f *model.Feature) error {
		matchedFeature = f
		return nil
	}
	if err := r.featureFinder.Find(ctx, model.Rectangle{Hi: point, Lo: point}, onFeatureFound); err != nil {
		return err
	}

	if err := onPointAdded(point, matchedFeature, previousPoint.Distance(point)); err != nil {
		return err
	}

	return nil
}

// MessagePoster is an interface for posting a message to a point.
type MessagePoster interface {
	// PostMessage posts the given message to the given point, then calls onPostedMessage for each message posted on the point.
	PostMessage(
		ctx context.Context,
		point model.Point,
		message string,
		onPostedMessage OnPostedMessage,
	) error
}

type OnPostedMessage func(message string) error

type RouteMessagePoster struct {
	feature  repository.Features
	messages repository.Messages
}

func NewRouteMessagePoster(
	feature repository.Features,
	messages repository.Messages,
) *RouteMessagePoster {
	return &RouteMessagePoster{
		feature:  feature,
		messages: messages,
	}
}

func (r *RouteMessagePoster) PostMessage(
	ctx context.Context,
	point model.Point,
	message string,
	onPostedMessage OnPostedMessage,
) error {
	m := &model.Message{
		Message: generated.Message{
			Body: message,
		},
	}

	f, err := r.feature.FindByPoint(ctx, point)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return xerrors.Errorf("failed to find the point: %w", err)
		}
	}

	if f != nil {
		m.Feature.ID = f.ID
	} else {
		m.Feature.Latitude = point.Latitude
		m.Feature.Longitude = point.Longitude
	}

	if _, err := r.messages.Save(ctx, m); err != nil {
		return xerrors.Errorf("failed to save a message: %w", err)
	}

	messages, err := r.messages.FindByFeature(ctx, &m.Feature, "")
	if err != nil {
		return xerrors.Errorf("failed to find messages: %w", err)
	}

	for _, message := range messages {
		if err := onPostedMessage(message.Body); err != nil {
			return xerrors.Errorf("callback error: %w", err)
		}
	}

	return nil
}
