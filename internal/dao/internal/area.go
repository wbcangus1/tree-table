// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AreaDao is the data access object for table area.
type AreaDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns AreaColumns // columns contains all the column names of Table for convenient usage.
}

// AreaColumns defines and stores column names for table area.
type AreaColumns struct {
	Id        string // 区域ID
	Name      string // 名称
	ParentId  string // 父区域ID
	NodePath  string // 节点路径, 根节点id/../上级节点id/当前节点id
	Sort      string // 同层级排序
	Level     string // 节点所在层级
	CreatedAt string // 创建时间
	UpdatedAt string // 最后更新时间
}

// areaColumns holds the columns for table area.
var areaColumns = AreaColumns{
	Id:        "id",
	Name:      "name",
	ParentId:  "parent_id",
	NodePath:  "node_path",
	Sort:      "sort",
	Level:     "level",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAreaDao creates and returns a new DAO object for table data access.
func NewAreaDao() *AreaDao {
	return &AreaDao{
		group:   "default",
		table:   "area",
		columns: areaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AreaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AreaDao) Columns() AreaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AreaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AreaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
