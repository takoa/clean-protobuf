package converter

import (
	"github.com/takoa/clean-protobuf/internal/entity/model"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
)

func ToGRPCFeature(f *model.Feature) *routeguidev1.Feature {
	if f == nil {
		return nil
	}
	return &routeguidev1.Feature{
		Name:     f.Name,
		Location: ToGRPCPoint(model.Point{Latitude: f.Latitude, Longitude: f.Longitude}),
	}
}
