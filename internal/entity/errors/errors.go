package errors

import "golang.org/x/xerrors"

var (
	ErrNilArgument = xerrors.New("nil argument")
)
