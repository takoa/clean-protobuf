package converter

import (
	"github.com/takoa/clean-protobuf/api"
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	"golang.org/x/xerrors"
)

func ToModelRectangle(r *api.Rectangle) (model.Rectangle, error) {
	if r == nil {
		return model.Rectangle{}, xerrors.Errorf("r: %w", errors.ErrNilArgument)
	}

	hi, err := ToModelPoint(r.Hi)
	if err != nil {
		return model.Rectangle{}, xerrors.Errorf("r.Hi: %w", errors.ErrNilArgument)
	}

	lo, err := ToModelPoint(r.Lo)
	if err != nil {
		return model.Rectangle{}, xerrors.Errorf("r.Lo: %w", errors.ErrNilArgument)
	}
	return model.Rectangle{
		Hi: hi,
		Lo: lo,
	}, nil
}

func ToGRPCRectangle(f model.Rectangle) *api.Rectangle {
	return &api.Rectangle{
		Hi: ToGRPCPoint(f.Hi),
		Lo: ToGRPCPoint(f.Lo),
	}
}
