package view

import (
	"ciel-admin/utility/utils/xurl"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
	"strings"
)

func QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func HasPrefix(str interface{}, suffix string) bool {
	s := gconv.String(str)
	return strings.HasPrefix(s, suffix)
}
func BindFuncMap() gview.FuncMap {
	return gview.FuncMap{
		"hasPrefix":   HasPrefix,
		"toUrlParams": xurl.ToUrlParams,
		"tag":         tag,
		"nodeLevel":   nodeLevel,
		"nodeTime":    nodeTime,
		"nodeWeek":    nodeWeek,
		"nav":         nav,
		"option":      option,
		"chooseSpan":  chooseSpan,
		"img":         img,
		"md":          MD,
	}
}
