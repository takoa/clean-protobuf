package grpc

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/model"
)

func ToModelRectangle(r *api.Rectangle) *model.Rectangle {
	if r == nil {
		return nil
	}
	return &model.Rectangle{
		Hi: ToModelPoint(r.Hi),
		Lo: ToModelPoint(r.Lo),
	}
}

func ToGRPCRectangle(f *model.Rectangle) *api.Rectangle {
	if f == nil {
		return nil
	}
	return &api.Rectangle{
		Hi: ToGRPCPoint(f.Hi),
		Lo: ToGRPCPoint(f.Lo),
	}
}
