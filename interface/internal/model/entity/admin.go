// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        uint64      `json:"id"        description:""`
	Rid       uint64      `json:"rid"       description:""`
	Uname     string      `json:"uname"     description:""`
	Icon      string      `json:"icon"      description:""`
	Pwd       string      `json:"pwd"       description:""`
	Status    int         `json:"status"    description:""`
	Desc      string      `json:"desc"      description:""`
	Ex1       string      `json:"ex1"       description:"other info"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
