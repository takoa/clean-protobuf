package converter

import (
	"github.com/takoa/clean-protobuf/internal/entity/errors"
	"github.com/takoa/clean-protobuf/internal/entity/model"
	routeguidev1 "github.com/takoa/clean-protobuf/internal/pkg/protobuf/routeguide/v1"
	"golang.org/x/xerrors"
)

func ToModelRectangle(r *routeguidev1.Rectangle) (model.Rectangle, error) {
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

func ToGRPCRectangle(f model.Rectangle) *routeguidev1.Rectangle {
	return &routeguidev1.Rectangle{
		Hi: ToGRPCPoint(f.Hi),
		Lo: ToGRPCPoint(f.Lo),
	}
}
