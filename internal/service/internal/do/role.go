// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-09 17:01:42
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table s_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:s_role, do:true"`
	Id        interface{} //
	Name      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
