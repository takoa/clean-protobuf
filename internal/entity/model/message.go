package model

import (
	"github.com/takoa/clean-protobuf/internal/entity/model/generated"
)

type Message struct {
	generated.Message

	Feature Feature
}
