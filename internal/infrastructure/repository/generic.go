package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Count(ctx context.Context) (count int64, err error) {
	if err = r.DB.WithContext(ctx).Count(&count).Error; err != nil {
		return
	}
	return
}

func (r *Repository[T]) Delete(ctx context.Context, models ...*T) (rowsAffected int64, err error) {
	tx := r.DB.WithContext(ctx).Delete(&models)
	return tx.RowsAffected, tx.Error
}

// Find records without conditions while applying a limit and offset.
// `limit` and `offset` can be cancelled with -1.
func (r *Repository[T]) Find(ctx context.Context, limit int, offset int, orderBy string) (result []*T, count int64, err error) {
	return r.FindByStruct(ctx, limit, offset, orderBy, nil)
}

func (r *Repository[T]) FindByID(ctx context.Context, id string) (result *T, err error) {
	tx := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (r *Repository[T]) FindByIDs(ctx context.Context, orderBy string, ids []string) (result []*T, err error) {
	tx := r.DB.WithContext(ctx).
		Where("id in ?", ids).
		Order(orderBy).
		Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (r *Repository[T]) FindByMap(ctx context.Context, limit int, offset int, orderBy string, conditions map[string]interface{}) (result []*T, count int64, err error) {
	tx := r.DB.WithContext(ctx).
		Where(conditions).
		Order(orderBy).
		Limit(limit).
		Offset(offset).
		Find(&result)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	count, err = r.lazyCount(ctx, result, limit, offset)
	return
}

func (r *Repository[T]) FindByStruct(ctx context.Context, limit int, offset int, orderBy string, condition *T) (result []*T, count int64, err error) {
	tx := r.DB.WithContext(ctx).
		Where(condition).
		Order(orderBy).
		Limit(limit).
		Offset(offset).
		Find(&result)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	count, err = r.lazyCount(ctx, result, limit, offset)
	return
}

func (r *Repository[T]) First(ctx context.Context, orderBy string) (result *T, err error) {
	tx := r.DB.WithContext(ctx).
		Order(orderBy).
		First(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (r *Repository[T]) Last(ctx context.Context, orderBy string) (result *T, err error) {
	tx := r.DB.WithContext(ctx).
		Order(orderBy).
		Last(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return
}

func (r *Repository[T]) Save(ctx context.Context, models ...*T) (rowsAffected int64, err error) {
	tx := r.DB.WithContext(ctx).Save(models)
	return tx.RowsAffected, tx.Error
}

func (r *Repository[T]) lazyCount(ctx context.Context, models []*T, limit int, offset int) (count int64, err error) {
	if limit < 0 {
		count = int64(len(models))
		return
	}

	if size := len(models); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Count(ctx)
	return
}
