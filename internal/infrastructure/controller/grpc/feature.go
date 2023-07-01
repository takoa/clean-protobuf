package grpc

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
)

func ToModelFeature(f *api.Feature) *model.Feature {
	if f == nil {
		return nil
	}
	return &model.Feature{
		Location: ToModelPoint(f.Location),
		Name:     f.Name,
	}
}

func ToGRPCFeature(f *model.Feature) *api.Feature {
	if f == nil {
		return nil
	}
	return &api.Feature{
		Name:     f.Name,
		Location: ToGRPCPoint(f.Location),
	}
}
