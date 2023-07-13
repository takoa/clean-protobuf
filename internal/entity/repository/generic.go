package repository

import (
	"context"
)

type Repository[T any] interface {
	Count(ctx context.Context) (count int64, err error)
	Delete(ctx context.Context, models ...*T) (rowsAffected int64, err error)
	Find(ctx context.Context, limit int, offset int, orderBy string) (result []*T, count int64, err error)
	FindByID(ctx context.Context, id string) (result *T, err error)
	FindByIDs(ctx context.Context, orderBy string, ids []string) (result []*T, err error)
	FindByMap(ctx context.Context, limit int, offset int, orderBy string, conditions map[string]interface{}) (result []*T, count int64, err error)
	FindByStruct(ctx context.Context, limit int, offset int, orderBy string, condition *T) (result []*T, count int64, err error)
	First(ctx context.Context, orderBy string) (result *T, err error)
	Last(ctx context.Context, orderBy string) (result *T, err error)
	Save(ctx context.Context, models ...*T) (rowsAffected int64, err error)
}
