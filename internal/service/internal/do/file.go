// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-09 17:01:42
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure of table s_file for DAO operations like Where/Data.
type File struct {
	g.Meta    `orm:"table:s_file, do:true"`
	Id        interface{} //
	Url       interface{} //
	Group     interface{} //
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
