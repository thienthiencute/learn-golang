package generics

import (
	"context"
	"errors"
	"gocas/pkg/golangviet/utils"

	"github.com/teoit/gosctx/core"
	"gorm.io/gorm"
)

var (
	ErrFilterEmpty = errors.New("filters not empty")
)

type genericDB struct {
	DB *gorm.DB
}

func BaseFilters(ctx context.Context, dbpostgres *gorm.DB, tableName string) *genericDB {
	db := dbpostgres.Table(tableName).WithContext(ctx)
	return &genericDB{
		DB: db,
	}
}

func (gen *genericDB) Preloads(conds *map[string]interface{}) *genericDB {
	if conds == nil {
		return gen
	}
	db := gen.DB
	for query, arg := range *conds {
		if args, ok := arg.([]interface{}); ok {
			db = db.Preload(query, args...)
		} else {
			if arg == nil {
				db = db.Preload(query)
			} else {
				db = db.Preload(query, arg)
			}
		}
	}
	gen.DB = db
	return gen
}

func (gen *genericDB) Columns(column *[]string) *genericDB {
	if column != nil {
		gen.DB = gen.DB.Select(*column)
	}
	return gen
}

func (gen *genericDB) Conditions(conds *map[string]interface{}) *genericDB {
	if conds == nil {
		return gen
	}
	db := gen.DB
	for query, arg := range *conds {
		if args, ok := arg.([]interface{}); ok {
			db = db.Where(query, args...)
		} else {
			db = db.Where(query, arg)
		}

	}
	gen.DB = db
	return gen
}

func (gen *genericDB) OrderBy(orderBys *[]string) *genericDB {
	if orderBys != nil {
		db := gen.DB
		for _, order := range *orderBys {
			db = db.Order(order)
		}
		gen.DB = db
	}
	return gen
}

func (gen *genericDB) Limit(limit int) *genericDB {
	if limit > 0 {
		gen.DB = gen.DB.Limit(limit)
	}
	return gen
}

func (gen *genericDB) Offset(offset int) *genericDB {
	if offset >= 0 {
		gen.DB = gen.DB.Offset(offset)
	}
	return gen
}

/**
 * Find with : column, where, in, notin, like, limit)
 */
func FindGeneric[T any](ctx context.Context, dbpostgres *gorm.DB, tableName string, filters *utils.Filters) ([]*T, error) {
	if filters == nil {
		return nil, ErrFilterEmpty
	}
	db := BaseFilters(ctx, dbpostgres, tableName).
		Preloads(filters.Preloads).
		Columns(filters.Columns).
		Conditions(filters.Conds).
		OrderBy(filters.OrderBy).
		Offset(filters.Offset).
		Limit(filters.PageSize)

	var datas []*T
	result := db.DB.Find(&datas)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, core.ErrRecordNotFound
	}
	return datas, nil
}

/**
 * First with : column, where, in, notin, like)
 */
func FirstGeneric[T any](ctx context.Context, dbpostgres *gorm.DB, tableName string, filters *utils.Filters) (*T, error) {
	if filters == nil {
		return nil, ErrFilterEmpty
	}
	db := BaseFilters(ctx, dbpostgres, tableName).
		Preloads(filters.Preloads).
		Columns(filters.Columns).
		Conditions(filters.Conds).
		OrderBy(filters.OrderBy)

	var datas *T

	result := db.DB.Limit(1).Find(&datas)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, core.ErrRecordNotFound
	}
	return datas, nil
}

/**
 * Max field
 * filter := Filters {
 * 	Conds : &[]string{
 * 		"MAX(level) as maxlevel"
 *  }
 * }
 */
func MaxFieldGeneric(ctx context.Context, dbpostgres *gorm.DB, tableName string, filters *utils.Filters) (*int64, error) {
	if filters == nil {
		return nil, ErrFilterEmpty
	}
	db := BaseFilters(ctx, dbpostgres, tableName).
		Columns(filters.Columns).
		Conditions(filters.Conds)

	var max int64
	db.DB.Scan(&max)
	return &max, nil
}

/**
 *
 */
func UpdatesMapGeneric(ctx context.Context, dbpostgres *gorm.DB, tableName string, filters *utils.UpdateFilters) error {
	if filters == nil {
		return ErrFilterEmpty
	}
	db := BaseFilters(ctx, dbpostgres, tableName).
		Columns(filters.Columns).
		Conditions(filters.Conds)
	if err := db.DB.Updates(*filters.UpdateMap).Error; err != nil {
		return err
	}
	return nil
}

/**
 * Count with: where, in, notin, like
 */
func CountGeneric(ctx context.Context, dbpostgres *gorm.DB, tableName string, filters *utils.Filters) (*int64, error) {
	db := BaseFilters(ctx, dbpostgres, tableName)

	if filters != nil {
		db = db.Conditions(filters.Conds)
	}
	var count int64
	if err := db.DB.Count(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}
