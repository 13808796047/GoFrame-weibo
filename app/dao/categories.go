// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"GoFrame-weibo/app/dao/internal"
)

// categoriesDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type categoriesDao struct {
	*internal.CategoriesDao
}

var (
	// Categories is globally public accessible object for table categories operations.
	Categories = categoriesDao{
		internal.NewCategoriesDao(),
	}
)

// Fill with you ideas below.
