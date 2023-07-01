package repository

import (
	"context"
	"sync"

	"github.com/takoa/clean-protobuf/internal/entity/model"
)

type Messages struct {
	mu              sync.Mutex
	pointMessageMap map[string][]string
}

func NewMessages() *Messages {
	return &Messages{
		pointMessageMap: make(map[string][]string),
	}
}

func (r *Messages) Find(ctx context.Context, p *model.Point) ([]string, error) {
	key := ""
	if p != nil {
		key = p.Serialize()
	}

	r.mu.Lock()
	messages := make([]string, len(r.pointMessageMap[key]))
	copy(messages, r.pointMessageMap[key])
	r.mu.Unlock()

	return messages, nil
}

func (r *Messages) Create(ctx context.Context, p *model.Point, message string) error {
	key := ""
	if p != nil {
		key = p.Serialize()
	}

	r.mu.Lock()
	r.pointMessageMap[key] = append(r.pointMessageMap[key], message)
	r.mu.Unlock()

	return nil
}
