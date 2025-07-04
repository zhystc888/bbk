// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"bbk/app/user/internal/dao/internal"
)

// adminUserDao is the data access object for the table bbk_admin_user.
// You can define custom methods on it to extend its functionality as needed.
type adminUserDao struct {
	*internal.AdminUserDao
}

var (
	// AdminUser is a globally accessible object for table bbk_admin_user operations.
	AdminUser = adminUserDao{internal.NewAdminUserDao()}
)

// Add your custom methods and functionality below.
