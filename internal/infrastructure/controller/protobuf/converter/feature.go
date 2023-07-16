package converter

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
)

func ToGRPCFeature(f *model.Feature) *api.Feature {
	if f == nil {
		return nil
	}
	return &api.Feature{
		Name:     f.Name,
		Location: ToGRPCPoint(model.Point{Latitude: f.Latitude, Longitude: f.Longitude}),
	}
}
