// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure for table login_log.
type LoginLog struct {
	Id        int         `json:"id"        description:""`
	Uid       int         `json:"uid"       description:""`
	Ip        string      `json:"ip"        description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
