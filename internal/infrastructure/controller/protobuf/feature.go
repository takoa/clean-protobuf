package protobuf

import (
	"context"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/entity/model/generated"
	"github.com/takoa/clean-protobuf/internal/usecase/feature"
)

type GetFeatureHandler struct {
	featureFinder feature.Finder
}

func NewGetFeatureHandler(
	featureFinder feature.Finder,
) *GetFeatureHandler {
	return &GetFeatureHandler{
		featureFinder: featureFinder,
	}
}

func (h *GetFeatureHandler) Invoke(ctx context.Context, point model.Point) (*model.Feature, error) {
	var feature *model.Feature
	onFeatureFound := func(f *model.Feature) error {
		feature = f
		return nil
	}
	if err := h.featureFinder.Find(ctx, model.Rectangle{Hi: point, Lo: point}, onFeatureFound); err != nil {
		return nil, err
	} else if feature == nil {
		return &model.Feature{
			Feature: generated.Feature{
				Latitude:  point.Latitude,
				Longitude: point.Longitude,
			},
		}, nil
	}

	return feature, nil
}

type ListFeaturesHandler struct {
	featureFinder feature.FeatureFinder
}

func NewListFeaturesHandler(
	featureFinder feature.FeatureFinder,
) *ListFeaturesHandler {
	return &ListFeaturesHandler{
		featureFinder: featureFinder,
	}
}

func (s *ListFeaturesHandler) Invoke(ctx context.Context, rect model.Rectangle, send func(*model.Feature) error) error {
	onFeatureFound := func(f *model.Feature) error {
		if err := send(f); err != nil {
			return err
		}
		return nil
	}
	if err := s.featureFinder.Find(ctx, rect, onFeatureFound); err != nil {
		return err
	}

	return nil
}
