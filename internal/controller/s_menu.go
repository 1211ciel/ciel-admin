// =================================================================================
// This is auto-generated by FreeKey Admin at 2022-7-7 https://github.com/1211ciel/ciel-admin
// =================================================================================

package controller

import (
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type menu struct{ *config.Search }

var Menu = &menu{Search: &config.Search{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []*config.Field{
		{Name: "pid"},
		{Name: "status"},
		{Name: "name", SearchType: 2},
		{Name: "path", SearchType: 2},
	},
}}

func (c *menu) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/menu2.html", g.Map{"icon": icon})
}
func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *menu) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *menu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
