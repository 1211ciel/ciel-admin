// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-09 17:01:42
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id        int         `json:"id"        description:""`
	Pid       int         `json:"pid"       description:""`
	Name      string      `json:"name"      description:""`
	Path      string      `json:"path"      description:""`
	Type      int         `json:"type"      description:"1normal 2group"`
	Sort      float64     `json:"sort"      description:""`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
