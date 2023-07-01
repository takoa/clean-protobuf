package route

import (
	"context"
	"math"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
	"github.com/takoa/clean-protobuf/internal/usecase/feature"
	"golang.org/x/xerrors"
)

// InformationGetter is an interface for recording a route composited of a sequence of points.
type InformationGetter interface {
	// GetInformation finds and calculates some information on the given point, then calls onPointRecorded passing the information to it.
	GetInformation(
		ctx context.Context,
		previousPoint *model.Point,
		point *model.Point,
		onPointAdded OnPointAdded,
	) error
}

type OnPointAdded func(point *model.Point, matchedFeature *model.Feature, addedDistance int32) error

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
	previousPoint *model.Point,
	point *model.Point,
	onPointAdded OnPointAdded,
) error {
	if point == nil {
		return xerrors.New("nil point")
	}

	var matchedFeature *model.Feature
	onFeatureFound := func(f *model.Feature) error {
		matchedFeature = f
		return nil
	}
	if err := r.featureFinder.Find(ctx, &model.Rectangle{Hi: point, Lo: point}, onFeatureFound); err != nil {
		return err
	}

	var addedDistance int32
	if previousPoint != nil {
		addedDistance = calcDistance(previousPoint, point)
	}

	if err := onPointAdded(point, matchedFeature, addedDistance); err != nil {
		return err
	}

	return nil
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// calcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
func calcDistance(p1 *model.Point, p2 *model.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// MessagePoster is an interface for posting a message to a point.
type MessagePoster interface {
	// PostMessage posts the given message to the given point, then calls onPostedMessage for each message posted on the point.
	PostMessage(
		ctx context.Context,
		point *model.Point,
		message string,
		onPostedMessage OnPostedMessage,
	) error
}

type OnPostedMessage func(message string) error

type RouteMessagePoster struct {
	messages repository.Messages
}

func NewRouteMessagePoster(
	messages repository.Messages,
) *RouteMessagePoster {
	return &RouteMessagePoster{
		messages: messages,
	}
}

func (r *RouteMessagePoster) PostMessage(
	ctx context.Context,
	point *model.Point,
	message string,
	onPostedMessage OnPostedMessage,
) error {
	if err := r.messages.Create(ctx, point, message); err != nil {
		return err
	}

	messages, err := r.messages.Find(ctx, point)
	if err != nil {
		return err
	}

	for _, message := range messages {
		if err := onPostedMessage(message); err != nil {
			return err
		}
	}

	return nil
}
