// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_demo/internal/dao/internal"
)

// internalEmpDao is internal type for wrapping internal DAO implements.
type internalEmpDao = *internal.EmpDao

// empDao is the data access object for table emp.
// You can define custom methods on it to extend its functionality as you wish.
type empDao struct {
	internalEmpDao
}

var (
	// Emp is globally public accessible object for table emp operations.
	Emp = empDao{
		internal.NewEmpDao(),
	}
)

// Fill with you ideas below.
