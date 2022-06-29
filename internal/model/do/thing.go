// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Thing is the golang structure of table f_thing for DAO operations like Where/Data.
type Thing struct {
	g.Meta    `orm:"table:f_thing, do:true"`
	Id        interface{} //
	Name      interface{} //
	Type      interface{} //
	BeginTime *gtime.Time //
	Status    interface{} // 1
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}