// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_demo/internal/dao/internal"
)

// internalHobbyDao is internal type for wrapping internal DAO implements.
type internalHobbyDao = *internal.HobbyDao

// hobbyDao is the data access object for table hobby.
// You can define custom methods on it to extend its functionality as you wish.
type hobbyDao struct {
	internalHobbyDao
}

var (
	// Hobby is globally public accessible object for table hobby operations.
	Hobby = hobbyDao{
		internal.NewHobbyDao(),
	}
)

// Fill with you ideas below.
