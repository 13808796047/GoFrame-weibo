// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"GoFrame-weibo/app/dao/internal"
)

// failedJobsDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type failedJobsDao struct {
	*internal.FailedJobsDao
}

var (
	// FailedJobs is globally public accessible object for table failed_jobs operations.
	FailedJobs = failedJobsDao{
		internal.NewFailedJobsDao(),
	}
)

// Fill with you ideas below.
