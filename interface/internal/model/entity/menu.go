// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id        uint64      `json:"id"        description:""`
	Pid       int64       `json:"pid"       description:""`
	Name      string      `json:"name"      description:""`
	Path      string      `json:"path"      description:""`
	Icon      string      `json:"icon"      description:""`
	Type      uint        `json:"type"      description:"0 normal,1 group,2 divide"`
	Sort      float64     `json:"sort"      description:""`
	Status    int         `json:"status"    description:"-1 off 1 ok"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
