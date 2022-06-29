// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThingRecord is the golang structure for table thing_record.
type ThingRecord struct {
	Id         int64       `json:"id"         description:""`
	Tid        int64       `json:"tid"        description:""`
	RecordDate *gtime.Time `json:"recordDate" description:""`
	Score      float64     `json:"score"      description:""`
	Day        int         `json:"day"        description:""`
	Desc       string      `json:"desc"       description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
}