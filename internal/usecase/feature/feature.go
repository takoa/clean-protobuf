package feature

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/entity/repository"
)

// Finder is an interface for finding a feature.
type Finder interface {
	// Find looks for features within rect, then calls onFeatureFound for each point found.
	// If rect points to an exact point, it looks for the exact match.
	Find(
		ctx context.Context,
		rect model.Rectangle,
		onFeatureFound OnFeatureFound,
	) error
}

type OnFeatureFound func(feature *model.Feature) error

type FeatureFinder struct {
	features repository.Features
}

func NewFeatureFinder(features repository.Features) *FeatureFinder {
	return &FeatureFinder{
		features: features,
	}
}

func (f *FeatureFinder) Find(
	ctx context.Context,
	rect model.Rectangle,
	onFeatureFound OnFeatureFound,
) error {
	features, _, err := f.features.Find(ctx, -1, -1, "")
	if err != nil {
		return err
	}

	if rect.Hi.Equals(rect.Lo) {
		// If rect points to an exact point, treat it as Point and look for the exact match.
		point := rect.Hi
		for _, feature := range features {
			featurePoint := model.Point{Latitude: feature.Latitude, Longitude: feature.Longitude}
			if point.Equals(featurePoint) {
				if err := onFeatureFound(feature); err != nil {
					return err
				}
			}
		}
	} else {
		// If not, look for all features within the area.
		for _, feature := range features {
			featurePoint := model.Point{Latitude: feature.Latitude, Longitude: feature.Longitude}
			if featurePoint.In(rect) {
				if err := onFeatureFound(feature); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
