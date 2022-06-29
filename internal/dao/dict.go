// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao/internal"
	"ciel-admin/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// internalDictDao is internal type for wrapping internal DAO implements.
type internalDictDao = *internal.DictDao

// dictDao is the data access object for table s_dict.
// You can define custom methods on it to extend its functionality as you wish.
type dictDao struct {
	internalDictDao
}

var (
	// Dict is globally public accessible object for table s_dict operations.
	Dict = dictDao{
		internal.NewDictDao(),
	}
)

// Fill with you ideas below.

func (d dictDao) GetByKey(ctx context.Context, key string) (*entity.Dict, error) {
	var data entity.Dict
	one, err := d.Ctx(ctx).One("k", key)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}

func (d dictDao) GetValueWithDefault(ctx context.Context, k string, defaultV string) (gdb.Value, error) {
	value, err := d.Ctx(ctx).Value("v", "k", k)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if value.IsEmpty() {
		g.Log().Warningf(ctx, "k%v value is empty", k)
		return gvar.New(defaultV), nil
	}
	return value, nil
}