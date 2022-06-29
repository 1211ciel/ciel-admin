// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao/internal"
	"ciel-admin/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// internalMenuDao is internal type for wrapping internal DAO implements.
type internalMenuDao = *internal.MenuDao

// menuDao is the data access object for table s_menu.
// You can define custom methods on it to extend its functionality as you wish.
type menuDao struct {
	internalMenuDao
}

var (
	// Menu is globally public accessible object for table s_menu operations.
	Menu = menuDao{
		internal.NewMenuDao(),
	}
)

// Fill with you ideas below.

func (d menuDao) GetByPath(ctx context.Context, path string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := d.Ctx(ctx).One("path", path)
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

func (d menuDao) GetByName(ctx context.Context, name string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := d.Ctx(ctx).One("name", name)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, nil
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