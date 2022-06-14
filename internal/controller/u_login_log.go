// =================================================================================
// This is auto-generated by FreeKey Admin at 2022-06-14 12:27:52. For more information please go to https://github.com/1211ciel/ciel-admin
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

type cLoginLog struct{ *config.Search }

var LoginLog = &cLoginLog{Search: &config.Search{
	T1: "u_login_log", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*config.Field{},
}}

func (c *cLoginLog) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/u/u_login_log.html", g.Map{"icon": icon})
}
func (c *cLoginLog) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cLoginLog) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cLoginLog) Post(r *ghttp.Request) {
	d := entity.LoginLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cLoginLog) Put(r *ghttp.Request) {
	d := entity.LoginLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cLoginLog) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
