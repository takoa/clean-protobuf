package grpc

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
)

func ToModelPoint(p *api.Point) *model.Point {
	if p == nil {
		return nil
	}
	return &model.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}

func ToGRPCPoint(p *model.Point) *api.Point {
	if p == nil {
		return nil
	}
	return &api.Point{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
}
