// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// LinksDao is the manager for logic model data accessing and custom defined data operations functions management.
type LinksDao struct {
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns LinksColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// LinksColumns defines and stores column names for table links.
type LinksColumns struct {
	Id        string //
	Title     string // 资源的描述
	Link      string // 资源的链接
	CreatedAt string //
	UpdatedAt string //
}

//  linksColumns holds the columns for table links.
var linksColumns = LinksColumns{
	Id:        "id",
	Title:     "title",
	Link:      "link",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewLinksDao creates and returns a new DAO object for table data access.
func NewLinksDao() *LinksDao {
	return &LinksDao{
		Group:   "default",
		Table:   "links",
		Columns: linksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LinksDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LinksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LinksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
