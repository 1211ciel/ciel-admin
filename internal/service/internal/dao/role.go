// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ciel-admin/internal/service/internal/dao/internal"
)

// roleDao is the data access object for table s_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	*internal.RoleDao
}

var (
	// Role is globally public accessible object for table s_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// Fill with you ideas below.