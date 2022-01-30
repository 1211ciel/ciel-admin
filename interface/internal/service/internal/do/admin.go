// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-30 15:53:54
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table s_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:s_admin, do:true"`
	Id        interface{} //
	Rid       interface{} //
	Uname     interface{} //
	Icon      interface{} //
	Pwd       interface{} //
	Status    interface{} //
	Desc      interface{} //
	Ex1       interface{} // other info
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}