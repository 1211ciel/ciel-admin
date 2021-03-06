// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMessage is the golang structure of table s_admin_message for DAO operations like Where/Data.
type AdminMessage struct {
	g.Meta    `orm:"table:s_admin_message, do:true"`
	Id        interface{} //
	FromUid   interface{} //
	ToUid     interface{} //
	Group     interface{} // {"Label":"分组","FieldType":"select","SearchType":1,"Required":1,"Options":"1:系统:tag-info,2:管理员:tag-warning,3:其他:tag-danger"}
	Type      interface{} // {"Label":"类型","FieldType":"select","SearchType":1,"Required":1,"Options":"1:文字:tag-info,2:图片:tag-primary,3:语音:tag-warning,4:视频:tag-success,5:链接:tag-danger"}
	Content   interface{} // {"Label":"内容","SearchType":2,"Required":1}
	Link      interface{} // {"Label":"链接","SearchType":2}
	Status    interface{} // {"Label":"状态","FieldType":"select","Options":"1:未读:tag-warning,2已读:tag-info"}
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
