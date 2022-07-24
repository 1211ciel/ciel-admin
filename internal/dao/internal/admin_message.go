// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMessageDao is the data access object for table s_admin_message.
type AdminMessageDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AdminMessageColumns // columns contains all the column names of Table for convenient usage.
}

// AdminMessageColumns defines and stores column names for table s_admin_message.
type AdminMessageColumns struct {
	Id        string //
	FromUid   string //
	ToUid     string //
	Group     string // {"Label":"分组","FieldType":"select","SearchType":1,"Required":1,"Options":"1:系统:tag-info,2:管理员:tag-warning,3:其他:tag-danger"}
	Type      string // {"Label":"类型","FieldType":"select","SearchType":1,"Required":1,"Options":"1:文字:tag-info,2:图片:tag-primary,3:语音:tag-warning,4:视频:tag-success,5:链接:tag-danger"}
	Content   string // {"Label":"内容","SearchType":2,"Required":1}
	Link      string // {"Label":"链接","SearchType":2}
	Status    string // {"Label":"状态","FieldType":"select","Options":"1:未读:tag-warning,2已读:tag-info"}
	CreatedAt string //
	UpdatedAt string //
}

//  adminMessageColumns holds the columns for table s_admin_message.
var adminMessageColumns = AdminMessageColumns{
	Id:        "id",
	FromUid:   "from_uid",
	ToUid:     "to_uid",
	Group:     "group",
	Type:      "type",
	Content:   "content",
	Link:      "link",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminMessageDao creates and returns a new DAO object for table data access.
func NewAdminMessageDao() *AdminMessageDao {
	return &AdminMessageDao{
		group:   "default",
		table:   "s_admin_message",
		columns: adminMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminMessageDao) Columns() AdminMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}