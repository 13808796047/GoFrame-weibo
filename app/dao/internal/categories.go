// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CategoriesDao is the manager for logic model data accessing and custom defined data operations functions management.
type CategoriesDao struct {
	Table   string            // Table is the underlying table name of the DAO.
	Group   string            // Group is the database configuration group name of current DAO.
	Columns CategoriesColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CategoriesColumns defines and stores column names for table categories.
type CategoriesColumns struct {
	Id          string //
	Name        string // 名称
	Description string // 描述
	PostCount   string // 帖子数
	CreatedAt   string //
	UpdatedAt   string //
}

//  categoriesColumns holds the columns for table categories.
var categoriesColumns = CategoriesColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	PostCount:   "post_count",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewCategoriesDao creates and returns a new DAO object for table data access.
func NewCategoriesDao() *CategoriesDao {
	return &CategoriesDao{
		Group:   "default",
		Table:   "categories",
		Columns: categoriesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CategoriesDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CategoriesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CategoriesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
